package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go say("Hi", 10)
	go say("Hola", 5)

	wg.Wait()
	fmt.Println("DONE!")
}

func say(text string, times int) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		fmt.Printf("[%03d] %s\n", i+1, text)
		time.Sleep(time.Millisecond * 500)
	}
}
