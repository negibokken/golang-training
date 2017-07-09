package bank

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name   string
		args   args
		expect int
	}{
		{"test deposit", args{10}, 10},
		{"test deposit", args{20}, 20},
		{"test deposit", args{30}, 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Deposit(tt.args.amount)
			if b := Balance(); b != tt.expect {
				t.Errorf("Balance() = %v expect %v", b, tt.expect)
			}
			Withdraw(tt.args.amount)
		})
	}
}

func TestBalance(t *testing.T) {
	tests := []struct {
		name string
		bal  int
		want int
	}{
		{"test balance", 10, 10},
		{"test balance", 20, 20},
		{"test balance", 30, 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Deposit(tt.bal)
			if got := Balance(); got != tt.want {
				t.Errorf("Balance() = %v, want %v", got, tt.want)
			}
			Withdraw(tt.bal)
		})
	}
}

func TestWithdraw(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name   string
		bal    int
		args   args
		expect int
		want   bool
	}{
		{"test withdraw", 20, args{10}, 10, true},
		{"test withdraw", 30, args{10}, 20, true},
		{"test withdraw", 00, args{10}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Deposit(tt.bal)
			if got := Withdraw(tt.args.amount); got != tt.want {
				t.Errorf("Withdraw() = %v, want %v", got, tt.want)
			}
			if got := Balance(); got != tt.expect {
				t.Errorf("Balance() = %v, want %v", got, tt.expect)
			}
			Withdraw(tt.expect)
		})
	}
}
