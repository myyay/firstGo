package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./superstar/_demo/web/views", ".html").Reload(true))

	app.Get("/", func(context context.Context) {
		context.WriteString("h from iris")
	})

	app.Get("/hello", func(c context.Context) {
		c.ViewData("Title", "我的页面")
		c.ViewData("Content", "Hello World!")
		c.View("hello.html")
	})

	app.Run(iris.Addr("localhost:8888"), iris.WithCharset("UTF-8"))
}
