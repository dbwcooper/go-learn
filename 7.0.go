// 数组
package main

import (
	"fmt"
)


func main() {

	// 1. 声明
	// 使用 切片和指针作为函数参数, 避免参数为大数组时消耗内存
	var arr1 [5]int // 值类型
	var arr2 = new([5]int) // 指针类型(引用类型)
	aaa := [...]string{"a", "b", "c"} // 切片命名法
	aaa2 := [10]string {1: "a", 2: "b", 3: "c" } 

	var multipleArr [5][5]int // 二维数组

	// 2. 循环数组 for
	//    数组不能越界, 否则会出现越界错误, 数组从索引 0 开始
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	// 3. 循环数组 range
	for i, v := range arr1 {
		fmt.Println(i, v)
	}
	for i, _ := range aaa {
		fmt.Println("Array item", i, " is ", aaa[i])
	}


	fmt.Println("arr2: ", arr2)
	fmt.Println("arr1: ", arr1)
	// 4. 引用类型和值类型 数组
	// 4.1 引用类型数组 作为函数参数传递时可以传递指针 func(&arr)

	p := arr2
  array := arr1

	p[0] = 1;
	array[0] = 1

	// 引用类型赋值后 p === arr2
	fmt.Println("p:  ", p)
	fmt.Println("arr2:  ", arr2)

	// 值类型赋值后 array !== arr1
	fmt.Println("array: ", array)
	fmt.Println("arr1:  ", arr1)
}
