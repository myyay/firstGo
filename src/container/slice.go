package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println(arr)

	//左闭右开区间， 取第2个到第5个  slice是一个视图
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:]=", s1)
	s2 := arr[:]
	fmt.Println("arr[:]=", s2)

	updateSlice(s1)
	updateSlice(s2)

	fmt.Println(arr)

	s3 := arr[2:6]
	s4 := s3[3:5]
	//神奇的地方 映射的原数组arr
	fmt.Println(s3, s4)
	fmt.Println(len(s3), len(s4))
	fmt.Println(cap(s3), cap(s4))

	s5 := append(s1, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)
	s8 := append(s7, 13)

	fmt.Println(s5, s6, s7, s8)

	fmt.Println(arr)

}
