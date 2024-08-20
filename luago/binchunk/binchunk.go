package binchunk

type binaryChunk struct {
	header                  //头部信息
	sizeUpValues byte       //主函数的upvalue数量
	mainFunc     *ProtoType //主函数院校
}

const (
	LUA_SIGNATURE    = "\x1bLua"
	LUAC_VERSION     = 0x53
	LUAC_FORMAT      = 0
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSZIET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

type header struct {
	signature       [4]byte //签名信息
	version         byte    //版本号
	formate         byte    //格式号
	luacData        [6]byte //luac data Ox1993回车符（ OxOD ）、 换行符 （ OxOA ）、 替换符（ OxlA ）和另一个换行符
	cintSize        byte    //cint
	sizetSize       byte    //size_t
	instructionSize byte    //lua 虚拟机指令
	luaIntegerSize  byte    //lua整数
	luaNumberSize   byte    //lua浮点数
	luacInt         int64   //n个字节存储 Ox5678，检测大小端模式
	luacNum         float64 //存放浮点数 370.5
}

type ProtoType struct {
	Source          string        //源文件名称
	LineDefined     uint32        //起行号
	LastLineDefined uint32        //止行号
	NumParams       byte          //固定参数个数
	IsVararg        byte          //是否式vararg函数 是否含有变长参数
	MaxStackSize    byte          //寄存器个数
	Code            []uint32      //指令表
	Constants       []interface{} //常量表
	Upvalues        []Upvalue     // upvalues表
	Protos          []*ProtoType  //子函数原型表
	LineInfo        []uint32      //行号表
	LocVars         []LocVar      //局部变量表
	UpvaluesNames   []string      //upvalue 名列表
}

const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x02
	TAG_INTEGER   = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x14
)

type Upvalue struct {
	Instack byte
	Idx     byte
}
type LocVar struct {
	VarName string
	StartPC uint32
	EndPC   uint32
}

func Updump(data []byte) *ProtoType {
	reader := &reader{data: data}
	reader.checkHeader()
	reader.readByte()
	return reader.readProto("")
}
