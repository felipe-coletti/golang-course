package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg    sync.WaitGroup
	count int
	mu    sync.Mutex
)

func goroutine(i int) {
	wg.Add(i)

	for j := 0; j < i; j++ {
		go func() {
			runtime.Gosched()

			mu.Lock()

			count++

			mu.Unlock()

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
