package main
import (
 "fmt"
 "time"
)

func ping1 (msg chan string) {
 time.Sleep(time.Second*1)
 msg <- "Ping1"
}

func ping2 (msg chan string) {
 time.Sleep(time.Second*2)
 msg <- "Ping2"
}

func main () {
 ch1 := make(chan string)
 ch2 := make(chan string)

 go ping1(ch1)
 go ping2(ch2)

 select {
  case msg := <-ch1:
   fmt.Println("Got message from ch1: "+msg)
  case msg := <-ch2:
   fmt.Println("Got msg from ch2: "+ msg)
  case <-time.After(5 * time.Millisecond):
   fmt.Println("no messages received. giving up.")
  }
}

