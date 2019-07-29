package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/** 转换成二进制 */
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, e := os.Open(filename)
	if e != nil {
		panic(e)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(256),
		convertToBin(0),
	)

	printFile("src/abc.txt")
}
