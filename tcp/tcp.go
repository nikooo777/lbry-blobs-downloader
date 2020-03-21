package tcp

import (
	"blobdownloader/shared"
	"io/ioutil"
	"time"

	"github.com/lbryio/errors.go"
	"github.com/lbryio/lbry.go/v2/stream"
	"github.com/lbryio/reflector.go/peer"
	"github.com/lbryio/reflector.go/store"
	"github.com/sirupsen/logrus"
)

func DownloadBlob(hash string) (*stream.Blob, error) {
	bStore := GetQuicBlobStore()
	start := time.Now()
	blob, err := bStore.Get(hash)
	if err != nil {
		err = errors.Prefix(hash, err)
		return nil, errors.Err(err)
	}
	elapsed := time.Since(start)
	logrus.Infof("[T] download time: %d ms\tSpeed: %.2f MB/s", elapsed.Milliseconds(), (float64(len(blob))/(1024*1024))/elapsed.Seconds())
	err = ioutil.WriteFile(hash, blob, 0644)
	if err != nil {
		return nil, errors.Err(err)
	}
	elapsed = time.Since(start) - elapsed
	//logrus.Infof("save time: %d us\tSpeed: %.2f MB/s", elapsed.Microseconds(), (float64(len(blob))/(1024*1024))/elapsed.Seconds())
	return &blob, nil
}

// GetQuicBlobStore returns default pre-configured blob store.
func GetQuicBlobStore() store.BlobStore {
	return peer.NewStore(peer.StoreOpts{
		Address: shared.ReflectorPeerServer,
		Timeout: 30 * time.Second,
	})
}

// downloads a stream and returns the speed in bytes per second
func DownloadStream(blob *stream.SDBlob) float64 {
	hashes := shared.GetStreamHashes(blob)
	totalSize := 0
	milliseconds := int64(0)
	for _, hash := range hashes {
		logrus.Info(hash)
		begin := time.Now()
		b, err := DownloadBlob(hash)
		milliseconds += time.Since(begin).Milliseconds()
		if err != nil {
			logrus.Error(errors.FullTrace(err))
		}
		totalSize += b.Size()
	}
	return float64(totalSize) / (float64(milliseconds) / 1000.0)
}
