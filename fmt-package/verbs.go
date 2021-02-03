package main

import "fmt"

/*
The default format for %v is:
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
*/

func main() {
	p := struct {
		name string
		age  int
	}{
		"Hunter",
		4,
	}

	/* general verbs */
	fmt.Println("---GENERAL-VERBS---")
	// %v: the value in a default format, when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%%v: %v \n", p)
	// %#v: a Go-syntax representation of the value
	fmt.Printf("%%#v: %#v \n", p)
	// %T: a Go-syntax representation of the type of the value
	fmt.Printf("%%T: %T \n", p)
	// %%: a literal percent sign; consumes no value
	fmt.Printf("%% \n")

	/* boolean */
	fmt.Println("---BOOLEAN---")
	var b bool
	// %t: the word true or false
	fmt.Printf("%%t: %t \n", b)

	/* integer */
	fmt.Println("---INTERGER---")
	i := 10
	// %b: base2 / binary
	fmt.Printf("%%b: %b \n", i)
	// %c: the character represented by the corresponding Unicode code point
	fmt.Printf("%%c: %c \n", i)
	// %d: base10
	fmt.Printf("%%d: %d \n", i)
	// %o: base8
	fmt.Printf("%%o: %o \n", i)
	// %O: base10 with 0o prefix
	fmt.Printf("%%O: %O \n", i)
	// %q: a single-quoted character literal safely escaped with Go syntax.
	fmt.Printf("%%q: %q \n", i)
	// %x: base 16, with lower-case letters for a-f
	fmt.Printf("%%x: %x \n", i)
	// %X: base 16, with upper-case letters for A-F
	fmt.Printf("%%X: %X \n", i)
	// %U: Unicode format: U+1234; same as "U+%04X"
	fmt.Printf("%%U: %U \n", i)

	/* floating-point values */
	fmt.Println("---FLOATING-POINT---")
	f := 10.1234
	// %b: decimalless scientific notation with exponent a power of two,
	fmt.Printf("%%b: %b \n", f)
	// %e: scientific notation, e.g. -1.234456e+78
	fmt.Printf("%%e: %e \n", f)
	// %E: scientific notation, e.g. -1.234456e+78
	fmt.Printf("%%E: %E \n", f)
	// %f: decimal point but no exponent, e.g. 123.456
	fmt.Printf("%%f: %f \n", f)
	// %[width].[precision]f: width and precision
	fmt.Printf("%%2.2f: %2.2f \n", f)
	// %F: synonym for %f
	fmt.Printf("%%F: %F \n", f)
	// %g: for large exponents, %f otherwise. Precision is discussed below.
	fmt.Printf("%%g: %g \n", f)
	// %G: %E for large exponents, %F otherwise
	fmt.Printf("%%G: %g \n", f)
	// %x: hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
	fmt.Printf("%%x: %x \n", f)

	/* string */
	fmt.Println("---STRING---")
	s := "hello"
	// %x: the uninterpreted bytes of the string or slice
	fmt.Printf("%%s: %s \n", s)
	// %q: the uninterpreted bytes of the string or slice
	fmt.Printf("%%q: %q \n", s)
	// %x: base 16, lower-case, two characters per byte
	fmt.Printf("%%x: %x \n", s)
	// %X: base 16, upper-case, two characters per byte
	fmt.Printf("%%X: %X \n", s)

	/* slices */
	fmt.Println("---SLICE---")
	si := []int{1, 2, 3}
	// %p: address of 0th element in base 16 notation, with leading 0x
	fmt.Printf("%%p: %p \n", si)

	/* pointer */
	fmt.Println("---POINTER---")
	ptr := &p
	// %p: base 16 notation, with leading 0x
	fmt.Printf("%%p: %p \n", ptr)
}
