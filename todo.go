package main 
import "fmt"

type Todo struct {
	id int
	todo string
	priority string
}
var i = 1;
func (tsk *Todo) createTodo(todo string, priority string) {
	tsk.id = i
	tsk.todo = todo
	tsk.priority = priority
	i++;
}

 func (tsk Todo) showTask(){
	// fmt.Println("Id:", tsk.id)
	// fmt.Println("Todo: ", tsk.todo)
	// fmt.Println("Priority: ", tsk.priority);
	fmt.Printf("Task\n Id: %v\n Todo: %v\n Priority: %v\n", tsk.id, tsk.todo, tsk.priority)

}
func main(){
	task1 := &Todo{}
	task1.createTodo( "Call Abhishek", "Low");
	task2 := &Todo{}
	task2.createTodo( "Complete Go by Example", "Medium");
	task3 := &Todo{}
	task3.createTodo( "Start Open Source", "High");
	task1.showTask()
	task2.showTask()
	task3.showTask()
	
	
}