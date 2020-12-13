/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-11-19 14:46:01
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-11-19 17:05:41
 */
package string

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	var str string = "abcabcbb"

	m, max, left := make(map[rune]int), 0, 0
	for idx, c := range str {

		if _, okay := m[c]; okay == true && m[c] >= left {
			if idx-left > max {
				max = idx - left
			}
			left = m[c] + 1
		}
		m[c] = idx
		fmt.Printf("index %d,left %d,max %d \n", idx, left, max)
		fmt.Println(m)
	}

	if len(str)-left > max {
		max = len(str) - left
	}

	fmt.Println(max)
}
