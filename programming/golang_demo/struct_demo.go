package golang_demo

import "fmt"

// 实参接收者和形参接收者是同一类型，比如都是T或者都是*T。（1，4，5，7）
// 实参接收者是T类型的变量而形参接收者是*T类型，编译器会隐式的获取变量的地址（3）。
// 实参接收者是T类型而形参接收者是T类型，编译器会隐式的获取实际的取值。（2，6）
// 其中8编译过程报错的原因是：编译器对T类型转化为T类型的隐式转化，只有实参接收者是变量才可以成功，因为无法获取临时变量的地址。

// 作者：百味纯净水
// 链接：https://www.jianshu.com/p/b6ae3f85c683
// 来源：简书
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

type Point struct {
	X int
	Y int
}

func (p Point) Print() {
	fmt.Println(p.X, p.Y)
}
func (p *Point) ScaleBy(factor int) {
	p.X *= factor
	p.Y *= factor
}
func main_struct() {
	p := Point{1, 1}
	ptr := &p
	p.Print()                 //1. 正确
	ptr.Print()               //2. 正确
	p.ScaleBy(2)              //3. 正确
	ptr.ScaleBy(2)            //4. 正确
	Point{1, 1}.Print()       //5. 正确
	(&Point{1, 1}).Print()    //6. 正确
	(&Point{1, 1}).ScaleBy(2) //7. 正确
	// Point{1, 1}.ScaleBy(2)    //8. 错误
}

// https://code.tutsplus.com/zh-hans/tutorials/lets-go-object-oriented-programming-in-golang--cms-26540
