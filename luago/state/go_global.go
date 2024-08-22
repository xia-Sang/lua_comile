package state

import "luago/api"

func (ls *luaState) PushGlobalTable() {
	global := ls.registry.get(api.LUA_RIDX_GLOBALS)
	ls.stack.push(global)
}
func (ls *luaState) GetGlobal(name string) api.LuaType {
	t := ls.registry.get(api.LUA_RIDX_GLOBALS)
	return ls.getTable(t, name)
}
func (ls *luaState) SetGlobal(name string) {
	t := ls.registry.get(api.LUA_RIDX_GLOBALS)
	v := ls.stack.pop()
	ls.setTable(t, name, v)
}
func (ls *luaState) Register(name string, fn api.GoFunction) {
	ls.PushGoFunction(fn)
	ls.SetGlobal(name)
}
