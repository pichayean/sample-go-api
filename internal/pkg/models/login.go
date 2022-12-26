package models

type LogIn struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
