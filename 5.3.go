// switch 
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	k :=6 
	switch k {
	case 4:
		fmt.Println(" 4")
		fallthrough
	case 5:
		fmt.Println(" 5")
		fallthrough
	case 6:
		fmt.Println(" 6")
		fallthrough
	case 7:
		fmt.Println(" 7")
		fallthrough
	case 8:
		fmt.Println(" 8")
		fallthrough
	default:
		fmt.Println("default")

	}
}


func season(num int) string {
	switch num {
	case 12, 1, 2:
		return "winter"
	case 3, 4, 5:
		return "spring"
	case 6, 7, 8:
		return "summer"
	case 9, 10, 11:
		return "autumn"
	}
	return ""
}