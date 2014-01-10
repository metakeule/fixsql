package fixsql

import (
	"fmt"
)

type UnknownDriver string

func (u UnknownDriver) Error() string {
	return fmt.Sprintf("unknown driver \"%s\"", string(u))
}

type ConnectionError string

func (ce ConnectionError) Error() string {
	return fmt.Sprintf("connection error \"%s\"", string(ce))
}

type ConnectionClosed struct{}

func (ce ConnectionClosed) Error() string {
	return "sql: database is closed"
}

type InvalidStatement string

func (i InvalidStatement) Error() string {
	return fmt.Sprintf("invalid statement %#v", string(i))
}
