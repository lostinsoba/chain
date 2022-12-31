# chain [![Doc][doc-badge]][doc-uri] [![Coverage][coverage-badge]][coverage-uri]

Lower and Upper Bounds Generator

## Description

It works like Python's [range()](https://docs.python.org/3/library/stdtypes.html#range) generator:
```python
>>> for i in range(0, 15, 5):
...     print(i)
... 
0
5
10
```

but returns the first and last values for every iteration instead:

```go
var c chain.Chain
c.SetStop(15)
c.SetStep(5)

for c.Next() {
    left, right := c.Bounds()
    fmt.Println(left, right)
}

// 0 5
// 5 10
// 10 15
```

and it's also possible to iterate in the opposite direction:

```go
c.Reverse()
for c.Next() {
    left, right := c.Bounds()
    fmt.Println(left, right)
}

// 10 15
// 5 10
// 0 5
```

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

func performTasks(wg *sync.WaitGroup, tasks []int) {
    fmt.Println(tasks)
    wg.Done()
}
```

[doc-badge]: https://pkg.go.dev/badge/github.com/lostinsoba/chain.svg
[doc-uri]: https://pkg.go.dev/github.com/lostinsoba/chain
[coverage-badge]: https://codecov.io/gh/lostinsoba/chain/branch/main/graph/badge.svg
[coverage-uri]: https://codecov.io/gh/lostinsoba/chain