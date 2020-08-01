package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		output := wallet.Balance()
		expected := Bitcoin(10)

		assertBTCEquals(t, output, expected)

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		output := wallet.Balance()
		expected := Bitcoin(10)

		assertBTCEquals(t, output, expected)
		assertNoError(t, err)

	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(30))

		output := wallet.Balance()

		assertBTCEquals(t, output, startingBalance)

		assertError(t, err, ErrInsufficientFunds)

	})
}

func assertBTCEquals(t *testing.T, output, expected Bitcoin) {
	t.Helper()
	if expected != output {
		t.Errorf("Expected = %q\nOutput = %q", expected, output)
	}
}

func assertError(t *testing.T, output error, expected error) {
	t.Helper()
	if output == nil {
		t.Fatal("Expected error, did not catch one")
	}

	if output != expected {
		t.Errorf("Expected = %s\nOutput = %s", expected, output)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Did not expect error, but received one")
	}
}
