package tcp

import (
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/nikooo777/lbry-blobs-downloader/shared"

	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/lbry.go/v2/stream"
	"github.com/lbryio/reflector.go/server/peer"

	"github.com/sirupsen/logrus"
)

type TcpStore struct {
	server string
	port   string
	store  *peer.Store
}

func NewTcpBlobDownloader(server, port string) *TcpStore {
	newStore := &TcpStore{
		server: server,
		port:   port,
		store: peer.NewStore(peer.StoreOpts{
			Address: server + ":" + port,
			Timeout: 30 * time.Second,
		}),
	}
	return newStore
}

func (s *TcpStore) DownloadBlob(hash string, fullTrace bool, downloadPath string) (*stream.Blob, error) {
	bStore := s.store
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
	err = ioutil.WriteFile(path.Join(downloadPath, hash), blob, 0644)
	if err != nil {
		return nil, errors.Err(err)
	}
	elapsed = time.Since(start) - elapsed
	//logrus.Infof("save time: %d us\tSpeed: %.2f MB/s", elapsed.Microseconds(), (float64(len(blob))/(1024*1024))/elapsed.Seconds())
	return &blob, nil
}

// DownloadStream downloads a stream and returns the speed in bytes per second
func (s *TcpStore) DownloadStream(blob *stream.SDBlob, fullTrace bool, downloadPath string) float64 {
	hashes := shared.GetStreamHashes(blob)
	totalSize := 0
	milliseconds := int64(0)
	for _, hash := range hashes {
		logrus.Debugln(hash)
		begin := time.Now()
		b, err := s.DownloadBlob(hash, fullTrace, downloadPath)
		milliseconds += time.Since(begin).Milliseconds()
		if err != nil {
			logrus.Error(errors.FullTrace(err))
		}
		totalSize += b.Size()
	}
	return float64(totalSize) / (float64(milliseconds) / 1000.0)
}
