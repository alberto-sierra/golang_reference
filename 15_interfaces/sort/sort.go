package main

import (
  "fmt"
  "sort"
)

type people []string

// methods to implement sort.Interface interface
func (p people) Len() int { return len(p) }
func (p people) Less(i, j int) bool { return p[i] < p[j] }
func (p people) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {

  // using the interface from type people
  studyGroup := people{"zeno", "Zeno", "John", "Al", "Jenny"}
  fmt.Println(studyGroup)
  sort.Sort(studyGroup)
  fmt.Println(studyGroup)
  sort.Sort(sort.Reverse(studyGroup))
  fmt.Println(studyGroup)
  fmt.Println("")

  // using the sort string methods from the sort package
  anotherStudyGroup := []string{"zeno", "Zeno", "John", "Al", "Jenny"}
  fmt.Println(anotherStudyGroup)
  sort.Strings(anotherStudyGroup)
  fmt.Println(anotherStudyGroup)
  sort.Sort(sort.Reverse(sort.StringSlice(anotherStudyGroup)))
  fmt.Println(anotherStudyGroup)
  fmt.Println("")

  // using the sort int methods from the sort package
  scoreCard := []int{7,4,8,2,9,19,12,32,3}
  fmt.Println(scoreCard)
  sort.Ints(scoreCard)
  fmt.Println(scoreCard)
  sort.Sort(sort.Reverse(sort.IntSlice(scoreCard)))
  fmt.Println(scoreCard)

}
