package state

type luaStack struct {
	slots   []luaValue  //存储数值信息
	top     int         //存储栈索引
	prev    *luaStack   //之前的调用栈
	closure *luaClosure //闭包
	varArgs []luaValue  //参数
	pc      int         //程序计数器
}

func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

// 检查stack空间大小，是否可以容量n个数值，不满足进行扩容
func (ls *luaStack) check(n int) {
	free := len(ls.slots) - ls.top
	for i := free; i < n; i++ {
		ls.slots = append(ls.slots, n)
	}
}

// push操作
func (ls *luaStack) push(value luaValue) {
	if ls.top == len(ls.slots) {
		panic("stack overflow")
	}
	ls.slots[ls.top] = value
	ls.top++
}

// pop数据
func (ls *luaStack) pop() luaValue {
	if ls.top < 1 {
		panic("stack underflow")
	}
	ls.top--
	value := ls.slots[ls.top]
	ls.slots[ls.top] = nil
	return value
}

// pop数据
func (ls *luaStack) popN(n int) []luaValue {
	vals := make([]luaValue, n)
	for i := n - 1; i >= 0; i-- {
		vals[i] = ls.pop()
	}
	return vals
}

// pop数据
func (ls *luaStack) pushN(vals []luaValue, n int) {
	nVals := len(vals)
	if n < 0 {
		n = nVals
	}
	for i := 0; i < n; i++ {
		if i < nVals {
			ls.push(vals[i])
		} else {
			ls.push(nil)
		}
	}
}

// 索引转换
func (ls *luaStack) absIndex(index int) int {
	if index >= 0 {
		return index
	}
	return index + ls.top + 1
}

// 判断索引是否有效
func (ls *luaStack) isValid(index int) bool {
	absIndex := ls.absIndex(index)
	return absIndex > 0 && index <= ls.top
}

// get数据
func (ls *luaStack) get(index int) luaValue {
	absIndex := ls.absIndex(index)
	if absIndex > 0 && absIndex <= ls.top {
		return ls.slots[absIndex-1]
	}
	return nil
}

// set数据
func (ls *luaStack) set(index int, value luaValue) {
	absIndex := ls.absIndex(index)
	if absIndex > 0 && absIndex <= ls.top {
		ls.slots[absIndex-1] = value
		return
	}
	panic("invalid index!")
}

// 反转操作
func (ls *luaStack) reverse(from, to int) {
	slots := ls.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
