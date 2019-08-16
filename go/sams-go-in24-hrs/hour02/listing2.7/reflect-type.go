package main

import (
 "fmt"
 "reflect"
)

func main (){
 var i int = 10
 var s string = "Kiran"

 fmt.Println("Type of i: " + reflect.TypeOf(i).Name())
 fmt.Println("Type of s: " + reflect.TypeOf(s).Name())
}
