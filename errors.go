package fixsql

import (
	"fmt"
)

type UnknownDriver string

func (u UnknownDriver) Error() string {
	return fmt.Sprintf("unknown driver \"%s\"", string(u))
}

type InvalidDataSource string

func (i InvalidDataSource) Error() string {
	return fmt.Sprintf("invalid data source \"%s\"", string(i))
}
