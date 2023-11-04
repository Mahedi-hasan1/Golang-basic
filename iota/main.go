package main

import "fmt"

func main(){
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b,c) // 0 1 2

	const (
		aa = iota
		bb 
		cc
	)
	fmt.Println(aa , bb, cc ) // 0 1 2
	
	const(
		a1 = (iota)*2+1
		b1
		c1
	)
	fmt.Println(a1, b1, c1) // 1,3, 5

	const(
		a2 = -(iota+2)
		_
		b2 
		c2
	)
	fmt.Println(a2, b2, c2)  // -2, -4, -5
}