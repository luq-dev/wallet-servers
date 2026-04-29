package main

import "fmt"

type Account struct {
	number  int
	name    string
	book    []Transaction
	balance int
}

type Transaction struct {
	from   string
	to     string
	amount int
}

func main() {
	var amount int

	account1 := Account{number: 9876543, name: "Lady1", balance: 10000}
	account2 := Account{number: 21224, name: "Lady2", balance: 1223}

	fmt.Printf("Logged in as %s\n\n", account1.name)

	for true {
		fmt.Printf("Transfering to %s\n\n", account2.name)
		fmt.Printf("Amount: ")
		fmt.Scanf("%d", &amount)
		fmt.Println()

		transfer(&account1, &account2, amount)

		fmt.Printf("Current Balance: %d\n\n", account1.balance)
		printBook(account1)
		fmt.Println()

		fmt.Println("Restarting Console..\n")
	}

}

func transfer(from *Account, to *Account, amount int) {
	if amount < from.balance {
		from.book = append(from.book, Transaction{from: from.name, to: to.name, amount: -amount})
		from.balance -= amount

		to.book = append(to.book, Transaction{from: from.name, to: to.name, amount: amount})
		to.balance += amount

	} else {
		fmt.Println("Insufficient balance\n")
	}
}

func printBook(acc Account) {
	for i := 0; i < len(acc.book); i++ {
		fmt.Println(acc.book[i])
	}
}
