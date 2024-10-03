package main

import (
	"fmt"
	"sync"
)

//Using mutex locks
//func main() {
//	sum := 0
//	var wg sync.WaitGroup
//	var lock sync.Mutex
//	for i := 0; i < 1000; i++ {
//		wg.Add(1)
//		go func() {
//			lock.Lock()
//			defer lock.Unlock()
//			sum = sum + 1
//			wg.Done()
//		}()
//	}
//
//	wg.Wait()
//	fmt.Println("Sum is: ", sum)
//}

//Using channels
func main() {
	sum := 0
	var wg sync.WaitGroup
	channel := make(chan int)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(sumChan chan int) {
			sumChan <- 1
			wg.Done()
		}(channel)
		sum += <-channel
	}

	wg.Wait()

	fmt.Println("Sum is: ", sum)
}
