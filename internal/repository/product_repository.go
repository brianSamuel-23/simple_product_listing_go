package repository

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"simple_product_listing_go/internal/customerror"
	"simple_product_listing_go/internal/dto"
	"simple_product_listing_go/internal/model"
	"simple_product_listing_go/internal/utils"
	"time"
)

type productRepository struct {
	collection *mongo.Collection
}

type ProductRepository interface {
	GetProducts(filter bson.M, page uint16, limit uint8, sortField string, sortOrder string) ([]model.Product, uint16, error)
}

func NewProductRepository(client *mongo.Client) ProductRepository {
	db := client.Database("example_db")
	collection := db.Collection("product")
	return &productRepository{collection}
}

func (r *productRepository) GetProducts(filter bson.M, page uint16, limit uint8, sortField string, sortOrder string) ([]model.Product, uint16, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//calculate how many elements to skip
	skip := (page - 1) * uint16(limit)

	pipeline := mongo.Pipeline{
		{{"$match", filter}}, // Filter documents
		{{"$facet", bson.M{
			"metadata": []bson.M{
				{"$count": "totalRecords"},
			},
			"data": bson.A{
				bson.M{"$sort": bson.M{sortField: utils.ConvertSort(sortOrder)}}, // Sorting
				bson.M{"$skip": skip},   // Pagination: skip
				bson.M{"$limit": limit}, // Pagination: limit
			},
		}}},
	}

	//sortOptions := bson.D{{Key: sortField, Value: utils.ConvertSort(sortOrder)}}
	//findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)).SetSort(sortOptions)
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Error(err)
		return nil, 0, customerror.ErrDatabase
	}
	defer cursor.Close(ctx)

	var result []dto.ProductAggregateResult
	if err = cursor.All(ctx, &result); err != nil {
		log.Error(err)
		return nil, 0, customerror.ErrParse
	}
	if len(result[0].Metadata) == 0 {
		log.Error(customerror.ErrEmptyResult)
		return nil, 0, customerror.ErrEmptyResult
	}
	totalRecords := result[0].Metadata[0].TotalRecords
	totalPages := (totalRecords + uint64(limit) - 1) / uint64(limit)
	return result[0].Data, uint16(totalPages), nil
}
