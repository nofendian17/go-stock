package handler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"go-stock/internal/model"
	"go-stock/internal/shared/response"
	"go-stock/internal/usecase"
	"net/http"
	"strings"
)

type StockSummaryHandler interface {
	FindStockSummaries(w http.ResponseWriter, r *http.Request)
}
type stockSummaryHandler struct {
	stockSummaryUseCase usecase.StockSummaryUseCase
	validate            *validator.Validate
}

func NewStockSummaryHandler(stockSummaryUseCase usecase.StockSummaryUseCase, validate *validator.Validate) StockSummaryHandler {
	return &stockSummaryHandler{
		stockSummaryUseCase: stockSummaryUseCase,
		validate:            validate,
	}
}

// FindStockSummaries find stock summaries
// @Summary Find stock summaries
// @Description Find stock summaries by stock code, start date, and end date
// @Tags Stock
// @Produce json
// @Param request query model.StockSummaryRequest true "query params"
// @Success 200 {array} model.StockSummaryResponse
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /api/v1/stock/summaries [get]
func (s *stockSummaryHandler) FindStockSummaries(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	stockCode := r.URL.Query().Get("stock_code")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	request := model.StockSummaryRequest{
		StockCode: stockCode,
		StartDate: startDate,
		EndDate:   endDate,
	}
	if err := s.validate.Struct(request); err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			errs := make([]response.Error, 0, len(validationErrs))
			for _, fieldError := range validationErrs {
				errs = append(errs, response.Error{
					Field:   fieldError.Field(),
					Message: fieldError.Error(),
				})
			}
			response.BadRequest(w, "", errs)
			return
		}
		response.InternalError(w, err.Error())
		return
	}

	results, err := s.stockSummaryUseCase.FindSummaries(r.Context(), strings.ToUpper(stockCode), startDate, endDate)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}

	data := make([]model.StockSummaryResponse, 0, len(results))
	for _, result := range results {
		data = append(data, model.StockSummaryResponse{
			IDStockSummary:      result.IDStockSummary,
			Date:                result.Date,
			StockCode:           result.StockCode,
			StockName:           result.StockName,
			Remarks:             result.Remarks,
			Previous:            result.Previous,
			OpenPrice:           result.OpenPrice,
			FirstTrade:          result.FirstTrade,
			High:                result.High,
			Low:                 result.Low,
			Close:               result.Close,
			Change:              result.Change,
			Volume:              result.Volume,
			Value:               result.Value,
			Frequency:           result.Frequency,
			IndexIndividual:     result.IndexIndividual,
			Offer:               result.Offer,
			OfferVolume:         result.OfferVolume,
			Bid:                 result.Bid,
			BidVolume:           result.BidVolume,
			ListedShares:        result.ListedShares,
			TradebleShares:      result.TradebleShares,
			WeightForIndex:      result.WeightForIndex,
			ForeignSell:         result.ForeignSell,
			ForeignBuy:          result.ForeignBuy,
			DelistingDate:       result.DelistingDate,
			NonRegularVolume:    result.NonRegularVolume,
			NonRegularValue:     result.NonRegularValue,
			NonRegularFrequency: result.NonRegularFrequency,
			Persen:              result.Persen,
			Percentage:          result.Percentage,
		})
	}

	response.Success(w, data, "")
	return
}
