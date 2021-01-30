package main

import (
	"fmt"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {
	defer recovery()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go divide(a, b, done) // since divide is ran on a separate go routine it is not aware of the defer recovery that's why it will panic
	// divide(a, b, done) without go keyword divide will run in the main go routine which also triggers the recover
	<-done
}

func divide(a int, b int, done chan bool) {
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	done <- true
}

func main() {
	sum(5, 0) // integer is divided by zero which cause implicit panic
	fmt.Println("normally returned from main")
}
