package main

import "fmt"

func oddoreven(number int) {
  if number%2 == 0 {
    fmt.Println(string(number) + " is odd")
  } else {
    fmt.Println(string(number) + " is even")
  }
}

func whatis(thing interface{}) {
  /* Type assertion ( from "A tour of Go"):
     A type assertion provides access to an interface value's
     underlying concrete value.
     t := i.(T)
     If i does not hold a T, the statement will trigger a panic.
     To test whether an interface value holds a specific type,
     a type assertion can return two values:
     the underlying value and a boolean value that reports
     whether the assertion succeeded.
     t, ok := i.(T)
     If i holds a T, then t will be the underlying value and ok will be true.
     If not, ok will be false and t will be the zero value of type T,
     and no panic occurs.
  */
  if _, ok := thing.(int); ok { // <- Variables can also be initialized here
    fmt.Println("It's an integer!")
  } else if _, ok := thing.(string); ok {
    fmt.Println("It's a string!")
  } else {
    fmt.Println("I have no idea what this is.")
  }
}

func main() {
  oddoreven(3)
  whatis("this")
}
