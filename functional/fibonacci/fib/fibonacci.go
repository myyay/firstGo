package fib

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

/*func main() {

	f := main2.Fibonacci()
	//是不是有点像文件
	fmt.Println(f()) //1
	fmt.Println(f()) //1
	fmt.Println(f()) //2
	fmt.Println(f()) //3
	fmt.Println(f()) //5
	fmt.Println(f()) //8
	fmt.Println(f()) //13
	fmt.Println(f()) //21
	fmt.Println(f()) //34

}
*/
