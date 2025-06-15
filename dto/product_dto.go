package dto

type ProductDTO struct {
	Name *string `json:"name"`
	Cost uint32  `json:"cost"`
}
