package model

type UserRegister struct {
	Name           string `json:"name" binding:"required"`
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required,min=8"`
	Email          string `json:"email" binding:"required,email"`
	ProfilePicture string `json:"profilePicture"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"jwtToken"`
}

type UpdateUser struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	Password       string `json:"password" binding:"omitempty,min=8"`
	Email          string `json:"email" binding:"omitempty,email"`
	ProfilePicture string `json:"profilePicture"`
}
