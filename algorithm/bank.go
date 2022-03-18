package main

func main() {

}

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance: balance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.checkAccount(account1) || !this.checkAccount(account2) || this.balance[account1-1] < money {
		return false
	}

	this.balance[account1-1] -= money
	this.balance[account2-1] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.checkAccount(account) {
		return false
	}

	this.balance[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.checkAccount(account) || this.balance[account-1] < money {
		return false
	}

	this.balance[account-1] -= money
	return true
}

func (this *Bank) checkAccount(account int) bool {
	return account >= 1 && account <= len(this.balance)
}
