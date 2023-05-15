package main

import "fmt"

func main(){
	
	// Basic for loop with a sinle condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// Classic for loop with intial/condition/after 

	for j := 0; j<5; j++ {
		fmt.Println(j*j)
	}

	// Break statement in for loop
	for k:= 1; k<3; k++ {
		fmt.Println("Value of k: ", k)
		break
	}

	// Continue statement in for loop

	m := 1
	for m<=10 {
		if m%2 == 0 {
			m++
			continue
		}
		fmt.Println(m, " is a odd no.")
		m++
	}
}