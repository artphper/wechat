package main

import (
	_ "wachat-demo/boot"
	_ "wachat-demo/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
