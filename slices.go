// Slice is another data structure like an array but with more powerful interface.
// we can append elements in slice

package main
import "fmt"

func main(){
	s := make([]string, 3)   // can be declared as var s = make([]string, 3);
	fmt.Println("slice s : ",s);

	// adding, and getting elemets from slices is same as array
	s[0] = "FirstElement"
	s[1] = "SecondElement"
	s[2] = "ThirdElement"
	fmt.Println("Slice s is: ", s)
	fmt.Println("2nd element of slice is: ", s[1])

	// Appending in slices

	fmt.Println("Len of slice s is: ", len(s));
	
	s = append(s, "4th"); // we need to accept the return value from append method because it may return new slice

	s = append(s, "5th", "6th");
	fmt.Println("After adding some element", s);

	// Copying one slices to another

	c := make([] string, len(s));
	copy(c, s);						// There is no sence to return any value from copy method since we already declared c as our new slice
	fmt.Println("Slice c become:", c)

	// Slice methods on slice data structure

	var l = s[2:5]
	fmt.Println("l:", l)

	t := s[2:]
	fmt.Println("t:", t)

	y := c[:6]
	fmt.Println("y:", y)
	
	// we can create 2D slices also. The length of inner slice can vary unlike array

}