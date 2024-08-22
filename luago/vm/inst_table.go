package vm

import "luago/api"

func newTable(i Instruction, vm api.LuaVm) {
	a, b, c := i.ABC()
	a += 1

	vm.CreateTable(Fb2int(b), Fb2int(c))
	vm.Replace(a)
}
func getTable(i Instruction, vm api.LuaVm) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	vm.GetRK(c)
	vm.GetTable(b)
	vm.Replace(a)
}
func setTable(i Instruction, vm api.LuaVm) {
	a, b, c := i.ABC()
	a += 1
	vm.GetRK(b)
	vm.GetRK(c)
	vm.SetTable(a)
}

const LFIELDS_PER_FLUSH = 50

func setList(i Instruction, vm api.LuaVm) {
	a, b, c := i.ABC()
	a += 1
	bIsZero := b == 0
	if bIsZero {
		b = int(vm.ToInteger(-1)) - a - 1
		vm.Pop(1)
	}
	if c > 0 {
		c = c - 1
	} else {
		c = Instruction(vm.Fetch()).Ax()
	}
	index := int64(c * LFIELDS_PER_FLUSH)
	for j := 1; j <= b; j++ {
		index++
		vm.PushValue(a + j)
		vm.SetI(a, index)
	}
	if bIsZero {
		for j := vm.RegisterCount() + 1; j <= vm.GetTop(); j++ {
			index++
			vm.PushValue(j)
			vm.SetI(a, index)
		}
		vm.SetTop(vm.RegisterCount())
	}
}
