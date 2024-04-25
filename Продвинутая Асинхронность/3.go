package main

import (
	"root/util"
	"sync"
	"sync/atomic"
)

func main() {
	util.Pench(Atomic)
	util.Pench(Mutex)
	util.Pench(ChanMutex)
}

func Atomic() {
	var (
		wg      sync.WaitGroup
		counter int64
	)
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
}

func Mutex() {
	var (
		counter int64
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
}

func ChanMutex() {
	var (
		counter int64
		wg      sync.WaitGroup
		mu      sync.Mutex
		done    = make(chan struct{})
	)

	wg.Add(1000)
	go func() {
		wg.Wait()
		close(done)
	}()

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	<-done
}
