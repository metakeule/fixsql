package fixsql

import (
	"fmt"
	"testing"
)

func TestPrepareClosed(t *testing.T) {
	_ = fmt.Print
	db := open()
	db.Close()
	_, err := Prepare(db, "select 1")

	if err == nil {
		t.Errorf("Prepare on a closed connection should return an error")
	}

	_, ok := err.(ConnectionClosed)

	if !ok {
		t.Errorf("Prepare on a closed connection should return an error of Type ConnectionClosed")
	}

}

func TestPrepareSyntax(t *testing.T) {
	db := open()
	defer db.Close()
	_, err := Prepare(db, "sel ect 1")

	if err == nil {
		t.Errorf("Prepare with invalid syntax should return an error")
	}

	_, ok := err.(InvalidStatement)

	if !ok {
		t.Errorf("Prepare with invalid syntax should return an error of Type InvalidStatement")
	}
}
