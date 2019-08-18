package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal,1)
	done := make(chan bool, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig1 := <-sig
		fmt.Println()
		fmt.Println(sig1)
		done <- true
	}()

	fmt.Println("Waiting for signal")
	<-done
	fmt.Println("Exiting..")
}
