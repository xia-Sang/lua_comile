package luago

import (
	"fmt"
	"luago/api"
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
	ls := state.New()
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
func TestLuaShow2(t *testing.T) {
	// filename := "./hw.luac"
	filename := "./lua/luac.out"
	data, err := os.ReadFile(filename)
	assert.Nil(t, err)
	ls := state.New()
	ls.Load(data, filename, "b")
	ls.Call(0, 0)
}

func print(ls api.LuaState) int {
	nArgs := ls.GetTop()
	for i := 1; i <= nArgs; i++ {
		if ls.IsBoolean(i) {
			fmt.Printf("%t", ls.ToBoolean(i))
		} else if ls.IsString(i) {
			fmt.Print(ls.ToString(i))
		} else {
			fmt.Print(ls.TypeName(ls.Type(i)))
		}
		if i < nArgs {
			fmt.Print("\t")
		}
	}
	fmt.Println()
	return 0
}
func TestLuaShow3(t *testing.T) {
	filename := "./lua/h.luac"
	data, err := os.ReadFile(filename)
	assert.Nil(t, err)
	ls := state.New()
	ls.Register("print", print)
	ls.Load(data, filename, "b")
	ls.Call(0, 0)
}
