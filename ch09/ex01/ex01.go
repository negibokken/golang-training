package bank

type withdraw struct {
	withdraw int
	result   chan bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan *withdraw)

// Deposit is add deposit
func Deposit(amount int) { deposits <- amount }

// Balance returns your balance
func Balance() int { return <-balances }

// Withdraw returns the results
func Withdraw(amount int) bool {
	res := make(chan bool)
	w := withdraw{amount, res}
	withdraws <- &w
	return <-res
}

func teller() {
	var balance int // closed variable
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case with := <-withdraws:
			if balance >= with.withdraw {
				balance -= with.withdraw
				with.result <- true
			} else {
				with.result <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
