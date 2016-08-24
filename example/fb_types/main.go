package main

import (
	"github.com/funny/fastbin"
	"github.com/funny/fastbin/example/fb_types/module"
)

func main() {
	fastbin.Register(&module.MyStruct{})
	fastbin.GenCode()
}
