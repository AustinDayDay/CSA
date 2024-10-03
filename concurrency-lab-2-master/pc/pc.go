package main

import (
	"fmt"
	"github.com/ChrisGora/semaphore"
	"sync"

	//"github.com/ChrisGora/semaphore"
	"math/rand"
	"time"
)

type buffer struct {
	b                 []int
	size, read, write int
}

func newBuffer(size int) buffer {
	return buffer{
		b:     make([]int, size),
		size:  size,
		read:  0,
		write: 0,
	}
}

func (buffer *buffer) get() int {
	x := buffer.b[buffer.read]
	fmt.Println("Get\t", x, "\t", buffer)
	buffer.read = (buffer.read + 1) % len(buffer.b)
	return x
}

func (buffer *buffer) put(x int) {

	buffer.b[buffer.write] = x
	fmt.Println("Put\t", x, "\t", buffer)
	buffer.write = (buffer.write + 1) % len(buffer.b)
}

func producer(buffer *buffer, spaceAvailable semaphore.Semaphore, workAvailable semaphore.Semaphore, lock *sync.Mutex, start, delta int) {
	x := start
	for {
		spaceAvailable.Wait()
		lock.Lock()

		buffer.put(x)
		x = x + delta
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		workAvailable.Post()
		lock.Unlock()
	}
}

func consumer(buffer *buffer, spaceAvailable semaphore.Semaphore, workAvailable semaphore.Semaphore, lock *sync.Mutex) {
	for {
		workAvailable.Wait()
		lock.Lock()

		_ = buffer.get()
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		spaceAvailable.Post()
		lock.Unlock()
	}
}

func main() {
	buffer := newBuffer(5)

	lock := new(sync.Mutex)

	spaceAvailable := semaphore.Init(5, 5)
	workAvailable := semaphore.Init(5, 0)

	go producer(&buffer, spaceAvailable, workAvailable, lock, 1, 1)
	go producer(&buffer, spaceAvailable, workAvailable, lock, 1000, -1)

	consumer(&buffer, spaceAvailable, workAvailable, lock)
}
