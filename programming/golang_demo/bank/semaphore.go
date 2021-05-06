package bank

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit2(amout int) {
	sema <- struct{}{} // acquire token
	balance = balance + amout
	<-sema // release token
}

func Balance2() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}
