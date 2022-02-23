package main

import (
	"fmt"
	"sync"
)

func onetime() {
	fmt.Printf("one time\n")
}

func main() {
	o := sync.Once{}

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		i := i
		go func(key int) {
			o.Do(onetime)
			fmt.Printf("routine %d\n", key)
			wg.Done() // <=> wg.Add(-1)
		}(i)
	}

	wg.Wait()
}
