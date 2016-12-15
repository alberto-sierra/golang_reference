package main

import "fmt"

func print(text int) { fmt.Println(text) }

func main() {
  number := 1000
  text := "Formatting and printing with Go."
  fmt.Println(text)
  fmt.Printf("%d - %b", number, number)
}
