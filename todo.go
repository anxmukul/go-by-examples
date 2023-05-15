package main 
import "fmt"

type task struct {
	id int
	todo string
	priority string
}
var i = 1;
func newTask(todo string, priority string) *task {
	var t task
	t.id = i
	i++
	t.todo = todo
	t.priority = priority

	return &t;
}

// func showTask(n task){
// 	fmt.Println("Id:", n.task)
// 	fmt.Println("Todo: ", n.todo)
// 	fmt.Println("Priority: ", n.priority);
// }
func main(){
	var task1 task;
	task1 = newTask("Call Abhisek", "Low")
	fmt.Println(task1)
	// newTask("Comple Go by examples", "High")
	// newTask("Learn more about structure in go", "Medium")
}