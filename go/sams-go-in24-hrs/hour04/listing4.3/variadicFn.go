package main
import (
 "fmt"
 "strconv"
)

func sumNumbers (numbers ...int) int {
 
 total :=0

 for _,number := range numbers {
  total += number
  }
 return total
}

func main () {

 fmt.Println ("Sum of numbers 1,2,3,4 is: " + strconv.Itoa (sumNumbers (1, 2, 3, 4) ) )
}
