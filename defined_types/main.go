package main
import ("fmt")

func main(){
	type age int
	type oldAge  age
	type veryOldAge oldAge  // int is underlying types

	var v1 veryOldAge = 100
	var v2 float32 = 10.1
	v1 = veryOldAge(v2) // precision point will be removed

	fmt.Printf("%T, %v\n", v1, v1) // main.veryOldAge , 10

	// alice declaration

	var a uint8 = 8
	var b byte  
	b = a;    // uint8 and byte same , thats the buildin alice declared
	_ = b;

	// try a user define
	type second = uint
	var mn second = 3600
	fmt.Printf("%v\n", mn);






}