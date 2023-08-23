// 工厂函数

package main

import "strings"
func main() {
	
	var jpg = MakeAddSuffix(".jpg")
	var png = MakeAddSuffix(".png")

	var name1 = jpg("bob")
	var name2 = png("bob")

	print(name1)
	print(name2)
}

// type Add func(a b int) int

func MakeAddSuffix(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}