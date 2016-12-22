package main

import (
  "fmt"
  "encoding/json"
)

type person struct {
  // to use a field with JSON it must be exported
  First string
  // JSON key names can be overriden using struct tags
  Last string `json:"LastName"`
  // keys can also be suppressed from being included in the JSON data
  Age uint8 `json:"-"`
  // unexported fields won't be included in the JSON data either.
  nickname string
}

func main() {
  p1 := person{}
  p1.First = "Joao"
  p1.Last = "Valls"
  p1.Age = 35
  p1.nickname = "Zico"

  p2 := person{}

  // json.Marshall converts a struct to JSON (returns a []byte)
  data, _ := json.Marshal(p1)
  fmt.Println("Marshal (to JSON)")
  fmt.Println(data)
  fmt.Println(string(data))

  // json.Unmarshall converts JSON to a struct (a pointer to a struct)
  somejson := []byte(`{"First":"Michael","LastName":"Perez","Age":27}`)
  json.Unmarshal(somejson, &p2)
  fmt.Println("Unmarshal (from JSON)")
  fmt.Println(p2)
}
