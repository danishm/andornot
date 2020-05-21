package gates

import (
	"fmt"
	"sync"
)

// Board defines the interface for a board on which a circuit
// could be built. Currently, a boards responsibility is to
// keep track of all the gates running and be able to centralize
// stopping them
type Board interface {
	Stop()
	Run(Gate)
	AddGate(Gate)
	Connect(from chan int, to ...chan int)
}

type board struct {
	running bool
	mu      sync.Mutex
	gates   []Gate
}

// Run gets the gate running
func (b *board) Run(g Gate) {
	go func() {
		for b.running {
			g.Compute()
		}
		fmt.Println("Board stopped")
	}()
}

// Connect connects output from one pin to multiple input pins
func (b *board) Connect(from chan int, to ...chan int) {
	go func() {
		for b.running {
			val, ok := <-from
			if ok {
				for _, pin := range to {
					pin <- val
				}
			}
		}
	}()
}

func (b *board) Stop() {
	fmt.Println("Board stop initiated...")
	b.running = false
	for _, gate := range b.gates {
		gate.Stop()
	}
}

func (b *board) AddGate(g Gate) {
	b.mu.Lock()
	b.gates = append(b.gates, g)
	b.mu.Unlock()
}

// DefaultBoard creates a new circuit board to work with
func DefaultBoard() Board {
	return &board{
		running: true,
	}
}
