package models

type RequireHeader struct {
	XAuth  string ` binding:"required"`
	XToken string ` binding:"required"`
}
