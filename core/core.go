package core

// Component defines the interface a component needs to implement
type Component interface {

	// Stops the component from doing it's job and fee up resources
	Stop()

	// Do all the actions a component needs to do to do it's job
	Run()
}
