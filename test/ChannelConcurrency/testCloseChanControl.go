package main

import (
	"fmt"
	"time"
)

func worker(cancel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
			// 正常工作
		case <-cancel:
			// 退出
		}
	}
}

func testCloseChanControl() {
	cancel := make(chan bool)

	for i := 0; i < 10; i++ {
		go worker(cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
	time.Sleep(time.Second)
}
