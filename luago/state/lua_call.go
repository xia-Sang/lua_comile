package state

import (
	"fmt"
	"luago/binchunk"
	"luago/vm"
)

func (ls *luaState) Load(chunk []byte, chunkName, mode string) int {
	proto := binchunk.Updump(chunk)
	c := newLuaCloser(proto)
	ls.stack.push(c)
	data, ok := ls.stack.slots[0].(*luaClosure)
	fmt.Println(data.proto, ok)
	return 0
}

func (ls *luaState) Call(nArgs, nResults int) {
	val := ls.stack.get(-(nArgs + 1))
	if c, ok := val.(*luaClosure); ok {
		fmt.Printf("call %s <%d,%d>\n", c.proto.Source, c.proto.LineDefined, c.proto.LastLineDefined)
		ls.callLuaClosure(nArgs, nResults, c)
	} else {
		panic("not function!")
	}
}
func (self *luaState) callLuaClosure(nArgs, nResults int, c *luaClosure) {
	nRegs := int(c.proto.MaxStackSize)
	nParams := int(c.proto.NumParams)
	isVararg := c.proto.IsVararg == 1

	newStack := newLuaStack(nRegs + 20)
	newStack.closure = c

	funcAndArgs := self.stack.popN(nArgs + 1)
	newStack.pushN(funcAndArgs[1:], nParams)
	newStack.top = nRegs
	if nArgs > nParams && isVararg {
		newStack.varArgs = funcAndArgs[nParams+1:]
	}

	self.pushLuaStack(newStack)
	self.runLuaClosure()
	self.popLuaStack()

	// return results
	if nResults != 0 {
		results := newStack.popN(newStack.top - nRegs)
		self.stack.check(len(results))
		self.stack.pushN(results, nResults)
	}
}
func (ls *luaState) runLuaClosure() {
	for {
		inst := vm.Instruction(ls.Fetch())
		inst.Execute(ls)
		if inst.Opcode() == vm.OP_RETURN {
			break
		}
	}
}
