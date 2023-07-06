package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printName(name string, times int, startChan, nextChan chan struct{}) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		<-startChan
		fmt.Println(name, i)
		nextChan <- struct{}{}
	}

}

func main() {
	nameList := []string{"张三", "李四", "王五", "赵六"}
	workTimes := 100
	head := make(chan struct{}, 1)
	startChan := head
	nextChan := head
	for index, name := range nameList {
		wg.Add(1)
		startChan = nextChan
		if index == len(nameList)-1 {
			nextChan = head
		} else {
			nextChan = make(chan struct{})
		}
		go printName(name, workTimes, startChan, nextChan)
	}
	head <- struct{}{}
	wg.Wait()
}
