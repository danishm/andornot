Logic Gates Using Go Channels
=============================

I've been using **Go** as my primary development language for over 2 years now. I thought the language's concept of [channels](https://golang.org/doc/effective_go.html#concurrency) was pretty cool. I used to wonder if I can make simple logic gates using channels and if they can be put together to form circuits.

This is my attempt!

Examples
-------

### Trying a simple AND gate

#### Code

```go
package main

import (
	"fmt"

	"github.com/danishm/andornot/gates"
)

func main() {
    board := gates.Board()

    gate := board.AND()
    gate.Pin1 <- 0
    gate.Pin2 <- 1
    out := <-gate.Out

    fmt.Println(out)

    board.Stop()
}
```

#### Output

```
0
```

### Construct a NAND gate using AND and NOT

#### Code
```go
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
```

#### Output

```
0
```