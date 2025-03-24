package helper

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"simple_product_listing_go/internal/customerror"
	"simple_product_listing_go/internal/dto"
	"strconv"
	"strings"
	"time"
)

func ConstructPricesFilter(query []dto.QueryParam) (bson.M, error) {
	filter := bson.M{}

	for _, param := range query {
		switch param.Field {
		case "sku":
			filter["sku"] = param.Value
		case "current_price":
			filter["price"], _ = strconv.Atoi(param.Value)
		case "created_date":
			if strings.Split(param.Value, "~~~")[0] == "" {
				filter["created_date"] = bson.M{
					"$lte": getDateFromString(strings.Split(param.Value, "~~~")[1]),
				}
			} else if strings.Split(param.Value, "~~~")[1] == "" {
				filter["created_date"] = bson.M{
					"$gte": getDateFromString(strings.Split(param.Value, "~~~")[0]),
				}
			} else {
				filter["created_date"] = bson.M{
					"$gte": getDateFromString(strings.Split(param.Value, "~~~")[0]),
					"$lte": getDateFromString(strings.Split(param.Value, "~~~")[1]),
				}
			}
		case "updated_date":
			if strings.Split(param.Value, "~~~")[0] == "" {
				filter["updated_date"] = bson.M{
					"$lte": getDateFromString(strings.Split(param.Value, "~~~")[1]),
				}
			} else if strings.Split(param.Value, "~~~")[1] == "" {
				filter["updated_date"] = bson.M{
					"$gte": getDateFromString(strings.Split(param.Value, "~~~")[0]),
				}
			} else {
				filter["updated_date"] = bson.M{
					"$gte": getDateFromString(strings.Split(param.Value, "~~~")[0]),
					"$lte": getDateFromString(strings.Split(param.Value, "~~~")[1]),
				}
			}
		case "price":
			filter["history.price"], _ = strconv.Atoi(param.Value)
		case "valid_from":
			if strings.Split(param.Value, "~~~")[0] == "" {
				filter["history.from"] = bson.M{
					"$lte": getDateFromString(strings.Split(param.Value, "~~~")[1]),
				}
			} else if strings.Split(param.Value, "~~~")[1] == "" {
				filter["history.from"] = bson.M{
					"$gte": getDateFromString(strings.Split(param.Value, "~~~")[0]),
				}
			} else {
				filter["history.from"] = bson.M{
					"$gte": getDateFromString(strings.Split(param.Value, "~~~")[0]),
					"$lte": getDateFromString(strings.Split(param.Value, "~~~")[1]),
				}
			}
		case "valid_to":
			if strings.Split(param.Value, "~~~")[0] == "" {
				filter["history.to"] = bson.M{
					"$lte": getDateFromString(strings.Split(param.Value, "~~~")[1]),
				}
			} else if strings.Split(param.Value, "~~~")[1] == "" {
				filter["history.to"] = bson.M{
					"$gte": getDateFromString(strings.Split(param.Value, "~~~")[0]),
				}
			} else {
				filter["history.to"] = bson.M{
					"$gte": getDateFromString(strings.Split(param.Value, "~~~")[0]),
					"$lte": getDateFromString(strings.Split(param.Value, "~~~")[1]),
				}
			}
		}
	}
	fmt.Println(filter)
	defer func() (bson.M, error) {
		if r := recover(); r != nil {
			return bson.M{}, customerror.ErrParse
		}
		return bson.M{}, nil
	}()
	return filter, nil
}

func getDateFromString(stringDate string) time.Time {
	timeFormat := "2006-01-02 15:04:05"
	date, _ := time.Parse(timeFormat, stringDate)
	return date
}
