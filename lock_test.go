package myMuetx

import (
	"fmt"
	"testing"
	"time"
)

func TestTryLock(t *testing.T) {
	var mu Mutex
	go func() {
		mu.Lock()
		time.Sleep(3 * time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)
	ok := mu.TryLock()
	if ok {
		fmt.Println("got the lock")
		mu.Unlock()
		return
	}
	fmt.Println("can`t get the lock")
}
