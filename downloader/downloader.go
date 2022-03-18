package downloader

import (
	"io/ioutil"

	"github.com/nikooo777/lbry-blobs-downloader/protocols/http"
	"github.com/nikooo777/lbry-blobs-downloader/protocols/http3"
	"github.com/nikooo777/lbry-blobs-downloader/protocols/tcp"
	"github.com/nikooo777/lbry-blobs-downloader/shared"

	"github.com/lbryio/lbry.go/v2/extras/errors"
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

type ServerParams struct {
	TcpServerAddress   string
	TcpServerPort      string
	Http3ServerAddress string
	Http3ServerPort    string
	HttpServerAddress  string
	HttpServerPort     string
}

func DownloadStream(sdHash string, fullTrace bool, mode Mode, downloadPath string, servers ServerParams) (*stream.SDBlob, error) {
	var blob *stream.Blob
	var err error

	tcpDownloader := tcp.NewTcpBlobDownloader(servers.TcpServerAddress, servers.TcpServerPort)
	http3Downloader := http3.NewHttp3BlobDownloader(servers.Http3ServerAddress, servers.Http3ServerPort)
	HttpDownloader := http.NewHttpBlobDownloader(servers.HttpServerAddress, servers.HttpServerPort)
	switch mode {
	case UDP:
		blob, err = http3Downloader.DownloadBlob(sdHash, fullTrace, downloadPath)
	case TCP:
		blob, err = tcpDownloader.DownloadBlob(sdHash, fullTrace, downloadPath)
	case HTTP, ALL:
		blob, err = HttpDownloader.DownloadBlob(sdHash, fullTrace, downloadPath)
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
		speed := http3Downloader.DownloadStream(sdb, fullTrace, downloadPath)
		logrus.Debugf("QUIC protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
	case TCP:
		speed := tcpDownloader.DownloadStream(sdb, fullTrace, downloadPath)
		logrus.Debugf("TCP protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
	case HTTP:
		speed := HttpDownloader.DownloadStream(sdb, fullTrace, downloadPath)
		logrus.Debugf("HTTP protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
	case ALL:
		speed := http3Downloader.DownloadStream(sdb, fullTrace, downloadPath)
		logrus.Debugf("QUIC protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
		speed = tcpDownloader.DownloadStream(sdb, fullTrace, downloadPath)
		logrus.Debugf("TCP protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
		speed = HttpDownloader.DownloadStream(sdb, fullTrace, downloadPath)
		logrus.Debugf("HTTP protocol downloaded at an average of %.2f MiB/s", speed/1024/1024)
	}
	return sdb, nil
}

func DownloadAndBuild(sdHash string, fullTrace bool, mode Mode, fileName string, destinationPath string, serverParams ServerParams) error {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return errors.Err(err)
	}

	sdBlob, err := DownloadStream(sdHash, fullTrace, mode, tmpDir, serverParams)
	if err != nil {
		return err
	}
	return shared.BuildStream(sdBlob, fileName, destinationPath, tmpDir)
}
