package models

type Trip struct {
	Id        int             `json:"id_trip" form:"id_trip" gorm:"primary_key:auto_increment"`
	Title     string          `json:"title" form:"title" gorm:"type: varchar(255)"`
	IdCountry int             `json:"-" form:"id_country"`
	Country   CountryResponse `json:"country" form:"country" gorm:"foreignKey:IdCountry"`
}
