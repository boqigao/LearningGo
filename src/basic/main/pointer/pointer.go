package pointer

import (
	"fmt"
)

func main() {
	var num int = 9
	var ptr *int
	ptr = &num
	*ptr = 10

	fmt.Println(num)

}
