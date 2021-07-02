package main

import (
	"fmt"
	"time"
)

var total int

func main() {
	go say("Hi", 10)
	say("Hola", 5)

	fmt.Println("DONE!")
}

func say(text string, times int) {
	for i := 0; i < times; i++ {
		total++
		fmt.Printf("[%03d] %s\n", i+1, text)
		time.Sleep(time.Millisecond * 500)
	}
}
