package gates

// Gate defines the interface a gate should support
type Gate interface {
	Compute()
	Stop()
	Pin1() chan int
	Pin2() chan int
	Out() chan int
}
