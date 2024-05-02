package data

import (
	"Chat/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func CreateUser(user models.User) (int, error) {
	res, err := db.Exec(`INSERT  INTO Users (Name, Gender, Age, Waifu, Password, PasswordSalt)
     VALUES ($1, $2, $3, $4, $5, $6)`,
		user.Name, user.Gender, user.Age, user.Waifu, user.Password, user.PasswordSalt)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return int(id), nil

}

func GetUserById(id int) (*models.User, error) {
	var user models.User
	row, err := db.Query(`SELECT * FROM Users WHERE Id = ${id}`)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		err = row.Scan(&user.Id, &user.Name, &user.Gender, &user.Age, &user.Waifu, &user.Password, &user.PasswordSalt)

		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
func GetUserByName(name string) (*models.User, error) {
	var user models.User
	row, err := db.Query(`SELECT * FROM Users WHERE Name = $1`, name)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		err = row.Scan(&user.Id, &user.Name, &user.Gender, &user.Age, &user.Waifu, &user.Password, &user.PasswordSalt)

		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
