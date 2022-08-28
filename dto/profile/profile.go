package profiledto

type ProfileRequest struct {
	Image   string `json:"image" gorm:"type: varchar(255)"`
	Phone   string `json:"phone" gorm:"type: varchar(255)"`
	Address string `json:"address" gorm:"type: varchar(255)"`
	Gender  string `json:"gender" gorm:"type: varchar(255)"`
	UserID  int    `json:"user_id" gorm:"type: int"`
}
