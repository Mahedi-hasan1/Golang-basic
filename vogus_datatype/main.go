package main

import "fmt"

func main() {
	var x int8 = 127 // max value for int8 and its 8 bits
	fmt.Println(x)   // 127

	var y int = 100 // int is platform dependent, 32 bits on 32-bit systems and 64 bits on 64-bit systems
	fmt.Println(y)  // 100
}
