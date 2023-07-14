package main

import (
	"fmt"
	"time"
)

func main() {

	var intChannel chan int

	intChannel = make(chan int, 10)

	go func() {
		for {
			time.Sleep(time.Second)
			intChannel <- 1
		}
	}()

	stringChannel := make(chan string, 10)
	go func() {
		for {
			time.Sleep(time.Second)
			stringChannel <- "string"
		}
	}()
	for {
		time.Sleep(time.Second)
		select {
		case v := <-intChannel:
			fmt.Println(v)
		case v := <-stringChannel:
			fmt.Println(v)
		default:
			fmt.Println("nil")
		}
	}

}
