package data

type AccountDAO struct {
	Account_user int64	`json:"user_id"`
	Account_type int64	`json:"account_type"`
	Account_name string	`json:"account_name"`
}

type UserDAO struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone_number string `json:"phone_number"`
	Password     string `json:"password"`
}

func (user *UserDAO) HasMinDetails() bool {
	return user.Email != "" && user.Name != ""
}
