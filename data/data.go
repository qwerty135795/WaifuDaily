package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var db *sql.DB

func InitDatabase() {
	err := os.Remove("chat.db")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("chat.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	con, err := sql.Open("sqlite3", "chat.db")
	if err != nil {
		log.Fatal(err)
	}
	db = con
	_, err = db.Exec(`CREATE  TABLE Users (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Name TEXT NOT NULL ,
    Gender TEXT NOT NULL,
    Age integer,
    Waifu TEXT,
    Password BLOB,
    PasswordSalt BLOB
)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT into Users (Name, Gender, Age, Waifu) VALUES 
                                                 ('Kostya','Male',20,'Hoshino'),
                                                 ('Nastya','Female',21,'Ruby'),
                                                 ('Victory','Female',23,'Nastya')`)
	if err != nil {
		log.Fatal(err)
	}

}
func Close() {
	db.Close()
}
