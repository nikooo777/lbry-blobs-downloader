package downloader

import (
	"blobdownloader/http"
	"blobdownloader/quic"
	"blobdownloader/shared"
	"blobdownloader/tcp"

	"github.com/lbryio/lbry.go/v2/stream"
	"github.com/sirupsen/logrus"
)

type Mode int

const (
	UDP Mode = iota
	TCP
	BOTH
)

func DownloadStream(sdHash string, fullTrace bool, mode Mode) (*stream.SDBlob, error) {
	var blob *stream.Blob
	var err error
	switch mode {
	case 0, 3:
		blob, err = quic.DownloadBlob(sdHash, fullTrace)
	case 1:
		blob, err = tcp.DownloadBlob(sdHash)
	case 2:
		blob, err = http.DownloadBlob(sdHash, fullTrace)
	}
	if err != nil {
		return nil, err
	}
	sdb := &stream.SDBlob{}
	err = sdb.FromBlob(*blob)

	if err != nil {
		return nil, err
	}

	switch mode {
	case 0:
		speed := quic.DownloadStream(sdb, fullTrace)
		logrus.Printf("QUIC protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
	case 1:
		speed := tcp.DownloadStream(sdb)
		logrus.Printf("TCP protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
	case 2:
		speed := http.DownloadStream(sdb, fullTrace)
		logrus.Printf("HTTP protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
	case 3:
		speed := quic.DownloadStream(sdb, fullTrace)
		logrus.Printf("QUIC protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
		speed = tcp.DownloadStream(sdb)
		logrus.Printf("TCP protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
		speed = http.DownloadStream(sdb, fullTrace)
		logrus.Printf("HTTP protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
	}
	return sdb, nil
}

func DownloadAndBuild(sdHash string, fullTrace bool, mode Mode, fileName string, destinationPath string) error {
	sdBlob, err := DownloadStream(sdHash, fullTrace, mode)
	if err != nil {
		return err
	}
	return shared.BuildStream(sdBlob, fileName, destinationPath, "./downloads/")
}
