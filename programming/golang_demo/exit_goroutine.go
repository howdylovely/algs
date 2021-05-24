// func doSomethingContext(ctx context.Context) error {
// 	for i := 0; i < 3000; i++ {
// 		time.Sleep(time.Second * 1)
// 		fmt.Printf("a")

// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("退出 goroutine")
// 			return nil
// 		default:
// 			fmt.Println("监控...")
// 			time.Sleep(1 * time.Second)

// 		}
// 	}

// 	return nil
// }
// func longRunningOperation(ctx context.Context, signal chan<- struct{}) {
// 	if err := doSomethingContext(ctx); err != nil {
// 		return
// 	}

// 	fmt.Println("退出 signal")
// 	signal <- struct{}{} // signal that this operation has finished (successfully)

// 	fmt.Println("退出 signal")
// }

// 	// go longRunningOperation(c.Request.Context(), signal)