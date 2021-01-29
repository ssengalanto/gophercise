package main

import "fmt"

func main() {
	run()
}

/*
delete: deletes the element with the specified key (m[key]) from the map. If m is nil or there is no such element, delete is a no-op.
func delete(m map[Type]Type1, key Type)
*/

func run() {

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	delete(m, "b")

	fmt.Println(m)
}
