package main
import (
	"fmt"
	"net/http"
)


func main(){
	fmt.Println("This is a rest-api......");
	http.HandleFunc("/", homePage);

	http.ListenAndServe(":3000", nil)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to homepage!!");
}