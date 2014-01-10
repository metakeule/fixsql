package fixsql

import (
	"testing"
)

func errString(err error) string {
	return err.Error()
}

func TestErrorStrings(t *testing.T) {

	testcases := map[error]string{
		UnknownDriver("xyz"):     `unknown driver "xyz"`,
		InvalidDataSource("xyz"): `invalid data source "xyz"`,
	}

	for err, expected := range testcases {
		eStr := errString(err)
		if eStr != expected {
			t.Errorf("wrong error string expected: %#v, got: %#v", expected, eStr)
		}
	}
}
