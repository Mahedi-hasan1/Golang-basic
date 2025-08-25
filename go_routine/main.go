package main

import (
	"fmt"
	"time"
)

func hello(cnt int) {
	fmt.Println("Hello, hasan", cnt)
}

func main() {
	fmt.Println("Hello, hasan", 0)
	go hello(1) // Starting a goroutine to run hello function concurrently
	// The main function will continue executing without waiting for the goroutine to finish
	go hello(2)
	go hello(3)
	go hello(4)
	fmt.Println("Hello, hasan", 0)
	time.Sleep(5 * time.Second) // Sleep for a while to allow goroutines to complete
}
// While a go program starts to run , 
// at first go runtime create a main goroutine in heap 
// where main&init starts executes and 
// after that it creates more goroutine in heap if needed. 
// In a goroutine every functions inside this executes in stack framework 
// and it can access data in current scope , and global scope(code&data segment). 
// Cant access data from another running goroutine.