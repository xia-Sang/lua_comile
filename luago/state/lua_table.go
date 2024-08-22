package state

import (
	"luago/number"
	"math"
)

// 表的内部实现
type luaTable struct {
	arr      []luaValue            //存储数组部分
	innerMap map[luaValue]luaValue //存放哈希表部分
}

// 创建一个空的表
func newLuaTable(nArr, nRec int) *luaTable {
	t := &luaTable{}
	if nArr > 0 {
		t.arr = make([]luaValue, 0, nArr)
	}
	if nRec > 0 {
		t.innerMap = make(map[luaValue]luaValue, nRec)
	}
	return t
}

// get数据元素
func (t *luaTable) get(key luaValue) luaValue {
	key = floatToInteger(key)
	if index, ok := key.(int64); ok {
		if index >= 1 && index <= int64(len(t.arr)) {
			return t.arr[index-1]
		}
	}
	return t.innerMap[key]
}
func floatToInteger(key luaValue) luaValue {
	if f, ok := key.(float64); ok {
		if i, ok := number.FloatToInterger(f); ok {
			return i
		}
	}
	return key
}
func (t *luaTable) put(key, value luaValue) {
	if key == nil {
		panic("table index is nil!")
	}
	if f, ok := key.(float64); ok && math.IsNaN(f) {
		panic("table index is NaN!")
	}
	key = floatToInteger(key)

	if index, ok := key.(int64); ok && index >= 1 {
		arrLength := int64(len(t.arr))
		if index <= arrLength {
			t.arr[index-1] = value
			if index == arrLength && value == nil {
				t.shrinkArray()
			}
			return
		}
		if index == arrLength+1 {
			delete(t.innerMap, key)
			if value != nil {
				t.arr = append(t.arr, value)
				t.expandArray()
			}
			return
		}
	}
	if value != nil {
		if t.innerMap == nil {
			t.innerMap = make(map[luaValue]luaValue, 8)
		}
		t.innerMap[key] = value
	} else {
		delete(t.innerMap, key)
	}
}
func (t *luaTable) shrinkArray() {
	for i := len(t.arr) - 1; i >= 0; i-- {
		if t.arr[i] == nil {
			t.arr = t.arr[:i]
		}
	}
}
func (t *luaTable) expandArray() {
	for index := int64(len(t.arr)) + 1; true; index++ {
		if val, ok := t.innerMap[index]; ok {
			delete(t.innerMap, index)
			t.arr = append(t.arr, val)
		} else {
			break
		}
	}
}
func (t *luaTable) len() int {
	return len(t.arr)
}
