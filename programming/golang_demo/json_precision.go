/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-10-22 15:01:24
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-10-22 15:06:11
 */
package golang_demo

import "strconv"

// float64 精度处理

type KeepZero float64

func (f KeepZero) MarshalJSON() ([]byte, error) {
	if float64(f) == float64(int(f)) {
		return []byte(strconv.FormatFloat(float64(f), 'f', 1, 32)), nil
	}
	return []byte(strconv.FormatFloat(float64(f), 'f', -1, 32)), nil
}

type Pt struct {
	Value KeepZero
	Unit  string
}

func PrintKeepZero() {

}
