package main

import (
 "fmt"
 "flag"
)

func main(){
 s:=flag.String("s","This is a string.","enter a string")
 i:=flag.Int("i", 100,"Enter an int")

 flag.Parse()
 fmt.Println("String:%s, Int:%d",*s,*i)
}

