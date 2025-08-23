package main

//
import "fmt"

// defer keeps all the deferred function calls in a stack and executes them in LIFO order
// defer uses link-list data structure to keep the function calls but uses stack behavior to execute them
func example1() {
	i := 0
	fmt.Println("first ", i)
	i++
	defer fmt.Println("second ", i)
	i++
	defer fmt.Println("third ", i)
	i++
	fmt.Println("fourth ", i)

}

func calculate() (result int) {
	fmt.Println("first ", result)

	show := func() {
		result += 10
		fmt.Println("defer ", result)
	}
	defer show()
	result += 5
	fmt.Println("second ", result)
	return // result value will be updated while deferred function is called
}

func calculate2() int {
	result := 0
	fmt.Println("first ", result)

	show := func() {
		result += 10
		fmt.Println("defer ", result)
	}
	defer show()
	result += 5
	fmt.Println("second ", result)
	return result // result value evoluated before the deferred function is called
}

func main() {
	example1()
	// Output:
	// first  0
	// fourth  3
	// third  2
	// second  1

	fmt.Println(calculate())
	// Output:
	// first  0
	// second  5
	// defer  15
	// 15
	fmt.Println(calculate2())
	// Output:
	// first  0
	// second  5
	// defer  15
	// 5

}
