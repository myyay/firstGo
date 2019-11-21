package main

import "fmt"

func main() {
	//map[K]V  复合map map[K}map[K2]V

	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) //m2 == empty map

	var m3 map[string]int

	//添加
	m2["cc"] = 9
	m2["dd"] = 4

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")

	for k, v := range m {
		fmt.Println(k, v)

	}

	fmt.Println("Getting values")

	//courseName, ok := m["course"]
	//fmt.Println(courseName, ok)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key dose not exists")
	}

	fmt.Println("Deleting values")

	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
	//打印是空值， 非nil
	fmt.Println("m[name]:", m["name"])

}
