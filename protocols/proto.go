package protocols

import (
	"github.com/lbryio/lbry.go/v2/stream"
)

type Proto interface {
	DownloadBlob(hash string, fullTrace bool, downloadPath string) (*stream.Blob, error)

	//DownloadStream downloads a stream and returns the speed in bytes per second
	DownloadStream(blob *stream.SDBlob, fullTrace bool, downloadPath string) float64
}
