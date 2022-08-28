package models

import "time"

type User struct {
	ID          int             `json:"id" gorm:"primary_key:auto_increment"`
	Name        string          `json:"fullname" gorm:"type: varchar(255)"`
	Email       string          `json:"email" gorm:"type: varchar(255)"`
	Password    string          `json:"password" gorm:"type: varchar(255)"`
	Status      string          `json:"status" gorm:"type: varchar(255)"`
	Transaction Transaction     `json:"transaction"`
	Profile     ProfileResponse `json:"profile"`
	CreatedAt   time.Time       `json:"-"`
	UpdatedAt   time.Time       `json:"-"`
}

type UsersProfileResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UsersTransactionResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}

func (UsersTransactionResponse) TableName() string {
	return "users"
}
