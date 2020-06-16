package identity

import "fmt"

var counts = map[string]int{}

// Get returns an unique ID given a name of a type
func Get(name string) string {
	if _, ok := counts[name]; !ok {
		counts[name] = 0
	}

	next := counts[name] + 1
	counts[name] = next
	return fmt.Sprintf("%s%d", name, next)
}
