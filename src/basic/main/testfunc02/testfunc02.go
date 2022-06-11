package main

import "fmt"

// get the sum of 1 or more
func sum(n1 int, args ...int) int {
	sum := n1
	// travers args
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	res := sum(1, 2, 3, 4, 5)
	fmt.Println(res)

	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println(a, b)
}
