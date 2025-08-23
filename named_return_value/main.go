package main

import "fmt"

func sum(a, b int) (result int) {
	result = a + b
	return // named return value, result will be returned
}

func main() {
	res := sum(1, 3)
	fmt.Println("Sum:", res) // Sum: 4
}
