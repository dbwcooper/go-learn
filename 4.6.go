// 字符串类型和操作
package main

import "fmt"
import "unicode/utf8"

func main() {
	str1 := "welcome to the rock world"
	
	fmt.Printf("str length: %d \n", len(str1))

	fmt.Printf("string count: %d \n", utf8.RuneCountInString(str1))
}