package golang_demo

import (
	"fmt"
	"strconv"
)

// string

// string to int,非常快
func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return i, nil
}




// 注意没有32,最快
func StirngToInt64(s string) (i int64, err error) {
	if i, err = strconv.ParseInt(s, 10, 64); err == nil {
		return i, nil
	}
	return 0, err
}

func StingToFloat32(s string) (f float64, err error) {
	if f, err := strconv.ParseFloat(s, 32); err == nil {
		return f, nil
	}
	return 0, err
}

// int
func IntToString(i int) string {
	s := strconv.Itoa(i) // s == "97"
	return s
}
func FormatIntToString(i int64) string {
	s := strconv.FormatInt(i, 10) // s == "97" (decimal)
	return s
}
func IntToFloat64(i int) float64 {
	return float64(i)
}

// int64
func IntToInt64(i int) int64 {
	return int64(i)
}

// 会被截断
func Int64ToInt(i int64) int {
	// var m int64 = 2 << 32
	return int(i)
}

// float
func FloatToString(f float64) string {
	s := fmt.Sprintf("%f", 123.456) // s == "123.456000"
	return s
}

// 舍去小数点的内容
func FloatToInt(f float64) int {
	return int(f)
}
