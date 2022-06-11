package main

import "fmt"

func main() {
	//var max int = 100
	var counter int = 0
	var sum int = 0

	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			counter++
			sum += i
		}
	}
	fmt.Printf("sum is %d, counter is %d", sum, counter)
}
