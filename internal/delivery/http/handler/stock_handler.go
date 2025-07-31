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

type StockHandler interface {
	ListStock(w http.ResponseWriter, r *http.Request)
	FindStock(w http.ResponseWriter, r *http.Request)
	FindHistoricalStock(w http.ResponseWriter, r *http.Request)
}
type stockHandler struct {
	stockUsecase usecase.StockUseCase
	validate     *validator.Validate
}

func NewStockHandler(stockUseCase usecase.StockUseCase, validate *validator.Validate) StockHandler {
	return &stockHandler{
		stockUsecase: stockUseCase,
		validate:     validate,
	}
}

// ListStock list all stocks
// @Summary List all stocks
// @Description List all stocks
// @Tags Stock
// @Produce json
// @Success 200 {array} model.StockResponse
// @Failure 500 {object} response.Error
// @Router /api/v1/stocks [get]
func (s *stockHandler) ListStock(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	lists, err := s.stockUsecase.ListStocks(r.Context())
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}

	data := make([]model.StockResponse, 0, len(lists))
	for _, list := range lists {
		profiles := make([]model.Profile, 0, len(list.Profiles))
		for _, profile := range list.Profiles {
			profiles = append(profiles, model.Profile{
				Address:      profile.Address,
				BAE:          profile.BAE,
				Industry:     profile.Industry,
				SubIndustry:  profile.SubIndustry,
				Email:        profile.Email,
				Fax:          profile.Fax,
				MainBusiness: profile.MainBusiness,
				StockCode:    profile.StockCode,
				StockName:    profile.StockName,
				TIN:          profile.TIN,
				Sector:       profile.Sector,
				SubSector:    profile.SubSector,
				ListingDate:  profile.ListingDate,
				Phone:        profile.Phone,
				Website:      profile.Website,
				Status:       profile.Status,
				Logo:         profile.Logo,
			})
		}

		secretaries := make([]model.Secretary, 0, len(list.Secretaries))
		for _, secretary := range list.Secretaries {
			secretaries = append(secretaries, model.Secretary{
				Name:         secretary.Name,
				PhoneNumber:  secretary.PhoneNumber,
				Website:      secretary.Website,
				Email:        secretary.Email,
				Fax:          secretary.Fax,
				MobileNumber: secretary.MobileNumber,
			})
		}

		directors := make([]model.Director, 0, len(list.Directors))
		for _, director := range list.Directors {
			directors = append(directors, model.Director{
				Name:         director.Name,
				Position:     director.Position,
				IsAffiliated: director.IsAffiliated,
			})
		}

		commissioners := make([]model.Commissioner, 0, len(list.Commissioners))
		for _, commissioner := range list.Commissioners {
			commissioners = append(commissioners, model.Commissioner{
				Name:          commissioner.Name,
				Position:      commissioner.Position,
				IsIndependent: commissioner.IsIndependent,
			})
		}

		auditCommittees := make([]model.AuditCommittee, 0, len(list.AuditCommittees))
		for _, committee := range list.AuditCommittees {
			auditCommittees = append(auditCommittees, model.AuditCommittee{
				Name:     committee.Name,
				Position: committee.Position,
			})
		}

		shareholders := make([]model.Shareholder, 0, len(list.Shareholders))
		for _, shareholder := range list.Shareholders {
			shareholders = append(shareholders, model.Shareholder{
				Share:        shareholder.Share,
				Category:     shareholder.Category,
				Name:         shareholder.Name,
				IsController: shareholder.IsController,
				Percentage:   shareholder.Percentage,
			})
		}

		subsidiaries := make([]model.Subsidiary, 0, len(list.Subsidiaries))
		for _, subsidiary := range list.Subsidiaries {
			subsidiaries = append(subsidiaries, model.Subsidiary{
				BusinessFields:  subsidiary.BusinessFields,
				TotalAsset:      subsidiary.TotalAsset,
				Location:        subsidiary.Location,
				Currency:        subsidiary.Currency,
				Name:            subsidiary.Name,
				Percentage:      subsidiary.Percentage,
				Units:           subsidiary.Units,
				OperationStatus: subsidiary.OperationStatus,
				CommercialYear:  subsidiary.CommercialYear,
			})
		}

		dividends := make([]model.Dividend, 0, len(list.Dividends))
		for _, dividend := range list.Dividends {
			dividends = append(dividends, model.Dividend{
				Name:                         dividend.Name,
				Type:                         dividend.Type,
				Year:                         dividend.Year,
				TotalStockBonus:              dividend.TotalStockBonus,
				CashDividendPerShareCurrency: dividend.CashDividendPerShareCurrency,
				CashDividendPerShare:         dividend.CashDividendPerShare,
				CumDate:                      dividend.CumDate,
				ExDate:                       dividend.ExDate,
				RecordDate:                   dividend.RecordDate,
				PaymentDate:                  dividend.PaymentDate,
				Ratio1:                       dividend.Ratio1,
				Ratio2:                       dividend.Ratio2,
				CashDividendCurrency:         dividend.CashDividendCurrency,
				CashDividendTotal:            dividend.CashDividendTotal,
			})
		}

		data = append(data, model.StockResponse{
			Code:            list.StockCode,
			Name:            list.StockName,
			Share:           list.Share,
			ListingDate:     list.ListingDate,
			Board:           list.Board,
			Profiles:        profiles,
			Secretaries:     secretaries,
			Directors:       directors,
			Commissioners:   commissioners,
			AuditCommittees: auditCommittees,
			Shareholders:    shareholders,
			Subsidiaries:    subsidiaries,
			Dividends:       dividends,
		})
	}

	response.Success(w, data, "")
	return
}

