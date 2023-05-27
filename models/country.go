package models

import "time"

type Country struct {
<<<<<<< Updated upstream
	Id   int    `json:"id_country" form:"id_country" gorm:"primary_key:auto_increment"`
	Name string `json:"name_country" form:"name" gorm:"type: varchar(255)"`
	// User      UserResponse `json:"-"`
=======
	Id        int       `json:"id_country" form:"id_country" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name_country" form:"name" gorm:"type: varchar(255)"`
>>>>>>> Stashed changes
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type CountryResponse struct {
	Id   int    `json:"id_country"`
	Name string `json:"name_country"`
}

func (CountryResponse) TableName() string {
	return "countries"
}
