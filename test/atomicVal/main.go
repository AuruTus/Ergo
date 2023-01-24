package main

import (
	"fmt"
	"sync/atomic"
)

type valS struct {
	val uint32
}

func main() {
	v := valS{}
	fmt.Printf("val: %d\n", atomic.LoadUint32(&v.val))
	atomic.StoreUint32(&v.val, 12)
	fmt.Printf("val: %d\n", atomic.LoadUint32(&v.val))
}