func (s *stockHandler) FindHistoricalStock(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	stockCode := r.URL.Query().Get("stock_code")
	if stockCode == "" {
		response.BadRequest(w, "stock_code is required", nil)
		return
	}

	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		response.BadRequest(w, "Invalid start_date format. Use YYYY-MM-DD", nil)
		return
	}

	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		response.BadRequest(w, "Invalid end_date format. Use YYYY-MM-DD", nil)
		return
	}

	summaries, err := s.stockUsecase.FindHistorical(r.Context(), strings.ToUpper(stockCode), startDate, endDate)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}

	response.Success(w, summaries, "")
}

// FindStock find stock by stock code
// @Summary Find stock by stock code
// @Description Find stock by stock code
// @Tags Stock
// @Produce json
// @Param request query model.StockRequest true "query params"
// @Success 200 {object} model.StockResponse
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /api/v1/stock [get]
func (s *stockHandler) FindStock(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	stockCode := r.URL.Query().Get("stock_code")

	request := model.StockRequest{
		StockCode: stockCode,
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

	result, err := s.stockUsecase.FindStock(r.Context(), strings.ToUpper(request.StockCode))
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}

	if result == nil {
		response.NotFound(w, "")
		return
	}

	profiles := make([]model.Profile, 0, len(result.Profiles))
	for _, profile := range result.Profiles {
		profiles = append(profiles, model.Profile{
			Address:      profile.Address,
			BAE:          profile.BAE,
			Industry:     profile.Industry,
			SubIndustry:  profile.SubIndustry,
			Email:        profile.Email,
			Fax:          profile.Fax,
			MainBusiness: profile.MainBusiness,
			StockCode:    profile.StockCode,
			StockName:    profile.StockName,
			TIN:          profile.TIN,
			Sector:       profile.Sector,
			SubSector:    profile.SubSector,
			ListingDate:  profile.ListingDate,
			Phone:        profile.Phone,
			Website:      profile.Website,
			Status:       profile.Status,
			Logo:         profile.Logo,
		})
	}

	secretaries := make([]model.Secretary, 0, len(result.Secretaries))
	for _, secretary := range result.Secretaries {
		secretaries = append(secretaries, model.Secretary{
			Name:         secretary.Name,
			PhoneNumber:  secretary.PhoneNumber,
			Website:      secretary.Website,
			Email:        secretary.Email,
			Fax:          secretary.Fax,
			MobileNumber: secretary.MobileNumber,
		})
	}

	directors := make([]model.Director, 0, len(result.Directors))
	for _, director := range result.Directors {
		directors = append(directors, model.Director{
			Name:         director.Name,
			Position:     director.Position,
			IsAffiliated: director.IsAffiliated,
		})
	}

	commissioners := make([]model.Commissioner, 0, len(result.Commissioners))
	for _, commissioner := range result.Commissioners {
		commissioners = append(commissioners, model.Commissioner{
			Name:          commissioner.Name,
			Position:      commissioner.Position,
			IsIndependent: commissioner.IsIndependent,
		})
	}

	auditCommittees := make([]model.AuditCommittee, 0, len(result.AuditCommittees))
	for _, committee := range result.AuditCommittees {
		auditCommittees = append(auditCommittees, model.AuditCommittee{
			Name:     committee.Name,
			Position: committee.Position,
		})
	}

	shareholders := make([]model.Shareholder, 0, len(result.Shareholders))
	for _, shareholder := range result.Shareholders {
		shareholders = append(shareholders, model.Shareholder{
			Share:        shareholder.Share,
			Category:     shareholder.Category,
			Name:         shareholder.Name,
			IsController: shareholder.IsController,
			Percentage:   shareholder.Percentage,
		})
	}

	subsidiaries := make([]model.Subsidiary, 0, len(result.Subsidiaries))
	for _, subsidiary := range result.Subsidiaries {
		subsidiaries = append(subsidiaries, model.Subsidiary{
			BusinessFields:  subsidiary.BusinessFields,
			TotalAsset:      subsidiary.TotalAsset,
			Location:        subsidiary.Location,
			Currency:        subsidiary.Currency,
			Name:            subsidiary.Name,
			Percentage:      subsidiary.Percentage,
			Units:           subsidiary.Units,
			OperationStatus: subsidiary.OperationStatus,
			CommercialYear:  subsidiary.CommercialYear,
		})
	}

	dividends := make([]model.Dividend, 0, len(result.Dividends))
	for _, dividend := range result.Dividends {
		dividends = append(dividends, model.Dividend{
			Name:                         dividend.Name,
			Type:                         dividend.Type,
			Year:                         dividend.Year,
			TotalStockBonus:              dividend.TotalStockBonus,
			CashDividendPerShareCurrency: dividend.CashDividendPerShareCurrency,
			CashDividendPerShare:         dividend.CashDividendPerShare,
			CumDate:                      dividend.CumDate,
			ExDate:                       dividend.ExDate,
			RecordDate:                   dividend.RecordDate,
			PaymentDate:                  dividend.PaymentDate,
			Ratio1:                       dividend.Ratio1,
			Ratio2:                       dividend.Ratio2,
			CashDividendCurrency:         dividend.CashDividendCurrency,
			CashDividendTotal:            dividend.CashDividendTotal,
		})
	}

	data := model.StockResponse{
		Code:            result.StockCode,
		Name:            result.StockName,
		Share:           result.Share,
		ListingDate:     result.ListingDate,
		Board:           result.Board,
		Profiles:        profiles,
		Secretaries:     secretaries,
		Directors:       directors,
		Commissioners:   commissioners,
		AuditCommittees: auditCommittees,
		Shareholders:    shareholders,
		Subsidiaries:    subsidiaries,
		Dividends:       dividends,
	}
	response.Success(w, data, "")
	return
}
