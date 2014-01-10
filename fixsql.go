package fixsql

import (
	"database/sql"
	"fmt"
)

/*
	Open runs database/sql.Open and returns it result.
	The returned errors are typed and a Ping is
	run, so that an error of type UnknownDriver is returned, if
	the given driverName is not registered and an error of type
	ConnectionError is returned, if a connection could not be established (Ping fails)
*/
func Open(driverName, dataSourceName string) (db *sql.DB, err error) {
	db, err = sql.Open(driverName, dataSourceName)

	if err != nil {
		err = UnknownDriver(driverName)
		return
	}

	err = db.Ping()
	if err != nil {
		err = ConnectionError(err.Error())
	}
	return
}

func MustOpen(driverName, dataSourceName string) (db *sql.DB) {
	var err error
	db, err = Open(driverName, dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("Error %T: %s", err, err))
	}
	return
}

func interpretError(in error) (out error) {
	if in == nil {
		return nil
	}
	if in.Error() == "sql: database is closed" {
		return ConnectionClosed{}
	}
	return InvalidStatement(in.Error())
}

/*
	runs *database/sql.DB.Exec() and returns the result
	The returned errors are typed, so that an error caused by a closed
	database returns an error of type ConnectionClosed and every other
	error is of type InvalidStatement
*/
func Exec(db *sql.DB, query string, args ...interface{}) (res sql.Result, err error) {
	res, err = db.Exec(query, args...)
	err = interpretError(err)
	return
}

/*
	runs *database/sql.DB.Query() and returns the result
	The returned errors are typed, so that an error caused by a closed
	database returns an error of type ConnectionClosed and every other
	error is of type InvalidStatement
*/
func Query(db *sql.DB, query string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = db.Query(query, args...)
	err = interpretError(err)
	return
}

/*
	runs *database/sql.DB.Prepare() and returns the result
	The returned errors are typed, so that an error caused by a closed
	database returns an error of type ConnectionClosed and every other
	error is of type InvalidStatement
*/
func Prepare(db *sql.DB, query string) (st *sql.Stmt, err error) {
	st, err = db.Prepare(query)
	err = interpretError(err)
	return
}

/*
	Each scans through all rows and calls fn to get the destinations
	It stopps on the first error, returning the number of succellfully scanned rows and
	the first error.
	Each makes sure the given rows are closed, so that there is no leakage
*/
func Each(rows *sql.Rows, fn func() (dest []interface{})) (num int, err error) {
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(fn()...)
		if err != nil {
			return
		}
		num++
	}
	return
}
