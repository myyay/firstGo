package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score : %d", score))

	case score < 60:
		g = "F"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g

}

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("contents: %s \n", contents)
	}

	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("dir: %s \n", dir)

	}

	fmt.Println(grade(70))
	fmt.Println(grade(80))
	fmt.Println(grade(90))
	fmt.Println(grade(101))

}
