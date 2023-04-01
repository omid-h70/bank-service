package domain

import "testing"

func Test_should_return_error_when_amount_is_not_in_rules_range(t *testing.T) {
	testRule := AccountRule{
		MinAmount: 1000,
		MaxAmount: 500000,
	}
	var amount int64 = 100

	Transaction := NewTransaction()
	err := Transaction.isTransactionValid(amount, testRule)
	if err != ErrInvalidMinTransactionAmount {
		t.Error("ErrInvalidMinTransactionAmount Test Failed")
	}

	amount = 500000 + 1
	err = Transaction.isTransactionValid(amount, testRule)
	if err != ErrInvalidMaxTransactionAmount {
		t.Error("ErrInvalidMinTransactionAmount Test Failed")
	}
}
