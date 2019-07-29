package main

import (
	"testing"
)

/** 性能测试 */
func BenchmarkSubStr(b *testing.B) {

	s, ans := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8

	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr2(s); actual != ans {
			b.Errorf("got %d for input %s ; expected %d", actual, s, ans)
		}
	}

}
