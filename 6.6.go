// 递归调用
package main
import (
	"time"
	"fmt"
)

func main() {

	start := time.Now()
	result, _ := fibo(40)
	end := time.Now()

	detail := end.Sub(start)
	fmt.Printf("%d \n", result)
	fmt.Printf("longCalculation took this amount of time: %s\n", detail)


	start2 := time.Now()
	res := fibonacci(40)
	end2 := time.Now()

	dur2 := end2.Sub(start2)
	fmt.Printf("%d \n", res)
	fmt.Printf("time: %s\n", dur2)

}

func fibo(n int) (result uint64, pos int) {
	
	if n < 2 {
		result = 1
	} else {
		v1, _ := fibo(n - 1)
		v2, _ := fibo(n - 2)
		result = v1 + v2
	}

	pos = n
	// print(result)
	return;
}


var fibs [41]uint64

func fibonacci(n int) (result uint64) {

	if fibs[n] != 0 {
		result = fibs[n]
		return
	}

	if n < 2  {
		result = 1
		// result = n
	} else {
		result = fibonacci(n - 1) + fibonacci(n - 2)
	}

	fibs[n] = result

	return
}