package main

import (
	"log"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		for i := range make([]int, 100) {
			log.Println(i)
		}
		chan1 <- "Done"
	}()

	go func() {
		//time.Sleep(1 * time.Second)
		chan2 <- "Done"
	}()

	for  range make([]int,2) {
		select {
		case msg1 := <-chan1 :
			log.Println("Go routine 1 finished.",msg1)
		case msg2 := <-chan2 :
			log.Println("Channel2 messaged.",msg2)
		}
	}
}
