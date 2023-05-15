package main
import "fmt"

func main(){
	if 19 % 2 == 1{
		fmt.Println("19 is an odd no")
	} else {
		fmt.Println("19 is an even no")
	}

	if 16 % 3 == 0 {
		fmt.Println("16 is divisible by 3")
	} else{
		fmt.Println("16 is not divisible by 3")
	}
	
	i := 13
	if i < 0 {
		fmt.Println(i, "is a negative no")
	}else if i >= 10 {
		fmt.Println(i, "is a multiple digit no")
	}else {
		fmt.Println(i, "is a single digit no")
	}
}