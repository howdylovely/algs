/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-09-20 13:00:52
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-09-21 09:50:55
 */
package main

import (
	"algs/utils/tree"
	"fmt"
	"io"
	"os"
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

func main() {
	treea := &tree.BinaryTree{}
	data := []int64{50, 25, 75, 11, 33, 61, 89, 30, 40, 52, 82, 95}
	for _, v := range data {
		treea.Insert(v)
	}
	print(os.Stdout, treea.Root, 0, 'M')
	// tree.Root.PrintInorder()
	treea.Root.Delete(50, treea.Root)
	print(os.Stdout, treea.Root, 0, 'M')
	// fmt.Println(tree.Root.Find(51))

	treea.Root.Traverse(treea.Root, func(n *tree.BinaryNode) { fmt.Print(n.Data, " | ") })
}
