package main

import (
 "fmt"
 "strconv"
 )

func main () {
 var str = "1234"
 var str2 = "1234aqwe5"

 fmt.Println (strconv.ParseInt(str,10,32))
 fmt.Println (strconv.ParseInt(str2,10,32))
}
