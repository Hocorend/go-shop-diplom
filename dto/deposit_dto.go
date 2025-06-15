package dto

type DepositDTO struct {
	Amount uint32  `json:"amount"`
	Login  *string `json:"login"`
}
