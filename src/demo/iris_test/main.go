// iris_test project main.go
package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		//		ctx.HTML("<b>Hello Iris<b/>")
		var reply map[string]interface{}
		reply = make(map[string]interface{})
		reply["c"] = 1
		reply["m"] = "ok"
		reply["d"] = "ok"

		//		ctx.XML(reply)
		ctx.JSON(reply)
	})
	app.Run(iris.Addr(":8080"))
}
