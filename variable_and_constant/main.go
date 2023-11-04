package main
import "fmt"

func main(){
	var age int = 30
	fmt.Println("Age: ", age)

	var name string = "Mahedi Hasan"
	fmt.Println("Name: ", name)

	_ = name // used for muting 'not used error'

	s := "Learning Golang"
	fmt.Println(s)

	car, cost := "audi" , 5000
	car, year := "BMW" , 2024 // := can be used when at lest one new variable

	fmt.Println(car,"price:",cost ,"buying year:",year)

	fmt.Print(car,cost, "\n")	// no space between two perametter

	open , file := false, "a.txt"
	_ , _ = open, file // muted not used error

	var (
		salary float64 = 20
		firstName string = "M Hasan"
		remote bool = false 
	)
	fmt.Println(salary, firstName, remote)

	var a, b, c int
	fmt.Println(a,b,c)

	// multiple assignment
	a, b, c = 10, 15, 16

	a, b = b, a  // swapping variables
	fmt.Println(a,b); // 15 10
	 
	sum := 5 + 2.3 

	fmt.Println(sum) // prints 7.3

	

	var x = 5
	var y = 6.4

	// x = int("10") // error
	// var sum = x + y // cant assign float64 to int type
	var sum1 = x + int(y)
	fmt.Println(sum1)

	fmt.Printf(" \"Hello world\"  \n") // print "" inside hello world

	const pi float64 = 3.1416 // need to assign value while declaring
	const (
		dis  = 500 
		max = 100
	)
	// cant change const 
	// dis = 150 // error

	// cont initialized with variables and function which returns variables
	// const val = math.Pow(2,4) // error
	// const tc = sum // error

	const lenth = len("string") // its okey 

	const conA, conB int = 100, 200
	const conC  = conA*conB // okey

	const conAA, conBB int = 100, 200
	//const conCC float64 = conAA * conBB 

	
	




	// zero values 
	// int 0
	// string "" ( empty string)
	// pointer nil
	// bool false

	// comments :
	// this is single line comment 
	/* this is
		i m in love with golang
		multiline comment */
	

	// fmt package 
	// %d for int 
	// %f for float64
	// %T for type
	// %t for bool 
	// %b for base 2 (binary)
	// %08b for 8 bit represetation in binary
	// %x in base 16

	



}
