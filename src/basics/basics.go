// Continue the Go tour from here: https://tour.golang.org/flowcontrol/1

// The package name of this file
// package main
package basics

// By convention, the package name is the same as the last element of the import path
// "factored" import statement below, could also put them on multiple "import x" lines

// func main() {
// 	var a = "hello"    // Type is inferred from the right side
// 	b := "world"       // Declare and initialize with :=
// 	var c, d string    // Only declare, uninitialized vars will hold 0 value ie. ""
// 	numA, numB := 5, 6 // Standard naming convention: camelCase

// 	// A name is exported if it begins with a capital letter, math.pi will NOT work
// 	fmt.Println("Math.Pi = ", math.Pi)

// 	// Calling a function and using sprintf
// 	fmt.Println(fmt.Sprintf("add %d %d = %d", numA, numB, add(numA, numB)))

// 	// Return tuple
// 	c, d = Swap(a, b)
// 	fmt.Printf("Swap `%s` and `%s`: `%s` `%s`\n", a, b, c, d)

// 	numC, numD := split(17)
// 	fmt.Printf("Split 17 into %d and %d\n", numC, numD)

// 	// Type conversion, printing types
// 	fmt.Printf("Type of numA: %T\tNow converted: %T\n", numA, float32(numA))
// }

// (x, y int) : only the last param needs to be type defined if they share the same type
func add(x, y int) int {
	return x + y
}

// Swap returns a tuple. Since it is capitalized it will be exported as part of this package
func Swap(x, y string) (string, string) {
	return y, x
}

// Named return values, (x, y int) names the return values so they can be used in the function. 'return' returns all named variables
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/* 	Setup and Workspace

The Go workspace has 2 folders, bin and src.
src typically holds many difference source controlled packages, bin holds executables (binaries)
Go will always look in packages under the $GOPATH variable, which we can see with 'go env GOPATH'
If we want to write our own packages we must include the Go workspace in the GOPATH
This can be done by adding `export GOPATH="path/to/go/workspace:$HOME/go"`
The path/to/go/workspace should have a src and bin folders
GOPATH is colon separated
$HOME/go is the default GOPATH
*/

/*	Imports and Packages
	A package is a module that can be imported.
	All Go files under the same folder must be part of the same package
	Therefore, go filenames do not need to be the same as the package name
*/

/*	Types
	Basic Types:
	bool
	string	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // alias for uint8
	rune // alias for int32, represents a Unicode code point
	float32 float64
	complex64 complex128
*/
