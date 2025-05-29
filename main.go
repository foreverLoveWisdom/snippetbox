package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", home)
	// mux.HandleFunc("/snippet/view", snippetView)
	// mux.HandleFunc("/snippet/create", snippetCreate)
	//
	// log.Print("starting server on :4000 ")
	// err := http.ListenAndServe(":4000", mux)
	// log.Fatal(err)

	var dilbert Employee
	dilbert.Salary -= 500
	dilbert.Name = "Tien Do"
	dilbert.Position = "Software engineer"
	log.Println("Dilbert salary: ", dilbert.Salary)
	position := dilbert.Position
	log.Println("Dilbert position pointer: ", position)
	position = "CEO"

	str := "Hello"
	runeCount := len([]rune(str)) // Convert string to runes and get the length
	fmt.Printf("The word '%s' has %d runes.\n", str, runeCount)
	// log.Println("Dilbert position: ", *position)
	// var employeeOfTheMonth *Employee = &dilbert
	//
	// log.Println("Award to: ", employeeOfTheMonth.Name)
	//

	// var a [5]int

	// var q [3]int = [3]int{1, 2, 3}
	// var r [3]int = [3]int{1, 2}
	// log.Println(r[2])

	// Print the indices and elements
	// for i, v := range a {
	// 	log.Printf("%d %d\n", i, v)
	// }

	// Print the indices only
	// for i := range a {
	// 	log.Printf("%d\n", i)
	// }
}

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}
