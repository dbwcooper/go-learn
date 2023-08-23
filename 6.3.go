// 函数参数切片
package main

import "fmt"
func main() {
	Greeting("hello: ", "tom", "alice")

	minNumber(1 ,2,3,4,5)
}



// 1. 切片的参数类型都一致
func Greeting(prefex string, who ...string) string {

	if len(who) ==0 {
		return ""
	}

	var str = prefex;
	for _, v := range who {
		
		str += v + ","
	}
	fmt.Println(str)
	return str
}

func minNumber(n ...int) int {

	if len(n) == 0 {
		return 0
	}

	min := n[0]
	for _, v := range n {
		if v < min {
			min = v
		}
	}
	fmt.Println(min)

	return min
}

// 2. 切片的参数类型不一致 TODO?

type Person struct {
	name string
	age int
	address string
	country string
}

func findWho(param ...) {

}