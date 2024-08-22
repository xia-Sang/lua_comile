package api

type LuaType = int
type ArithOp = int   //类型别名
type CompareOp = int //类型别名

type GoFunction func(LuaState) int

const (
	LUA_MINSTACK            = 20
	LUA_MAXSTACK            = 1000000
	LUA_REGISTRYINDEX       = -LUA_MAXSTACK - 1000
	LUA_RIDX_GLOBALS  int64 = 2
)

type LuaVm interface {
	LuaState             //基本状态类型
	PC() int             //返回当前pc
	AddPC(n int)         //修改pc
	Fetch() uint32       //取出当前指令
	GetConst(index int)  //将指定常量存放栈顶
	GetRK(rk int)        //将指定常量或者栈数值推入栈顶
	LoadProto(index int) //导入proto信息
	RegisterCount() int  //注册计数
	LoadVararg(n int)    //导入参数信息
}
