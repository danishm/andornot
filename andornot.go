package main

import (
	"fmt"

	"github.com/danishm/andornot/gates"
)

func main() {
	board := gates.DefaultBoard()

	// Create a NAND gate
	gate := gates.NAND(board)

	// See if it works
	gate.Pin1() <- 1
	gate.Pin2() <- 1
	out := <-gate.Out()

	fmt.Println(out)

	board.Stop()
}

func simpleANDTest() {
	board := gates.DefaultBoard()

	gate := gates.AND(board)
	gate.Pin1() <- 0
	gate.Pin2() <- 1
	out := <-gate.Out()

	fmt.Println(out)

	board.Stop()
}
