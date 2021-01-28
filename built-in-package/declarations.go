// package declaration
package main

// import declaration
import "fmt"

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
	x, y := 1, 0 // multiple short hand variable declaration
	fmt.Println(x, y)
	fmt.Println(s, i, b, si, m, c, p, arr, pt)
}
