package model

type PaginationRequest struct {
	Page  int64 `json:"page" validate:"min=1"`
	Limit int64 `json:"limit" validate:"min=1,max=100"`
}

type PaginationResponse struct {
	Data       []StockResponse `json:"data"`
	Total      int64           `json:"total"`
	Page       int64           `json:"page"`
	Limit      int64           `json:"limit"`
	TotalPages int64           `json:"total_pages"`
}
