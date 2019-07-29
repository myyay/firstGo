package main

import (
	"firstGo/src/retriever/mock"
	"firstGo/src/retriever/real"
	"fmt"
	"time"
)

const url = "http://www.baidu.com"

type Retriever interface {
	//interface里 全是函数  没有属性
	Get(url string) string
}
type Poster interface {
	//interface里 全是函数  没有属性
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)

}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "me",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	return s.Post(url, map[string]string{
		"name":   "me",
		"course": "golang",
	})
}

func main() {

	var r Retriever
	r = &mock.Retriever{Contents: "this is fake baidu"}
	retriever := mock.Retriever{Contents: "this is fake baidu"}
	inspect(r)

	r = &real.Retriever{UserAgent: "Mozilla/5.0", Timeout: time.Minute}
	inspect(r)

	if retriever, ok := r.(*real.Retriever); ok {
		fmt.Println("UserAgent", retriever.UserAgent)
	} else {
		fmt.Println(" not a *real.Retriever")
	}

	fmt.Println("try a session")
	fmt.Println(session(&retriever))
	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting ", r)

	fmt.Printf("> %T %v\n", r, r)
	fmt.Printf("%T %v\n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}
}
