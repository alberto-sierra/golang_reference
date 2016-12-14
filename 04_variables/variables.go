package main

import "fmt"

// Standard declaration (zero value)
var a int
var b float64
var c string
var d bool

func stdeclaration() {
	fmt.Printf("a = %d\nb = %f\nc = %s\nd = %t", a, b, c, d)
}

// Shorthand declaration (with assignment)
func shorthand() { // only inside functions
	e := 1
	f := 1.0
	g := "qwerty"
	h := true
	fmt.Printf("e = %d\nf = %f\ng = %s\nh = %t\n", e, f, g, h)
}

func main() {
	stdeclaration()
	shorthand()
}
