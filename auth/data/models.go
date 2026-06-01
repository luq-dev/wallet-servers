package data

type Account struct {
	UserID int64  `json:"user_id"`
	Type   int64  `json:"account_type"`
	Name   string `json:"account_name"`
}

type User struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
