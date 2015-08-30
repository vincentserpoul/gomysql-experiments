package main

import (
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

func main() {

	db := sqlx.MustConnect(
		"mysql",
		"root:dev@tcp(172.19.8.101:3306)/test?parseTime=true",
	)

	// testuuid := Insert(db)

	// Select(db, testuuid.String())
	tetuint := uint64(9223372036854775826)
	SelectuInt(db, tetuint)

}

type UserTest struct {
	ID uuid.UUID
}

func InsertBin(dbx *sqlx.DB) uuid.UUID {
	// Creating UUID Version 4
	u1 := uuid.NewV4()

	_, err := dbx.Exec("INSERT INTO test_table_bin (id) VALUES (?)", u1.Bytes())

	if err != nil {
		fmt.Printf("%v", err)
	}

	return u1
}

func InsertIntAuto(dbx *sqlx.DB) uint64 {

	_, _ = dbx.Exec("INSERT INTO test_table_int() values()")

	row := dbx.QueryRow("SELECT MAX(id) FROM test_table_int")

	var testuint uint64
	// var testuintX int64
	_ = row.Scan(&testuint)
	// testuintX, _ = res.LastInsertId()

	return testuint

}

func Select(dbx *sqlx.DB, testuuid string) {
	uuid, _ := uuid.FromString(testuuid)
	var user UserTest
	row := dbx.QueryRow("SELECT id from test_table WHERE ID = ?", uuid.Bytes())

	err := row.Scan(&user.ID)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v", user.ID.String())
}

func SelectuInt(dbx *sqlx.DB, testuint uint64) {

	type userS struct {
		ID          uint64
		TimeCreated time.Time `db:"time_created"`
	}

	rows, err := dbx.Queryx("SELECT id, time_created from test_table_int WHERE ID=" + strconv.FormatUint(testuint, 10))

	for rows.Next() {
		var user userS
		err = rows.StructScan(&user)
		idstr := strconv.FormatUint(user.ID, 10)
		fmt.Printf("%#v", user.TimeCreated.Location())
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%#v", idstr)
	}

}
