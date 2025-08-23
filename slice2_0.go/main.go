package main

import "fmt"

// variadic function , // it can take any number of arguments
func add(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	var x []int      //[], len=0, cap=0
	x = append(x, 1) //[1], len=1, cap=1
	x = append(x, 2) //[1, 2], len=2, cap=2
	x = append(x, 3) //[1, 2, 3], len=3, cap=4
	x = append(x, 4) //[1, 2, 3, 4], len=4, cap=4
	x = append(x, 5) //[1, 2, 3, 4, 5], len=5 cap=8
	fmt.Println(x)   // [1 2 3 4 5]
	//until 1024 it will double the capacity after reaching the limit
	// after that it will increase by 25% of the current capacity

	y := x           // y starting with same backing array as x
	x = append(x, 6) // [1 2 3 4 5 6], len=6, cap=8
	y = append(y, 7) // [1 2 3 4 5 7], len=7, cap=8
	x[0] = 10
	fmt.Println(x)      // [10 2 3 4 5 7]
	fmt.Println(y)      // [10 2 3 4 5 7]
	fmt.Println(x[0:8]) //[10 2 3 4 5 7 0 0]

	//a slice keeps 3 things
	// 1. pointer to the backing array
	// 2. length of the slice
	// 3. capacity of the slice

	//variadic function , // it can take any number of arguments
	fmt.Println(add(1, 2, 3))       // 6
	fmt.Println(add(1, 2, 3, 4, 5)) // 15
}
