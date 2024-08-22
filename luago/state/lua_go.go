package state

import "luago/api"

func (ls *luaState) PushGoFunction(fn api.GoFunction) {
	ls.stack.push(newGoCloser(fn))
}
func (ls *luaState) IsGoFunction(index int) bool {
	val := ls.stack.get(index)
	if c, ok := val.(*closure); ok {
		return c.goFunc != nil
	}
	return false
}
func (ls *luaState) ToGoFunction(index int) api.GoFunction {
	val := ls.stack.get(index)
	if c, ok := val.(*closure); ok {
		return c.goFunc
	}
	return nil
}
