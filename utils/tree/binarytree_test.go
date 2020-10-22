/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-09-21 09:46:07
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-10-22 15:09:16
 */
package tree

import (
	"fmt"
	"testing"
)

func TestTraverse(t *testing.T) {
	treea := &BinaryTree{}
	data := []int64{50, 25, 75, 11, 33, 61, 89, 30, 40, 52, 82, 95}
	for _, v := range data {
		treea.Insert(v)
	}
	treea.Root.Traverse(treea.Root, func(n *BinaryNode) { fmt.Print(n.Data, " | ") })
}

func TestRescursionPrint(t *testing.T) {
	// w io.Writer, node *tree.BinaryNode, ns int, ch rune
	// if node == nil {
	// 	return
	// }

	// for i := 0; i < ns; i++ {
	// 	fmt.Fprint(w, " ")
	// }
	// fmt.Fprintf(w, "%c:%v\n", ch, node.Data)
	// print(w, node.Left, ns+2, 'L')
	// print(w, node.Right, ns+2, 'R')
}
