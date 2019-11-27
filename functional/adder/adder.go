package main

import "fmt"

//返回值是函数  闭包的概念 返回一个闭包  这个数据里sum也在闭包里
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

//正统函数式编程不能有状态sum ，因此要把状态放在一个新的函数里面
//递归定义一个函数
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)

	}
}

func main() {

	fmt.Println("adder()...")

	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}

	fmt.Println("adder2()...")

	b := adder2(0)
	for i := 0; i < 10; i++ {
		//fmt.Println(a(i))
		var s int
		//把b指向下一个
		s, b = b(i)
		fmt.Println(i, s)
	}

}
