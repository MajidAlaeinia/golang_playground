package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	empJsonData := `{"Name":"George Smith","Age":30,"Address":"New York, USA"}`
	fmt.Printf("%T\n", empJsonData)
	empBytes := []byte(empJsonData)
	var emp Response
	err := json.Unmarshal(empBytes, &emp)
	if err != nil {
		return
	}
	fmt.Println(emp.Name)
	fmt.Println(emp.Age)
	fmt.Println(emp.Address)
}
