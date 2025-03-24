package dto

type Page struct {
	Data      any    `json:"data"`
	Page      uint16 `json:"page"`
	Limit     uint8  `json:"limit"`
	TotalPage uint16 `json:"totalPage"`
}
