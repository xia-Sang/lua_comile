package state

import (
	"luago/api"
	"luago/number"
	"math"
)

var (
	iadd  = func(a, b int64) int64 { return a + b }
	fadd  = func(a, b float64) float64 { return a + b }
	isub  = func(a, b int64) int64 { return a - b }
	fsub  = func(a, b float64) float64 { return a - b }
	imul  = func(a, b int64) int64 { return a * b }
	fmul  = func(a, b float64) float64 { return a * b }
	imod  = number.IMod
	fmod  = number.FMod
	pow   = math.Pow
	div   = func(a, b float64) float64 { return a / b }
	iidiv = number.IFloorDiv
	fidiv = number.FFloorDiv
	band  = func(a, b int64) int64 { return a & b }
	bor   = func(a, b int64) int64 { return a | b }
	bxor  = func(a, b int64) int64 { return a ^ b }
	shl   = number.ShiftLeft
	shr   = number.ShiftRight
	iunm  = func(a, _ int64) int64 { return -a }
	funm  = func(a, _ float64) float64 { return -a }
	bnot  = func(a, _ int64) int64 { return ^a }
)

type operator struct {
	integerFunc func(int64, int64) int64
	floatFunc   func(float64, float64) float64
}

var operators = []operator{
	operator{iadd, fadd},
	operator{isub, fsub},
	operator{imul, fmul},
	operator{imod, fmod},
	operator{nil, pow},
	operator{nil, div},
	operator{iidiv, fidiv},
	operator{band, nil},
	operator{bor, nil},
	operator{bxor, nil},
	operator{shl, nil},
	operator{shr, nil},
	operator{iunm, funm},
	operator{bnot, nil},
}

func (ls *luaState) Arith(op api.ArithOp) {
	var a, b luaValue
	b = ls.stack.pop()
	if op != api.LUA_OPUNM && op != api.LUA_OPBNOT {
		a = ls.stack.pop()
	} else {
		a = b
	}

	operator := operators[op]
	if result := arith(a, b, operator); result != nil {
		ls.stack.push(result)
	} else {
		panic("arithmetic error!")
	}
}
func arith(a, b luaValue, op operator) luaValue {
	if op.floatFunc == nil {
		if x, ok := convertToInteger(a); ok {
			if y, ok := convertToInteger(b); ok {
				return op.integerFunc(x, y)
			}
		}
	} else {
		if op.integerFunc != nil {
			if x, ok := a.(int64); ok {
				if y, ok := b.(int64); ok {
					return op.integerFunc(x, y)
				}
			}
		}
		if x, ok := convertToFloat(a); ok {
			if y, ok := convertToFloat(b); ok {
				return op.floatFunc(x, y)
			}
		}
	}
	return nil
}
func (ls *luaState) Len(index int) {
	val := ls.stack.get(index)
	if s, ok := val.(string); ok {
		ls.stack.push(int64(len(s)))
	} else {
		panic("length error!")
	}
}
func (ls *luaState) Compare(index1, index2 int, op api.CompareOp) bool {
	a := ls.stack.get(index1)
	b := ls.stack.get(index2)
	switch op {
	case api.LUA_OPEQ:
		return eq(a, b)
	case api.LUA_OPLT:
		return lt(a, b)
	case api.LUA_OPLE:
		return le(a, b)
	default:
		panic("invalid compare op!")
	}
}
func eq(a, b luaValue) bool {
	switch x := a.(type) {
	case nil:
		return b == nil
	case bool:
		y, ok := b.(bool)
		return ok && x == y
	case string:
		y, ok := b.(string)
		return ok && x == y
	case int64:
		switch y := b.(type) {
		case int64:
			return x == y
		case float64:
			return float64(x) == y
		default:
			return false
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x == float64(y)
		case float64:
			return x == y
		default:
			return false
		}
	default:
		return a == b
	}
}
func lt(a, b luaValue) bool {
	switch x := a.(type) {
	case string:
		if y, ok := b.(string); ok {
			return x < y
		}
	case int64:
		switch y := b.(type) {
		case int64:
			return x < y
		case float64:
			return float64(x) < y
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x < float64(y)
		case float64:
			return x < y
		}
	}
	panic("comparison error!")
}
func le(a, b luaValue) bool {
	switch x := a.(type) {
	case string:
		if y, ok := b.(string); ok {
			return x > y
		}
	case int64:
		switch y := b.(type) {
		case int64:
			return x > y
		case float64:
			return float64(x) > y
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x > float64(y)
		case float64:
			return x > y
		}
	}
	panic("comparison error!")
}
func (ls *luaState) ConCat(n int) {
	if n == 0 {
		ls.stack.push("")
	} else if n >= 2 {
		for i := 1; i < n; i++ {
			if ls.IsString(-1) && ls.IsString(-2) {
				s2 := ls.ToString(-1)
				s1 := ls.ToString(-2)
				ls.stack.pop()
				ls.stack.pop()
				ls.stack.push(s1 + s2)
				continue
			}
			panic("concatenation error!")
		}
	} //n=1不处理
}
