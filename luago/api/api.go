package api

type LuaType = int
type ArithOp = int   //类型别名
type CompareOp = int //类型别名

type LuaVm interface {
	LuaState            //基本状态类型
	PC() int            //返回当前pc
	AddPC(n int)        //修改pc
	Fetch() uint32      //取出当前指令
	GetConst(index int) //将指定常量存放栈顶
	GetRK(rk int)       //将指定常量或者栈数值推入栈顶
}

// 基本状态类型
type LuaState interface {
	/* basic stack manipulation */
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
	/* access functions (stack -> Go) */
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
	/* push functions (Go -> stack) */
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
}
