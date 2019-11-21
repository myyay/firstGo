package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	//每个字母最后出现的位置
	lastOccurred := make(map[rune]int)

	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength

}

//存储上次出现的位置+1   0表示没出现过
var lastOccurred = make([]int, 0xffff)

//清理内存会消耗时间  但如果串特别长的情况下，map的存取还是比数组要慢的
func lengthOfNonRepeatingSubStr2(s string) int {

	//会被编译成mem clear 通过pprof图可以看出
	for i := range lastOccurred {
		lastOccurred[i] = 0
	}

	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {

		if lastI := lastOccurred[ch]; lastI >= start {
			start = lastI
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i + 1
	}
	return maxLength

}

func main() {

	fmt.Println([]rune("abcabcbb"))
	fmt.Println([]rune("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))

	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("这里是我家！"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(lengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))

}
