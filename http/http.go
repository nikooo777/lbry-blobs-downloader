package http

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/nikooo777/lbry-blobs-downloader/shared"

	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/lbry.go/v2/stream"
	"github.com/lbryio/reflector.go/store"

	"github.com/sirupsen/logrus"
)

// DownloadStream downloads a stream, returns the speed in bytes per second
// or halts and returns the first download error encountered.
func DownloadStream(blob *stream.SDBlob, fullTrace bool, downloadPath string, threads int) (float64, error) {
	hashes := shared.GetStreamHashes(blob)
	totalSize := int64(0)
	milliseconds := int64(0)
	maxThreads := threads

	var wg sync.WaitGroup
	ch := make(chan string, maxThreads)
	errChan := make(chan error, len(hashes))

	for _, hash := range hashes {
		wg.Add(1)

		ch <- hash
		go func(hash string) {
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			defer wg.Done()
			defer func() { <-ch }()

			logrus.Debugln(hash)
			begin := time.Now()
			blob, err := downloadBlobWithRetry(ctx, hash, fullTrace, &downloadPath)
			if err != nil {
				errChan <- fmt.Errorf("failed to download blob: %w", err)
				return
			}
			atomic.AddInt64(&milliseconds, time.Since(begin).Milliseconds())
			atomic.AddInt64(&totalSize, int64(blob.Size()))
		}(hash)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		return float64(totalSize) / (float64(milliseconds) / 1000.0), err
	}

	return float64(totalSize) / (float64(milliseconds) / 1000.0), nil
}

func DownloadBlob(hash string, fullTrace bool, downloadPath *string) (*stream.Blob, error) {
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
	if downloadPath != nil {
		err = os.MkdirAll(*downloadPath, os.ModePerm)
		if err != nil {
			return nil, errors.Err(err)
		}
		err = os.WriteFile(path.Join(*downloadPath, hash), blob, 0644)
		if err != nil {
			return nil, errors.Err(err)
		}
	}
	return &blob, nil
}

// GetHttpBlobStore returns default pre-configured blob store.
// EdgeToken can be set to bypass restrictions for protected content.
func GetHttpBlobStore() *store.UpstreamStore {
	return store.NewUpstreamStore(store.UpstreamParams{
		Upstream:  "http://" + shared.ReflectorHttpServer,
		EdgeToken: shared.EdgeToken,
	})
}

func downloadBlobWithRetry(ctx context.Context, hash string, fullTrace bool, path *string) (*stream.Blob, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			blob, err := DownloadBlob(hash, fullTrace, path)
			if err != nil {
				if isRetriable(err) {
					logrus.Debugln("failed to download blob in time. retrying...")
				} else {
					logrus.Error(errors.FullTrace(err))
					return nil, err
				}
			} else {
				return blob, nil
			}
		}
	}
}

func isRetriable(err error) bool {
	return strings.Contains(err.Error(), "No recent network activity")
}
