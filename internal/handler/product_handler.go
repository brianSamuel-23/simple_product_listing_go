package handler

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"simple_product_listing_go/internal/customerror"
	"simple_product_listing_go/internal/dto"
	"simple_product_listing_go/internal/service"
	"strconv"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService}
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {

	var queryParams []dto.QueryParam
	filterStr := c.Query("querySearch", "[]") // Default to empty array if not provided

	// Decode URL-encoded JSON string
	decodedFilterStr, err := url.QueryUnescape(filterStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid URL encoding"})
	}

	// Parse JSON into []QueryParam
	if err := json.Unmarshal([]byte(decodedFilterStr), &queryParams); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid filter format"})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortField := c.Query("sortField", "id")
	sortOrder := c.Query("sortOrder", "asc")

	// Call service
	pageResult, err := h.productService.GetProducts(queryParams, uint16(page), uint8(limit), sortField, sortOrder)

	if err != nil && errors.Is(err, customerror.ErrEmptyResult) {
		return c.Status(204).JSON(fiber.Map{"error": err.Error()})
	} else if err != nil && errors.Is(err, customerror.ErrInternalServerError) {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	} else if err != nil && errors.Is(err, customerror.ErrParse) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid query format"})
	}

	return c.JSON(fiber.Map{
		"data":       pageResult.Data,
		"page":       pageResult.Page,
		"limit":      pageResult.Limit,
		"totalPages": pageResult.TotalPage,
	})

}

func (h *ProductHandler) GetProductPrices(c *fiber.Ctx) error {

	var queryParams []dto.QueryParam
	filterStr := c.Query("querySearch", "[]") // Default to empty array if not provided

	// Decode URL-encoded JSON string
	decodedFilterStr, err := url.QueryUnescape(filterStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL encoding"})
	}

	// Parse JSON into []QueryParam
	if err := json.Unmarshal([]byte(decodedFilterStr), &queryParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid filter format"})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortField := c.Query("sortField", "id")
	sortOrder := c.Query("sortOrder", "asc")

	// Call service based on role, that is admin or user
	var pageResult dto.Page

	if c.Locals("admin").(bool) {
		pageResult, err = h.productService.GetProductPricesAdmin(queryParams, uint16(page), uint8(limit), sortField, sortOrder)
	} else {
		pageResult, err = h.productService.GetProductPricesUser(queryParams, uint16(page), uint8(limit), sortField, sortOrder)
	}

	if err != nil && errors.Is(err, customerror.ErrEmptyResult) {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"error": err.Error()})
	} else if err != nil && errors.Is(err, customerror.ErrInternalServerError) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} else if err != nil && errors.Is(err, customerror.ErrParse) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid query format"})
	}

	return c.JSON(fiber.Map{
		"data":       pageResult.Data,
		"page":       pageResult.Page,
		"limit":      pageResult.Limit,
		"totalPages": pageResult.TotalPage,
	})

}
