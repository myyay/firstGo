package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {

	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation:" + op)

	}
}

// q为
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {

	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()

	fmt.Printf("Calling Function %s with args "+"( %d , %d) \n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b *int) {
	//值传递  以下无效
	//b, a = a, b

	*b, *a = *a, *b

}

func main() {

	fmt.Println(eval(3, 4, "*"))

	fmt.Println(div(13, 3))

	q, r := div(15, 3)

	fmt.Println(q, r)

	if i, e := eval(3, 4, "x"); e != nil {
		fmt.Println("Error: ", e)
	} else {
		fmt.Println(i)
	}

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	swap(&a, &b)

	fmt.Println(a, b)
}
