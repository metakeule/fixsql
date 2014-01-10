package fixsql

import (
	"fmt"
	"testing"
)

func TestTransactionError(t *testing.T) {
	_ = fmt.Print
	db := open()
	defer db.Close()

	var secondCall bool

	f1 := func(d DB) error {
		return fmt.Errorf("transactionError")
	}

	f2 := func(d DB) error {
		secondCall = true
		return nil
	}

	err := Transaction(db, f1, f2)

	if err == nil {
		t.Errorf("Transaction with error should return error")
	}

	if err.Error() != "transactionError" {
		t.Errorf("Transaction with error should return \"transactionError\", but returns %#v", err.Error())
	}

	if secondCall {
		t.Errorf("second query should not be executed if first has error in transaction, but is executed")
	}

}

func TestTransactionSuccess(t *testing.T) {
	db := open()
	defer db.Close()

	var i, j int

	f1 := func(d DB) error {
		rows, e := Query(d, "Select 1")
		if e != nil {
			return e
		}

		f := func() []interface{} {
			return []interface{}{&i}
		}
		_, e = Each(rows, f)
		return e
	}

	f2 := func(d DB) error {
		rows, e := Query(d, "Select 2")
		if e != nil {
			return e
		}

		f := func() []interface{} {
			return []interface{}{&j}
		}
		_, e = Each(rows, f)
		return e
	}

	err := Transaction(db, f1, f2)

	if err != nil {
		t.Errorf("Transaction has error: %#v", err)
		return
	}

	if i != 1 {
		t.Errorf("i should be 1 but is: %d", i)
	}

	if j != 2 {
		t.Errorf("j should be 2 but is: %d", j)
	}
}
