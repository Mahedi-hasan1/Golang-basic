package main

import "fmt"

func main() {
	price , inStock :=  19, false;


	// start else from the closing brackets line

	if price >= 20 && inStock {
		fmt.Println("Available")
	}else if inStock == false {
		fmt.Println("Unavailable");
	}else {
		fmt.Println("Unknown")
	}
	
	

}
 