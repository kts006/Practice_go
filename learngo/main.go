package main

import (
	"fmt"
	"strings"

	"github.com/kts006/learngo/something"
)

func main() {
	const msg string = "Hello" // Const 상수
	// msg = "ABCD"  안됨

	var var_msg string = "Say"
	// 축약형 var_msg := "Say", ':=' type을 찾아서 변수를 '생성'해줌, func 안에서만 사용가능)
	fmt.Println(msg, var_msg)
	something.SayHello()

	fmt.Println(multiply(11, 13, "a", "b"))
	print_str(var_msg, msg)

	nameLen, uppderName := lenAndUpper("patrick")
	fmt.Println(nameLen, uppderName)

	nameLen2, _ := lenAndUpper("patrick_abcd")
	fmt.Println(nameLen2)
	// 변수를 선언만 하고 사용하지 않을 경우 컴파일 에러발생
	// '_' 사용하면 무시가능

	repeatString("A", "BB", "CCC", "DDDD", "EEEEE")

	total_num := superAdd(1, 2, 3, 4, 5, 7, 8, 9, 10)
	fmt.Println("Total Num is :", total_num)
}

// - main.go 는 EntryPoint 컴파일이 필요할 경우 있어야 하는 파일, 다른 사람이 import 해서 쓰게 할거면 필요 없음.
// - 다른 패키지 혹은 모듈의 func을 쓰고 싶을 경우, func 이름의 첫 글자는 대문자로 시작 해야함.

// func func_name(args_name args_type) return_type {}
// 연속된 args의 타입이 같다면 앞에 타입은 무시해도 됨
func multiply(a, b int, c, d string) int {
	fmt.Println("You input ", a, b, c, d)
	return a * b
}

func print_str(a, b string) {
	fmt.Println(a, b)
}

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }
// 아래랑 같은 함수 같은 동작을 함....
func lenAndUpper(name string) (length int, upperName string) { // 리턴할 변수를 미리 생성
	defer print_str("After Done of", "lenAndUpper func") // 현재 func가 끝나고 난후 실행 되는 func
	length = len(name)                                   // 생성된 변수에 대입.
	upperName = strings.ToUpper(name)
	return
}

func repeatString(words ...string) { /// ... 해당 타입의 변수를 무한대로 받을 수 있음
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	// numbers -> array
	total := 0
	for idx, number := range numbers { // range -> loop를 만듬..
		fmt.Print(number)
		if idx < len(numbers)-1 {
			fmt.Print("+")
		} else {
			fmt.Print("\n")
		}

		total += number // total = total + number
	}
	// for i := 0; i < len(numbers): i++{
	// }
	return total
}
