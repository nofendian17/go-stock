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

type BrokerHandler interface {
	Find(w http.ResponseWriter, r *http.Request)
}

type brokerHandler struct {
	brokerUsecase usecase.BrokerUseCase
	validate      *validator.Validate
}

func NewBrokerHandler(brokerUsecase usecase.BrokerUseCase, validate *validator.Validate) BrokerHandler {
	return &brokerHandler{
		brokerUsecase: brokerUsecase,
		validate:      validate,
	}
}

// Find brokers
// @Summary Find brokers
// @Description Find brokers by code
// @Tags Broker
// @Produce json
// @Param request query model.BrokerRequest true "query params"
// @Success 200 {array} model.BrokerResponse
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /api/v1/brokers [get]
func (h *brokerHandler) Find(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	code := r.URL.Query().Get("code")

	request := model.BrokerRequest{
		Code: code,
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

	results, err := h.brokerUsecase.Find(r.Context(), strings.ToUpper(request.Code))
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}

	data := make([]model.BrokerResponse, 0, len(results))
	for _, result := range results {
		data = append(data, model.BrokerResponse{
			Code:    result.Code,
			Name:    result.Name,
			License: result.License,
		})
	}

	response.Success(w, data, "")
	return
}
