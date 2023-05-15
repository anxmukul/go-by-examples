package main
import (
	"fmt"
	"time"
)

func main(){
	i := 2
	switch(i){				// we cam write switch i {} also without parenthese
	case 1:
		fmt.Println(i, " is one")
	case 2:
		fmt.Println(i, "is two")
	case 3:
		fmt.Println(i, "is three")
	default:
		fmt.Println("Given no is neither 1, nor 2, nor 3")
	}

	// lets play with some functions
	switch time.Now().Weekday() {		// This is same as switch(time.Now().Weekday())
	case time.Saturday, time.Sunday :
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// var weekday string = time.Now().Weekday()   //It doesnot return string....
	// fmt.Println(weekday)
}