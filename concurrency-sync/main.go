package main

import (
	"fmt"
	"sync"
)

func main() {
	// Mutex example
	var count int
	var mu sync.Mutex
	increment := func() {
		mu.Lock()
		defer mu.Unlock()
		count++
		fmt.Println("Incremented count to", count)
	}

	// RWMutex example
	var data int
	var rwMu sync.RWMutex
	readData := func() {
		rwMu.RLock()
		defer rwMu.RUnlock()
		fmt.Println("Read data:", data)
	}
	writeData := func(d int) {
		rwMu.Lock()
		defer rwMu.Unlock()
		data = d
		fmt.Println("Wrote data:", data)
	}

	// WaitGroup example
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		increment()
	}()
	go func() {
		defer wg.Done()
		writeData(42)
	}()

	// Once example
	var once sync.Once
	initialize := func() {
		fmt.Println("Initialized only once")
	}
	doInitialize := func() {
		once.Do(initialize)
	}

	// Pool example
	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance")
			return "New instance"
		},
	}

	// Cond example
	cond := sync.NewCond(&sync.Mutex{})
	ready := false
	go func() {
		cond.L.Lock()
		for !ready {
			cond.Wait()
		}
		fmt.Println("Go signal to proceed")
		cond.L.Unlock()
	}()

	// Map example
	var m sync.Map
	m.Store("hello", "world")

	// Main logic
	go doInitialize()
	go doInitialize()
	instance := pool.Get()
	fmt.Println("Got from pool:", instance)
	pool.Put(instance)

	cond.L.Lock()
	ready = true
	cond.Broadcast()
	cond.L.Unlock()

	wg.Wait()
	readData()
	if val, ok := m.Load("hello"); ok {
		fmt.Println("Value from map:", val)
	}
}
