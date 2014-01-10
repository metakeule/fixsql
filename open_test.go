package fixsql

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"testing"
)

var pg_url = "postgres://docker:docker@172.17.0.2:5432/pgsqltest?schema=public"
var dataSourceName, _ = pq.ParseURL(pg_url)

func open() *sql.DB {
	return MustOpen("postgres", dataSourceName)
}

//var db = open()

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

	_, err = Open("postgres", dataSourceName+"x")

	if err == nil {
		t.Errorf("Open with an invalid dataSourceName should return an error")
	}

	_, ok = err.(ConnectionError)

	if !ok {
		t.Errorf("Open with an invalid dataSourceName should return an error of Type ConnectionError")
	}

	c, e := Open("postgres", dataSourceName)

	if e != nil {
		t.Errorf("Open with an valid dataSourceName should return no error, but we got: %s (%T)", e, e)
	}

	defer c.Close()
}
