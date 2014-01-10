package fixsql

import (
	"database/sql"
)

/*
	Open runs database/sql.Open and returns it result.
	However the returned errors are typed and a test query is
	run, so that an error of type UnknownDriver is returned, if
	the given driverName is not registered and an error of type
	InvalidDataSource is returned, if a connection could not be established
*/
func Open(driverName, dataSourceName string) (db *sql.DB, err error) {
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		err = UnknownDriver(driverName)
		return
	}

	_, err = db.Exec("select 1")
	if err != nil {
		err = InvalidDataSource(err.Error())
	}
	return
}
