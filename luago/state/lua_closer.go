package state

import (
	"luago/api"
	"luago/binchunk"
)

type closure struct {
	proto  *binchunk.ProtoType //lua 闭包
	goFunc api.GoFunction      //go 闭包
}

func newLuaCloser(proto *binchunk.ProtoType) *closure {
	return &closure{proto: proto}
}
func newGoCloser(fn api.GoFunction) *closure {
	return &closure{goFunc: fn}
}
