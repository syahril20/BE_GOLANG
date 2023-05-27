package userdto

type CreateUser struct {
	Name      string `json:"name" form:"name" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Phone     string `json:"phone" form:"phone" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	IdCountry int    `json:"id_country" form:"id_country" validate:"required"`
}

type UpdateUser struct {
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Phone     string `json:"phone" form:"phone"`
	Address   string `json:"address" form:"address"`
	IdCountry int    `json:"id_country" form:"id_country"`
}
