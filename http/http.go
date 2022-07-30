package http

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/nikooo777/lbry-blobs-downloader/shared"

	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/lbry.go/v2/stream"
	"github.com/lbryio/reflector.go/store"

	"github.com/sirupsen/logrus"
)

func DownloadBlob(hash string, fullTrace bool, downloadPath string) (*stream.Blob, error) {
	bStore := GetHttpBlobStore()
	start := time.Now()
	blob, trace, err := bStore.Get(hash)
	if fullTrace {
		logrus.Debugln(trace.String())
	}
	if err != nil {
		err = errors.Prefix(hash, err)
		return nil, errors.Err(err)
	}
	elapsed := time.Since(start)
	logrus.Debugf("[H] download time: %d ms\tSpeed: %.2f MB/s", elapsed.Milliseconds(), (float64(len(blob))/(1024*1024))/elapsed.Seconds())
	err = os.MkdirAll(downloadPath, os.ModePerm)
	if err != nil {
		return nil, errors.Err(err)
	}
	err = ioutil.WriteFile(path.Join(downloadPath, hash), blob, 0644)
	if err != nil {
		return nil, errors.Err(err)
	}
	elapsed = time.Since(start) - elapsed
	return &blob, nil
}

// GetHttpBlobStore returns default pre-configured blob store.
// edgeToken can be set to bypass restrictions for protected content
func GetHttpBlobStore() *store.HttpStore {
	return store.NewHttpStore(shared.ReflectorHttpServer, shared.EdgeToken)
}

//DownloadStream downloads a stream and returns the speed in bytes per second
func DownloadStream(blob *stream.SDBlob, fullTrace bool, downloadPath string) float64 {
	hashes := shared.GetStreamHashes(blob)
	totalSize := 0
	milliseconds := int64(0)
	for _, hash := range hashes {
		logrus.Debugln(hash)
		begin := time.Now()
		var b *stream.Blob
		var err error
		for {
			b, err = DownloadBlob(hash, fullTrace, downloadPath)
			milliseconds += time.Since(begin).Milliseconds()
			if err != nil {
				if strings.Contains(err.Error(), "No recent network activity") {
					logrus.Debugln("failed to download blob in time. retrying...")
				} else {
					logrus.Error(errors.FullTrace(err))
					return 0
				}
			} else {
				break
			}
		}
		totalSize += b.Size()
	}
	return float64(totalSize) / (float64(milliseconds) / 1000.0)
}
