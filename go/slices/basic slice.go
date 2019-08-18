package main

import (
	"fmt"
	"log"
)

func main() {
	s := make([]string, 3)
	s = []string{"ABC"}
	fmt.Println("String:", s)
	s = append(s, " is a company name")
	log.Println("String appended:", s)
	//log.Fatal("This is Fatal")
	for index, value := range "Kiran" {
		log.Printf("%d,%#U",index, value)
	}
}
