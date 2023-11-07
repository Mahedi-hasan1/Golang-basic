package main

import "fmt"

func main() {
	price, inStock := 23, false

	if price > 20 && inStock {
		fmt.Println("Available")
	}

}
