package dto

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType bool   `json:"userType"`
	Dni      int    `json:"dni"`
}

type UsersDto []UserDto
