package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var chars [26]byte

	chars[0] = 'A'
	for i := 1; i < 26; i++ {
		chars[i] = chars[i-1] + 1
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c\n", chars[i])
	}

	rand.Seed(time.Now().UnixNano())
	rand.Int()
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = rand.Int()
		time.Sleep(10 * time.Nanosecond)
	}
	fmt.Println(arr)

	for i := 0; i < len(arr)/2; i++ {
		tmp := arr[len(arr)-1-i]
		arr[len(arr)-1-i] = arr[i]
		arr[i] = tmp
	}

	fmt.Println(arr)
}
