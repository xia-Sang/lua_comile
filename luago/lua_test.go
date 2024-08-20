package luago

import (
	"luago/binchunk"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLuaShow(t *testing.T) {
	filename := "./hw.luac"
	data, err := os.ReadFile(filename)
	assert.Nil(t, err)
	proto := binchunk.Updump(data)
	list(proto)
}
