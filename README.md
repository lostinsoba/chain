# chain [![Go Reference](https://pkg.go.dev/badge/github.com/lostinsoba/chain.svg)](https://pkg.go.dev/github.com/lostinsoba/chain)

Lower and Upper Bounds Generator  
It works like Python's [range()](https://docs.python.org/3/library/stdtypes.html#range) but returns only the first and last values

## Installation

```shell
go get -u github.com/lostinsoba/chain
```

## Usage

```go
package main

import (
	"fmt"
	"sync"

	"github.com/lostinsoba/chain"
)

func main() {

	tasks := []int{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10,
	}

	var wg sync.WaitGroup
	var c chain.Chain

	c.SetStop(len(tasks))
	c.SetStep(4)

	for c.Next() {
		left, right := c.Bounds()
		subtasks := tasks[left:right]

		wg.Add(1)
		go performTasks(&wg, subtasks)
	}

	wg.Wait()

	// Output:
	// [9 10]
	// [5 6 7 8]
	// [1 2 3 4]
}

func performTasks(wg *sync.WaitGroup, subtasks []int) {
	fmt.Println(subtasks)
	wg.Done()
}
```