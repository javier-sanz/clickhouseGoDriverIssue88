package main

import (
	"database/sql"
	"time"

	_ "github.com/kshvakov/clickhouse"
)

func createTestTable(dbConn *sql.DB) {
	_, err := dbConn.Exec(`CREATE TABLE IF NOT EXISTS default.test
		(
		   EventDate Date,
		   Value  UInt64
		)
		ENGINE = MergeTree(EventDate, EventDate, 8192);`)
	if err != nil {
		panic("Could not create default table " + err.Error())
	}
}

func main() {

	dbConn, err := sql.Open("clickhouse", "tcp://localhost:9000?username=default")
	defer dbConn.Close()
	if err != nil {
		panic("Could not connect with clickhouse db")
	}

	createTestTable(dbConn)
	transaction, err := dbConn.Begin()
	if err != nil {
		panic("Error creating transaction: " + err.Error())
	}
	statement, err := transaction.Prepare("INSERT INTO default.test " +
		" (EventDate, Value) VALUES (?, ?)")
	if err != nil {
		panic("Error preparing insert: " + err.Error())
	}
	_, err = statement.Exec(time.Now(), "123")
	//_, err = statement.Exec(time.Now(), 123)

	err = transaction.Commit()
	if err != nil {
		panic("Error committing: " + err.Error())
	}

}
