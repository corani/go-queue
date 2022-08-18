package main

import (
	"fmt"

	"github.com/corani/go-queue/queue"
)

func main() {
	q := queue.New[string]()
	q.Push("Hello")
	q.Push("World")

	for !q.Empty() {
		fmt.Println(q.Pop())
	}
}
