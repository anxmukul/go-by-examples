// We can return multiple return value from a function since this is an inbuilt feature in go

package main
import "fmt"

func swap(m, n int)(int, int){
	m = m*n
	n = m/n
	m = m/n
	return m,n

}
// We can also reurnt different type of variable form the same function.
func stringandnumber(a int, b int)(string, int){
	s := "Returned from stringandnumber function that"
	res := a+b;
	return s, res
}	

func ssmd(q, r int)(int, int, int, int){
	return q+r, q-r, q*r, q/r
}
func main(){
	a, b := swap(4,5);
	fmt.Println("a is: ",a, "b is: ", b)

	x := 6;
	y := 5;

	u,v := stringandnumber(x,y);
	fmt.Println(u, "sum of given two no is:", v);

	// if we want subset of the retured value we can do that like this

	_,_,_, divide := ssmd(24, 7);

	_,substration,_,divide := ssmd(8,2)
	fmt.Println("Divion of 24/7 gives: ", divide);
	fmt.Println("Substraction of 8 & 2 gives: ",substration, "Divion of 24/7 gives: ", divide);
}