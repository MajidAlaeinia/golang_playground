package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//aa := strings.Join(strings.Fields(strings.TrimSpace("Hello    World")), " ")
	aa := RandomStringWithLength(10)
	fmt.Println(aa)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomStringWithLength(ln int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, ln)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
