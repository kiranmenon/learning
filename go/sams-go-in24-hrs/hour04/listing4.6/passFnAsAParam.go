package main
import (
 "fmt"
 "strconv"
)

func caller (fn func() string) int {
 i, _ := strconv.ParseInt(fn(),10,32)
 return int(i)
}

func theRealFn ()  string {
 fmt.Println ("Called theRealFn.")
 return "1234"
}

func main () {
 fmt.Printf("%d\n",caller(theRealFn))
}

