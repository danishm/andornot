package main

import (
	"fmt"
	"time"
)

func Do(ch chan int) {
	ch <- 1
	fmt.Println("Done 1")
	ch <- 3
	fmt.Println("Done 3")
}

func main() {
	ch := make(chan int)

	go Do(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)

	time.Sleep(time.Second)
	fmt.Println("Done")
}
