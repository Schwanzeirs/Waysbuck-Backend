package productdto

type ProductRequest struct {
	Title  string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price  int    `json:"price" form:"price" gorm:"type: int"`
	UserID int    `json:"user_id" gorm:"type: int"`
	Image  string `json:"image" form:"image" gorm:"type: varchar(255)"`
}
