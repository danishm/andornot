package main

import (
	"fmt"

	"github.com/danishm/andornot/gates"
)

func main() {
	board := gates.Board()

	gate := board.AND()
	gate.Pin1 <- 0
	gate.Pin2 <- 1
	out := <-gate.Out

	fmt.Println(out)

	board.Stop()
}
