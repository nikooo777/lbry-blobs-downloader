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
