package main
import (
 "fmt"
)


func multipleValReturns () (int,string) {
 return 10, "Kiran"
}

func main () {
 i, s := multipleValReturns()
 fmt.Printf ("returned values %T %v, %T %v\n",i,i,s,s)
}

