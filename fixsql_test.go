package fixsql

import (
	"fmt"
	"github.com/lib/pq"
	"testing"
)

func TestOpen(t *testing.T) {
	_ = fmt.Printf
	_, err := Open("xyz", "p")

	if err == nil {
		t.Errorf("Open with an unknown driver should return an error")
	}

	_, ok := err.(UnknownDriver)

	if !ok {
		t.Errorf("Open with an unknown driver should return an error of Type UnknownDriver")
	}

	pg_url := "postgres://docker:docker@172.17.0.2:5432/pgsqltest?schema=public"

	var dataSourceName string

	dataSourceName, err = pq.ParseURL(pg_url)
	if err != nil {
		panic(err.Error())
	}

	_, err = Open("postgres", dataSourceName+"x")

	if err == nil {
		t.Errorf("Open with an invalid dataSourceName should return an error")
	}

	_, ok = err.(InvalidDataSource)

	if !ok {
		t.Errorf("Open with an invalid dataSourceName should return an error of Type InvalidDataSource")
	}

	_, err = Open("postgres", dataSourceName)

	if err != nil {
		t.Errorf("Open with an valid dataSourceName should return no error, but we got: %s (%T)", err, err)
	}

}
