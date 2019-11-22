package main

import "testing"

func BenchmarkIncrement(b *testing.B) {

	var s []int

	//for i := 0; i < b.N; i++ {
	for j := 0; j < 100000000; j++ {
		s = append(s, j, j+1, j+2, j+3, j+4, j+5, j+6, j+7, j+8, j+9)
		j = j + 10

		//fmt.Println(s)
		s = s[10:]
	}
	//}

}
