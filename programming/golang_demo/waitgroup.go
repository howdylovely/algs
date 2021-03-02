package golang_demo

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

var Cnt int

func MyAdd(iter int) {
	for i := 0; i < iter; i++ {
		Cnt++
	}
}

// 注意 需要用go test trace,生成trace.out文件,然后执行命令
func main_waitgroup() {

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	// Your program here

	wg := &sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			MyAdd(100000)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(Cnt)
}
