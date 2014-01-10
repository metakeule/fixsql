package fixsql

import (
	"fmt"
	"testing"
)

func TestExecClosed(t *testing.T) {
	db := open()
	_ = fmt.Print
	db.Close()
	_, err := Exec(db, "select 1")

	if err == nil {
		t.Errorf("Exec on a closed connection should return an error")
	}

	_, ok := err.(ConnectionClosed)

	if !ok {
		t.Errorf("Exec on a closed connection should return an error of Type ConnectionClosed")
	}

}

func TestExecSyntax(t *testing.T) {
	db := open()
	defer db.Close()
	_, err := Exec(db, "sel ect 1")

	if err == nil {
		t.Errorf("Exec with invalid syntax should return an error")
	}

	_, ok := err.(InvalidStatement)

	if !ok {
		t.Errorf("Exec with invalid syntax should return an error of Type InvalidStatement")
	}
}
