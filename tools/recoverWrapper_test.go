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
func _namedFunc_2() {
	var arr []int
	fmt.Printf("%d", arr[0])
}

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

	wg.Wait()
	t.Logf("happy endding\n")
}
