package main

import (
	"fmt"

	"github.com/kts006/learngo/something"
)

func main() {
	fmt.Printf("hello, world\n")
	something.SayHello()
}

/*
- main.go 는 EntryPoint 컴파일이 필요할 경우 있어야 하는 파일, 다른 사람이 import 해서 쓰게 할거면 필요 없음.
- 다른 패키지 혹은 모듈의 func을 쓰고 싶을 경우, func 이름의 첫 글자는 대문자로 시작 해야함.
*/
