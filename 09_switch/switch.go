package main

import "fmt"

func simpleswitch(choice int) {
  /*
     the case statement has no default fallthrough
     (a break is not required).
  */
  fmt.Println("Simple switch")
  switch choice {
  case 1:
    fmt.Println("First choice.")
  case 2:
    fmt.Println("Second choice.")
  case 3:
    fmt.Println("Third choice.")
  default:
    fmt.Println("Nothing matched.")
  }
}

func fallthroughswitch(word string) {
  /*
     Fallthrough is explicit
  */
  fmt.Println("Second example:")
  switch { // <- no switch variable here
  case word == "alpha": // <- comparison built in the case statement
    fmt.Println(string(945))
    fallthrough // <- explicit fallthrough
  case word == "beta":
    fmt.Println(string(946))
  case word == "gamma", word == "delta": // <- multiple cases
    fmt.Println(string(947) + " " + string(948))
  default:
    fmt.Println("Nothing matched.")
  }
}

func complexcases(number int) {
  switch {
  case number%2 == 0:
    fmt.Println("Even.")
    fallthrough
  case number > 0 && number < 10: // <- AND/OR
    fmt.Println("Positive and less than ten.")
  }
}

func swithbytype(thing interface{}) {
  switch thing.(type) { // <- type assertion
  case int:
    fmt.Println("Integer.")
  case string:
    fmt.Println("String.")
  case rune:
    fmt.Println("Rune.")
  default:
    fmt.Println("Unknown.")
  }
}

func main() {
  simpleswitch(2)
  fallthroughswitch("alpha")
  complexcases(4)
  swithbytype('x')
}
