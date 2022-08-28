package repositories

import (
	"bewaysbuck/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCarts(ID int) ([]models.Cart, error)
	FindToppingsById(ToppingID []int) ([]models.Topping, error)
	GetTransactionID(ID int) (models.Transaction, error)
	GetCart(ID int) (models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart) (models.Cart, error)
	FindCartsTransaction(TrsID int) ([]models.Cart, error)
	GetIDTransaction() (models.Transaction, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCarts(ID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product").Preload("Topping").Preload("User").Find(&carts, "user_id = ?", ID).Error

	return carts, err
}

func (r *repository) FindToppingsById(ToppingID []int) ([]models.Topping, error) {
	var toppings []models.Topping
	err := r.db.Find(&toppings, ToppingID).Error

	return toppings, err
}

func (r *repository) GetTransactionID(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Find(&transaction, "user_id =? AND status = ?", ID, "Active").Error
	return transaction, err
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Product").Preload("Topping").First(&cart, ID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}

func (r *repository) FindCartsTransaction(TrsID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product").Preload("Topping").Find(&carts, "transaction_id = ?", TrsID).Error

	return carts, err
}

func (r *repository) GetIDTransaction() (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Cart").Preload("Cart.Product").Preload("Cart.Topping").Find(&transaction, "status = ?", "Active").Error

	return transaction, err
}
