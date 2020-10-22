/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-10-21 20:56:11
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-10-22 13:42:02
 */
package heap

import (
	"fmt"
	"testing"

	"github.com/facebookgo/pqueue"
)

func TestBinomialHeap(t *testing.T) {
	q := pqueue.New(7)

	q.Push(&pqueue.Item{Priority: 2, Value: 2, Index: 1})
	q.Push(&pqueue.Item{Priority: 3, Value: 3, Index: 2})
	q.Push(&pqueue.Item{Priority: 5, Value: 5, Index: 3})
	q.Push(&pqueue.Item{Priority: 4, Value: 4, Index: 4})
	q.Push(&pqueue.Item{Priority: 6, Value: 6, Index: 5})
	q.Push(&pqueue.Item{Priority: 1, Value: 1, Index: 6})

	q.PeekAndShift(6)

	x := q.Pop()

	q.PeekAndShift(5)
	x = q.Pop()

	q.PeekAndShift(5)
	x = q.Pop()

	fmt.Println(x)
}
