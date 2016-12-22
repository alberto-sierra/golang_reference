package main

import "fmt"

/*
  A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type.
  The value of an uninitialized map is nil.
  A map is a reference type.
*/

func makemap() map[string]string {
  /*
    The preferred way to create a map is with the make instruction. It takes an
    optional length parameter.
    make(map[string]int)
  */
  m := make(map[string]string)
  // The map is ready right after creation
  m["EN"] = "Good Morning"
  m["ES"] = "Buenos Días"
  m["FR"] = "Bonjour"
  m["PT"] = "Bon Dia"
  m["IT"] = "Buongiorno"
  return m
}

func emptymap() (map[string]int, map[string]int) {
  // Creates an empty map
  m := map[string]int{}
  var n = map[string]int{}
  m["m"] = 1
  return m, n
}

func nilmap() map[string]int {
  // A nil map is equivalent to an empty map except that no elements may be added.
  var m map[string]int
  // This would fail (panic: assignment to entry in nil map)
  m["a"] = 1
  return m
}

func deletefrommap(m map[string]int, k string) map[string]int {
  if _, exists := m[k]; exists {
      delete(m,k)
      fmt.Printf("Key %s deleted\n", k)
  } else {
    fmt.Println("Specified key not found: ",k)
  }
  return m
}

func main() {
  fmt.Println("Contents: ", makemap())
  fmt.Println("Map lenght: ", len(makemap()))
  fmt.Print("Contents: ")
  fmt.Println(emptymap())
  fmt.Printf("\nDelete from maps\n")
  m := map[string]int{"a":1, "b":2, "c":3}
  fmt.Println("Before: ", m)
  deletefrommap(m, "x")
  m = deletefrommap(m, "b")
  fmt.Println("After: ", m)
  fmt.Printf("Assigning to a nil map fails:\n\n")
  fmt.Println(nilmap())
}
