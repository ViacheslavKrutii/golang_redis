package database

import (
	"database/sql"
	"fmt"
	"log"
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

func (m *MysqlInviteHistoryState) CheckHistoryExist(p1 string) (bool, string) {

	querySelect := "SELECT history FROM inviteHistory WHERE Player = ?"

	var existingHistory string
	err := m.db.QueryRow(querySelect, p1).Scan(&existingHistory)
	switch {
	case err == sql.ErrNoRows:
		// Запис не існує
		return false, ""

	case err != nil:
		// Інша помилка при виконанні SELECT
		log.Println("Error executing SQL query:", err)
		return false, ""
	default:
		return true, existingHistory
	}
}

func (m *MysqlInviteHistoryState) WriteHistory(p1, p2 string) {
	msg := fmt.Sprintf("%v invite %v", p1, p2)
	querySelect := "SELECT history FROM inviteHistory WHERE Player = ?"
	queryInsert := "INSERT INTO inviteHistory (Player, history) VALUES (?, ?) ON DUPLICATE KEY UPDATE history = ?"

	var existingHistory string
	err := m.db.QueryRow(querySelect, p1).Scan(&existingHistory)
	switch {
	case err == sql.ErrNoRows:
		// Запис не існує, виконуємо INSERT
		_, err = m.db.Exec(queryInsert, p1, msg, msg)
		if err != nil {
			log.Println("Error executing SQL query:", err)
			return
		}
	case err != nil:
		// Інша помилка при виконанні SELECT
		log.Println("Error executing SQL query:", err)
		return
	default:
		// Запис існує, виконуємо UPDATE
		newHistory := existingHistory + ";" + msg
		_, err := m.db.Exec(queryInsert, p1, newHistory, newHistory)
		if err != nil {
			log.Println("Error executing SQL query:", err)
			return
		}
	}

	log.Printf("%s write history", p1)
}
