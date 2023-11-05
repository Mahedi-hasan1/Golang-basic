package  main
import (
	"fmt"
)

func main(){
	// int8 , int16, int32(rune), int64 avaiable also uint8(byte), uint16 ....
	var i1 int8 = 127 // max of int8 127
	var i2 uint8 = 255 // max of uint8 255
	fmt.Printf("%T  %T \n", i1, i2)  //int8 uint8

	//float32, float64 
	var i3 float32 = -3.3
	fmt.Printf("%T\n" , i3)

	// rune and byte
	var i4 rune = 100
	var i5 byte = 100
	fmt.Printf("%T %T\n" ,i4, i5) //int32 uint8

	// bool 
	var b bool = true
	fmt.Printf("%T\n", b) // bool
	
	//string
	var s string = "Mr Hasan"
	fmt.Printf("%T\n",s) //string

	//array
	var num  = [4] int {1,2,-3,4}
	fmt.Printf("%T\n", num)  // [4]int

	//slice
	var cities = [] string {"London" , "Tokyo"}
	fmt.Printf("%T\n", cities) // []string

	//map
	blances := map[string]float64{
		"taka" : 12.4,
		"Dollar" : 1.1,	

	}
	_ = blances

	//struct
	type person struct{
		name string
		age int
	}
	var me person
	me.age = 25
	me.name = "Hasan"
	fmt.Println(me, me.name) // {Hasan 25} Hasan

	//pointer 
	var x int = 5
	ptr := &x
	fmt.Println(x, "located in", ptr) //5 located in 0xc000012150
	
	//function
	fmt.Printf("%T", f); //func()

		
	

}
func f(){

}