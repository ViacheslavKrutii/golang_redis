package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlInviteHistoryState struct {
	db *sql.DB
}

func NewMysqlConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/inviteHistory")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func NewMysql() *MysqlInviteHistoryState {
	return &MysqlInviteHistoryState{
		NewMysqlConnection(),
	}
}

func (m *MysqlInviteHistoryState) WriteHistory(p1, p2 string) {

}
