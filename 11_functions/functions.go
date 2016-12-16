package main

import "fmt"

func dummy(parameter string)  { // <- a function declares a parameter
  return // does nothing
}

func greet(first, last string) { // Or FunctionName(y T, x T)
  fmt.Println(fmt.Sprint("Howdy!, ",first," ",last,"!"))
}

func countchars(first, last string) int {
  return len(fmt.Sprint(first,last))
}

func sumchars(first, last string) (sum int) { // <- named return value
  n := []byte(fmt.Sprint(first,last))
  sum = 0 // pre-declared in the function signature
  for i := 0; i < len(n); i++ {
    sum += int(n[i])
  }
  return // returns the declared variable (sum)
}

func reversename(first, last string) (string, string) { // Returns 2 values
  return last, first
}

// Variadic functions can take 0 or more parameters.
func average(n ...float64) float64 { // <- n becomes a []float64
  var total float64 // <- auto-initialized to zero
  /* range loops over an array, slice, string, or map and can also
  read from a channel. It returns a k,v pair.
  */
  for _, value := range n { // <- discard the key, because n is a slice.
    total += value
  }
  return total / float64(len(n)) // division by zero returns NaN!!
}

// Functions can be returned from another function
func addtwo() func() int { // Return type: "func() int"
  var x int
  return func() int { // Returns a "func() int" as defined above
    x += 2
    return x
  }
}

// And have functions as parameters (callback)
func isodd(n int, callback func(string)) bool {
  if n%2 != 0 {
    callback("The number is odd")
    return true
  }
  return false
}


func main() {
  dummy("argument") // <- an argument is passed to a function (not a parameter)
  greet("Harry","Potter")
  fmt.Printf("Number of characters: %d\n", countchars("Harry", "Potter"))
  fmt.Printf("Sum of characters: %d\n", sumchars("Harry", "Potter"))
  /*
    This fails, because Printf cannot mix a string with a multi-return function:
    fmt.Println("Reversed: %s %s\n", reversename("Harry","Potter"))
  */
  last, first := reversename("Harry","Potter")
  fmt.Printf("Reversed: %s %s\n", last, first)
  /*
    defer delays execution of a function to before its parent process finishes
    It is useful for keeping I/O open() and close() statements near but
    allowing the resources to be used until the process ends.
  */
  // This will execute at the end of main()
  defer fmt.Printf("Average: %.2f\n", average(10, 20, 30, 40, 50))

  // It's not possible to have a function inside another function, but...
  nestedfunc := func() { // you can assign a function to a variable
    fmt.Println("A function is a Type!")
  }
  nestedfunc()

  // This fails: fmt.Println(addtwo())
  even := addtwo()
  fmt.Println(even())
  fmt.Println(even())

  // or pass a function as an argument (callback).
  number := 5
  odd := isodd(number, func(s string) { fmt.Println(s) })
  if odd {
    fmt.Printf("The remainder is: %d\n", number%2)
  }

  // Anonymous, self-executing functions.
  func() { // <- no name, no signature
    fmt.Println("Look ma, no hands!")
  }() // <- auto-execution call


}
