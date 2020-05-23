package core

// Runnable represents that has some work to do
type Runnable interface {
	// Do all the actions a component needs to do to do it's job
	Run()
}

// Stoppable represents that's something that runs and can be stopper
type Stoppable interface {
	// Stops the component from doing it's job and fee up resources
	Stop()
}

// Countable represents something that can count the number of basic gates it has
type Countable interface {
	// Return the count of basic gates in use i.e. either AND, OR & NOT
	CoreGatesCount() int
}

// Component defines the interface a component needs to implement
type Component interface {
	Runnable
	Stoppable
	Countable
}
