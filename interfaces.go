//  Lets say we have a companywho wants to calculate total expances based on their employee salary
// and the company have two type of emp. 1. Permanent, 2. Custom

package main
import "fmt"

type salaryCalcultor interface{
	calculateSalary() int
}

type Permanent struct {
	id int
	basicSalary int
	pf int
}

type Contract struct {
	id  int
	basicSalary int
}

func (p Permanent) calculateSalary() int {
	return p.basicSalary + p.pf
}

func (c Contract) calculateSalary() int {
	return c.basicSalary
}

func totalExpances(s []salaryCalcultor) int {
	expances := 0
	for _, emp := range s {
		expances += emp.calculateSalary()
	}

	// fmt.Println("Total Company Expances is: ", expances)
	return expances
}
func main(){
	// pemp1 := Permanent{1, 250000, 23000}
	// pemp2 := Permanent{2, 450000, 13000}
	// cemp1 := Contract{3, 65000}
	// cemp2 := Contract{4, 89000}
	// pemp3 := Permanent{5, 650000, 6500}
	// Now if we want to calculate the total expances we have to read all the emp salary individually and add them
	
	// totalExpances := pemp1.calculateSalary() +pemp2.calculateSalary() + cemp1.calculateSalary() + cemp2.calculateSalary();
	// fmt.Println("Total expances: ", totalExpances);
	
	// Although Go automatically handles conversion between values and pointers for method calls. It's still lengthy

	// Here comes the interface :- Now we will crete an interface (an interface is a set of method signatures.)
	// create a single calculateexpances method and pass slice of employee object to that method and interface will
	// Automatically calculate the total expances
	
	var emp salaryCalcultor
	emp1 := Permanent{1, 45000, 2300}
	emp = emp1;
	fmt.Println(emp.calculateSalary())
	// employees := [] salaryCalcultor {pemp1, pemp2, cemp1, cemp2, pemp3}
	// fmt.Println("Total Company Expances is: ", totalExpances(employees))

}