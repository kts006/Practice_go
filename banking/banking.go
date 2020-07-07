package banking

import (
	"errors"
	"fmt"
	"log"
)

// BankAccount struct
// type BankAccount struct {
// 	// 첫글자 대문자 -> Public,
// 	// Owner   string // Public struct의 value를 외부에서 지정하기 위해서 Key도 대문자로 시작
// 	owner   string
// 	balance int
// }

// BankingMain main method
func BankingMain() {
	account := NewAccount("AAAA")
	fmt.Println(account) // &{AAAA 0}

	account.Deposit(10)
	fmt.Println(account)

	fmt.Println(account.GetBalance())

	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.GetBalance())
}

// BankAccount account
type BankAccount struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("NO Money")

// NewAccount account maker
func NewAccount(owner string) *BankAccount {
	// *BankAccount -> struct 를 반환하기 위해서... 객체 반환
	account := BankAccount{owner: owner, balance: 0}

	return &account
}

// Deposit x amount on your account
func (a *BankAccount) Deposit(amount int) {
	// func (receiverName *Type)
	// * BankAccount 에 직접 접근해서 수정.. CallByReference
	// 없으면 받은 BankAccount의 복사본에 수정됨, 실제 수정 X CallByValue
	a.balance += amount
}

//GetBalance get yout Balance
func (a BankAccount) GetBalance() int {
	return a.balance
}

//Withdraw withdraw on your account
func (a *BankAccount) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
