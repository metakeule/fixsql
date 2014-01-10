package fixsql

import (
	"fmt"
	"testing"
)

func TestQueryClosed(t *testing.T) {
	_ = fmt.Print
	db := open()
	db.Close()
	_, err := Query(db, "select 1")

	if err == nil {
		t.Errorf("Query on a closed connection should return an error")
	}

	_, ok := err.(ConnectionClosed)

	if !ok {
		t.Errorf("Query on a closed connection should return an error of Type ConnectionClosed")
	}

}

func TestQuerySyntax(t *testing.T) {
	db := open()
	defer db.Close()
	_, err := Query(db, "sel ect 1")

	if err == nil {
		t.Errorf("Query with invalid syntax should return an error")
	}

	_, ok := err.(InvalidStatement)

	if !ok {
		t.Errorf("Query with invalid syntax should return an error of Type InvalidStatement")
	}
}

func TestEach(t *testing.T) {
	db := open()
	defer db.Close()
	rows, _ := Query(db, "select 1, 'one'")

	var i int
	var s string

	fn := func() []interface{} { return []interface{}{&i, &s} }

	num, err := Each(rows, fn)

	if err != nil {
		t.Errorf("valid Query should not return an error, but returns %#v", err.Error())
	}

	if num != 1 {
		t.Errorf("Query should return 1 row, but returns %d", num)
	}

	if i != 1 {
		t.Errorf("i should be 1, but is %d", i)
	}

	if s != "one" {
		t.Errorf("s should be `one`, but is %#v", s)
	}
}

func TestEachError(t *testing.T) {
	db := open()
	defer db.Close()
	rows, _ := Query(db, "select 1, 'one'")

	var i int
	var s int

	fn := func() []interface{} { return []interface{}{&i, &s} }

	num, err := Each(rows, fn)

	if err == nil {
		t.Errorf("invalid Scan should return an error")
	}

	_, ok := err.(ScanError)

	if !ok {
		t.Errorf("invalid Scan should return an error of Type ScanError")
	}

	if num != 0 {
		t.Errorf("invalid Scan should return 1 row, but returns %d", num)
	}
}
