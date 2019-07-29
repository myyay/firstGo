package main

import "fmt"

//数组是【值类型】 数组是有界的
func printArray(arr [5]int) {

	arr[0] = 100

	for i, v := range arr {
		fmt.Println(i, v)

	}

}

func main() {
	var arr1 [5]int

	arr2 := [3]int{1, 3, 5}

	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)

	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])

	}

	printArray(arr1)
	printArray(arr3)

	fmt.Println(arr3)

}
