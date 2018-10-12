package main

import (
  "fmt"
  "encoding/json"
)

func main() {
  var name string
  var address string

  m := make(map[string]string)
  
  
  fmt.Println("Please enter a name.")
  fmt.Scan(&name)

  fmt.Println("Please enter an address.")
  fmt.Scan(&address)
  
  
  m["address"] = address
  m["name"] = name

  mapping, _ := json.Marshal(m)
  fmt.Println(string(mapping))
/*
  fmt.Println("Please a name.")
  fmt.Scan(&name)

  fmt.Println("Please an address.")
  fmt.Scan(&address)
  
  m[name] = address
*/
/*
  fmt.Println("name: ", name)
  fmt.Println("address: ", address)
  fmt.Println("map: ", m)
  */
}