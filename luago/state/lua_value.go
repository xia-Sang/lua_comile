package state

import (
	"luago/api"
	"luago/number"
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
	case *luaTable:
		return api.LUA_TTABLE
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
func convertToFloat(val luaValue) (float64, bool) {
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	case string:
		return number.ParseFloat(x)
	default:
		return 0, false
	}
}
func convertToInteger(val luaValue) (int64, bool) {
	switch x := val.(type) {
	case float64:
		return number.FloatToInterger(x)
	case int64:
		return x, true
	case string:
		return stringToInterger(x)
	default:
		return 0, false
	}
}
func stringToInterger(s string) (int64, bool) {
	if i, ok := number.ParseInteger(s); ok {
		return i, true
	}
	if f, ok := number.ParseFloat(s); ok {
		return number.FloatToInterger(f)
	}
	return 0, false
}
