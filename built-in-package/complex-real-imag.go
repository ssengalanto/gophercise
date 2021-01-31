package main

import (
	"fmt"
)

/*
	complex: constructs a complex value from two floating-point values.
	The real and imaginary parts must be of the same size, either float32 or float64 (or assignable to them),
	and the return value will be the corresponding complex type (complex64 for float32, cmplx128 for float64).

	func complex(r, i FloatType) ComplexType
*/

/*
	real: returns the real part of the complex number c.
	The return value will be floating point type corresponding to the type of c.

	func real(c ComplexType) FloatType
*/

/*
	imag: returns the imaginary part of the complex number c.
	The return value will be floating point type corresponding to the type of c.

	func imag(c ComplexType) FloatType
*/

func main() {
	cmplx1 := complex(10, 4)
	cmplx2 := 1 + 4i
	fmt.Println("complex num 1", cmplx1)
	fmt.Println("complex num 2:", cmplx2)
	fmt.Println("real num 1:", real(cmplx1))
	fmt.Println("imaginary num 1:", imag(cmplx1))
	fmt.Println("real num 1:", real(cmplx2))
	fmt.Println("imaginary num 2:", imag(cmplx2))
	fmt.Println("addition ", cmplx1+cmplx2)
	fmt.Println("substraction:", cmplx1-cmplx2)
	fmt.Println("division ", cmplx1/cmplx2)
	fmt.Println("multiplication:", cmplx1*cmplx2)
}
