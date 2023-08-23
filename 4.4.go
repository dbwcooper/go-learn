// 变量 - var

package main
import (
	"fmt"
	"os"
)


var name = "tom"

func main() {
	// 1. 变量的初始值
	var a int     // 0
  var b float64 // 0.0
  var c string  // ''
  var d *int    // nil


	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)


	// 2 变量的命名方式: 驼峰

	var startDate string


	// 3. 局部变量, 块级变量

	var name = "bob"

	fmt.Println("startDate: ", startDate)
	fmt.Println("name: ", name)
	
	if(true) {
		var age = 12
	  fmt.Println("age: ", age)
   }
	//fmt.Println("age: ", age)

	// 4. 声明方式 :=

	path := os.Getenv("PATH")
	fmt.Println("path: ", path)

	test()


	// 5. 并行赋值
	// aa: 1, bb: 2, cc: 3
	var aa, bb, cc = 1, 2, 3 


	fmt.Println(" aa: ", aa)
	fmt.Println(" bb: ", bb)
	fmt.Println(" cc: ", cc)

	// 交换值
	aa, bb = bb, aa
	print(" aa: ", aa)
	fmt.Println(" bb: ", bb)

	// 6.init 函数, 无法被单独调用, 它在当前文件被初始化之后自动执行
}


func test(){
	// undefined: path
  // fmt.Printf("path: ", path)
}