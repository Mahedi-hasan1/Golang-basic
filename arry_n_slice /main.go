package main

import (
	"fmt"
)

func main() {
	num := [3]int{3, 2, 1}
	fmt.Println(num[2])
	fmt.Println(num)

	for v := range num {
		fmt.Println(num[v])
	}

	names := [...]string{
		5: "Hasan",
		"Mahi",
		//"Rahim", // will be error as index 7 already asssigned
		2: "Imran",
		"Nirob",
		7: "Karim",
	}
	fmt.Println(names) //   Imran Nirob Hasan Mahi Karim

	/// slices , is dynamic type array
	cities := []string{
		"dka", "ctg",
	} // slice literal(no size needed , declare like array)
	cities = append(cities, "khl", "rng")
	fmt.Println(cities) // dka, ctg , khl , rng

	cities2 := []string{
		"syt", "cox",
	}
	cities = append(cities, cities2...)
	fmt.Println(cities) // dka, ctg ,khl , rng, syt , cox

	copy(cities, cities2)

	fmt.Println(cities) //syt, cox, khl rng, syl cox

	cities3 := cities[1:3]

	fmt.Println(cities3) // cox,khl (1 to 3 excluding)
	cities3 = cities[1:] // same as [1 : len(cities)]
	cities3 = cities[:]
	fmt.Println(cities3) // same as [0:len(cities)]
	cities3 = cities[:2] // same as [0, 2]
	fmt.Println(cities3) // syl cox

	//slices internal  backing array
	//slices takes lower memory than array
	s1 := []int{10, 20, 30, 40, 50}
	s2, s3 := s1[0:2], s1[1:3]
	s2[1] = 200
	fmt.Println(s1, s2, s3) // 10 200 30 40 50 , 10 200, 200 30
	// s1, s2,s3 all slices changes as slices uses same internal baking array

	cars := []string{
		"audi", "ford", "honda", "range rover",
	}
	new_cars := []string{}
	//new_cars = cars[0:2]
	new_cars = append(new_cars, cars[0:2]...) // append wont use same backing array

	cars[1] = "nissan"
	fmt.Println(cars, new_cars)

	s5 := make([]int, 3, 3) //declare slice with length 3, capacity 3
	s5 = []int{1, 2, 3, 4, 5}
	newS := s5[2:5]

	fmt.Println(newS, s5)             // 3 4 5 , 1 2 3 4 5
	fmt.Println(len(newS), cap(newS)) // 3,3  //capacity is lenth of backing array
	// from straring of new  slice

	newS = s5[1:2]                                    //starting from index 1 to before 2
	fmt.Println(newS)                                 // 2 ,   1 2 3 4 5
	fmt.Println("len ", len(newS), "cap ", cap(newS)) //1,4

	//append increses capacity with power of two
	newS = append(newS, 3, 4, 5)

	fmt.Println(len(newS), cap(newS)) // cap 4

	newS = append(newS, 6)
	fmt.Println(len(newS), cap(newS)) // cap 8


	var s6 []int //empty slice
	fmt.Println(s6, len(s6), cap(s6)) // [] 0 0
}
