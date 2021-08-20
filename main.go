package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

const size = 100000000

func main() {

	for i := 20; i > 0; i-- {

		debug.SetGCPercent(-1)

		src := make([]int, size)

		for idx := range src {
			src[idx] = 42
		}

		fmt.Println("before    ", src[:4], src[size-4:])
		t1 := time.Now()
		multiply2(src, i)
		dur := time.Since(t1)
		fmt.Println("after ", src[:4], src[size-4:])
		fmt.Printf("goroutine's number: %d, time: %s\n", i, dur)
		fmt.Println()

		runtime.GC()

	}
}

func multiply2(src []int, workers int) {

	if workers <= 0 {
		return
	}
	wg := new(sync.WaitGroup)

	last := 0

	for i := 0; i < workers; i++ {

		idx1 := len(src) / workers * i
		idx2 := len(src) / workers * (i + 1)
		last = idx2

		wg.Add(1)
		go func(idx1, idx2 int) {
			defer wg.Done()
			for idx := idx1; idx < idx2; idx++ {
				src[idx] = src[idx] * 2
			}
		}(idx1, idx2)

	}

	if last < len(src)-1 {
		idx1 := last
		idx2 := len(src)
		wg.Add(1)
		go func(idx1, idx2 int) {
			defer wg.Done()
			for idx := idx1; idx < idx2; idx++ {
				src[idx] = src[idx] * 2
			}
		}(idx1, idx2)
	}

	wg.Wait()

}

