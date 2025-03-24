package service

import (
	"errors"
	"simple_product_listing_go/internal/customerror"
	"simple_product_listing_go/internal/dto"
	"simple_product_listing_go/internal/helper"
	"simple_product_listing_go/internal/repository"
)

type ProductService interface {
	GetProducts(filters []dto.QueryParam, page uint16, limit uint8, sortField string, sortOrder string) (pageResponse dto.Page, err error)
	GetProductPricesAdmin(filters []dto.QueryParam, page uint16, limit uint8, sortField string, sortOrder string) (pageResponse dto.Page, err error)
	GetProductPricesUser(filters []dto.QueryParam, page uint16, limit uint8, sortField string, sortOrder string) (pageResponse dto.Page, err error)
}

type productService struct {
	productRepository      repository.ProductRepository
	productPriceRepository repository.ProductPriceRepository
}

func NewProductService(productRepository repository.ProductRepository, productPriceRepository repository.ProductPriceRepository) ProductService {
	return &productService{productRepository, productPriceRepository}
}

func (s *productService) GetProducts(filters []dto.QueryParam, page uint16, limit uint8, sortField string, sortOrder string) (pageResponse dto.Page, err error) {
	filter, err := helper.ConstructProductFilter(filters)
	if err != nil {
		return dto.Page{}, customerror.ErrParse
	}

	//make sure that we only fetch active products
	filter["active"] = true
	// fetch from repository
	products, totalPages, err := s.productRepository.GetProducts(filter, page, limit, sortField, sortOrder)
	//if err, return error
	if err != nil {
		if errors.Is(err, customerror.ErrEmptyResult) {
			return dto.Page{}, customerror.ErrEmptyResult
		} else {
			return dto.Page{}, customerror.ErrInternalServerError

		}
	}

	return dto.Page{
		Data:      products,
		Page:      page,
		Limit:     limit,
		TotalPage: totalPages,
	}, err

}

func (s *productService) GetProductPricesAdmin(filters []dto.QueryParam, page uint16, limit uint8, sortField string, sortOrder string) (pageResponse dto.Page, err error) {

	var priceResponse []dto.ProductPriceResponse
	filter, err := helper.ConstructPricesFilter(filters)
	if err != nil {
		return dto.Page{}, customerror.ErrParse
	}

	// fetch from repository
	productPrices, totalPages, err := s.productPriceRepository.GetProductPrices(filter, page, limit, sortField, sortOrder)
	//if err, return error
	if err != nil {
		if errors.Is(err, customerror.ErrEmptyResult) {
			return dto.Page{}, customerror.ErrEmptyResult
		} else {
			return dto.Page{}, customerror.ErrInternalServerError

		}
	}
	for _, prices := range productPrices {
		var historyResponse []dto.ProductPriceHistoryResponse
		for _, history := range prices.History {
			historyResponse = append(historyResponse, dto.ProductPriceHistoryResponse{
				Price:     history.Price,
				ValidFrom: history.From.Format("2006-01-02 15:04:05"),
				ValidTo:   history.To.Format("2006-01-02 15:04:05"),
			})
		}
		priceResponse = append(priceResponse, dto.ProductPriceResponse{
			Id:           prices.Id,
			Sku:          prices.Sku,
			CurrentPrice: prices.Price,
			History:      historyResponse,
			CreatedDate:  prices.CreatedDate.Format("2006-01-02 15:04:05"),
			UpdatedDate:  prices.UpdatedDate.Format("2006-01-02 15:04:05"),
		})

	}
	return dto.Page{
		Data:      priceResponse,
		Page:      page,
		Limit:     limit,
		TotalPage: totalPages,
	}, err
}

func (s *productService) GetProductPricesUser(filters []dto.QueryParam, page uint16, limit uint8, sortField string, sortOrder string) (pageResponse dto.Page, err error) {

	var priceResponse []dto.ProductPriceResponseUser
	filter, err := helper.ConstructPricesFilter(filters)

	if err != nil {
		return dto.Page{}, customerror.ErrParse
	}

	// fetch from repository
	productPrices, totalPages, err := s.productPriceRepository.GetProductPrices(filter, page, limit, sortField, sortOrder)
	//if err, return error
	if err != nil {
		if errors.Is(err, customerror.ErrEmptyResult) {
			return dto.Page{}, customerror.ErrEmptyResult
		} else {
			return dto.Page{}, customerror.ErrInternalServerError

		}
	}
	for _, prices := range productPrices {
		priceResponse = append(priceResponse, dto.ProductPriceResponseUser{
			Id:           prices.Id,
			Sku:          prices.Sku,
			CurrentPrice: prices.Price,
		})

	}
	return dto.Page{
		Data:      priceResponse,
		Page:      page,
		Limit:     limit,
		TotalPage: totalPages,
	}, err
}
