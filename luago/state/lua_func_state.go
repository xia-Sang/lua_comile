package state

// 函数调用栈
func (ls *luaState) pushLuaStack(stack *luaStack) {
	stack.prev = ls.stack
	ls.stack = stack
}
func (ls *luaState) popLuaStack() {
	stack := ls.stack
	ls.stack = stack.prev
	stack.prev = nil
}

// IsFunction implements api.LuaVm.
func (ls *luaState) IsFunction(idx int) bool {
	panic("unimplemented")
}

// IsTable implements api.LuaVm.
func (ls *luaState) IsTable(idx int) bool {
	panic("unimplemented")
}

// IsThread implements api.LuaVm.
func (ls *luaState) IsThread(idx int) bool {
	panic("unimplemented")
}

// PC 实现luaVm接口
func (ls *luaState) PC() int {
	return ls.stack.pc
}

// AddPC 增加地址信息
func (ls *luaState) AddPC(n int) {
	ls.stack.pc += n
}

// Fetch PC索引从函数原型的指令表里取出当前指令
func (ls *luaState) Fetch() uint32 {
	i := ls.stack.closure.proto.Code[ls.stack.pc]
	ls.stack.pc++
	return i
}

// GetConst 索引从函数原型的常量表里取出一个常量值
func (ls *luaState) GetConst(index int) {
	c := ls.stack.closure.proto.Constants[index]
	ls.stack.push(c)
}
