package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 100)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Println("发送:", i)
			time.Sleep(time.Millisecond * 100)
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println("接收:", num)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	wg.Wait()
}
