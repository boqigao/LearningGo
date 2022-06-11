package main

import "fmt"

func main() {
	var char byte
	fmt.Scanf("%c", &char)
	if char >= 97 && char <= 101 {
		fmt.Printf("%c", char-32)
	} else {
		fmt.Println("other")
	}

	switch char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("other")
	}
}
