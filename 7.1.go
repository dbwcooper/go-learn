// 切片
package main

import (
	"fmt"
)


func array() {

	// 一维数组的三种定义方式
	// 1. 指定长度
	var arrFixedLen1 [2]int;
	fmt.Println(arrFixedLen1) // [0 0]
	var arrFixedLen2 = [2]int{1,2}
	fmt.Println(arrFixedLen2) // [1 2]

	// 2. 使用 ... 做长度推断
	var arrGoLen = [...]int{1, 2, 3}
	fmt.Println(arrGoLen) // [1 2 3]

	// 3. 指定下标
	var arrIndex = [...]int{1: 10, 2, 3}
	fmt.Println(arrIndex) // [0 10 2 3]

	// 定义二维数组的两种方式
	// 1. 使用 ... 长度推断
	var arr4 =  [...][2]int {
		{1,2},
		{3,4},
	}
	// 2. 指定长度
	var arr5 = [1][5]int{
		{1,2,3,4,5},
	}
	fmt.Println(arr4, arr5)

	// invalid operation: arrIndex == arrGoLen (mismatched types [4]int and [3]int)
	// fmt.Println(arrIndex == arrGoLen)

}

// 数组是一个值类型, 修改数组中的值之后, 不会影响到原数组
func f(a [3]int) {
	a[0] = 10;
	fmt.Println("f a: ", a)
}

// 传递一个数组的引用, 修改数组中的值之后, 会影响到原数组
func fp(a *[3]int) {
	a[0] = 0;
	fmt.Println("fp a: ", a)
}

func function() {
	var ccc [3]int = [3]int{1,2,3}
	f(ccc) // 值
	fp(&ccc) // 引用 fp 内部修改传入参数的值会直接反映到 ccc 上
	fmt.Println("main ccc: ", ccc)
}


func main() {

	array()
	function()

	// 1. 切片的声明, 不需要指定长度
	var a []int
	var aNil []int
	a = []int{1,2,3,4,5}
	
	var a2 []int = a[1: 5] // [2,3,4,5]
	var a3 []int = a[:] // [1,2,3,4,5]
	var a4 []int = a[1:] // [2,3,4,5]
	var a5 []int = a[:4] // [1,2,3,4]

	fmt.Println(a)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(aNil == nil) // true


	// 2. 切片作为函数参数
	var arr = make([]int, 2, 5)	

 // 3. make 和 new 的区别
 //   make 返回值
 //   new 返回指针
	var arrr = new([5] int)
	fmt.Println(arr)
	fmt.Println(arrr)

	fmt.Println(sum(arr))

}





func sum(a []int) int {
	s :=0
	for i:=0; i< len(a); i++ {
		s += a[i]
	}
	return s
}