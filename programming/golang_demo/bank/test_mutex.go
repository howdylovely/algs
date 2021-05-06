package bank

import (
	"sync"
)

var (
	mu       sync.Mutex
	balance3 int
)

func Deposit3(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance3 = balance3 + amount
}

func Balance3() int {
	mu.Lock()
	defer mu.Unlock()
	return balance3
}

// 为什么 Balance 也需要互斥，多线程同时读取一个变量有问题吗？？？ Balance 也需要互斥是为了防止在写的过程中进行读取。如果在读的过程中没有线程在写多线程同时读取是没有问题的。所以为了提高读取的并发量可以用读写锁改写

// 作者：lesliefang
// 链接：https://www.jianshu.com/p/c804a5e70743
// 来源：简书
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// var mu sync.RWMutex
// var balance int
//  func Balance() int {
//      mu.RLock() // readers lock
//      defer mu.RUnlock()
//      return balance
// }
