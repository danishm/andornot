package gates

import (
	"fmt"
	"sync"

	"github.com/danishm/andornot/core"
)

// Board defines the interface for a board on which a circuit
// could be built. Currently, a boards responsibility is to
// keep track of all the gates running and be able to centralize
// stopping them
type Board interface {
	core.Stoppable
	core.Countable
	RunComponent(core.Component)
	AddComponent(core.Component)
	Connect(from chan int, to ...chan int)
}

type board struct {
	running    bool
	mu         sync.Mutex
	components []core.Component
}

// Run gets the gate running
func (b *board) RunComponent(c core.Component) {
	go func() {
		for b.running {
			c.Run()
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
	for _, component := range b.components {
		component.Stop()
	}
}

func (b *board) AddComponent(c core.Component) {
	b.mu.Lock()
	b.components = append(b.components, c)
	b.mu.Unlock()
}

// DefaultBoard creates a new circuit board to work with
func DefaultBoard() Board {
	return &board{
		running: true,
	}
}

func (b *board) CoreGatesCount() int {
	total := 0
	for _, c := range b.components {
		total += c.CoreGatesCount()
	}
	return total
}
