package main

import "fmt"

type BankAccount struct {
	Owner   *string
	Balance float64
	like    map[string]string
}

func checkBankAccount(account BankAccount) {
	fmt.Println("Owner:", *account.Owner)
	fmt.Println("Balance:", account.Balance)
	fmt.Println("Like:", account.like)
}

func (account *BankAccount) addFunds(amount float64) {
	account.Balance += amount

}

func main() {
	name := "Bob"
	newMap := make(map[string]string)
	newMap["favoriteColor"] = "blue"
	newMap["favoriteFood"] = "pizza"

	account := BankAccount{Owner: &name, Balance: 100.0, like: newMap}
	checkBankAccount(account)
	account.addFunds(50.0)
	checkBankAccount(account)

	var emptyAcc *BankAccount
	emptyAcc = new(BankAccount)
	emptyAcc.Owner = &name
	emptyAcc.Balance = 200.0
	emptyAcc.like = make(map[string]string)
	emptyAcc.addFunds(100.0)
	checkBankAccount(*emptyAcc)

}
