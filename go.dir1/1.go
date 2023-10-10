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

//срезы
// package main

// import "fmt"

// func main() {
// 	var nested = [4][]int{
// 		{1, 2, 3, 4, 5},
// 		{6, 7, 8, 9, 10},
// 		{11, 12, 13, 13, 15},
// 	}
// 	fmt.Println(nested)
// }

// рандомный срез рандомного массива
// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	array := []int32{}

// 	for i := 0; i <= 10; i++ {
// 		array = append(array, rand.Int31()*2)
// 	}
// 	subSclice := array[2:]
// 	fmt.Println(array)
// 	fmt.Println(subSclice)
// }

// package main

// import "fmt"

// func main() {
// 	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	n1 := n[:6]
// 	n2 := n[3:8]
// 	n3 := n[4:10]

// 	fmt.Println(n1, len(n1), cap(n1)) // => [1 2 3 4 5 6]  6 10
// 	fmt.Println(n2, len(n2), cap(n2)) // => [4 5 6 7 8]  5 7
// 	fmt.Println(n3, len(n3), cap(n3)) // => [5 6 7 8 9 10] 6 6

// 	// change n1 at index 4 to 15
// 	n1[4] = 15

// 	fmt.Println(n)  // => [1 2 3 4 15 6 7 8 9 10]
// 	fmt.Println(n1) // => [1 2 3 4 15 6]
// 	fmt.Println(n2) // => [4 15 6 7 8]
// 	fmt.Println(n3) // => [15 6 7 8 9 10]

// }
