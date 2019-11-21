package main

func main() {

	var s []int

	var i = 1
	for i < 1000000000 {
		s = append(s, i, i+1, i+2, i+3, i+4, i+5, i+6, i+7, i+8, i+9)
		i = i + 10

		//fmt.Println(s)
		s = s[10:]
	}

}
