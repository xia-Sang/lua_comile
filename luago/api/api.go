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
