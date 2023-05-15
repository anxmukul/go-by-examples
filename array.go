package main
import "fmt"

func main() {
	var array [5] int
	fmt.Println(array)
	array[4] = 400
	fmt.Println(array)
	fmt.Println("Get a[4]", array[4]);

	fmt.Println("length of array", len(array))

	b := [5] int {2, 4, 6, 8, 10}
	fmt.Println("declared array b:", b)

	friends := [3] string {"Abhisek", "Manash"}
	fmt.Println("List of friends:", friends)

	// 2dimension array
	var twoD [2][3] int 
	for i :=0; i<2; i++{
		for j:=0; j<3; j++{
			twoD[i][j] = (i+j)*2;
		}
	}
	fmt.Println("Length of twoD array is:", len(twoD));
	fmt.Println("TwoD array :", twoD);
}