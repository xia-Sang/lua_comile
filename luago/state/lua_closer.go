package state

import (
	"luago/binchunk"
)

type closure struct {
	proto *binchunk.ProtoType //lua 闭包
	// goFunc GoFunction          //go 闭包
}

func newLuaCloser(proto *binchunk.ProtoType) *closure {
	return &closure{proto: proto}
}
