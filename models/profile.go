package models

import "time"

type Profile struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Image     string               `json:"image" gorm:"type: varchar(255)"`
	Phone     string               `json:"phone" gorm:"type: varchar(255)"`
	Gender    string               `json:"gender" gorm:"type: varchar(255)"`
	Address   string               `json:"address" gorm:"type: text"`
	Postcode  int                  `json:"postcode" gorm:"type: int"`
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type ProfileResponse struct {
	Image    string `json:"image"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
	Postcode int    `json:"postcode"`
	UserId   string `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
