package vm

import "luago/api"

func closure(i Instruction, vm api.LuaVm) {
	a, bx := i.ABx()
	a += 1

	vm.LoadProto(bx)
	vm.Replace(a)
}
func call(i Instruction, vm api.LuaVm) {
	a, b, c := i.ABC()
	a += 1

	nArgs := pushFuncAndArgs(a, b, vm)
	vm.Call(nArgs, c-1)
	popResults(a, c, vm)
}
func luaReturn(i Instruction, vm api.LuaVm) {
	a, b, _ := i.ABC()
	a += 1
	if b == 1 {

	} else if b > 1 {
		vm.CheckStack(b - 1)
		for i := a; i <= a+b-2; i++ {
			vm.PushValue(i)
		}
	} else {
		fixStack(a, vm)
	}
}
func varArg(i Instruction, vm api.LuaVm) {
	a, b, _ := i.ABC()
	a += 1
	if b != 1 {
		vm.LoadVararg(b - 1)
		popResults(a, b, vm)
	}
}
func tailCall(i Instruction, vm api.LuaVm) {
	a, b, _ := i.ABC()
	a += 1
	c := 0
	nArgs := pushFuncAndArgs(a, b, vm)
	vm.Call(nArgs, c-1)
	popResults(a, c, vm)
}
func self(i Instruction, vm api.LuaVm) {
	a, b, c := i.ABC()
	a += 1
	b += 1

	vm.Copy(b, a+1)
	vm.GetRK(c)
	vm.GetTable(b)
	vm.Replace(a)
}
func pushFuncAndArgs(a, b int, vm api.LuaVm) (nArgs int) {
	if b >= 1 {
		vm.CheckStack(b)
		for i := a; i < a+b; i++ {
			vm.PushValue(i)
		}
		return b - 1
	} else {
		fixStack(a, vm)
		return vm.GetTop() - vm.RegisterCount() - 1
	}
}
func popResults(a, c int, vm api.LuaVm) {
	if c == 1 {

	} else if c > 1 {
		for i := a + c - 2; i >= a; i-- {
			vm.Replace(i)
		}
	} else {
		vm.CheckStack(1)
		vm.PushInteger(int64(a))
	}
}
func fixStack(a int, vm api.LuaVm) {
	x := int(vm.ToInteger(-1))
	vm.Pop(1)

	vm.CheckStack(x - a)
	for i := a; i < x; i++ {
		vm.PushValue(i)
	}
	vm.Rotate(vm.RegisterCount()+1, x-a)
}
