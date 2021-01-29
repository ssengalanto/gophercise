// package declaration
package main

// import declaration
import "fmt"

/*
 Different Types in Go
	type bool
	type byte
	type complex128
	type complex64
	type error
	type float32
	type float64
	type int
	type int16
	type int32
	type int64
	type int8
	type rune
	type string
	type uint
	type uint16
	type uint32
	type uint64
	type uint8
	type uintptr
*/

// type declaration
type human interface {
	speak()
}

// type alias
type someType = string

// type of person struct
type person struct {
	name string
	age  int
}

/* package scope variables */
// const declaration (must be of type string bool and numbers(int, float etc.))
const (
	cs = "string"
	ci = 1
	cb = true
)

// variable declaration
var s string         // "" : zero value of string | equivalent to var s = ""
var i int            // 0 : zero value of numbers (int, float64 etc.)
var b bool           // false : zero value of boolean
var si []int         // nil: zero value of slice
var m map[string]int // nil : zero value of map
var c chan string    // nil : zero value of channel
var pt *string       // nil : zero value of pointers
var p person         // { name: "" , age: 0 } : zero value of each field of struct (in this case the struct has two field of type string and int which defaults to "" and 0)
var arr [3]int       // [0, 0, 0] : zero value of array (in this case type of array is int so it defaults to 0)

func main() {
	printAll()
}

// func declaration
func printAll() {
	/* local/block scope variables */
	// short-hand variable declaration can only be used locally while var declaration can be used locally or package level scope
	x, y := 1, 2 // multiple short hand declaration
	z := 3       // single short hand declaration
	fmt.Println(x, y, z)
	fmt.Println(s, i, b, si, m, c, p, arr, pt)
}
