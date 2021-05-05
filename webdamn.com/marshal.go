package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string
	Age     int
	Address string
	//Ids     [2]string //TODO
}

func main() {
	//emp := Employee{Name: "George Smith", Age: 30, Address: "New York, USA", Ids: ["aaa", "bbb"]} //TODO
	emp := Employee{Name: "George Smith", Age: 30, Address: "New York, USA"}
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(empData))
}
