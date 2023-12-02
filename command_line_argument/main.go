package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	//go run main.go Mahedi Hasan  23
	fmt.Println("os.Args", os.Args) 	// os.Args [/tmp/go-build781759733/b001/exe/main Mahedi Hasan]
	fmt.Println("path: ", os.Args[0])	//path:  /tmp/go-build781759733/b001/exe/main
	fmt.Println("path: ", os.Args[2])	//path:  Hasan
	fmt.Println("length ", len(os.Args)) // length 3

	val , err := strconv.Atoi(os.Args[ len(os.Args)-1]);
	if(err != nil){
		fmt.Printf("Last argument must be a number\n");
	}else {
		fmt.Printf("Double of %d = %d\n", val, val*2);    //  Double of 23 = 46
	}

}