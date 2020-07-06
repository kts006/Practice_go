package something

import "fmt"

func MemoryMain() {
	a := 2
	b := &a // b는 a의 메모리 위치를 가르키고 있다. b는 포인터
	fmt.Println(a, b)
	fmt.Println(&a, &b)
	fmt.Println(*b)
	a = 3
	fmt.Println(*b)
	*b = 100 // 현재 b는 포인터, b가 가르키는 곳의 값을 변경한다.
	fmt.Println(*b, a)

	// & 주소를 표현
	// * 해당주소의 값을 표현
}
