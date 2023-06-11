package userdto

import "dumbmerch/models"

type UserResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Image    string `json:"image"`
	RoleId   int    `json:"role_id"`
	// Role        models.RoleResponse  `json:"role"`
	Transaction []models.Transaction `json:"transaction"`
	// TransactionId int                `json:"transaction_id"`
}
