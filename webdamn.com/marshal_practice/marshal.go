package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name string
	Age  int
}

func main() {
	emp := Employee{Name: "Majid Alaeinia", Age: 31}
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(empData))
}
