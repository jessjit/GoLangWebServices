package main

import (
	"fmt"
	"time"
)

func main() {
	InputChannel1 := make(chan int)
	InputChannel2 := make(chan int)
	InputChannel3 := make(chan int)
	InputChannel4 := make(chan int)
	InputChannel5 := make(chan int)

	go func() {
		for {
			InputChannel1 <- 1
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			InputChannel1 <- 2
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			InputChannel1 <- 3
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			InputChannel1 <- 4
			time.Sleep(time.Second * 4)
		}
	}()

	go func() {
		for {
			InputChannel1 <- 5
			time.Sleep(time.Second * 5)
		}
	}()
	for {
		select {
		case out1 := <-InputChannel1:
			fmt.Println(out1)
		case out2 := <-InputChannel2:
			fmt.Println(out2)
		case out3 := <-InputChannel3:
			fmt.Println(out3)
		case out4 := <-InputChannel4:
			fmt.Println(out4)
		case out5 := <-InputChannel5:
			fmt.Println(out5)
		}
	}
}
