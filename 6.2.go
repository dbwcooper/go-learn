// 函数参数和返回值
package main

import (
	"fmt"
)
type Int []int

// func (st Stack) Pop() int {
// 	v := 0
// 	for ix := len(st) - 1; ix >= 0; ix-- {
// 			if v = st[ix]; v != 0 {
// 					st[ix] = 0
// 					// return v
// 					break;
// 			}
// 	}

// 	return v;
// }  


var num int = 10
var numx2, numx3 int

func main() {
	// var stack = []int{1, 2, 3, 4, 5}
	// var a = Pop(stack)
	// println(a)

	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()

}

func PrintValues() {
	fmt.Printf(" num = %d, 2x num = %d, 3x num = %d", num, numx2, numx3)
}

func getX2AndX3(input1 int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2, x3 = 2 * input, 3 * input
	return
}


func MinMax(a int, b int) (min int, max int){

	if(a < b) {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}

func multiply(a int, b int, reply *int) {
	*reply = a * b
}