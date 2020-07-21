package main

import (
	"fmt"
	"time"
)

func main() {
	// go countFunc("B")
	// countFunc("A") // 병렬적으로 동작을 하지만, 프로그램이 돌아가는 동안만 main에서 돌아가는 것이 없으면 돌아가지 않음

	channel := make(chan string) // Goroutine의 Return값을 받을 수 있는 채널, 전달 받을 Return type 지정필요
	people := [2]string{"AAA", "BBB"}

	for _, person := range people {
		go isSexy(person, channel) // goroutine에 Channel 전달, 리턴값 전달 받음
	}
	// result := <-channel // 1개의 Goroutine에서 보낸것만 받음

	for i := 0; i < len(people); i++ {
		fmt.Println("Received this message.", <-channel) //<-channel : Blocked message, 1개만 전달 받음
	}
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

func isSexy(person string, ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- person + " ABCDE" // 전달받은 채널로 값 전달
}
