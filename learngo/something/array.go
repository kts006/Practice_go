package something

import "fmt"

func ArrayMain() {
	charArray := [5]string{"A", "B", "C", "D", "E"}
	//Array_name := [길이]타입{"내용"}
	fmt.Println(charArray)

	charSlice := []string{"a", "b", "c", "d", "e"}
	//길이를 입력하지 않음, Slice 길이제한 없음.
	fmt.Println(charSlice)
	charSlice = append(charSlice, "f") //append 주어진 값을 추가한 Slice를 반환
	fmt.Println(charSlice)

}
