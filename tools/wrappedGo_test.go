package tools

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func namedFunc_1() { panic(fmt.Errorf("named func with error content")) }
func namedFunc_2() { panic("named func with string content") }

func TestPrintPanicFunction(t *testing.T) {
	wg := sync.WaitGroup{}

	Go(func() {
		wg.Add(1)
		Go(func() { panic(fmt.Errorf("anonymous func with error content")) })
		t.Logf("goroutine sent\n")
		time.Sleep(time.Second)
		wg.Done()
	})
	Go(func() {
		wg.Add(1)
		Go(func() { panic("anonymous func with string content") })
		time.Sleep(time.Second)
		wg.Done()
	})

	Go(func() {
		wg.Add(1)
		Go(namedFunc_1)
		t.Logf("goroutine sent\n")
		time.Sleep(time.Second)
		wg.Done()
	})

	Go(func() {
		wg.Add(1)
		Go(namedFunc_2)
		t.Logf("goroutine sent\n")
		time.Sleep(time.Second)
		wg.Done()
	})

	testRecover := func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("panic captured")
			}
		}()
		t.Logf("send panic\n")
		panic("normal panic")
	}
	Go(func() {
		wg.Add(1)
		testRecover()
		t.Logf("panic recovered\n")
		time.Sleep(time.Second)
		wg.Done()
	})

	wg.Wait()
	time.Sleep(5 * time.Second)
	Log.Infof("happy endding\n")
}
