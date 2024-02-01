package main
import(
	"fmt"
)

func main(){
	for i := 0 ; i <10 ; i++ {
		fmt.Print(i, " ")   // 0 1 2 3 4 5 6 7 8 9 
	}
	fmt.Println()

	// There is no for loops in go lang
	// But we can make while loop using for loop

	// while loop 
	j  := 10 
	for j>=0 {
		fmt.Print(j, " ");   //10 9 8 7 6 5 4 3 2 1 0
		j--;
	}
	fmt.Println()

	// continue and break
	for i := 0 ; i<20 ; i++ {
		if i%2==1 {
			continue;
		}else if (i>10){
			break;
		}
		fmt.Print(i, " ");  		//	0 2 4 6 8 10
		  
	}
	fmt.Println()

	//levels
	
	people := [5]string{"Mahedi", "Hasan", "Mahim", "Fahim", "Rahim"}

	friends := [2]string{"Mahedi", "Hasan"}

	Loop1: 
	for  i:=0 ; i<5; i++{
		for j:=0 ; j<2 ; j++{
			if(people[i] == friends[j]){
				fmt.Print(people[i], " is common\n");       // Mahedi is common
				break Loop1 
			}
		}
	}
	fmt.Println("outside of loops")  //outside of loops

}



