/*
协程
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个无缓冲的channel
	ch := make(chan int)
	writec := (chan<- int)(ch)
	readc := (<-chan int)(ch)
	wg := new(sync.WaitGroup)
	go SetChannel(writec, 10, wg)
	for i := 0; i < 10; i++ {
		go GetChannel(readc, wg)
	}
	wg.Wait()
}


func SetChannel(ch chan<- int, num int, wg *sync.WaitGroup) {
	// 创建一个无缓冲的channel
	for i := 0; i < num; i++ {
		wg.Add(1)
		ch <- i
	}
}

func GetChannel(ch <-chan int, wg *sync.WaitGroup) {
	fmt.Println(<-ch)
	wg.Done()
}





