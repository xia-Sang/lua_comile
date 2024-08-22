package luago

import (
	"fmt"
	"luago/binchunk"
	"luago/state"
	"luago/vm"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLuaShow(t *testing.T) {
	// filename := "./hw.luac"
	filename := "./lua/for.luac"
	data, err := os.ReadFile(filename)
	assert.Nil(t, err)
	proto := binchunk.Updump(data)
	luaMain(proto)
}
func luaMain(pro *binchunk.ProtoType) {
	nRegs := int(pro.MaxStackSize)
	ls := state.NewLuaState(nRegs+8, pro)
	ls.SetTop(nRegs)
	for {
		pc := ls.PC()
		inst := vm.Instruction(ls.Fetch())
		if inst.Opcode() != vm.OP_RETURN {
			inst.Execute(ls)
			fmt.Printf("[%02d] %s ", pc+1, inst.OpName())
			state.PrintStack(ls)
		} else {
			break
		}
	}
}
func TestLuaShow1(t *testing.T) {
	// filename := "./hw.luac"
	filename := "./lua/table.luac"
	data, err := os.ReadFile(filename)
	assert.Nil(t, err)
	proto := binchunk.Updump(data)
	luaMain(proto)
}
