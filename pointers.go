package main 

import "fmt"

func passbyval(a int) {
	a = 1;
}


func passbyaddr(a *int){
	fmt.Println("value of a: ", a)
	fmt.Println("value at &a: ", &a)
	*a = 1;
}

func main(){
	n := 10;
	fmt.Println("value of intial n:", n);
	passbyval(n)
	fmt.Println("value of n after passing by val:", n);
	passbyaddr(&n)
	fmt.Println("value of n after passing by addr:", n);


}