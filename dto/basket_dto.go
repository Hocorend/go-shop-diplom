package dto

type BasketDTO struct {
	Login       *string `json:"user_login"`
	ProductName *string `json:"product_name"`
	Count       uint32  `json:"product_count"`
}
