package main

import "fmt"

type Rectangular struct {
	len    int
	width  int
	height int
}

func (rectangular *Rectangular) volume() int {
	return rectangular.len * rectangular.width * rectangular.height
}

func main() {
	var rec = Rectangular{
		len:    3,
		width:  4,
		height: 5,
	}

	fmt.Println(rec.volume())
}
