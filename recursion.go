package main
import "fmt"

func factorial(n int) int{
	if n == 0 {
		return 1
	}
	return n * factorial(n-1);
}

func fibbonaci(n int) int {
	if n < 2 {
		return n
	}
	return fibbonaci(n-1) + fibbonaci(n-2)
}

func main(){
	res := factorial(5);
	fmt.Println("Factorial of 5 is:", res)

	fib := fibbonaci(7)
	fmt.Println("Fibbonacci of 7 is ", fib)
}