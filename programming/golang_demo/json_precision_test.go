/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-10-22 15:01:24
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-10-22 15:04:30
 */
package golang_demo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPrintKeepZero(t *testing.T) {
	data, err := json.Marshal(Pt{40.0, "some_string"})
	fmt.Println(string(data), err)
}
