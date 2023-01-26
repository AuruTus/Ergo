package tools

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func _goRecover[F WrappedFunc](f F) {
	go func() {
		if err := WithRecover(f); err != nil {
			fmt.Printf("%v\n", err)
		}
	}()
}

func _namedFunc_1() { panic(fmt.Errorf("named func with error content")) }
func _namedFunc_2() { panic("named func with string content") }

func TestPrintPanicFunction(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		_goRecover(func() { panic(fmt.Errorf("anonymous func with error content")) })
		t.Logf("goroutine sent\n")
		time.Sleep(time.Second)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		_goRecover(func() { panic("anonymous func with string content") })
		time.Sleep(time.Second)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		_goRecover(_namedFunc_1)
		t.Logf("goroutine sent\n")
		time.Sleep(time.Second)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		_goRecover(_namedFunc_2)
		t.Logf("goroutine sent\n")
		time.Sleep(time.Second)
		wg.Done()
	}()

	testRecover := func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("panic captured")
			}
		}()
		t.Logf("send panic\n")
		panic("normal panic")
	}
	wg.Add(1)
	go func() {
		testRecover()
		t.Logf("panic recovered\n")
		time.Sleep(time.Second)
		wg.Done()
	}()

	wg.Wait()
	Log.Infof("happy endding\n")
}
