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
	return &UserDAO{DB: db}
}

func (u *UserDAO) GetUserByID(uid int64) (*data.User, error) {
	var user data.User
	err := u.DB.QueryRow("SELECT fullname, email, user_type, phone_number FROM users WHERE id = $1", uid).Scan(&user.Name, &user.Email, &user.Role, &user.PhoneNumber)

	return &data.User{
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		PhoneNumber: user.PhoneNumber,
	}, err
}

func (u *UserDAO) GetUserByEmail(email string) (*data.User, error) {
	var user data.User
	err := u.DB.QueryRow("SELECT fullname, email, user_type, phone_number FROM users WHERE email = $1", email).Scan(&user.Name, &user.Email, &user.Role, &user.PhoneNumber)

	return &data.User{
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		PhoneNumber: user.PhoneNumber,
	}, err
}

func (u *UserDAO) AddUser(user *data.User) (int64, error) {
	var userId int64
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("Password Generation Failed")
	}

	if user.Email != "" && user.Name != "" {
		err := u.DB.QueryRow("INSERT INTO users(fullname, email, password, phone_number) values ($1,$2,$3,$4) RETURNING id", user.Name, user.Email, string(password), user.PhoneNumber).Scan(&userId)
		if err != nil {
			return 0, fmt.Errorf("Failed to Add User: Internal Server Error")
		}
	} else {
		return 0, fmt.Errorf("Missing Data")
	}
	return userId, nil
}
