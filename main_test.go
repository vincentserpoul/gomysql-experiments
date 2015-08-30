package main

import (
	"testing"

	"github.com/jmoiron/sqlx"
)

func BenchmarkInsert(b *testing.B) {
	b.StopTimer()
	dbx := sqlx.MustConnect(
		"mysql",
		"root:dev@tcp(172.19.8.101:3306)/test",
	)

	_, _ = dbx.Exec("DROP TABLE IF EXISTS test_table_bin")
	_, _ = dbx.Exec("CREATE TABLE IF NOT EXISTS test_table_bin (`id` BINARY(16) NOT NULL, `time_created` INT(11) DEFAULT UNIX_TIMESTAMP(NOW()), PRIMARY KEY (`id`)) ENGINE = InnoDB")
	b.StartTimer()
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		InsertBin(dbx)
	}
}

func BenchmarkInsertIntAuto(b *testing.B) {
	b.StopTimer()
	dbx := sqlx.MustConnect(
		"mysql",
		"root:dev@tcp(172.19.8.101:3306)/test",
	)
	_, _ = dbx.Exec("DROP TABLE IF EXISTS test_table_int")
	_, _ = dbx.Exec("CREATE TABLE IF NOT EXISTS test_table_int (`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT, `time_created` INT(11) DEFAULT UNIX_TIMESTAMP(), PRIMARY KEY (`id`)) ENGINE = InnoDB AUTO_INCREMENT=9223372036854775807")
	b.StartTimer()
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		InsertIntAuto(dbx)
	}
}
