package main

import (
	"log"
)

func main() {

	items := []string{"A", "1", "B", "2", "C", "3"}

	// Missing Example
	found := IsStringInSlice(items, "A")
	log.Println(found, "\n")
}

// IsStringInSlice takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func IsStringInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
