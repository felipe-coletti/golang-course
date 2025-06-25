package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	wg    sync.WaitGroup
	count int32
)

func goroutine(i int) {
	wg.Add(i)

	for j := 0; j < i; j++ {
		go func() {
			runtime.Gosched()

			atomic.AddInt32(&count, 1)

			wg.Done()
		}()
	}
}
func main() {
	const routinesNumber = 3

	goroutine(routinesNumber)

	wg.Wait()

	fmt.Printf("Número de rotinas: %v\nValor no contador: %v\n", routinesNumber, count)
}
