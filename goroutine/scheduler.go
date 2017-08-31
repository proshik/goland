package main

import (
	"runtime"
	"time"
	"fmt"
)

//never end
func main() {
	var x int
	threads := runtime.GOMAXPROCS(0)
	for i := 0; i < threads; i ++ {
		go func() {
			for {
				x++
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("x = ", x)
}
