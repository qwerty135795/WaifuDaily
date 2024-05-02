package data

import (
	. "Chat/models"
	"log"
)

func NewMessage(message Message) (int, error) {
	res, err := db.Exec(`INSERT INTO Messages (Text, SenderId, ReceiverId) VALUES ($1,$2,$3)`,
		message.Text, message.SenderId, message.ReceiverId)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	_, err = db.Exec(`INSERT INTO UserMessages (UserId, MessageId) VALUES ($1,$2), ($3,$2)`,
		message.SenderId, id, message.ReceiverId)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return int(id), nil
}

func GetMessages(userId int) (*[]Message, error) {
	row, err := db.Query(`SELECT Messages.Id, Text, SenderId FROM UserMessages 
    JOIN Messages On UserMessages.MessageId = Messages.Id
	
	`, userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var messages []Message
	for row.Next() {
		var message Message
		err = row.Scan(&message.Id, &message.Text, &message.SenderId, &message.ReceiverId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		messages = append(messages, message)
	}
	return &messages, nil
}
func GetDialog(userId, receiverId int) ([]Message, error) {
	row, err := db.Query(`SELECT Messages.Id As MessageId, Text, SenderId, ReceiverId FROM UserMessages 
    JOIN Messages On UserMessages.MessageId = Messages.Id
	WHERE ReceiverId  In ($1, $2) And SenderId In ($1,$2) And ($1 = $2 and ReceiverId = SenderId) or ($1 != $2 and ReceiverId != Messages.SenderId)
	GROUP BY MessageId
	`, userId, receiverId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var messages []Message
	for row.Next() {
		var message Message
		err = row.Scan(&message.Id, &message.Text, &message.SenderId, &message.ReceiverId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
func GetMessageById(id int) (*Message, error) {
	row, err := db.Query(`SELECT * FROM Messages WHERE Id = $1`, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var message Message
	row.Next()
	err = row.Scan(&message.Id, &message.Text, &message.SenderId, &message.ReceiverId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &message, nil
}
