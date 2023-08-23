// defer 和追踪
package main

import (
	"fmt"
	"io"
	"log"
)

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
	test("Goland")
}

func test(p string) (n int, err error) {

	defer func(){
		log.Printf("test(%q) = %d, %v", p, n, err)
	}()

	return 0, io.EOF
}