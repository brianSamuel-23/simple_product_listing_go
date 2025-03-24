package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Sku    string             `json:"sku,omitempty" bson:"sku,omitempty"`
	Stock  int32              `json:"stock,omitempty" bson:"stock,omitempty"`
	Price  int64              `json:"price,omitempty" bson:"price,omitempty"`
	Active bool               `json:"active,omitempty" bson:"active,omitempty"`
}
