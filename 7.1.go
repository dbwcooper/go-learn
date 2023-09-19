// 切片
package main

import (
	"fmt"
)


func main() {

	// 1. 切片的声明, 不需要指定长度
	var a []int
	a = []int{1,2,3,4,5}
	
	var a2 []int = a[1: 5] // a2 = a[start, end-1]

	var a3 []int = a[:] // a3 = a

	fmt.Println(a)
	fmt.Println(a2)
	fmt.Println(a3)


}


