package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

type Names struct {
  fname string
  lname string
}


func main() {


  var fileName string
  

  
  fmt.Println("Please enter the name of a file.")
  fmt.Scan(&fileName)

  b, err := ioutil.ReadFile(fileName) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    str := string(b) // convert content to a 'string'
    str2 := strings.FieldsFunc(str, Split)
    tmp := make([]Names, len(str2)/2)
    j := 0
    for i := range str2 {
      if i%2==0{
        tmp[j].fname = str2[i]
      } else {
        tmp[j].lname = str2[i]
        j = j + 1
      }
    }

    for j := range tmp {
      fmt.Println(j, "fname: ", tmp[j].fname)
      fmt.Println(j, "lname: ", tmp[j].lname)
    }

    fmt.Println(tmp)
}

func Split(r rune) bool {
  return r == ' ' || r == '\n'
}