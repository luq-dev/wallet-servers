package dao

import (
	"database/sql"
	"fmt"
	"user/data"

	"golang.org/x/crypto/bcrypt"
)

type UserDAO struct {
	DB *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{ DB: db }
}

func (u *UserDAO) GetUserByID(uid int64) (*data.User, error){
	var user data.User
	err := u.DB.QueryRow("SELECT fullname, email, type, phone_number FROM users WHERE id = $1", uid).Scan(&user.Name, &user.Email, &user.Type, &user.PhoneNumber)

	return &data.User{
		Name: user.Name,
		Email: user.Email,
		Type: user.Type,
		PhoneNumber: user.PhoneNumber,
		}, err
}

func (u *UserDAO) GetUserByEmail(email string) (*data.User, error){
	var user data.User
	err := u.DB.QueryRow("SELECT fullname, email, type, phone_number FROM users WHERE email = $1", email).Scan(&user.Name, &user.Email, &user.Type, &user.PhoneNumber)

	return &data.User{
		Name: user.Name,
		Email: user.Email,
		Type: user.Type,
		PhoneNumber: user.PhoneNumber,
		}, err
}

func (u *UserDAO) AddUser(user *data.User)  error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Password Generation Failed")
	}

	if user.Email != "" && user.Name != "" {
		_, err := u.DB.Exec("INSERT INTO users(fullname, email, password, phone_number) values ($1,$2,$3,$4)", user.Name, user.Email, string(password), user.PhoneNumber)
		if err != nil {
			return fmt.Errorf("Failed to Add User: Internal Server Error")
		}
	} else {
		return fmt.Errorf("Missing Data")
	}
	return nil
}


