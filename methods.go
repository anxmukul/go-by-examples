package main
import "fmt"


type rect struct {
	length int
	width int
}

func (r *rect) area() int {
	return r.length * r.width
}
func (r rect) peri() int {
	return 2*(r.length + r.width)
}
func main(){
	r1 := rect{length:5, width: 7}
	r2 := &rect{length:7, width: 3}
	fmt.Printf("Rectangle1:\n%+v\n", r1)
	fmt.Printf("Rectangle2:\n Length: %v\n Width:%v\n", r2.length, r2.width)
	fmt.Println("Area r1: ", r1.area())
	fmt.Println("Perimeter r1: ", r1.peri())
	fmt.Println("Area r2: ", r2.area())
	fmt.Println("Perimeter r2: ", r2.peri())
	

	//   so how does go handles values and pointer for method call? Area wants pointer as a parameter for r1, but 
	// we are providing copy of the rect type struct ??
	
	// Go automatically handles conversion between values and pointers for method calls. 
	// You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.

}