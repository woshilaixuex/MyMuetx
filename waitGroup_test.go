package myMuetx

import (
	"fmt"
	"testing"
)

func TestCountCounter(t *testing.T) {
	var wt WaitGroup
	wt.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wt.Done()
			x := wt.CountCounter()
			y := wt.CountWaiter()
			fmt.Println(x, y)
		}()
	}
	wt.Wait()
}
