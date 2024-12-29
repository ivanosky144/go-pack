# Go Pack - A Comprehensive Data Structures Library

Go Pack is a Go library that provides a collection of commonly used data structures such as binary trees, hash sets, linked lists, n-ary trees, queues, stacks, etc. It is designed for developers who need well-implemented data structures in their Go projects.

## Usage

Add the library to your project by running
`
go get github.com/ivanosky144/go-pack
`

### Example: Queue

`
package main

import (
    "fmt"
    "github.com/ivanosky144/go-pack/queue"
)

func main() {
    q := queue.NewQueue[int]() // Create a new queue of integers
    q.Enqueue(1)
    q.Enqueue(2)
    fmt.Println(q.Dequeue()) // Output: 1
}
`

