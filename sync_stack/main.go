package main

import (
	"sync"
)

// Need to show solution

type Stack struct {
	mutex sync.Mutex
	data  []string
}

func NewStack() *Stack {
	return &Stack{}
}

func (b *Stack) Push(value string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.data = append(b.data, value)
}

func (b *Stack) Pop() string {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if len(b.data) == 0 {
		panic("pop: stack is empty")
	}

	value := b.data[len(b.data)-1]
	b.data = b.data[:len(b.data)-1]
	return value
}

var stack Stack

func producer() {
	for i := 0; i < 1000; i++ {
		stack.Push("message")
	}
}

func consumer() {
	for i := 0; i < 10; i++ {
		stack.Pop() //Тут могли встроиться, поэтому замена на 1 метод!!!
	}
}

func main() {
	producer()

	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			consumer()
		}()
	}

	wg.Wait()
}
