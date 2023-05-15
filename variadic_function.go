//  Variadic function canbe called with any no of arguments i.e it can accept varying no of arguments

package main
import "fmt"

// Suppose you want to make a function that add the numbers you provide, but the no of interger is not fixed 
// each time its vary so you can use variadic function

func sum(nums ...int){
	fmt.Println("Given numbers are: ", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Sum of given no", total);
}
func main(){
	sum(1,3)
	sum(7,8,9,10,-45)
	sum(0)
	// If we have to pass multiple arguments why we are passing one each 
	// we can store arguments in a slice/array or the pass the slice/array in function

	numbers := []int {1,2,3,4,5}		// This is a slice not an array, array need size at the time of declarations
	sum(numbers...)
}