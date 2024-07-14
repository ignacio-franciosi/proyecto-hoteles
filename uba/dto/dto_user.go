package dto

type UserDto struct {
	IdUser   int    `json:"id_user"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType bool   `json:"userType"`
	Dni      int    `json:"dni"`
}

type UsersDto []UserDto
