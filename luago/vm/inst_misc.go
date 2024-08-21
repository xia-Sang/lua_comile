package vm

import "luago/api"

func move(i Instruction, vm api.LuaVm) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.Copy(b, a)
}
func jmp(i Instruction, vm api.LuaVm) {
	a, sBx := i.AsBx()
	vm.AddPC(sBx)
	if a != 0 {
		panic("todo!")
	}
}
func loadNil(i Instruction, vm api.LuaVm) {
	a, b, _ := i.ABC()
	a += 1
	vm.PushNil()
	for i := a; i <= a+b; i++ {
		vm.Copy(-1, i)
	}
	vm.Pop(1)
}

func loadBool(i Instruction, vm api.LuaVm) {
	a, b, c := i.ABC()
	a += 1
	vm.PushBoolean(b != 0)
	vm.Replace(a)
	if c != 0 {
		vm.AddPC(1)
	}
}
func loadK(i Instruction, vm api.LuaVm) {
	a, bx := i.ABx()
	a += 1
	vm.GetConst(bx)
	vm.Replace(a)
}
func loadKx(i Instruction, vm api.LuaVm) {
	a, _ := i.ABx()
	a += 1
	ax := Instruction(vm.Fetch()).Ax()
	vm.GetConst(ax)
	vm.Replace(a)
}
