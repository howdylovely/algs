/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-09-21 09:46:07
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-09-21 09:51:21
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
