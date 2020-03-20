package shared

import (
	"encoding/hex"

	"github.com/lbryio/lbry.go/v2/stream"
)

var ReflectorPeerServer = "refractor.lbry.com:5567"
var ReflectorQuicServer = "refractor.lbry.com:5568"

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
