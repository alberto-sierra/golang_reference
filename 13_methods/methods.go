package main

import "fmt"

// methods are tied to structs through the function's receiver parameter

type person struct {
  first string
  last string
}

func (p person) getfullname() string { // declared as a method for type person
  return fmt.Sprint(p.first, " ", p.last)
}

func main() {
  bob := person{}
  bob.first = "Robert"
  bob.last = "Peters"
  // the method is only visible through the variable:
  fmt.Println(bob.getfullname())
}
