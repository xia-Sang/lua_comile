package number

import (
	"math"
	"strconv"
)

// 进行整除
func IFloorDiv(a, b int64) int64 {
	if a > 0 && b > 0 || a < 0 && b < 0 || a%b == 0 {
		return a / b
	} else {
		return a/b - 1
	}
}
func FFloorDiv(a, b float64) float64 {
	return math.Floor(a / b)
}

// 取模运算
func IMod(a, b int64) int64 {
	return a - IFloorDiv(a, b)*b
}
func FMod(a, b float64) float64 {
	return a - math.Floor(a/b)*b
}

// 移位运算
func ShiftLeft(a, n int64) int64 {
	if n >= 0 {
		return a << uint64(n)
	} else {
		return ShiftRight(a, -n)
	}
}
func ShiftRight(a, n int64) int64 {
	if n >= 0 {
		return int64(uint64(a) >> uint64(n))
	} else {
		return ShiftLeft(a, -n)
	}
}

// 浮点数解析
func FloatToInterger(f float64) (int64, bool) {
	i := int64(f)
	return i, float64(i) == f
}

// 字符串解析为数字
func ParseInteger(str string) (int64, bool) {
	i, err := strconv.ParseInt(str, 10, 64)
	return i, err == nil
}

func ParseFloat(str string) (float64, bool) {
	f, err := strconv.ParseFloat(str, 64)
	return f, err == nil
}

// 任意数值转换
