package shared

import (
	"encoding/hex"

	"github.com/lbryio/lbry.go/v2/stream"
)

var DefaultReflectorPeerServer = "cdn.reflector.lbry.com"
var DefaultReflectorPeerPort = "5567"
var DefaultReflectorQuicServer = "cdn.reflector.lbry.com"
var DefaultReflectorQuicPort = "5568"
var DefaultReflectorHttpServer = "cdn.reflector.lbry.com"
var DefaultReflectorHttpPort = "5569"

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
