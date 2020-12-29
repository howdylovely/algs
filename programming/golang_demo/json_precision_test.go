/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-10-22 15:01:24
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-10-27 14:55:51
 */
package golang_demo

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestPrintKeepZero(t *testing.T) {

	var value float64 = 10.111226
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", value), 64)
	fmt.Println(value)

	tdata := TestPecision{
		Atest: 0,
		Btest: 12.01,
		Ctest: 12.10,
	}

	data, _ := json.Marshal(tdata)

	fmt.Println(string(data))
}
