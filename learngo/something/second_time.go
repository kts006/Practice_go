package something

import "fmt"

func SecondMain() {
	fmt.Println(CanIDrink(18))
}
func CanIDrink(age int) bool {
	switch koreaAge := age + 2; koreaAge {
	case 18:
		return true
	case 16:
		return false
	}

	// #case1
	//switch {
	//case age < 18:
	//	return true
	//}
	// #case2
	//switch age{
	//case 18:
	//	return true
	//}

	return false
}
