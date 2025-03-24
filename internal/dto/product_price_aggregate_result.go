package dto

import "simple_product_listing_go/internal/model"

type ProductPriceAggregateResult struct {
	Metadata []struct {
		TotalRecords uint64 `bson:"totalRecords"`
	} `bson:"metadata"`
	Data []model.ProductPrice `bson:"data"`
}
