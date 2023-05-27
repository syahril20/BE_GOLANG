package models

// import "time"

// type Trip struct {
// 	Id             int             `json:"id" form:"id" gorm:"primary_key:auto_increment"`
// 	Title          string          `json:"title" form:"accomodation" gorm:"type: varchar(255)"`
// 	IdCountry      int             `json:"id_country"`
// 	Country        CountryResponse `json:"country" form:"id_country"`
// 	Accomodation   string          `json:"accomodation" form:"accomodation" gorm:"type: varchar(255)"`
// 	Transportation string          `json:"transportation" form:"transportation" gorm:"type: varchar(255)"`
// 	Eat            string          `json:"eat" form:"eat" gorm:"type: varchar(255)"`
// 	Day            int             `json:"day" form:"day"`
// 	Night          int             `json:"night" form:"night"`
// 	DateTrip       string          `json:"date_trip" form:"date_trip" gorm:"type: varchar(255)"`
// 	Price          int             `json:"price" form:"price"`
// 	Quota          int             `json:"quota" form:"quota"`
// 	Description    string          `json:"description" form:"description" gorm:"type: varchar(255)"`
// 	Image          string          `json:"image" form:"image" gorm:"type: varchar(255)"`
// 	CreatedAt      time.Time       `json:"-"`
// 	UpdatedAt      time.Time       `json:"-"`
// }

// type TripResponse struct {
// 	Id        int             `json:"id"`
// 	Title     string          `json:"title"`
// 	IdCountry int             `json:"id_country"`
// 	Country   CountryResponse `json:"country" form:"id_country"`
// }

// func (TripResponse) TableName() string {
// 	return "trips"
// }
