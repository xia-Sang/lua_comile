package state

import (
	"luago/api"
	"testing"
)

func TestArith(t *testing.T) {
	ls := NewLuaState(20, nil)
	ls.PushInteger(1)
	printStack(ls)
	ls.PushString("2.0")
	ls.PushString("3.0")
	ls.PushNumber(4.0)

	printStack(ls)

	ls.Arith(api.LUA_OPADD)
	printStack(ls)
	ls.Arith(api.LUA_OPBNOT)
	printStack(ls)
	ls.Len(2)
	printStack(ls)
	ls.Concat(3)
	printStack(ls)
	ls.PushBoolean(ls.Compare(1, 2, api.LUA_OPEQ))
	printStack(ls)
}
