package main

import (
	"fmt"
	"time"
)

func fetchData(id int, ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Task data %d loaded!", id)
}

func main() {
	channel := make(chan string)

	for i := range 3 {
		go fetchData(i, channel)
	}

	for _ = range 3 {
		resultado := <-channel
		fmt.Println(resultado)
	}
}
