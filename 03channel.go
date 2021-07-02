package main

import (
	"fmt"
	"time"
)

var total int

func main() {
	c := make(chan string)

	go say("Hi", 5, c)
	go say("Hola", 5, c)

	for msg := range c {
		fmt.Println(msg)
	}

	fmt.Println("DONE!")
}

func say(text string, times int, c chan string) {
	for i := 0; i < times; i++ {
		total++
		c <- fmt.Sprintf("[%03d] %s\n", i+1, text)
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
