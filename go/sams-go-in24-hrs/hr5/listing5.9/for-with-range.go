package main
import (
 "fmt"
)


func main () {
 array := []int {10,20,30,40,50}
 
 for i,n := range array {
  fmt.Printf ("index:%d, value:%d\n",i,n)
  }
}
