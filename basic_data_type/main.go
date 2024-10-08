package main

import "fmt"

// Numeric types

// uint8       the set of all unsigned  8-bit integers (0 to 255) = 256 değer
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

// int8        the set of all signed  8-bit integers (-128 to 127) =256
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

// float32     the set of all IEEE-754 32-bit floating-point numbers
// float64     the set of all IEEE-754 64-bit floating-point numbers

// complex64   the set of all complex numbers with float32 real and imaginary parts
// complex128  the set of all complex numbers with float64 real and imaginary parts

// byte        alias for uint8
// rune        alias for int32

func main() {

	// var ( // bu şekilde var kullanabılırız
	// 	firstName, lastName string = "bartu", "kurnaz"

	// 	age       uint8 = 255
	// 	isMarried bool  = true
	// 	weight          = 89.9
	// )

	// var name, age, isMarried, weight, height = "Arin", 40, true, 110, 180

	// name, age, isMarried, weight, height := "Arin", 40, true, 110, 180

	var name bool
	fmt.Println(name) // false Zero Values

	studentName := "john"
	studentName = "joooohn"

	fmt.Println(studentName)
	// fmt.Println(age)
	// fmt.Println(isMarried)
	// fmt.Println(weight)
	// fmt.Println(height)

	// fmt.Printf("Type of : %T\n", name)
	// fmt.Printf("Type of : %T\n", age)
	// fmt.Printf("Type of : %T\n", isMarried)
	// fmt.Printf("Type of : %T\n", weight)

}
