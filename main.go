package main

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/lbry.go/v2/stream"
	"github.com/lbryio/reflector.go/peer"
	"github.com/sirupsen/logrus"
)

var reflectorServer = "refractor.lbry.com:5567"

func main() {
	concurrentThreads := int64(20)
	isStream := false
	var err error
	if len(os.Args) < 2 {
		logrus.Errorln("you must specify a blob hash to download")
		os.Exit(1)
	}
	if len(os.Args) >= 3 {
		reflectorServer = os.Args[2]
	}
	if len(os.Args) >= 4 {
		isStream, err = strconv.ParseBool(os.Args[3])
		if err != nil {
			logrus.Errorln(err.Error())
			os.Exit(1)
		}
	}
	if len(os.Args) >= 5 {
		var err error
		concurrentThreads, err = strconv.ParseInt(os.Args[4], 10, 32)
		if err != nil {
			logrus.Errorln(err.Error())
			os.Exit(1)
		}
	}
	logrus.Println(concurrentThreads)
	wg := &sync.WaitGroup{}
	for i := int64(0); i < concurrentThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isStream {
				blob, err := DownloadBlob(os.Args[1])
				if err != nil {
					logrus.Error(errors.FullTrace(err))
					os.Exit(1)
				}
				sdb := &stream.SDBlob{}
				err = sdb.FromBlob(*blob)

				if err != nil {
					logrus.Error(errors.FullTrace(err))
					os.Exit(1)
				}
				DownloadStream(sdb)
			} else {
				_, err := DownloadBlob(os.Args[1])
				if err != nil {
					logrus.Error(errors.FullTrace(err))
					os.Exit(1)
				}
			}
		}()
	}
	wg.Wait()
}

func DownloadBlob(hash string) (*stream.Blob, error) {
	bStore := GetBlobStore()
	start := time.Now()
	blob, err := bStore.Get(hash)
	if err != nil {
		err = errors.Prefix(hash, err)
		return nil, errors.Err(err)
	}
	elapsed := time.Since(start)
	logrus.Infof("download time: %d ms\tSpeed: %.2f MB/s", elapsed.Milliseconds(), (float64(len(blob))/(1024*1024))/elapsed.Seconds())
	err = ioutil.WriteFile(hash, blob, 0644)
	if err != nil {
		return nil, errors.Err(err)
	}
	elapsed = time.Since(start) - elapsed
	//logrus.Infof("save time: %d us\tSpeed: %.2f MB/s", elapsed.Microseconds(), (float64(len(blob))/(1024*1024))/elapsed.Seconds())
	return &blob, nil
}

// GetBlobStore returns default pre-configured blob store.
func GetBlobStore() *peer.Store {
	return peer.NewStore(peer.StoreOpts{
		Address: reflectorServer,
		Timeout: 30 * time.Second,
	})
}

func DownloadStream(blob *stream.SDBlob) {
	hashes := GetStreamHashes(blob)
	for _, hash := range hashes {
		logrus.Info(hash)
		_, err := DownloadBlob(hash)
		if err != nil {
			logrus.Error(errors.FullTrace(err))
		}
	}
}

func GetStreamHashes(blob *stream.SDBlob) []string {
	blobs := make([]string, 0, len(blob.BlobInfos))
	for _, b := range blob.BlobInfos {
		hash := hex.EncodeToString(b.BlobHash)
		if hash == "" {
			continue
		}
		blobs = append(blobs, hex.EncodeToString(b.BlobHash))
	}
	return blobs
}
