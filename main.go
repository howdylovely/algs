/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-09-20 13:00:52
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-09-25 15:57:29
 */
package main

import (
	"algs/utils/tree"
	"fmt"
	"io"
)

func print(w io.Writer, node *tree.BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Data)
	print(w, node.Left, ns+2, 'L')
	print(w, node.Right, ns+2, 'R')
}

type Bdata struct {
	A int
	B int
}

type Adata struct {
	Bdata []Bdata
	C     int
}

func main() {

}
