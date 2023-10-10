//рекурсия

// package main

// import (
// 	"fmt"
// )

// func pow(x, n int) []int {
// 	var result int = 1

// 	for i := 0; i < n; i++ {
// 		result *= x
// 	}
// 	return []int{result}
// }

// func main() {
// 	fmt.Println(pow(2, 3))
// }

// package main

// import "fmt"

// func fact(a int) int {
// 	if a == 0 {
// 		return 1
// 	} else {
// 		return a * fact(a-1)
// 	}
// }

// func main() {
// 	fmt.Println(fact(5))
// }

//_______задания с массивами_________

// заполнить массив рандомными числами
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sort"
// )

// func arrApendRandom() []int {
// 	array := []int{}
// 	t := 40
// 	for i := 0; i < t; i++ {
// 		array = append(array, int(rand.Int31())*t)
// 	}
// 	sort.Ints(array)
// 	return array
// }

// func main() {
// 	fmt.Println(arrApendRandom())
// }

//заполнить массив рандомными числами
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sort"
// 	"time"
// )

// func arr() []int {
// 	array := []int{}
// 	c := 40
// 	time.Sleep(2 * time.Second)
// 	for i := 0; i < c; i++ {
// 		array = append(array, int(rand.Int31()))
// 		sort.Ints(array[:])
// 	}
// 	return array[:]

// }

// func main() {
// 	fmt.Println(arr())

// }
