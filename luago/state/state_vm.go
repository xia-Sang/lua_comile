package state

func (ls *luaState) LoadProto(index int) {
	proto := ls.stack.closure.proto.Protos[index]
	closure := newLuaCloser(proto)
	ls.stack.push(closure)
}
func (ls *luaState) LoadVararg(n int) {
	if n < 0 {
		n = len(ls.stack.varArgs)
	}
	ls.stack.check(n)
	ls.stack.pushN(ls.stack.varArgs, n)
}
func (ls *luaState) RegisterCount() int {
	return int(ls.stack.closure.proto.MaxStackSize)
}

// 获取栈顶索引
func (ls *luaState) GetTop() int {
	return ls.stack.top
}

// 调用 PushValue 方法把某个索引处的栈值推人栈顶
func (ls *luaState) GetRK(rk int) {
	if rk > 0xFF { // constant
		ls.GetConst(rk & 0xFF)
	} else { //register
		ls.PushValue(rk + 1)
	}
}
