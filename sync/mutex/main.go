package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Data struct {
	data sync.Map
	// mutex sync.RWMutex
}

func (d *Data) ReadRandom() (int, interface{}, bool) {
	// d.mutex.RLock()
	// defer d.mutex.RUnlock()

	key := rand.Intn(100)
	val, ok := d.data.Load(key)
	return key, val, ok
}

func (d *Data) Write(i int) {
	// d.mutex.Lock()
	// defer d.mutex.Unlock()
	d.data.Store(i, i)
}

func writeLoop(d *Data) {
	for {
		for i := 0; i < 100; i++ {
			d.Write(i)
		}
	}
}

func ReadLoop(d *Data, name string) {
	for {
		key, val, ok := d.ReadRandom()
		fmt.Printf("%s: %d (%d) %v\n", name, key, val, ok)
	}
}

func main() {
	var d = &Data{
		// mutex: sync.RWMutex{},
		data: sync.Map{},
	}
	go writeLoop(d)
	go ReadLoop(d, "first")
	go ReadLoop(d, "second")

	block := make(chan struct{})
	<-block
}
