package main

import (
  "fmt"
)

/*
  Arrays: a static list of literals of the same type.
*/

func emptyarray() [25]string {
  // an array is always initialized to the zero value of its type
  var arr [25]string
  return arr
}

func octet() ([4]uint8, int) {
	/*
    An array is static and its lenght and type
    must be explicitly defined on creation
  */
	var brd [4]uint8
	for i := range brd {
		brd[i] = 255
	}
  return brd, len(brd)
}

// An array can be multi-dimensional (a matrix)
func m2x2(list [4]int) (mx [2][2]int) {
  h := 0
  for i := 0; i < 2; i++ {
    for j := 0; j < 2; j++ {
      mx[i][j] = list[h]
      h++
    }
  }
  return
}


func main() {
  emptyarray := emptyarray()
  fmt.Printf("Contents: %v\nSize: %d\n", emptyarray, len(emptyarray))
  cont, length := octet()
  fmt.Printf("Contents: %v\nSize: %d\n", cont, length )
  fmt.Printf("Contents: %v\n", m2x2([4]int{2,4,6,8}) )
}
