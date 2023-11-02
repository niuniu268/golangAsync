package waitGroup

import (
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	taskNum := 10

	dataCh := make(chan interface{})
	resp := make([]interface{}, 0, taskNum)

	go func() {
		for data := range dataCh {
			resp = append(resp, data)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < taskNum; i++ {
		wg.Add(1)

		go func(ch chan<- interface{}) {
			defer wg.Done()
			ch <- time.Now().Nanosecond()
		}(dataCh)
	}

	wg.Wait()
	close(dataCh)

	t.Logf("resp: %v", resp)

}
