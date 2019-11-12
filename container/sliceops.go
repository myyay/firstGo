package main

import "fmt"

//slice的操作

func printSlice(s []int) {
	//即使s=nil也不会崩溃  cap是按2的倍数扩充的
	fmt.Printf("%v len=%d, cap=%d \n", s, len(s), cap(s))
}

func main() {
	//默认值 s=nil
	var s []int

	for i := 0; i < 100; i++ {
		//printSlice(s)
		s = append(s, i*2+1)
	}

	//fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)

	s3 := make([]int, 16, 32)
	printSlice(s3)

	fmt.Println(" copying slice")

	copy(s2, s1)
	printSlice(s2)

	s2[4] = 5
	//多向少拷贝， 超过cap的会被忽略
	copy(s1, s2)
	printSlice(s1)

	fmt.Println(" delete elements")

	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

}
