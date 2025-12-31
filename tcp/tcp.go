package tcp

import (
	"os"
	"path"
	"time"

	"github.com/nikooo777/lbry-blobs-downloader/shared"

	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/lbry.go/v2/stream"
	"github.com/lbryio/reflector.go/store"

	"github.com/sirupsen/logrus"
)

func DownloadBlob(hash, downloadPath string) (*stream.Blob, error) {
	bStore := GetTcpBlobStore()
	start := time.Now()
	blob, _, err := bStore.Get(hash)
	if err != nil {
		err = errors.Prefix(hash, err)
		return nil, errors.Err(err)
	}
	elapsed := time.Since(start)
	logrus.Debugf("[T] download time: %d ms\tSpeed: %.2f MB/s", elapsed.Milliseconds(), (float64(len(blob))/(1024*1024))/elapsed.Seconds())
	err = os.MkdirAll(downloadPath, os.ModePerm)
	if err != nil {
		return nil, errors.Err(err)
	}
	err = os.WriteFile(path.Join(downloadPath, hash), blob, 0644)
	if err != nil {
		return nil, errors.Err(err)
	}
	return &blob, nil
}

// GetTcpBlobStore returns default pre-configured blob store.
func GetTcpBlobStore() *store.PeerStore {
	return store.NewPeerStore(store.PeerParams{
		Address: shared.ReflectorPeerServer,
		Timeout: 30 * time.Second,
	})
}

// DownloadStream downloads a stream and returns the speed in bytes per second
func DownloadStream(blob *stream.SDBlob, downloadPath string) float64 {
	hashes := shared.GetStreamHashes(blob)
	totalSize := 0
	milliseconds := int64(0)
	for _, hash := range hashes {
		logrus.Debugln(hash)
		begin := time.Now()
		b, err := DownloadBlob(hash, downloadPath)
		milliseconds += time.Since(begin).Milliseconds()
		if err != nil {
			logrus.Error(errors.FullTrace(err))
		}
		totalSize += b.Size()
	}
	return float64(totalSize) / (float64(milliseconds) / 1000.0)
}
