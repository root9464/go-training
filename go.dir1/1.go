//рекурсия

package main

import (
	"fmt"
)

func pow(x, n int) []int {
	var result int = 1

	for i := 0; i < n; i++ {
		result *= x
	}
	return []int{result}
}

func main() {
	fmt.Println(pow(2, 3))
}
