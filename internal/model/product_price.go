package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductPrice struct {
	Id          primitive.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	Sku         string                `json:"sku,omitempty" bson:"sku,omitempty"`
	Price       int64                 `json:"price,omitempty" bson:"price,omitempty"`
	History     []ProductPriceHistory `json:"history,omitempty" bson:"history,omitempty"`
	CreatedDate time.Time             `json:"created_date,omitempty" bson:"created_date,omitempty"`
	UpdatedDate time.Time             `json:"updated_date,omitempty" bson:"updated_date,omitempty"`
}

type ProductPriceHistory struct {
	Price int64     `json:"price,omitempty" bson:"price,omitempty"`
	From  time.Time `json:"from,omitempty" bson:"from,omitempty"`
	To    time.Time `json:"to,omitempty" bson:"to,omitempty"`
}
