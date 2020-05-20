package main

import (
	"fmt"

	"github.com/danishm/andornot/gates"
)

func main() {
	board := gates.DefaultBoard()

	// Create a NAND gate
	and := gates.AND(board)
	not := gates.NOT(board)
	board.Connect(and.Out(), not.Pin1())

	// See if it works
	and.Pin1() <- 1
	and.Pin2() <- 0
	out := <-not.Out()

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
