package wallet

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	output := wallet.Balance()
	expected := 10

	if expected != output {
		t.Errorf("Expected = %d\nOutput = %d", expected, output)
	}
}
