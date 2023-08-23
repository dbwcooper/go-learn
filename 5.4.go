// for 循环

package main

import (
	"fmt"
	// "strings"
)
func main() {
	
	str := "G"

	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}
	// fmt.Printf(str)


	i := 5
	for i > 0 {
		i-- 
		println("i : ", i)
	}


	strA := "New York"

	for index, v:= range strA {
		fmt.Printf("Character on position %d is: %c \n", index, v)
		println("v: ", v)
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v) // 0 0 0 0 0 
		v = 5
	}

	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i is now:", i)
	// }

	// for i := 0; i < 3; {
	// 	fmt.Println("Value of i:", i)
	// }
	
	// s := ""
	// for ; s != "aaaaa"; {
	// 	fmt.Println("Value of s:", s)
	// 	s = s + "a"
	// }

// 	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
// 	s = i+1, j+1, s + "a" {
// 	fmt.Println("Value of i, j, s:", i, j, s)
// }

}