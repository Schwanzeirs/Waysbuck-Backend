package cartdto

import "bewaysbuck/models"

type CartRequest struct {
	ProductId     int   `json:"product_id" form:"product_id" gorm:"type: int"`
	ToppingID     []int `json:"topping_id" form:"topping_id" gorm:"-"`
	TransactionId int   `json:"transaction_id" form:"transaction_id" gorm:"type: int"`
	Qty           int   `json:"qty" form:"qty" gorm:"type: int"`
	SubAmount     int   `json:"sub_amount" form:"sub_amount" gorm:"type: int"`
}

type CartResponse struct {
	ID            int              `json:"id"`
	Product       models.Product   `json:"product"`
	Topping       []models.Topping `json:"topping"`
	TransactionId int              `json:"transaction_id"`
	Userid        int              `json:"user_id"`
	Qty           int              `json:"qty"`
	SubAmount     int              `json:"sub_amount"`
}
