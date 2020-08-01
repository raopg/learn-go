package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertBTCEquals := func(t *testing.T, output, expected Bitcoin) {
		t.Helper()
		if expected != output {
			t.Errorf("Expected = %s\nOutput = %s", expected, output)
		}
	}

	assertError := func(t *testing.T, output error, expected string) {
		t.Helper()
		if output == nil {
			t.Fatal("Expected error, did not catch one")
		}

		if output.Error() != expected {
			t.Errorf("Expected = %s\nOutput = %s", expected, output)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		output := wallet.Balance()
		expected := Bitcoin(10)

		assertBTCEquals(t, output, expected)

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		output := wallet.Balance()
		expected := Bitcoin(10)

		assertBTCEquals(t, output, expected)

	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(30))

		output := wallet.Balance()

		assertBTCEquals(t, output, startingBalance)

		assertError(t, err, "Insufficient funds")

	})
}
