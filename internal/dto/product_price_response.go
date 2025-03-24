package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductPriceResponse struct {
	Id           primitive.ObjectID            `json:"_id,omitempty" `
	Sku          string                        `json:"sku,omitempty"`
	CurrentPrice int64                         `json:"current_price,omitempty"`
	History      []ProductPriceHistoryResponse `json:"history,omitempty"`
	CreatedDate  string                        `json:"created_date,omitempty"`
	UpdatedDate  string                        `json:"updated_date,omitempty"`
}

type ProductPriceHistoryResponse struct {
	Price     int64  `json:"price,omitempty"`
	ValidFrom string `json:"valid_from,omitempty"`
	ValidTo   string `json:"valid_to,omitempty"`
}
