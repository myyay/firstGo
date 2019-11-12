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
	const filename = "basic/branch/abc.txt"

	if bytes, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("%s\n", bytes)
	}
	//os.Args[0] C:\Users\userName\AppData\Local\Temp\___go_build_firstGo_basic_branch.exe
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("dir: %s, os.Args[0]: %s \n", dir, os.Args[0])

	}
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		// Uncomment to see it panics.
		// grade(-3),
	)

}
