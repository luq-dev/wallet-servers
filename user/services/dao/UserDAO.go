package dao

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	Password    string `json:"password"`
}

type UserDAO struct {
	DB *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{DB: db}
}

func (u *UserDAO) GetUserByID(ctx context.Context, uid int64) (*UserDTO, error) {
	var user UserDTO
	err := u.DB.QueryRowContext(ctx, "SELECT fullname, email, user_type, phone_number FROM users WHERE id = $1", uid).Scan(&user.Name, &user.Email, &user.Role, &user.PhoneNumber)

	return &UserDTO{
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		PhoneNumber: user.PhoneNumber,
	}, err
}

func (u *UserDAO) GetUserByEmail(ctx context.Context, email string) (*UserDTO, error) {
	var user UserDTO
	err := u.DB.QueryRowContext(ctx, "SELECT fullname, email, user_type, phone_number FROM users WHERE email = $1", email).Scan(&user.Name, &user.Email, &user.Role, &user.PhoneNumber)

	return &UserDTO{
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		PhoneNumber: user.PhoneNumber,
	}, err
}

func (u *UserDAO) NewUser(ctx context.Context, user *UserDTO) (int64, error) {
	var userId int64
	var password string
	// 0Auth Insertion
	if user.Password == "" {
		password = ""
	} else { 
		pB, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return 0, fmt.Errorf("Password Generation Failed")
		}
		password = string(pB)
	}

	if user.Email != "" && user.Name != "" {
		err := u.DB.QueryRowContext(ctx, "INSERT INTO users(fullname, email, password, phone_number) values ($1,$2,$3,$4) RETURNING id", user.Name, user.Email, string(password), user.PhoneNumber).Scan(&userId)
		if err != nil {
			return 0, fmt.Errorf("Failed to Add User")
		}
	} else {
		return 0, fmt.Errorf("Missing Data")
	}
	return userId, nil
}
