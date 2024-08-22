package state

import "luago/api"

// table部分
func (ls *luaState) CreateTable(nArr, nRec int) {
	t := newLuaTable(nArr, nRec)
	ls.stack.push(t)
}
func (ls *luaState) NewTable() {
	ls.CreateTable(0, 0)
}
func (ls *luaState) GetTable(index int) api.LuaType {
	t := ls.stack.get(index)
	k := ls.stack.pop()
	return ls.getTable(t, k)
}
func (ls *luaState) getTable(t, k luaValue) api.LuaType {
	if tb, ok := t.(*luaTable); ok {
		v := tb.get(k)
		ls.stack.push(v)
		return typeOf(v)
	}
	panic("not a table!")
}
func (ls *luaState) GetField(index int, k string) api.LuaType {
	t := ls.stack.get(index)
	return ls.getTable(t, k)
}

//	func (ls *luaState) GetField(index int, k string) api.LuaType {
//		ls.PushString(k)
//		return ls.GetTable(index)
//	}
func (ls *luaState) GetI(index int, i int64) api.LuaType {
	t := ls.stack.get(index)
	return ls.getTable(t, i)
}
func (ls *luaState) SetTable(index int) {
	t := ls.stack.get(index)
	v := ls.stack.pop()
	k := ls.stack.pop()
	ls.setTable(t, k, v)
}
func (ls *luaState) setTable(t, k, v luaValue) {
	if table, ok := t.(*luaTable); ok {
		table.put(k, v)
		return
	}
	panic("not a table!")
}

func (ls *luaState) SetField(index int, k string) {
	t := ls.stack.get(index)
	v := ls.stack.pop()
	ls.setTable(t, k, v)
}

func (ls *luaState) SetI(index int, i int64) {
	t := ls.stack.get(index)
	v := ls.stack.pop()
	ls.setTable(t, i, v)
}
