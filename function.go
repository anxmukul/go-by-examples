package main
import "fmt"


func addition(a int, b int) int {
	return a+b
}
// func swap(m int, n int) {
// done in multiple-return-value.go file
// }
// func addmorethanone(a, b, c, d, int) int {

// }
func concatenate(fname string, lname string) string {
	return (fname + lname)
}
func main(){
	res := addition(3,4);
	fmt.Println("Sum of 3, 4 is: ", res)

	fmt.Println("Sum of -4, -5 is:", addition(-4, -5));
	
	var firstName string = "Mukul"
	var secondName string = "Kumar"

	fullName := concatenate(firstName, secondName);
	fmt.Println("Full name is:", fullName);

}