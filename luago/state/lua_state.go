package state

import (
	"fmt"
	"luago/api"
)

type luaState struct {
	stack *luaStack
}

func NewLuaState() *luaState {
	return &luaState{
		stack: newLuaStack(20),
	}
}

// 获取栈顶索引
func (ls *luaState) GetTop() int {
	return ls.stack.top
}

// 返回绝对索引
func (ls *luaState) AbsIndex(index int) int {
	return ls.stack.absIndex(index)
}

// 检查stack空间大小，是否可以容量n个数值，不满足进行扩容
func (ls *luaState) CheckStack(n int) bool {
	ls.stack.check(n)
	return true
}

// 从栈顶弹出n个数值
func (ls *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		ls.stack.pop()
	}
}

// 将一个位置的数值复制到另一个位置
func (ls *luaState) Copy(fromIndex, toIndex int) {
	val := ls.stack.get(fromIndex)
	ls.stack.set(toIndex, val)
}

// 将指定索引处的数值写入到栈顶
func (ls *luaState) PushValue(index int) {
	value := ls.stack.get(index)
	ls.stack.push(value)
}

// 将栈顶元素弹出 写入到指定位置
func (ls *luaState) Replace(index int) {
	value := ls.stack.pop()
	ls.stack.set(index, value)
}

// 实现旋转操作 实现[index,top]旋转n个位置
func (ls *luaState) Rotate(index, n int) {
	top := ls.stack.top - 1
	base := ls.stack.absIndex(index) - 1
	var m int

	if n >= 0 {
		m = top - n
	} else {
		m = base - n - 1
	}
	ls.stack.reverse(base, m)
	ls.stack.reverse(m+1, top)
	ls.stack.reverse(base, top)
}

// 删除索引处数值 将上面数值位置下移
func (ls *luaState) Remove(index int) {
	ls.Rotate(index, -1)
	ls.Pop(1)
}

// 将栈顶数值弹出 插入到指定位置
func (ls *luaState) Insert(index int) {
	ls.Rotate(index, 1)
}

// 将栈顶元素设置为指定数值
// 小于当前索引弹出 大于当前索引添加nil
func (ls *luaState) SetTop(index int) {
	newTop := ls.stack.absIndex(index)
	if newTop < 0 {
		panic("stack underflow!")
	}
	n := ls.stack.top - newTop
	if n > 0 {
		for i := 0; i < n; i++ {
			ls.stack.pop()
		}
	} else if n < 0 {
		for i := 0; i > n; i-- {
			ls.stack.push(nil)
		}
	}
}

// func (ls *luaState) Pop(n int) {
// ls.SetTop(-n-1)
// }
func (ls *luaState) PushNil() {
	ls.stack.push(nil)
}
func (ls *luaState) PushBoolean(data bool) {
	ls.stack.push(data)
}

func (ls *luaState) PushInteger(data int64) {
	ls.stack.push(data)
}
func (ls *luaState) PushNumber(data float64) {
	ls.stack.push(data)
}
func (ls *luaState) PushString(data string) {
	ls.stack.push(data)
}

func (ls *luaState) TypeName(ly api.LuaType) string {
	switch ly {
	case api.LUA_TNONE:
		return "no value"
	case api.LUA_TNIL:
		return "nil"
	case api.LUA_TBOOLEAN:
		return "boolean"
	case api.LUA_TNUMBER:
		return "number"
	case api.LUA_TSTRING:
		return "string"
	case api.LUA_TTABLE:
		return "table"
	case api.LUA_TFUNCTION:
		return "function"
	case api.LUA_TTHREAD:
		return "thread"
	default:
		return "userdata"
	}
}
func (ls *luaState) Type(index int) api.LuaType {
	if ls.stack.isValid(index) {
		val := ls.stack.get(index)
		return typeOf(val)
	}
	return api.LUA_TNONE
}

func (ls *luaState) IsNone(index int) bool {
	return ls.Type(index) == api.LUA_TNONE
}
func (ls *luaState) IsNil(index int) bool {
	return ls.Type(index) == api.LUA_TNIL
}
func (ls *luaState) IsNoneOrNil(index int) bool {
	return ls.Type(index) <= api.LUA_TNIL
}
func (ls *luaState) IsBoolean(index int) bool {
	return ls.Type(index) == api.LUA_TBOOLEAN
}

// 判断索引的数值时 字符
func (ls *luaState) IsString(index int) bool {
	ty := ls.Type(index)
	return ty == api.LUA_TSTRING || ty == api.LUA_TNUMBER
}
func (ls *luaState) IsNumber(index int) bool {
	ty := ls.Type(index)
	return ty == api.LUA_TSTRING || ty == api.LUA_TNUMBER
}
func (ls *luaState) IsInteger(index int) bool {
	val := ls.stack.get(index)
	_, ok := val.(int64)
	return ok
}
func (ls *luaState) ToBoolean(index int) bool {
	val := ls.stack.get(index)
	return convertToBoolean(val)
}
func (ls *luaState) ToNumber(index int) float64 {
	n, _ := ls.ToNumberX(index)
	return n
}
func (ls *luaState) ToNumberX(index int) (float64, bool) {
	val := ls.stack.get(index)
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	default:
		return 0, false
	}
}
func (ls *luaState) ToInteger(index int) int64 {
	i, _ := ls.ToIntegerX(index)
	return i
}
func (ls *luaState) ToIntegerX(index int) (int64, bool) {
	val := ls.stack.get(index)
	i, ok := val.(int64)
	return i, ok
}
func (ls *luaState) ToStringX(index int) (string, bool) {
	val := ls.stack.get(index)
	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		s := fmt.Sprintf("%v", x)
		ls.stack.set(index, s) //修改数据栈
		return s, true
	default:
		return "", false
	}
}
func (ls *luaState) ToString(index int) string {
	s, _ := ls.ToStringX(index)
	return s
}
func printStack(ls *luaState) {
	top := ls.GetTop()
	for i := 1; i <= top; i++ {
		t := ls.Type(i)
		switch t {
		case api.LUA_TBOOLEAN:
			fmt.Printf("[%t]", ls.ToBoolean(i))
		case api.LUA_TNUMBER:
			fmt.Printf("[%g]", ls.ToNumber(i))
		case api.LUA_TSTRING:
			fmt.Printf("[%q]", ls.ToString(i))
		default:
			fmt.Printf("[%s]", ls.TypeName(t))
		}
	}
	fmt.Println()
}
