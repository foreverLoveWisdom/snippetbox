package main

import (
	"log"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
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
	log.Printf("The word '%s' has %d runes.\n", str, runeCount)
}
