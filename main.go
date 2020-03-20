package main

import (
	"fmt"
	"os"
	"sync"

	"blobdownloader/quic"
	"blobdownloader/shared"
	"blobdownloader/tcp"

	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/lbry.go/v2/stream"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	hash              string
	reflectorAddr     string
	peerPort          string
	quicPort          string
	isStream          bool
	concurrentThreads int
	mode              int
)

func main() {
	cmd := &cobra.Command{
		Use:   "blobdownloader",
		Short: "download blobs or streams from reflector.",
		Run:   downloader,
		Args:  cobra.RangeArgs(0, 0),
	}
	cmd.Flags().StringVar(&hash, "hash", "58742ec8f86abbaadf11ad45e22a78c01e3f89ac3d9f3f1c0d1b77198d34b52672aad8f908a68c763d6767858761c247", "hash of the blob or sdblob")
	cmd.Flags().StringVar(&reflectorAddr, "reflector-address", "reflector.lbry.com", "the address of the reflector server (without port)")
	cmd.Flags().StringVar(&peerPort, "peer-port", "5567", "the port reflector listens to for TCP peer connections")
	cmd.Flags().StringVar(&quicPort, "quic-port", "5568", "the port reflector listens to for QUIC peer connections")
	cmd.Flags().BoolVar(&isStream, "stream", false, "whether the hash is for a stream or not (download whole file)")
	cmd.Flags().IntVar(&concurrentThreads, "concurrent-threads", 1, "Number of concurrent downloads to run")
	cmd.Flags().IntVar(&mode, "mode", 0, "0: only use QUIC, 1: only use TCP, 2: use both")

	shared.ReflectorPeerServer = reflectorAddr + ":" + peerPort
	shared.ReflectorQuicServer = reflectorAddr + ":" + quicPort
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func downloader(cmd *cobra.Command, args []string) {
	var err error

	wg := &sync.WaitGroup{}
	for i := 0; i < concurrentThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isStream {
				err := downloadStream(hash)
				if err != nil {
					panic(errors.FullTrace(err))
				}
			} else {
				switch mode {
				case 0:
					_, err = quic.DownloadBlob(hash)
				case 1:
					_, err = tcp.DownloadBlob(hash)
				case 2:
					logrus.Println("QUIC protocol:")
					_, err = quic.DownloadBlob(hash)
					if err != nil {
						logrus.Error(errors.FullTrace(err))
						os.Exit(1)
					}
					logrus.Println("TCP protocol:")
					_, err = tcp.DownloadBlob(hash)
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

func downloadStream(hash string) error {
	var blob *stream.Blob
	var err error
	switch mode {
	case 0, 2:
		blob, err = quic.DownloadBlob(hash)
	case 1:
		blob, err = tcp.DownloadBlob(hash)
	}
	if err != nil {
		return err
	}
	sdb := &stream.SDBlob{}
	err = sdb.FromBlob(*blob)

	if err != nil {
		return err
	}

	switch mode {
	case 0:
		speed := quic.DownloadStream(sdb)
		logrus.Printf("QUIC protocol downloaded at an average of %.2f MiB/s", speed/1024/104)
	case 1:
		speed := tcp.DownloadStream(sdb)
		logrus.Printf("TCP protocol downloaded at an average of %.2f MiB/s", speed/1024/104)

	case 2:
		speed := quic.DownloadStream(sdb)
		logrus.Printf("QUIC protocol downloaded at an average of %.2f MiB/s", speed/1024/104)
		speed = tcp.DownloadStream(sdb)
		logrus.Printf("TCP protocol downloaded at an average of %.2f MiB/s", speed/1024/104)
	}
	return nil
}
