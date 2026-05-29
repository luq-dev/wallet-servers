package data

type AccountDAO struct {
	id int64
	user_id int64
}

type UserDAO struct {
	Name string
	Email string
	Phone_number string
	Password string
}

func (user *UserDAO) HasMinDetails() bool{
	return user.Email != "" &&  user.Name != "" 
}
