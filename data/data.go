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

	_, err = db.Exec(` CREATE TABLE Messages (
	 Id INTEGER PRIMARY KEY AUTOINCREMENT,
	 Text TEXT,
	 MessageDate TEXT,
	 SenderId INTEGER,
	 ReceiverId INTEGER,
	 Foreign Key (ReceiverId) REFERENCES Users(Id)
	 FOREIGN KEY (SenderId) REFERENCES Users(Id)
	 );`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(` CREATE Table UserMessages (
	 Id INTEGER PRIMARY KEY AUTOINCREMENT ,
	UserId,
	MessageId,
	Foreign Key (UserId) REFERENCES Users(Id),
	FOREIGN KEY (MessageId) REFERENCES Messages(Id)
	 );`)
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
