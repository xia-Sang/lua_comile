package state

import (
	"luago/api"
)

type luaValue interface{}

func typeOf(value luaValue) api.LuaType {
	switch value.(type) {
	case nil:
		return api.LUA_TNIL
	case bool:
		return api.LUA_TBOOLEAN
	case int64, float64:
		return api.LUA_TNUMBER
	case string:
		return api.LUA_TSTRING
	default:
		panic("错误的数据类型！")
	}
}
func convertToBoolean(value luaValue) bool {
	switch x := value.(type) {
	case nil:
		return false
	case bool:
		return x
	default:
		return true
	}
}
