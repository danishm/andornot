package main

import (
	"fmt"

	"github.com/danishm/andornot/gates"
)

func main() {
	board := gates.Board()

	// Create a NAND gate
	and := board.AND()
	not := board.NOT()
	board.Connect(and.Out, not.Pin1)

	// See if it works
	and.Pin1 <- 1
	and.Pin2 <- 1
	out := <-not.Out

	fmt.Println(out)

	board.Stop()
}

func simpleANDTest() {
	board := gates.Board()

	gate := board.AND()
	gate.Pin1 <- 0
	gate.Pin2 <- 1
	out := <-gate.Out

	fmt.Println(out)

	board.Stop()
}
