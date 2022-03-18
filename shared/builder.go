package shared

import (
	"bufio"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path"

	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/lbry.go/v2/stream"
)

func BuildStream(sdBlob *stream.SDBlob, fileName string, destinationPath string, blobsDirectory string) error {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return errors.Err(err)
	}
	tmpName := path.Join(tmpDir, fileName+".tmp")
	finalName := path.Join(destinationPath, fileName)
	f, err := os.Create(tmpName)
	if err != nil {
		return errors.Err(err)
	}
	w := bufio.NewWriter(f)
	for _, info := range sdBlob.BlobInfos {
		if info.Length == 0 {
			continue
		}
		hash := hex.EncodeToString(info.BlobHash)
		blobPath := path.Join(blobsDirectory, hash)
		blobToDecrypt, err := ioutil.ReadFile(blobPath)
		_ = os.Remove(blobPath) //we don't need the blob anymore
		if err != nil {
			return errors.Err(err)
		}
		decryptedBlob, err := stream.DecryptBlob(blobToDecrypt, sdBlob.Key, info.IV)
		if err != nil {
			return errors.Err(err)
		}
		_, err = w.Write(decryptedBlob)
		if err != nil {
			return errors.Err(err)
		}
		err = w.Flush()
		if err != nil {
			return errors.Err(err)
		}
	}
	err = os.Rename(tmpName, finalName)
	if err != nil {
		return errors.Err(err)
	}
	return nil
}
