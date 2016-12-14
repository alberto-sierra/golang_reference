package main

import "fmt"

func ptrreceiver(ptr *int) {
	// Receives a pointer of an integer
	fmt.Print("(Function) Memory address: ")
	fmt.Println(ptr)
	fmt.Print("(Function) Value: ")
	// * dereferences a pointer to access its value
	fmt.Println(*ptr)
	*ptr = 200
}

func main() {
	// a pointer is a type: *type
	variable := 100
	// "&" converts to the memory address of a resource
	var varptr *int = &variable
	fmt.Print("(main) Memory address: ")
	fmt.Println(&variable)
	ptrreceiver(varptr) // or ptrreceiver(&variable)
	fmt.Print("(main) Value: ")
	fmt.Println(variable)
}
