package main

import ("fmt"
		"net/http"
		"strconv"
		"io/ioutil"
		"encoding/json"
		
)

type Person struct {
	Name string	
	UserId string
}

func main(){
	// http.HandleFunc("/", getHello)
	http.HandleFunc("/user", userDetail)
	http.ListenAndServe(":3000", nil)

}

func getHello(w http.ResponseWriter, r *http.Request){
	// fmt.Fprint(w, "Hello World!!");
	num := r.URL.Query().Get("no");
	number, err := strconv.Atoi(num)
	if err != nil {
		panic(err);
	}else{
	number = number*number;
	}

	fmt.Fprint(w, "Square of no is ", number);
}

func userDetail(w http.ResponseWriter, r *http.Request){
	var user Person
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		panic(err)
	}else{
		json.Unmarshal(body, &user)
		fmt.Fprint(w, user)

	}
	
}