package helper

import (
	"go.mongodb.org/mongo-driver/bson"
	"simple_product_listing_go/internal/customerror"
	"simple_product_listing_go/internal/dto"
	"strconv"
)

func ConstructProductFilter(query []dto.QueryParam) (bson.M, error) {
	filter := bson.M{}
	for _, param := range query {
		switch param.Field {
		case "name":
			filter["name"] = param.Value
		case "sku":
			filter["sku"] = param.Value
		case "stock":
			filter["stock"], _ = strconv.Atoi(param.Value)
		case "price":
			filter["price"], _ = strconv.Atoi(param.Value)
		case "active":
			filter["active"] = param.Value
		}
	}
	defer func() (bson.M, error) {
		if r := recover(); r != nil {
			return bson.M{}, customerror.ErrParse
		}
		return bson.M{}, nil
	}()
	return filter, nil
}
