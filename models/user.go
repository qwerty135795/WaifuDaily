package models

type User struct {
	Id           int
	Name         string
	Gender       string
	Age          int
	Waifu        string
	Password     []byte
	PasswordSalt []byte
}
