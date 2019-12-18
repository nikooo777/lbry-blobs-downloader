package main

import (
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/reflector.go/peer"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 2 {
		logrus.Errorln("you must specify a blob hash to download")
		os.Exit(1)
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := DownloadBlob(os.Args[1])
			if err != nil {
				logrus.Fatalf(errors.FullTrace(err))
				os.Exit(1)
			}
		}()
	}
	wg.Wait()
}

func DownloadBlob(hash string) error {
	bStore := GetBlobStore()
	start := time.Now()
	blob, err := bStore.Get(hash)
	if err != nil {
		return errors.Err(err)
	}
	elapsed := time.Since(start)
	logrus.Infof("download time: %d ms\tSpeed: %.2f MB/s", elapsed.Milliseconds(), (float64(len(blob))/(1024*1024))/elapsed.Seconds())
	err = ioutil.WriteFile(hash, blob, 0644)
	if err != nil {
		return errors.Err(err)
	}
	elapsed = time.Since(start) - elapsed
	//logrus.Infof("save time: %d us\tSpeed: %.2f MB/s", elapsed.Microseconds(), (float64(len(blob))/(1024*1024))/elapsed.Seconds())
	return nil
}

// GetBlobStore returns default pre-configured blob store.
func GetBlobStore() *peer.Store {
	return peer.NewStore(peer.StoreOpts{
		Address: "refractor.lbry.com:5567",
		Timeout: 30 * time.Second,
	})
}
