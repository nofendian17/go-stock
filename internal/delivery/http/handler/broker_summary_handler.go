package handler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"go-stock/internal/model"
	"go-stock/internal/shared/response"
	"go-stock/internal/usecase"
	"net/http"
	"strings"
	"time"
)

type BrokerSummaryHandler interface {
	Find(w http.ResponseWriter, r *http.Request)
}

type brokerSummaryHandler struct {
	brokerSummaryUsecase usecase.BrokerSummaryUseCase
	validate             *validator.Validate
}

func NewBrokerSummaryHandler(brokerSummaryUsecase usecase.BrokerSummaryUseCase, validate *validator.Validate) BrokerSummaryHandler {
	return &brokerSummaryHandler{
		brokerSummaryUsecase: brokerSummaryUsecase,
		validate:             validate,
	}
}

func (h *brokerSummaryHandler) Find(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	stockCode := r.URL.Query().Get("stock_code")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	investorType := r.URL.Query().Get("investor_type")
	transactionType := r.URL.Query().Get("transaction_type")

	request := model.BrokerSummaryRequest{
		StockCode:       stockCode,
		StartDate:       startDate,
		EndDate:         endDate,
		InvestorType:    investorType,
		TransactionType: transactionType,
	}
	if err := h.validate.Struct(request); err != nil {
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

	// format date string from yyyy-mm-dd to mm/dd/yyyy
	if t, err := time.Parse("2006-01-02", startDate); err == nil {
		startDate = t.Format("01/02/2006")
	}
	if t, err := time.Parse("2006-01-02", endDate); err == nil {
		endDate = t.Format("01/02/2006")
	}

	result, err := h.brokerSummaryUsecase.Find(r.Context(), strings.ToUpper(request.StockCode), startDate, endDate, investorType, transactionType)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}

	if result == nil {
		response.NotFound(w, "")
		return
	}

	buyers := make([]model.BrokerSummaryData, 0, len(result.Buyers))
	for _, buyer := range result.Buyers {
		buyers = append(buyers, model.BrokerSummaryData{
			BrokerCode: buyer.BrokerCode,
			Lot:        buyer.Lot,
			Val:        buyer.Val,
			Avg:        buyer.Avg,
		})
	}

	sellers := make([]model.BrokerSummaryData, 0, len(result.Sellers))
	for _, seller := range result.Sellers {
		sellers = append(sellers, model.BrokerSummaryData{
			BrokerCode: seller.BrokerCode,
			Lot:        seller.Lot,
			Val:        seller.Val,
			Avg:        seller.Avg,
		})
	}

	data := model.BrokerSummaryResponse{
		StockCode: result.StockCode,
		StartDate: result.StartDate,
		EndDate:   result.EndDate,
		Buyers:    buyers,
		Sellers:   sellers,
		Summary: model.Summary{
			TotalVal:      result.Summary.TotalVal,
			ForeignNetVal: result.Summary.ForeignNetVal,
			TotalLot:      result.Summary.TotalLot,
			Avg:           result.Summary.Avg,
		},
	}

	response.Success(w, data, "")
	return
}
