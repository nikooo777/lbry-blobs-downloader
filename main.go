package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/nikooo777/lbry-blobs-downloader/chainquery"
	"github.com/nikooo777/lbry-blobs-downloader/downloader"
	"github.com/nikooo777/lbry-blobs-downloader/http"
	"github.com/nikooo777/lbry-blobs-downloader/quic"
	"github.com/nikooo777/lbry-blobs-downloader/shared"
	"github.com/nikooo777/lbry-blobs-downloader/tcp"

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
	rename            bool
)

func main() {
	cmd := &cobra.Command{
		Use:   "blobsdownloader",
		Short: "download blobs or streams from reflector.",
		Run:   blobsDownloader,
		Args:  cobra.RangeArgs(0, 0),
	}
	cmd.Flags().StringVar(&hash, "hash", "c333e168b1adb5b8971af26ca2c882e60e7a908167fa9582b47a044f896484485df9f5a0ada7ef6dc976489301e8049d", "hash of the blob or sdblob")
	cmd.Flags().StringVar(&upstreamReflector, "upstream-reflector", "blobcache-eu.odycdn.com", "the address of the reflector server (without port)")
	cmd.Flags().StringVar(&peerPort, "peer-port", "5567", "the port reflector listens to for TCP peer connections")
	cmd.Flags().StringVar(&quicPort, "http3-port", "5568", "the port reflector listens to for HTTP3 peer connections")
	cmd.Flags().StringVar(&httpPort, "http-port", "5569", "the port reflector listens to for HTTP connections")
	cmd.Flags().BoolVar(&isStream, "stream", false, "whether the hash is for a stream or not (download whole file)")
	cmd.Flags().IntVar(&concurrentThreads, "concurrent-threads", runtime.NumCPU(), "Number of concurrent downloads to run")
	cmd.Flags().IntVar(&mode, "mode", 2, "0: HTTP3, 1: TCP (LBRY), 2: HTTP, 3: use all")
	cmd.Flags().BoolVar(&fullTrace, "trace", false, "print all traces")
	cmd.Flags().BoolVar(&build, "build", false, "build the file from the blobs")
	cmd.Flags().BoolVar(&rename, "rename", false, "attempt renaming the downloaded file to its original name")

	logrus.SetLevel(logrus.DebugLevel)
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func blobsDownloader(cmd *cobra.Command, args []string) {
	var err error
	shared.ReflectorPeerServer = upstreamReflector + ":" + peerPort
	shared.ReflectorQuicServer = upstreamReflector + ":" + quicPort
	shared.ReflectorHttpServer = upstreamReflector + ":" + httpPort
	logrus.Debugf("tcp server: %s", shared.ReflectorPeerServer)
	logrus.Debugf("http3 server: %s", shared.ReflectorQuicServer)
	logrus.Debugf("http server: %s", shared.ReflectorHttpServer)
	downloadPath := "./downloads"
	if isStream {
		if build {
			builtDir := "./built_downloads/"
			err = os.MkdirAll(builtDir, os.ModePerm)
			if err != nil {
				panic(errors.FullTrace(err))
			}
			fileName := hash
			if rename {
				name, err := chainquery.GetOriginalName(hash)
				if err == nil {
					fileName = name
				} else {
					logrus.Warnf("Failed to get original name for %s: %s", hash, err.Error())
				}
			}
			err = downloader.DownloadAndBuild(hash, fullTrace, downloader.Mode(mode), fileName, builtDir, concurrentThreads)
		} else {
			_, err = downloader.DownloadStream(hash, fullTrace, downloader.Mode(mode), downloadPath, concurrentThreads)
		}
		if err != nil {
			logrus.Error(errors.FullTrace(err))
			os.Exit(1)
		}
	} else {
		switch mode {
		case 0:
			_, err = quic.DownloadBlob(hash, fullTrace, downloadPath)
		case 1:
			_, err = tcp.DownloadBlob(hash, downloadPath)
		case 2:
			_, err = http.DownloadBlob(hash, fullTrace, &downloadPath)
		case 3:
			logrus.Debugln("HTTP3 protocol:")
			_, err = quic.DownloadBlob(hash, fullTrace, downloadPath)
			if err != nil {
				logrus.Error(errors.FullTrace(err))
				os.Exit(1)
			}
			logrus.Debugln("TCP protocol:")
			_, err = tcp.DownloadBlob(hash, downloadPath)
			if err != nil {
				logrus.Error(errors.FullTrace(err))
				os.Exit(1)
			}
			logrus.Debugln("HTTP protocol:")
			_, err = http.DownloadBlob(hash, fullTrace, &downloadPath)
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
}
