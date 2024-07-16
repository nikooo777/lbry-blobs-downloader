package downloader

import (
	"os"

	"github.com/nikooo777/lbry-blobs-downloader/http"
	"github.com/nikooo777/lbry-blobs-downloader/quic"
	"github.com/nikooo777/lbry-blobs-downloader/shared"
	"github.com/nikooo777/lbry-blobs-downloader/tcp"

	"github.com/lbryio/lbry.go/v2/stream"
	"github.com/sirupsen/logrus"
)

type Mode int

const (
	UDP Mode = iota
	TCP
	HTTP
	ALL
)

func DownloadStream(sdHash string, fullTrace bool, mode Mode, downloadPath string, threads int) (*stream.SDBlob, error) {
	var blob *stream.Blob
	var err error
	switch mode {
	case UDP:
		blob, err = quic.DownloadBlob(sdHash, fullTrace, downloadPath)
	case TCP:
		blob, err = tcp.DownloadBlob(sdHash, downloadPath)
	case HTTP, ALL:
		blob, err = http.DownloadBlob(sdHash, fullTrace, &downloadPath)
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
	case UDP:
		speed := quic.DownloadStream(sdb, fullTrace, downloadPath)
		logrus.Debugf("QUIC protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
	case TCP:
		speed := tcp.DownloadStream(sdb, downloadPath)
		logrus.Debugf("TCP protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
	case HTTP:
		speed := http.DownloadStream(sdb, fullTrace, downloadPath, threads)
		logrus.Debugf("HTTP protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
	case ALL:
		speed := quic.DownloadStream(sdb, fullTrace, downloadPath)
		logrus.Debugf("QUIC protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
		speed = tcp.DownloadStream(sdb, downloadPath)
		logrus.Debugf("TCP protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
		speed = http.DownloadStream(sdb, fullTrace, downloadPath, 0)
		logrus.Debugf("HTTP protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
	}
	return sdb, nil
}

func DownloadAndBuild(sdHash string, fullTrace bool, mode Mode, fileName string, destinationPath string, threads int) error {
	tmpDir := os.TempDir()
	sdBlob, err := DownloadStream(sdHash, fullTrace, mode, tmpDir, threads)
	if err != nil {
		return err
	}
	return shared.BuildStream(sdBlob, fileName, destinationPath, tmpDir)
}
