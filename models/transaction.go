package models

import "time"

type Transaction struct {
	ID        int                      `json:"id" gorm:"primary_key:auto_increment"`
	Status    string                   `json:"status" gorm:"type: varchar(255)"`
	UserId    int                      `json:"user_id"`
	User      UsersTransactionResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cart      []Cart                   `json:"cart" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount    int                      `json:"amount"`
	CreatedAt time.Time                `json:"-"`
	UpdatedAt time.Time                `json:"-"`
}
