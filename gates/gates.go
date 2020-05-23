package gates

import (
	"github.com/danishm/andornot/core"
)

// Gate defines the interface a gate should support
type Gate interface {
	core.Component
	Pin1() chan int
	Pin2() chan int
	Out() chan int
}
