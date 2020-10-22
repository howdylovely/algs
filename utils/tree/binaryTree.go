/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-09-20 13:00:35
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-10-22 15:08:11
 * @Doc:https://appliedgo.net/bintree/
 */

package tree

import (
	"errors"
)

// 二叉树中节点的定义
type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Data  int64
}

//插入节点操做
func (n *BinaryNode) Insert(data int64) {
	if n == nil {
		return
	} else if data <= n.Data {
		if n.Left == nil {
			n.Left = &BinaryNode{Data: data, Left: nil, Right: nil}
		} else {
			n.Left.Insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &BinaryNode{Data: data, Left: nil, Right: nil}
		} else {
			n.Right.Insert(data)
		}
	}
}

//查找最小的节点地址
func (n *BinaryNode) FindMinAdd(parent *BinaryNode) (*BinaryNode, *BinaryNode) {
	if n == nil {
		return nil, parent
	}
	if n.Left == nil {
		return n, parent
	}
	return n.Left.FindMinAdd(n)
}

//查找最小节点的值
func (n *BinaryNode) FindMinValue() int64 {
	if n.Left == nil {
		return n.Data
	}
	return n.Left.FindMinValue()
}

//最大的节点的值
func (n *BinaryNode) FindMaxValue() int64 {
	if n.Right == nil {
		return n.Data
	}
	return n.Right.FindMaxValue()
}

//查找节点的值是否存在并返回
func (n *BinaryNode) Find(s int64) (int64, bool) {
	if n == nil {
		return -1, false
	}
	switch {
	case s == n.Data:
		return n.Data, true
	case s < n.Data:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
}

//替换节点
func (n *BinaryNode) replaceNode(parent, replacement *BinaryNode) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}
	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

//删除节点
func (n *BinaryNode) Delete(s int64, parent *BinaryNode) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}
	switch {
	case s < n.Data:
		return n.Left.Delete(s, n)
	case s > n.Data:
		return n.Right.Delete(s, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}

		replacement, replParent := n.Right.FindMinAdd(n)
		n.Data = replacement.Data
		return replacement.Delete(replacement.Data, replParent)
	}
}

// 打印节点
func (t *BinaryNode) Traverse(n *BinaryNode, f func(*BinaryNode)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

//定义二叉树
type BinaryTree struct {
	Root *BinaryNode
}

//树中插入节点
func (t *BinaryTree) Insert(data int64) *BinaryTree {
	if t.Root == nil {
		t.Root = &BinaryNode{Data: data, Left: nil, Right: nil}
	} else {
		t.Root.Insert(data)
	}
	return t
}

// 二叉树中查找节点是否存在
func (t *BinaryTree) Find(s int64) (int64, bool) {
	if t.Root == nil {
		return -1, false
	}
	return t.Root.Find(s)
}

func (t *BinaryTree) Delete(s int64) error {
	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}

	fakeParent := &BinaryNode{Right: t.Root}
	err := t.Root.Delete(s, fakeParent)
	if err != nil {
		return err
	}
	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}

func RescursionPrint() {

}
