package net

import (
	"database/sql"
	. "message/models"
)

type MessageService struct {
	db *sql.DB
}

func NewMessageService(db *sql.DB) *MessageService {
	return &MessageService{
		db: db,
	}
}

/*
func (s *MessageService) FetchMessages(daysAgo int, id int64) ([]Message, error) {
	var notifications []Message

	res, err := s.db.Query( // Fetch all unread notifications from 3 days ago.. 3 is hard coded
		`SELECT title, notification_message FROM notifications
			WHERE is_read = FALSE AND user_id = $1 AND created_at >= NOW() - INTERVAL '$2 days'`,
		id, daysAgo)

	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var buff Message
		err := res.Scan(&buff.Title, &buff.Message)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, buff)
	}

	if err = res.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}
	*/

func (s *MessageService) AppendMessage(m Message){

}

func (s *MessageService) Send(m Message) {

}
