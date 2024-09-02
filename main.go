package main

import "e-dars/internals/db"

func main() {

	err := db.ConnectToDb()

	defer func() {
		err := db.CloseDbConnection()
		if err != nil {

		}
	}()

	err = db.MigrateTables()
	if err != nil {
		panic(err)
	}
}
