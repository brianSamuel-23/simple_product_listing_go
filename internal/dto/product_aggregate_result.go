package dto

import "simple_product_listing_go/internal/model"

type ProductAggregateResult struct {
	Metadata []struct {
		TotalRecords uint64 `bson:"totalRecords"`
	} `bson:"metadata"`
	Data []model.Product `bson:"data"`
}
