package shared

import (
	"encoding/hex"

	"github.com/lbryio/lbry.go/v2/stream"
)

var ReflectorPeerServer = "cdn.reflector.lbry.com:5567"
var ReflectorQuicServer = "cdn.reflector.lbry.com:5568"
var ReflectorHttpServer = "cdn.reflector.lbry.com:5569"

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
