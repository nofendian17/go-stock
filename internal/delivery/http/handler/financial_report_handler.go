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

type FinancialReportHandler interface {
	FindFinancialReport(w http.ResponseWriter, r *http.Request)
}
type financialReportHandler struct {
	financialReportUseCase usecase.FinancialReportUseCase
	validate               *validator.Validate
}

func NewFinancialReportHandler(financialReportUseCase usecase.FinancialReportUseCase, validate *validator.Validate) FinancialReportHandler {
	return &financialReportHandler{
		financialReportUseCase: financialReportUseCase,
		validate:               validate,
	}
}

func (h *financialReportHandler) FindFinancialReport(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	stockCode := r.URL.Query().Get("stock_code")
	reportPeriod := r.URL.Query().Get("report_period")
	reportYear := r.URL.Query().Get("report_year")

	request := model.FinancialReportRequest{
		StockCode:    stockCode,
		ReportPeriod: reportPeriod,
		ReportYear:   reportYear,
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

	result, err := h.financialReportUseCase.Find(r.Context(), strings.ToUpper(request.StockCode), request.ReportPeriod, request.ReportYear)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}

	if result == nil {
		response.NotFound(w, "")
		return
	}

	attachments := make([]model.Attachment, 0, len(result.Attachment))
	for _, attachment := range result.Attachment {
		attachments = append(attachments, model.Attachment{
			StockCode:    attachment.StockCode,
			StockName:    attachment.StockName,
			FileID:       attachment.FileID,
			FileModified: attachment.FileModified,
			FileName:     attachment.FileName,
			FilePath:     attachment.FilePath,
			FileSize:     attachment.FileSize,
			FileType:     attachment.FileType,
			ReportPeriod: attachment.ReportPeriod,
			ReportType:   attachment.ReportType,
			ReportYear:   attachment.ReportYear,
		})
	}

	data := model.FinancialReportResponse{
		StockCode:    result.StockCode,
		FileModified: result.FileModified,
		ReportPeriod: result.ReportPeriod,
		ReportYear:   result.ReportYear,
		StockName:    result.StockName,
		Attachment:   attachments,
	}

	response.Success(w, data, "")
	return
}
