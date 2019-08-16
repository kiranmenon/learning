package main
import (
 "fmt"
)

func main () {
 defer fmt.Println("First defer")
 fmt.Println("Normal println.")
 defer fmt.Println("2nd defer")
 defer fmt.Println("3rd defer")
 defer fmt.Println("4th defer")
 defer fmt.Println("5th defer")
 fmt.Println("Main finished.")
}

