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
	"testing"
)

func TestPrintKeepZero(t *testing.T) {

	tdata := TestPecision{
		Atest: 0,
		Btest: 12.01,
		Ctest: 12.10,
	}

	data, _ := json.Marshal(tdata)

	fmt.Println(string(data))
}
