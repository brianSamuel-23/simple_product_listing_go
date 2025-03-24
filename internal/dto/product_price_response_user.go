package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductPriceResponseUser struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Sku          string             `json:"sku,omitempty" bson:"sku,omitempty"`
	CurrentPrice int64              `json:"current_price,omitempty" bson:"price,omitempty"`
}
