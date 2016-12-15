package main

import "fmt"

func main() {
	// A string is a list of characters
	hello := "Hello"
	fmt.Println("1: " + string(hello[0]))
	fmt.Printf("2: %v\n", []byte(hello))

	// A character is an encoded number
	h := 104                       // <- initialized as a number
	fmt.Println("3: " + string(h)) // <- printed as a string

	/*
					   UTF-8 is an 4-byte encoding that fits
		         entirely in the int32 type.
				     A rune is an alias for the int32 type.
	*/
	e := 'e' // <- simple quotes for rune
	fmt.Printf("4: %s (%v)\n", string(e), e)

	fmt.Println(
		`5: Backticks can be used to
   "print" text maintaing the
   format ¯\_(ツ)_/¯`)

}
