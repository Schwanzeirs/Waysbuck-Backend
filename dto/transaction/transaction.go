package transactiondto

type TransactionRequest struct {
	ID     int    `json:"id" gorm:"type: int"`
	Status string `json:"status" gorm:"type: varchar(255)"`
	UserID int    `json:"user_id" gorm:"type: int"`
	Amount int    `json:"amount" gorm:"type: int"`
}

type TransactionResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status" gorm:"type: varchar(255)"`
	UserID int    `json:"user_id" gorm:"type: int"`
	Amount int    `json:"amount" gorm:"type: int"`
}

type UpdateTransaction struct {
	UserID int    `json:"user_id" form:"user_id"`
	Status string `json:"status"`
	Amount int    `json:"amount"`
}
