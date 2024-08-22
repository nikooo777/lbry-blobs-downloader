package chainquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOriginalName(t *testing.T) {
	name, err := GetOriginalName("7798b17ff60be8904df1e79bbe4f6b68546a3b4aab1478c54a8b243a7f3013525cb41e1b6604d6857b9151f3325bb50c")
	assert.Nil(t, err)
	assert.Equal(t, "i-tried-finding-hidden-gems-on.mp4", name)
}

func TestGetSdHash(t *testing.T) {
	sdHash, err := GetSdHash("05dbe782f1d8588251b80365610eda80920d8278")
	assert.Nil(t, err)
	assert.Equal(t, "323be2060c9f6c7877afc5feec4eb4c0a35eec00ec6439a47a910e6379f58090d77d0a26e6023fea2ea792244ded4e49", sdHash)
}

func TestGetChannelStreams(t *testing.T) {
	streams, err := GetChannelStreams("80d2590ad04e36fb1d077a9b9e3a8bba76defdf8")
	assert.Nil(t, err)
	assert.Greater(t, len(streams), 10)

}
func TestGetThumbnail(t *testing.T) {
	thumbnail, err := GetClaimThumbnail("4b8f25e74ea0faff844feb6b9c60a204c59fef90")
	assert.Nil(t, err)
	assert.Equal(t, thumbnail, "https://thumbnails.lbry.com/vQR7Ve846qY")
}
