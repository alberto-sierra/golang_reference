package main

import "fmt"

// Slicing: create a slice from parts of another array (or slice)
func slicefour(collection [10]int) []int {
  return collection[3:7]
}

/*
  A slice is a pointer to a section of an underlying array. The size of the
  underlying array becomes the slice's `capacity` which can be different from
  the slice's size or length. Because it is a pointer, it is a dynamic and
  can change its size during runtime.
*/

func makeslice() []int{
  /*
    The best way to create a slice is with the make instruction
    make(type, len[, cap])
    if capacity is not set, by default it will match length.
  */
  slc := make([]int, 10)
  // different length/capacity:
  // slc := make([int], 10, 20)
  return slc
}

func newslice() []int {
  // This is the exact same equivalent to using make (with length only)
  slc := new([10]int)[0:10]
  // different length/capacity:
  // slc := new([20]int)[0:10]
  return slc
}

func shorthand() []int {
  // In this way, the underlying array will be created with no size.
  slc := []int{}
  // Direct assignment using slc[n] is not possible (index out of range)
  return slc
}

func varslice() []int {
  // In this way, the new slice does not instantiate an underlying array.
  var slc []int
  // Direct assignment using slc[n] is not possible (index out of range)
  return slc
}

func addtoslice(slc []int, value int) []int {
  /*
    Slices created with an emtpy array (:=) or with var, need to use append
    to add capacity and length to the slice.

    s := make([]int, 1, 2)
    s[0] = 1 // <- this is okay
    s[1] = 2 // <- this will panic with index out of range (length is 1)
  */
  return append(slc,value)
}

func main() {
  fmt.Printf("Contents: %v\n", slicefour([10]int{1,2,3,4,5,6,7,8,9,10}))
  fmt.Printf("Slice with make: %v\n", makeslice())
  fmt.Printf("Slice is nil: %t\n", makeslice() == nil)
  fmt.Printf("Slice with shorthand: %v\n", shorthand())
  fmt.Printf("Slice is nil: %t\n", shorthand() == nil)
  fmt.Printf("Slice with var: %v\n", varslice())
  fmt.Printf("Slice is nil: %t\n", varslice() == nil)
  fmt.Println("Appending to slices:")
  fmt.Printf("Content (initialized to []): %v\n", addtoslice(shorthand(), 0))
  fmt.Printf("Content (uninitialized): %v\n", addtoslice(varslice(), 0))
  // Append must also be used with slices when the length has been reached:
  slc := make([]int, 1, 2)
  fmt.Printf("Length: %v\n", len(slc))
  fmt.Printf("Capacity: %v\n", cap(slc))
  fmt.Println("Modifying slices:")
  fmt.Printf("Contents: %v\n", slc)
  slc[0] = 1
  fmt.Printf("Contents: %v\n", slc)
  // would panic with index out of range (length is 1):
  // slc[1] = 0
  fmt.Printf("Appending beyond initial length and capacity:\n\n")
  slc = append(slc, 10)
  fmt.Println(slc)
  fmt.Printf("Length: %v\n", len(slc))
  fmt.Printf("Capacity: %v\n\n", cap(slc))
  slc = append(slc, 20)
  fmt.Println(slc)
  fmt.Printf("Length: %v\n", len(slc))
  fmt.Printf("Capacity: %v\n\n", cap(slc))
  slc = append(slc, 30)
  fmt.Println(slc)
  fmt.Printf("Length: %v\n", len(slc))
  fmt.Printf("Capacity: %v\n\n", cap(slc))
  slc = append(slc, 40)
  fmt.Println(slc)
  fmt.Printf("Length: %v\n", len(slc))
  fmt.Printf("Capacity: %v\n\n", cap(slc))
  fmt.Printf("Deleting from slices:\n\n")
  slc = append(slc[:0], slc[1:]...)
  fmt.Println(slc)
}
