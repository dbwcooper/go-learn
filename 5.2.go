// 多返回值的情况

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var origin string = "ABC"
	
	an, err := strconv.Atoi(origin)
	if err != nil {
	  fmt.Println(err)
		return
	}
	fmt.Println(an)

}