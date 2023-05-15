package main 
import "fmt"

type Todo struct {
	id int
	todo string
	priority string
}

func showtask(t *Todo){
	fmt.Printf("Task\n Id: %v\n Todo: %v\n Priority: %v\n", t.id, t.todo, t.priority)

}

func main (){
	task1 := &Todo{1, "Schedule a meeting", "High"}
	// fmt.Println(task1.todo) // we don't need to dereference like (*task1).todo , the given also works...
	
	task2 := &Todo{2, "Complete Go Tutorial", "Medium"}
	task3 := &Todo{3, "Call Abhisek", ""}
	// Create task 3 in a different way so that we donot need to pass 3rd variable as a empty string it will automatically set it to o
	task4 := &Todo{id:4, todo:"Take a break"}
	var task5 Todo = Todo{5, "Learn more about struct", "Low"}
	showtask(task1)
	showtask(task2)
	showtask(task3)
	showtask(task4)

	fmt.Printf("Task5\n Id: %v\n Todo: %v\n Priority: %v\n", task5.id, task5.todo, task5.priority)


}