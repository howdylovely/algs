package golang_demo

import (
	"fmt"
	"strconv"
	"testing"
)

// unicode utf8
// http://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html

// https://studygolang.com/articles/15455?fr=sidebar

// rune 、 byte 和 string 都是 Go 的内置类型
// byte
// byte是uint8的别名，在所有方面都等同于uint8
// 按惯例，它用于区分字节值和8位无符号整数值。
// rune
// rune是int32的别名，在所有方面都等同于int32
// 按惯例，它用于区分字符值和整数值。
// string
// string是所有8位字节字符串的集合，通常但不一定代表UTF-8编码的文本
// 字符串可能为空，但是不能为 nil
// 字符串类型的值是不可变的
// 由上面得解释我们大概可以明白
// rune 可以表示得比 byte 多
// string 类型的底层是一个byte 数组
// 以上解释都来此 Go 源码注释
// 刚刚上面标注了字节和字符，现在我们来梳理字符和字节的概念
// 存储单位 字节

// 计算机存储信息的最小单位，称之为位 bit，二进制的一个0或1叫一位
// 计算机存储容量基本单位是字节 Byte，8个二进制位组成 1 个字节
// 信息表示单位 字符

// 字符 是一种符号，像 英文a和中文阿 就是不同字符
// 不同的字符在不同的编码格式下，所需要的存储单位不一样
// ASCLII 编码中一个英文字母一字节，一个汉字两字节
// UTF-8 编码中 一个英文字母一字节，一个常见汉字3字节，不常用的超大字符集汉字4字节
// Go 源码文件默认采用Unicode字符集，Unicode码点和内存中字节序列的变换实现使用了UTF-8，这使得Go编程无需考虑编码转换的问题非常方便
// 从编码上来分析
// byte用来强调一个字节代表的数据（例如字符 a 就是 97），而不是数字；
// rune用来表示Unicode的码点，即一个字符
// 通俗一点
// byte 只能操作简单的字符，不支持中文操作
// rune 能操作任何字符

func TestBytesDemo(t *testing.T) {
	// str := "Hello world !你好啊兵"
	// strBytes := []byte(str)
	// strRune := []rune(str)

	// fmt.Printf("%+v", strBytes)
	// fmt.Println()

	// for _, v := range strBytes {
	// 	fmt.Println()
	// 	fmt.Printf("%+v", string(v))
	// }

	// for _, v := range strRune {
	// 	fmt.Println()
	// 	fmt.Printf("%+v", string(v))

	// }

	sText := "严"
	quoted := strconv.Quote(sText)
	textQuoted := strconv.QuoteToASCII(sText)
	// textUnquoted := textQuoted[1 : len(textQuoted)-1]
	fmt.Println(quoted)
	fmt.Println(textQuoted)

}

func TestTruth(t *testing.T) {
	if true != true {
		t.Fatalf("The world is crumbiling")
	}
}

func BenchmarkStringFromAssignment(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAssignment(100)
	}
}
func BenchmarkStringFromAppendJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAppendJoin(100)
	}
}
func BenchmarkStringFromBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromBuffer(100)
	}
}
