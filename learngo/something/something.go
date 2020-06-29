package something

import "fmt"

func sayBuy() { //private func
	fmt.Println("Buy")
}

func SayHello() { //public func
	fmt.Println("Hello")
	sayBuy()
}
