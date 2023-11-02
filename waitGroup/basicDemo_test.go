package waitGroup

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
	"time"
)

func Test_waitGroup(t *testing.T) {

	logrus.SetOutput(os.Stdout)

	tasksNum := 10

	ch := make(chan int, tasksNum)

	for i := 0; i < tasksNum; i++ {
		go func() {
			defer func() {
				ch <- i
				logrus.WithFields(logrus.Fields{"value": i}).Info("input i")
			}()

			<-time.After(time.Second)
		}()
	}

	for i := 10; i < tasksNum; i++ {
		receive := <-ch

		t.Logf("receive %v", receive)
	}
}
