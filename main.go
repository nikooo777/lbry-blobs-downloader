package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/nikooo777/lbry-blobs-downloader/downloader"
	"github.com/nikooo777/lbry-blobs-downloader/protocols/http"
	"github.com/nikooo777/lbry-blobs-downloader/protocols/http3"
	"github.com/nikooo777/lbry-blobs-downloader/protocols/tcp"
	"github.com/nikooo777/lbry-blobs-downloader/shared"

	"github.com/lbryio/lbry.go/v2/extras/errors"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	hash              string
	upstreamReflector string
	peerPort          string
	quicPort          string
	httpPort          string
	isStream          bool
	concurrentThreads int
	mode              int
	fullTrace         bool
	build             bool
)

func main() {
	cmd := &cobra.Command{
		Use:   "blobsdownloader",
		Short: "download blobs or streams from reflector.",
		Run:   blobsDownloader,
		Args:  cobra.RangeArgs(0, 0),
	}
	cmd.Flags().StringVar(&hash, "hash", "c333e168b1adb5b8971af26ca2c882e60e7a908167fa9582b47a044f896484485df9f5a0ada7ef6dc976489301e8049d", "hash of the blob or sdblob")
	cmd.Flags().StringVar(&upstreamReflector, "upstream-reflector", "reflector.lbry.com", "the address of the reflector server (without port)")
	cmd.Flags().StringVar(&peerPort, "peer-port", "5567", "the port reflector listens to for TCP peer connections")
	cmd.Flags().StringVar(&quicPort, "http3-port", "5568", "the port reflector listens to for HTTP3 peer connections")
	cmd.Flags().StringVar(&httpPort, "http-port", "5569", "the port reflector listens to for HTTP connections")
	cmd.Flags().BoolVar(&isStream, "stream", false, "whether the hash is for a stream or not (download whole file)")
	cmd.Flags().IntVar(&concurrentThreads, "concurrent-threads", 1, "Number of concurrent downloads to run")
	cmd.Flags().IntVar(&mode, "mode", 0, "0: HTTP3, 1: TCP (LBRY), 2: HTTP, 3: use all")
	cmd.Flags().BoolVar(&fullTrace, "trace", false, "print all traces")
	cmd.Flags().BoolVar(&build, "build", false, "build the file from the blobs")

	logrus.SetLevel(logrus.DebugLevel)
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func blobsDownloader(cmd *cobra.Command, args []string) {
	var err error
	serverParams := downloader.ServerParams{
		TcpServerAddress:   shared.DefaultReflectorPeerServer,
		TcpServerPort:      shared.DefaultReflectorPeerPort,
		Http3ServerAddress: shared.DefaultReflectorQuicServer,
		Http3ServerPort:    shared.DefaultReflectorQuicPort,
		HttpServerAddress:  shared.DefaultReflectorHttpServer,
		HttpServerPort:     shared.DefaultReflectorHttpPort,
	}

	tcpDownloader := tcp.NewTcpBlobDownloader(shared.DefaultReflectorPeerServer, shared.DefaultReflectorPeerPort)
	http3Downloader := http3.NewHttp3BlobDownloader(shared.DefaultReflectorQuicServer, shared.DefaultReflectorQuicPort)
	HttpDownloader := http.NewHttpBlobDownloader(shared.DefaultReflectorHttpServer, shared.DefaultReflectorHttpPort)

	wg := &sync.WaitGroup{}
	downloadPath := "./downloads"
	for i := 0; i < concurrentThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isStream {
				if build {
					builtDir := "./built_downloads/"
					err = os.MkdirAll(builtDir, os.ModePerm)
					if err != nil {
						panic(errors.FullTrace(err))
					}
					err = downloader.DownloadAndBuild(hash, fullTrace, downloader.Mode(mode), hash, builtDir, serverParams)

				} else {
					_, err = downloader.DownloadStream(hash, fullTrace, downloader.Mode(mode), downloadPath, serverParams)
				}
				if err != nil {
					logrus.Error(errors.FullTrace(err))
					os.Exit(1)
				}
			} else {
				switch mode {
				case 0:
					_, err = http3Downloader.DownloadBlob(hash, fullTrace, downloadPath)
				case 1:
					_, err = tcpDownloader.DownloadBlob(hash, fullTrace, downloadPath)
				case 2:
					_, err = HttpDownloader.DownloadBlob(hash, fullTrace, downloadPath)
				case 3:
					logrus.Debugln("HTTP3 protocol:")
					_, err = http3Downloader.DownloadBlob(hash, fullTrace, downloadPath)
					if err != nil {
						logrus.Error(errors.FullTrace(err))
						os.Exit(1)
					}
					logrus.Debugln("TCP protocol:")
					_, err = tcpDownloader.DownloadBlob(hash, fullTrace, downloadPath)
					if err != nil {
						logrus.Error(errors.FullTrace(err))
						os.Exit(1)
					}
					logrus.Debugln("HTTP protocol:")
					_, err = HttpDownloader.DownloadBlob(hash, fullTrace, downloadPath)
					if err != nil {
						logrus.Error(errors.FullTrace(err))
						os.Exit(1)
					}
				}
				if err != nil {
					logrus.Error(errors.FullTrace(err))
					os.Exit(1)
				}
			}
		}()
	}
	wg.Wait()
}
