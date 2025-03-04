package main

import (
	"fmt"
	"sync"
	"time"
)

// The call to Wait does the following under the hood
// 1. Calls Cond Unlock on the condition Locker
// 2. Notifies the list wait
// 3. Calls Cond Lock on the condition Locker

// The Cond type besides the Locker also has access to 2 important methods:
// 1. Signal - wakes up 1 go routine waiting on a condition (rendezvous point)
// 2. Broadcast - wakes up all go routines waiting on a condition (rendezvous point)
func main() {
	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})

	wg.Add(2)
	go func() {
		defer wg.Done()
		cond.L.Lock()
		defer cond.L.Unlock()
		cond.Wait()
		fmt.Println("go routine 1")
	}()
	go func() {
		defer wg.Done()
		cond.L.Lock()
		defer cond.L.Unlock()
		cond.Wait()
		fmt.Println("go routine 2")
	}()

	time.Sleep(time.Second)
	cond.Broadcast()

	wg.Wait()
}
