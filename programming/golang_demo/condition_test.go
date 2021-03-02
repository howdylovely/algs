package golang_demo

import (
	"log"
	"sync"
	"testing"
	"time"
)

// https://segmentfault.com/a/1190000037426003
// https://segmentfault.com/a/1190000025190291

func TestDemo3(t *testing.T) {
	cond := sync.NewCond(new(sync.Mutex)) // 初始化条件变量
	maxSize := 10
	counter := 0

	// 排水口
	go func() {
		for {
			cond.L.Lock()     // 上锁
			if counter == 0 { // 没水了
				cond.Wait() // 啥时候来水？等通知！
			}
			counter--
			log.Printf("OUTPUT counter = %d", counter)
			cond.Signal()               // 单发通知：已排水
			cond.L.Unlock()             // 解锁
			time.Sleep(5 * time.Second) // 为了演示效果，睡眠5秒
		}
	}()

	// 注水口
	for {
		cond.L.Lock()           // 上锁
		if counter == maxSize { // 水满了
			cond.Wait() // 啥时候排水？等待通知！
		}
		counter++
		log.Printf(" INPUT counter = %d", counter)
		cond.Signal()               // 单发通知：已来水
		cond.L.Unlock()             // 解锁
		time.Sleep(1 * time.Second) // 为了演示效果，睡眠1秒
	}
}

// go test -v ./condition_test.go -test.run TestDemo1

func TestDemo1(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	go func() {
		mut.Lock()
		log.Println("goroutine B Lock")
		counter = 1
		log.Println("goroutine B counter =", counter)
		time.Sleep(5 * time.Second)
		mut.Unlock()
		log.Println("goroutine B Unlock")
	}()
	time.Sleep(1 * time.Second)
	mut.Lock()
	log.Println("goroutine A Lock")
	counter = 2
	log.Println("goroutine A counter =", counter)
	mut.Unlock()
	log.Println("goroutine A Unlock")
}
