package golang_demo

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// chan 通讯 共享内存

// ok 判断是否关闭
// close 关闭管道，不向接收端关闭，仅在发送端关闭
// make 1,缓存
// 向一个关闭的管道中发送数据,是会panic
// 向一个已满的管道中放数据时，会阻塞
func NewChan() {
	strChan := make(chan struct{}, 1)
	defer func() {
		close(strChan)
	}()
	strChan <- struct{}{}
	elem, ok := <-strChan

	fmt.Println(elem, ok)
}

// 非缓冲channel,需启动一个协程等着接
func UnCacheChan() {
	unbufChan := make(chan int)
	//unbufChan := make(chan int, 1) 有缓冲容量

	//启用一个Goroutine接收元素值操作
	go func() {
		fmt.Println("Sleep a second...")
		time.Sleep(time.Second) //休息1s
		num := <-unbufChan      //接收unbufChan通道元素值
		fmt.Printf("Received a integer %d.\n", num)
	}()

	num := 1
	fmt.Printf("Send integer %d...\n", num)
	//发送元素值
	unbufChan <- num
	fmt.Println("Done.")
}

/**
* 阻塞的方法
 */
//信号
func SignalAsleep() {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
}

// 并发阻塞,确保多个协程执行完后执行wait后的内容
func WaitGroupAsleep() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i = i + 1 {
		wg.Add(1)
		go func(n int) {
			// defer wg.Done()
			defer wg.Add(-1)
			time.Sleep(3e9)
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

// channel阻塞,nil值也会阻塞
/**
* 这种会保证 协程执行
 */
func ChanAsleep() {
	ch := make(chan int, 1)
	go func() {
		log.Println("wait 3s...")
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	<-ch
}

// select需要业务
func SelectAsleep() {
	select {}
}

// for需要业务
func forAsleep() {
	for {
	}
}

// sync.Mutex
func MutexAsleep() {
	var m sync.Mutex
	m.Lock()
	m.Lock()
}
