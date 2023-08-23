

// 常量
package main
import (
	"fmt"
)
// 1. 隐式类型定义
const Pi = 3.14159

// 2. 显示类型定义
const Pi2 float64 = 3.14159

// 3. 并行定义
const (
	Monday, Tuesday, Wednesday = 1, 2, 3
)

// 4. 枚举定义
const (
	Unknown = 0
	Female = 1
	Male = 2
)

// const (
// 	a = Unknown
// 	b
// 	c 
// 	d = 5
// 	e
// )
// 4.1 iota 用于枚举时, 每当作用于一个新行, 它都将 ➕ 1
const (
	a = iota  // a = 0
	b         // b = 1
	c         // c = 2
	d = 5     // d = 5   
	e         // e = 5
)

func main() {

	fmt.Println(Pi)
	fmt.Println(Pi2)
	fmt.Println(e)
}