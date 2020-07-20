package main

import (
	"fmt"
	"time"
)

func main() {
	// go countFunc("B")
	// countFunc("A") // 병렬적으로 동작을 하지만, 프로그램이 돌아가는 동안만 main에서 돌아가는 것이 없으면 돌아가지 않음

	channel := make(chan bool) // Goroutine의 Return값을 받을 수 있는 채널.
	people := [2]string{"AAA", "BBB"}

	for _, person := range people {
		go isSexy(person, channel)
	}
	result := <-channel // 1개의 Goroutine에서 보낸것만 받음
	fmt.Println(result)

}

// func countFunc(person string) {
// 	i := 0
// 	if person == "A" {
// 		i = 5
// 	}

// 	for i := i; i < 10; i++ {
// 		fmt.Println(person, "Person, ", i)
// 		time.Sleep(time.Second)
// 	}
// }

func isSexy(person string, ch chan bool) {
	time.Sleep(time.Second * 5)
	ch <- true
}
