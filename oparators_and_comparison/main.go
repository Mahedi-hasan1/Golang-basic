package main
import (
	"fmt"
)

func main(){

	// oparators (same as c++)
	a, b := 6,5
	var r int = (a+b) / (a-b) *  (a%b)
	fmt.Println(r) // 11

	a += b // a = 11
	a /= b // a = 2
	a %= b // a = 2
	a *= b // a = 10
	a++ // a = 11
	a-- // a = 10
	fmt.Println(a)  // 10
	

	// comparison (same as c++)
	fmt.Println(a==b) // false
	fmt.Println(a!=b) //true
	fmt.Println(a>=b && a>b) //true
	fmt.Println(a>b || b>a) //true	

}