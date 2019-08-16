package main
import (
 "fmt"
)

func main () {
 mymap := make(map[string]int)
 mymap2 := map[string]string {
  "name": "Kiran",
  "age": "29",
  "sex":"male",
  }
 
 mymap["ten"] = 10
 value, ok := mymap2["age"]
 if ok {
  fmt.Println("Map contains age. "+value)
  }
}

