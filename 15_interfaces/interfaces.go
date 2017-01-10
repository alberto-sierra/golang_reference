package main

import "fmt"

/*
  Interfaces are types that just declare behaviour.
  It defines signatures to be implemented by methods.
*/
type userActions interface {
  whoami() string
}

// An empty interface is implemented by everything
type entity interface {}

type user struct {
  name string
}

type client struct {
  user
  email string
}

type admin struct {
  user
  secret string
}

// the method whoami() matches the signature from the interface
func (usr user) whoami() string {
  return fmt.Sprint("My name is ", usr.name)
}

/*
  when this happens, it is said that the method implements the interface
  a method with a pointer receiver can only receive pointer types
*/
func (adm *admin) whoami() string {
  return fmt.Sprint("I am ", adm.name, " and I rule this land.")
}

func whoami(ua userActions) string {
  return fmt.Sprint("The user says: ", ua.whoami())
}

func main() {
  p := user{}
  p.name = "Peter"

  d := admin{}
  d.name = "Donald"

  userlist := []interface{}{p,d}

  fmt.Println(p.whoami())
  fmt.Println(d.whoami())
  // whoami() method for the admin type only receives pointer types
  fmt.Println(whoami(&d))
  fmt.Println(userlist)
}
