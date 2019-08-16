package main
import (
 "fmt"
 "encoding/json"
 "log"
)

type Person struct {
 Name string `json:"name"`
 Age int `json:"age"`
 Sex string `json:"sex"`
 Hobbies []string `json:"hobbies"`
}

func main() {
 hobbies := []string {"Gym","HIIT"}
 p:= Person {
  Name: "Kiran",
  Age: 30,
  Hobbies: hobbies,
 }
 fmt.Printf("Person struct: %+v\n",p)
 pAsJsonByte,err := json.Marshal(p)
 if err != nil {
  log.Fatal(err)
 }
 pAsJsonString := string(pAsJsonByte)
 fmt.Println("Json: ",pAsJsonString)
}

