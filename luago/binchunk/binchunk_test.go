package binchunk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinChunk(t *testing.T) {
	filename := "./hw.luac"
	data, err := os.ReadFile(filename)
	assert.Nil(t, err)
	proto := Updump(data)
	list(proto)
}
