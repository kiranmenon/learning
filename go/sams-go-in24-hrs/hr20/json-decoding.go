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
 log.SetFlags(log.LstdFlags | log.Lshortfile)
 //creating new Person struct
 hobbies := []string {"Gym","HIIT"}
 p:= Person {
  Name: "Kiran",
  Age: 30,
  Hobbies: hobbies,
 }
 fmt.Printf("Person struct: %+v\n",p)
 //person struct encoded as json
 pAsJsonByte,err := json.Marshal(p)
 if err != nil {
  log.Fatal(err)
 }
 pAsJsonString := string(pAsJsonByte)
 fmt.Println("Json: ",pAsJsonString)

 //person struct is edited
 hobbies = append(hobbies,"Biking")
 p.Hobbies = hobbies
 pAsJsonByte,err = json.Marshal(p)
 if err != nil {
  log.Fatal(err)
 }
 pAsJsonString = string(pAsJsonByte)
 fmt.Println("Json: ",pAsJsonString)

 //New Json string for decoding to person struct
 newPerson := `
 {
  "name":"Karun",
  "age":26,
  "sex":"male",
  "hobbies":["swimming","karate"]
 }`

 newPersonAsByte := []byte(newPerson)
 err = json.Unmarshal(newPersonAsByte,&p)
 if err != nil {
  log.Fatal(err)
 }
 fmt.Printf("New Person struct: %+v\n",p)
}

