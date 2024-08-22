package api

// 基本状态类型
type LuaState interface {
	//基本的栈实现
	GetTop() int
	AbsIndex(idx int) int
	CheckStack(n int) bool
	Pop(n int)
	Copy(fromIdx, toIdx int)
	PushValue(idx int)
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx, n int)
	SetTop(idx int)
	//访问函数实现
	TypeName(tp LuaType) string
	Type(idx int) LuaType
	IsNone(idx int) bool
	IsNil(idx int) bool
	IsNoneOrNil(idx int) bool
	IsBoolean(idx int) bool
	IsInteger(idx int) bool
	IsNumber(idx int) bool
	IsString(idx int) bool
	IsTable(idx int) bool
	IsThread(idx int) bool
	IsFunction(idx int) bool
	ToBoolean(idx int) bool
	ToInteger(idx int) int64
	ToIntegerX(idx int) (int64, bool)
	ToNumber(idx int) float64
	ToNumberX(idx int) (float64, bool)
	ToString(idx int) string
	ToStringX(idx int) (string, bool)
	// push函数实现
	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)
	// 新添加四个方法
	Arith(op ArithOp)                          //执行算术和按位运算，
	Compare(idx1, idx2 int, op CompareOp) bool //执行比较运算
	Len(index int)                             //长度原酸
	Concat(n int)                              //拼接运算
	// 添加支持table的
	NewTable()
	CreateTable(nArr, nRec int)
	GetTable(index int) LuaType
	GetField(index int, k string) LuaType
	GetI(index int, i int64) LuaType
	// 设置函数
	SetTable(index int)
	SetField(index int, k string)
	SetI(index int, n int64)
	// 函数调用栈
	Load(chunk []byte, chunkName, mode string) int
	Call(nArgs, nResults int)
}
