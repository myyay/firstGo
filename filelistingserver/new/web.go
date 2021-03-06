package main

import (
	"firstGo/filelistingserver/new/handler"
	"log"
	"net/http"
	"os"
	//有了这个就可以监控服务器, 可以到这个文件里看都有哪些操作！！！！
	_ "net/http/pprof"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {

			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}

		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error occurred  handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusUnauthorized
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)

		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(handler.FileServerHandler))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
