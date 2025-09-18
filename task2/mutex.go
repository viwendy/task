package main

import (
	"fmt"
	"sync"
)

//编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var count int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(goroutineId int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
			fmt.Printf("goroutine %d finished\n", goroutineId)
		}(i)
		wg.Wait()
	}
	fmt.Println(count)
}
