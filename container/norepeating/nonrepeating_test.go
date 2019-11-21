package main

import (
	"testing"
)

/** 普通测试 */
func TestSubStr(t *testing.T) {

	//可以在test case上写上注释

	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"这里是天二一", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}

	//使用run with coverage 可以查看代码覆盖率

	//使用命令行可以 go test -coverprofile=c.out  把覆盖数据输出到文件中
	//然后使用go tool cover 可以看帮助文档
	//go tool cover -html=c.out -o coverage.html 可以指定到html
	//或者直接使用 go tool cover -html=c.out

}

/** 性能测试 */
func BenchmarkSubStr(b *testing.B) {

	s, ans := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8
	for i := 0; i < 13; i++ {
		s += s
	}
	b.Logf("len(s): %d", len(s))
	//删除准备输入数据的时间
	b.ResetTimer()
	//N不用指定 系统帮我们算
	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr(s); actual != ans {
			b.Errorf("got %d for input %s ; expected %d", actual, s, ans)
		}
	}

	//BenchmarkSubStr-8            507           2328233 ns/op
	//BenchmarkSubStr-8   	       198	         6052143 ns/op

	//可以使用命令 go test -bench .
	//运行benchmark
	//go test -bench . -cpuprofile cpu.out
	//生成cpu.out
	//go tool pprof cpu.out
	//交互式命令行  help 可以查看命令
	//web

}
