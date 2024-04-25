package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	taskCount := 10000

	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err == nil {
			taskCount = n
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < taskCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Second)
		}()
	}

	wg.Wait()
	fmt.Println("All tasks are finished.")
}
