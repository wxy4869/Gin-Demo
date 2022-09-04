package response

import "ginDemo/model/database"

type CommonA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type RegisterQ struct {
	Username  string `json:"username" binding:"required"`
	Password1 string `json:"password1" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
	Email     string `json:"email" binding:"omitempty"`
}

type LoginQ struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginA struct {
	CommonA
	Token string        `json:"token"`
	User  database.User `json:"user"`
}

type GetUserInfoQ struct {
	UserID uint64 `json:"user_id" binding:"required"`
}

type GetUserInfoA struct {
	CommonA
	User   database.User `json:"user"`
	Poster database.User `json:"poster"`
}
