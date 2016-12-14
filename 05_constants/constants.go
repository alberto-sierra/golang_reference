package main

import "fmt"

// constants can be TYPED or UNTYPED:
const c1 string = "The only constant is change"
const c2 = true

/*
An iota is a special kind of constant that
autoincrements a *group* of declared constants
The initial value for any iota is 0
*/

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

/*
The character "_" is called a blank identifier
works as a blackhole for discarding return from a function
*/

const (
	_ = iota // discards value
	d        // 1
	e        // 2
)

func main() {
	fmt.Println("First constant declaration:")
	fmt.Printf("%s, %t\n", c1, c2)
	fmt.Println("First iota declaration:")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("Second iota declaration:")
	fmt.Println(d)
	fmt.Println(e)
}
