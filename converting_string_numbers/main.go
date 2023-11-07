package main

import (
	"fmt"
	"strconv"
)
func main(){
	var s1 float64 = 33.4
	// fmt.Sprint converts any values to string 
	var s2 = fmt.Sprint(s1)
	fmt.Printf("%T\n", s2); //string

	var s3 int = 100
	s2 = string(s3) // converts to string
	fmt.Println(s2) // d (ascii char of 100)

	s2 = fmt.Sprint(s3)
	fmt.Println(s2) // 100

	s2 = fmt.Sprint(s1)
	var s4, err = strconv.ParseFloat(s2, 32) // converting string to float32,err for error
	_ = err
	fmt.Printf("%T, %v , %v\n", s4, s4,err) //float64, 33.400001525878906 , <nil>

	// string to int(Atoi) and int to string(Itoa) 
	val1, err := strconv.Atoi("-50")  // show error for floating numbers
	val2 := strconv.Itoa(50)
	fmt.Printf("%v , %v, %v\n" , val1, err, val2)

}