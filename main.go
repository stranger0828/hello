package main

import (
	"fmt"
	"hello/hello/modules/define"
	"hello/hello/modules/print"
)

func main() {
	fmt.Println("hello world!")
	print.Printdetail()
	fmt.Println(define.HeroName)
}
