// 指针

package main
import "fmt"

func main(){

	var a = 10;
	var p = &a;
	var b = *p; // b 占用新的内存地址, 它与 a 的值相等, 但是内存地址不一样.
	
	*p = 20


	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("a address is ", p)

}