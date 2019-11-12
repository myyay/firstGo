package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/** 转换成二进制 */
func convertToBin(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""

	for ; n > 0; n /= 2 {
		lsb := n % 2
		//fmt.Printf("cur n: %d, cur lsb: %d \n", n, lsb)
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	//文件只读
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printReaderContent(file)

}

func printReaderContent(file io.Reader) {
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

	printFile("abc.txt")

	fmt.Println("================")

	s := `abc"d"
	kkkk
    123

	p`
	printReaderContent(strings.NewReader(s))

}
