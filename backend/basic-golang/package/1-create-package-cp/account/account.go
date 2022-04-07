// TODO: answer here
package account

type Account struct {
	Name    string
	Balance int
}

func (a Account) GetBalance() int {
	// TODO: answer here
	return a.Balance
}

func (a *Account) Deposit(amount int) {
	// TODO: answer here
	if amount > 0 {
		a.Balance += amount
	}
	return
}
