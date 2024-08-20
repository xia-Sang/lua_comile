package vm

/*
IABC		 // [  B:9  ][  C:9  ][ A:8  ][OP:6]
IABx         // [      Bx:18     ][ A:8  ][OP:6]
IAsBx        // [     sBx:18     ][ A:8  ][OP:6]
IAx          // [           Ax:26        ][OP:6]
*/
type Instruction uint32

const MAXARG_Bx = 1<<18 - 1       // 262143
const MAXARG_sBx = MAXARG_Bx >> 1 // 131071

func (i Instruction) Opcode() int {
	return int(i & 0x3f)
}
func (i Instruction) ABC() (a, b, c int) {
	a = int(i >> 6 & 0xFF)
	c = int(i >> 14 & 0x1FF)
	b = int(i >> 23 & 0x1FF)
	return
}

func (i Instruction) ABx() (a, bx int) {
	a = int(i >> 6 & 0xFF)
	bx = int(i >> 14)
	return
}

func (i Instruction) AsBx() (a, sbx int) {
	// 只有这个会被解析为有符号整数的 所以要进行减法操作
	a, bx := i.ABx()
	return a, bx - MAXARG_sBx
}
func (i Instruction) Ax() int {
	return int(i >> 6)
}
func (i Instruction) OpName() string {
	return opcodes[i.Opcode()].name
}
func (i Instruction) OpMode() byte {
	return opcodes[i.Opcode()].opMode
}
func (i Instruction) BMode() byte {
	return opcodes[i.Opcode()].argBMode
}
func (i Instruction) CMode() byte {
	return opcodes[i.Opcode()].argCMode
}
