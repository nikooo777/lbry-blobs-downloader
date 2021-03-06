package main

import (
	"fmt"
	"os"
	"sync"

	"blobdownloader/downloader"
	"blobdownloader/http"
	"blobdownloader/quic"
	"blobdownloader/shared"
	"blobdownloader/tcp"

	"github.com/lbryio/lbry.go/v2/extras/errors"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	hash              string
	reflectorAddr     string
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
		Use:   "blobdownloader",
		Short: "download blobs or streams from reflector.",
		Run:   blobDownloader,
		Args:  cobra.RangeArgs(0, 0),
	}
	cmd.Flags().StringVar(&hash, "hash", "c333e168b1adb5b8971af26ca2c882e60e7a908167fa9582b47a044f896484485df9f5a0ada7ef6dc976489301e8049d", "hash of the blob or sdblob")
	cmd.Flags().StringVar(&reflectorAddr, "reflector-address", "cdn.reflector.lbry.com", "the address of the reflector server (without port)")
	cmd.Flags().StringVar(&peerPort, "peer-port", "5567", "the port reflector listens to for TCP peer connections")
	cmd.Flags().StringVar(&quicPort, "quic-port", "5568", "the port reflector listens to for QUIC peer connections")
	cmd.Flags().StringVar(&httpPort, "http-port", "5569", "the port reflector listens to for HTTP connections")
	cmd.Flags().BoolVar(&isStream, "stream", false, "whether the hash is for a stream or not (download whole file)")
	cmd.Flags().IntVar(&concurrentThreads, "concurrent-threads", 1, "Number of concurrent downloads to run")
	cmd.Flags().IntVar(&mode, "mode", 0, "0: only use QUIC, 1: only use TCP, 2: only use HTTP, 3: use all")
	cmd.Flags().BoolVar(&fullTrace, "trace", false, "print all traces")
	cmd.Flags().BoolVar(&build, "build", false, "build the file from the blobs")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func blobDownloader(cmd *cobra.Command, args []string) {
	var err error
	shared.ReflectorPeerServer = reflectorAddr + ":" + peerPort
	shared.ReflectorQuicServer = reflectorAddr + ":" + quicPort
	shared.ReflectorHttpServer = reflectorAddr + ":" + httpPort
	logrus.Println("tcp server: " + shared.ReflectorPeerServer)
	logrus.Println("quic server: " + shared.ReflectorQuicServer)
	logrus.Println("http server: " + shared.ReflectorHttpServer)
	wg := &sync.WaitGroup{}
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
					err = downloader.DownloadAndBuild(hash, fullTrace, downloader.Mode(mode), hash, builtDir)

				} else {
					_, err = downloader.DownloadStream(hash, fullTrace, downloader.Mode(mode))
				}
				if err != nil {
					panic(errors.FullTrace(err))
				}
			} else {
				switch mode {
				case 0:
					_, err = quic.DownloadBlob(hash, fullTrace)
				case 1:
					_, err = tcp.DownloadBlob(hash)
				case 2:
					_, err = http.DownloadBlob(hash, fullTrace)
				case 3:
					logrus.Println("QUIC protocol:")
					_, err = quic.DownloadBlob(hash, fullTrace)
					if err != nil {
						logrus.Error(errors.FullTrace(err))
						os.Exit(1)
					}
					logrus.Println("TCP protocol:")
					_, err = tcp.DownloadBlob(hash)
					if err != nil {
						logrus.Error(errors.FullTrace(err))
						os.Exit(1)
					}
					logrus.Println("HTTP protocol:")
					_, err = http.DownloadBlob(hash, fullTrace)
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
