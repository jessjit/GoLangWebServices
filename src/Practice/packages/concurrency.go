package packages

import (
	"time"
)

//BASIC CONCURRENCY
func Count(thing string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

//TRY FOR BUFFERED CHANNEL FUNCTION
func BufferConc() chan int {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	return c
}

//BASIC CONCURRENCY QUESTION
type slice []string

func CountVal(stringslice slice, channel1 chan string) {
	defer close(channel1)
	for index, _ := range stringslice {
		channel1 <- stringslice[index]
		time.Sleep(time.Second)
	}

}
