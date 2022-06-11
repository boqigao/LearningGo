package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("input name")
	fmt.Scanln(&name)

	fmt.Println("input age")
	fmt.Scanln(&age)

	fmt.Println("input sal")
	fmt.Scanln(&sal)

	fmt.Println("input isPass")
	fmt.Scanln(&isPass)

	fmt.Printf("name is %v, age is %v, sal is %v, isPass is %v", name, age, sal, isPass)

	fmt.Println("input name, age, sal, isPass")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
}
