package handlers

import (
	cartdto "bewaysbuck/dto/cart"
	dto "bewaysbuck/dto/result"
	"bewaysbuck/models"
	"bewaysbuck/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

var transimg = "http://localhost:5000/uploads/"

func (h *handlerCart) FindCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	carts, err := h.CartRepository.FindCarts(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var responseCart []cartdto.CartResponse
	for _, t := range carts {
		responseCart = append(responseCart, convertResponseCart(t))
	}

	for i, t := range responseCart {
		imagePath := transimg + t.Product.Image
		responseCart[i].Product.Image = imagePath
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: responseCart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var cart models.Cart
	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) CreateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// var toppingsId []int
	// for _, r := range r.FormValue("toppingId") {
	// 	if int(r-'0') >= 0 {
	// 		toppingsId = append(toppingsId, int(r-'0'))
	// 	}
	// }

	// productId, _ := strconv.Atoi(r.FormValue("product_id"))
	// transactionId, _ := strconv.Atoi(r.FormValue("transaction_id"))
	// qty, _ := strconv.Atoi(r.FormValue("qty"))
	// subAmount, _ := strconv.Atoi(r.FormValue("sub_amount"))

	// request := cartdto.CartRequest{
	// 	ProductId:     productId,
	// 	ToppingID:     toppingsId,
	// 	TransactionId: transactionId,
	// 	Qty:           qty,
	// 	SubAmount:     subAmount,
	// }

	request := new(cartdto.CartRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction, err := h.CartRepository.GetTransactionID(userId)

	requestForm := models.Cart{
		ProductId:     request.ProductId,
		ToppingID:     request.ToppingID,
		TransactionId: int(transaction.ID),
		Qty:           request.Qty,
		SubAmount:     request.SubAmount,
		UserID:        userId,
	}

	validate := validator.New()
	errs := validate.Struct(requestForm)
	if errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping, _ := h.CartRepository.FindToppingsById(request.ToppingID)

	cart := models.Cart{
		ProductId:     request.ProductId,
		Topping:       topping,
		TransactionId: int(transaction.ID),
		Qty:           request.Qty,
		SubAmount:     request.SubAmount,
		UserID:        userId,
	}

	cart, err = h.CartRepository.CreateCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) UpdateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var toppingsId []int
	for _, r := range r.FormValue("toppingId") {
		if int(r-'0') >= 0 {
			toppingsId = append(toppingsId, int(r-'0'))
		}
	}

	productId, _ := strconv.Atoi(r.FormValue("product_id"))
	transactionId, _ := strconv.Atoi(r.FormValue("transaction_id"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))
	subAmount, _ := strconv.Atoi(r.FormValue("sub_amount"))
	request := cartdto.CartRequest{
		ProductId:     productId,
		ToppingID:     toppingsId,
		TransactionId: transactionId,
		Qty:           qty,
		SubAmount:     subAmount,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var topping []models.Topping
	if len(toppingsId) != 0 {
		topping, _ = h.CartRepository.FindToppingsById(toppingsId)
	}

	cart, _ := h.CartRepository.GetCart(id)

	cart.ProductId = request.ProductId
	cart.Topping = topping
	cart.TransactionId = request.TransactionId
	cart.Qty = request.Qty
	cart.SubAmount = request.SubAmount

	cart, err = h.CartRepository.UpdateCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) DeleteCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deleteCart, err := h.CartRepository.DeleteCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: deleteCart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) FindCartsByTrans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaction, err := h.CartRepository.GetIDTransaction()
	cart, err := h.CartRepository.FindCartsTransaction(int(transaction.ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, t := range cart {
		cart[i].Product.Image = transimg + t.Product.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cart}
	json.NewEncoder(w).Encode(response)
}

func convertResponseCart(t models.Cart) cartdto.CartResponse {
	return cartdto.CartResponse{
		ID:            t.ID,
		Product:       t.Product,
		Topping:       t.Topping,
		TransactionId: t.TransactionId,
		Qty:           t.Qty,
		SubAmount:     t.SubAmount,
	}
}
