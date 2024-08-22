package state

import "luago/binchunk"

type luaClosure struct {
	proto *binchunk.ProtoType
}

func newLuaCloser(proto *binchunk.ProtoType) *luaClosure {
	return &luaClosure{proto: proto}
}
