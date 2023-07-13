package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432	
	user     = "postgres"
	password = "blah"
	dbname   = "employee"
)

func main() {
	fmt.Println("This is a rest-api......")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/insert", insertRow)
	http.HandleFunc("/getdata", readFromTable)
	http.HandleFunc("/updaterow", updateRow)
	http.HandleFunc("/deleterow", deleteRow)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error

	db, err = sql.Open("postgres", psqlconn)
	fmt.Println(db)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to db.....")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		panic(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to homepage!!")
}

type emp struct {
	Emp_id    int
	Emp_name  string
	Emp_email string
	Emp_dep   string
}

func readFromTable(w http.ResponseWriter, r *http.Request) {
	var returnedRow emp
	selectQuery := `select * from employeedetail where emp_id = 3`
	err := db.QueryRow(selectQuery).Scan(&returnedRow.Emp_id, &returnedRow.Emp_name, &returnedRow.Emp_email, &returnedRow.Emp_dep)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, "Employee table data fetched")
	fmt.Println(returnedRow)
}

func insertRow(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("in handler")

	insertQuery := `insert into employeedetail (emp_name, emp_email, emp_dep) values ('Ravi', 'ravi@email.com', 'eng');`

	// fmt.Println("db obj:", db);
	// fmt.Println("in handler 2")
	_, err := db.Query(insertQuery)
	if err != nil {
		// fmt.Println(err);
		panic(err)
	}
	fmt.Fprint(w, "Row Inserted with emp_no")
}

func updateRow(w http.ResponseWriter, r *http.Request) {
	updateQuery := `update employeedetail set emp_email = 'priyans@rediffmail.com' where emp_id = 3;`
	_, err := db.Query(updateQuery)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, "Row updated")
}

func deleteRow(w http.ResponseWriter, r *http.Request) {
	deleteQuery := `delete from employeedetail where emp_name = 'Ravi';`
	_, err := db.Query(deleteQuery)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, "Row Deleted")
}
