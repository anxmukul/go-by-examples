package main
import "fmt"

func binary_search(array []int, k int) int {
	n := len(array);
	var l, r int = 0, n-1;
	// fmt.Println(n, l, r);
	for l <= r {
		mid := l+(r-1)/2
		if(array[mid] == k){
			return mid;
		}

		if(array[mid] < k){
			l = mid+1;
		} else{
			r = mid+1;
		}
	}
	return -1;
}

func main(){
	var	arr =  [] int {3,  6, 19, 35, 75, 87, 98}
	fmt.Println("Given array:- ")
	for i := 0; i<7; i++ {
		fmt.Println(arr[i])
	}
	res := binary_search(arr,3)
	fmt.Println("Found at index: ", res)
}