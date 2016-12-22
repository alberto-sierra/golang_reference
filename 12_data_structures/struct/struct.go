package main

import "fmt"

/*
  A struct is a sequence of named elements named fields, each of which
  has a name and a type.
  A struct provides OOP encapsulation.
*/

type person struct {
  first string
  last string
  age uint8
}

// structs can contain other structs
type employee struct {
  person
  position string
  salary float64
}

type contractor struct {
  employee // <- anonymous field, takes the name of its type
  salary int // <- overrides inheritance
  org string
}

func main() {
  e1 := employee{
    person: person{
      first: "Mario",
      last: "Lucca",
      age: 32,
    },
    position: "manager",
    salary: 55000.0,
  }

  contractors := []contractor{}

  // define a struct variable from literals
  c1 := contractor{
    employee: employee {
      person: person {
       first: "Adrian",
       last: "Chen",
       age: 45,
     },
     salary: 3000.0,
    },
    salary: 60,
    org: "ContractMe ltd.",
  }

  c2 := contractor{}
  c2.first = "Dominic"
  c2.last = "Smith"
  c2.org = "Computers and Stuff"

  contractors = append(contractors, c1, c2)

fmt.Printf("Name: %s\nAge: %d\nPosition: %s\n\n",
  e1.first, e1.age, e1.position)

for c := range contractors {
  fmt.Printf("Name: %s %s\nOrganization: %s\nSalary: %.2f(%d)\n\n",
    contractors[c].first,
    contractors[c].last,
    contractors[c].org,
    // accessing overriden fields
    contractors[c].employee.salary,
    contractors[c].salary,
  )
}

// pointers behave different with struct types
p1 := &person{} // pointer to struct person
p1.first = "Carlos"
p1.last = "Ericsson"
fmt.Printf("Contents: %v\nType: %T\n", p1, p1)




}
