package main

import (
 "fmt"
)


func stringConcate(s string) string {

 return s+ " EOF"
}

func main () {

 str2Cat := stringConcate("Hello")
 fmt.Println (str2Cat)
}

