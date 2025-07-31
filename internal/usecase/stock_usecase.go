package usecase

import (
	"context"
	"fmt"
	"go-stock/internal/config"
	"go-stock/internal/entity"
	"go-stock/internal/infrastructure/idx"
	"go-stock/internal/repository"
	"go-stock/internal/shared/helper"
)

type StockUseCase interface {
	UpdateStock(ctx context.Context) error
	ListStocks(ctx context.Context) ([]entity.Stock, error)
	FindStock(ctx context.Context, code string) (*entity.Stock, error)
	ListStocksWithPagination(ctx context.Context, limit, offset int64) ([]entity.Stock, int64, error)
	SearchStocks(ctx context.Context, query string) ([]entity.Stock, error)
}
type stockUseCase struct {
	stockRepository repository.StockRepository
	idxClient       idx.IdxClient
	cfg             config.Config
}

func NewStockUsecase(cfg config.Config, idxClient idx.IdxClient, stockRepository repository.StockRepository) StockUseCase {
	return &stockUseCase{
		stockRepository: stockRepository,
		idxClient:       idxClient,
		cfg:             cfg,
	}
}

func (s *stockUseCase) UpdateStock(ctx context.Context) error {
	list, err := s.idxClient.GetStockList(ctx)
	if err != nil {
		return err
	}

	var stocks []entity.Stock

	for _, stock := range list.StockListData {
		company, err := s.idxClient.GetCompanyProfile(ctx, stock.Code)
		if err != nil {
			return fmt.Errorf("invalid stock company %s: %w", stock.Code, err)
		}

		profiles := make([]entity.Profile, 0, len(company.Profiles))
		for _, profile := range company.Profiles {
			profiles = append(profiles, entity.Profile{
				Address:      profile.Alamat,
				BAE:          profile.Bae,
				Industry:     profile.Industri,
				SubIndustry:  profile.SubIndustri,
				Email:        profile.Email,
				Fax:          profile.Fax,
				MainBusiness: profile.KegiatanUsahaUtama,
				StockCode:    profile.KodeEmiten,
				StockName:    profile.NamaEmiten,
				TIN:          profile.Npwp,
				Sector:       profile.Sektor,
				SubSector:    profile.SubSektor,
				ListingDate:  profile.TanggalPencatatan,
				Phone:        profile.Telepon,
				Website:      profile.Website,
				Status:       profile.Status,
				Logo:         fmt.Sprintf("%s%s", s.cfg.GetService().IDXService.BaseURL, profile.Logo),
			})
		}

		secretaries := make([]entity.Secretary, 0, len(company.Sekretaris))
		for _, secretary := range company.Sekretaris {
			secretaries = append(secretaries, entity.Secretary{
				Name:         secretary.Nama,
				PhoneNumber:  secretary.Telepon,
				Website:      secretary.Website,
				Email:        secretary.Email,
				Fax:          secretary.Fax,
				MobileNumber: secretary.Hp,
			})
		}

		directors := make([]entity.Director, 0, len(company.Direktur))
		for _, director := range company.Direktur {
			directors = append(directors, entity.Director{
				Name:         director.Nama,
				Position:     director.Jabatan,
				IsAffiliated: director.Afiliasi,
			})
		}

		commissioners := make([]entity.Commissioner, 0, len(company.Komisaris))
		for _, commissioner := range company.Komisaris {
			commissioners = append(commissioners, entity.Commissioner{
				Name:          commissioner.Nama,
				Position:      commissioner.Jabatan,
				IsIndependent: commissioner.Independen,
			})
		}

		auditCommittees := make([]entity.AuditCommittee, 0, len(company.KomiteAudit))
		for _, auditCommittee := range company.KomiteAudit {
			auditCommittees = append(auditCommittees, entity.AuditCommittee{
				Name:     auditCommittee.Nama,
				Position: auditCommittee.Jabatan,
			})
		}

		shareHolders := make([]entity.Shareholder, 0, len(company.PemegangSaham))
		for _, shareHolder := range company.PemegangSaham {
			shareHolders = append(shareHolders, entity.Shareholder{
				Share:        shareHolder.Jumlah,
				Category:     shareHolder.Kategori,
				Name:         shareHolder.Nama,
				IsController: shareHolder.Pengendali,
				Percentage:   shareHolder.Persentase,
			})
		}

		subsidiaries := make([]entity.Subsidiary, 0, len(company.AnakPerusahaan))
		for _, subsidiary := range company.AnakPerusahaan {
			subsidiaries = append(subsidiaries, entity.Subsidiary{
				BusinessFields:  subsidiary.BidangUsaha,
				TotalAsset:      subsidiary.JumlahAset,
				Location:        subsidiary.Lokasi,
				Currency:        subsidiary.MataUang,
				Name:            subsidiary.Nama,
				Percentage:      subsidiary.Persentase,
				Units:           subsidiary.Satuan,
				OperationStatus: subsidiary.StatusOperasi,
				CommercialYear:  subsidiary.TahunKomersil,
			})
		}

		dividends := make([]entity.Dividend, 0, len(company.Dividen))
		for _, dividend := range company.Dividen {
			dividends = append(dividends, entity.Dividend{
				Name:                         dividend.Nama,
				Type:                         dividend.Jenis,
				Year:                         dividend.TahunBuku,
				TotalStockBonus:              dividend.TotalSahamBonus,
				CashDividendPerShareCurrency: dividend.CashDividenTotalMU,
				CashDividendPerShare:         dividend.CashDividenPerSaham,
				CumDate:                      helper.StringToDate(dividend.TanggalCum),
				ExDate:                       helper.StringToDate(dividend.TanggalExRegulerDanNegosiasi),
				RecordDate:                   helper.StringToDate(dividend.TanggalDPS),
				PaymentDate:                  helper.StringToDate(dividend.TanggalPembayaran),
				Ratio1:                       dividend.Rasio1,
				Ratio2:                       dividend.Rasio2,
				CashDividendCurrency:         dividend.CashDividenTotalMU,
				CashDividendTotal:            dividend.CashDividenTotal,
			})
		}

		stocks = append(stocks, entity.Stock{
			StockCode:       stock.Code,
			StockName:       stock.Name,
			Share:           stock.Share,
			ListingDate:     helper.StringToDate(stock.ListingDate),
			Board:           stock.Board,
			Profiles:        profiles,
			Secretaries:     secretaries,
			Directors:       directors,
			Commissioners:   commissioners,
			AuditCommittees: auditCommittees,
			Shareholders:    shareHolders,
			Subsidiaries:    subsidiaries,
			Dividends:       dividends,
		})
	}

	if len(stocks) == 0 {
		return nil // no stock data to update
	}

	if err := s.stockRepository.BulkUpsert(ctx, stocks); err != nil {
		return fmt.Errorf("bulk upsert failed: %w", err)
	}

	return nil
}

func (s *stockUseCase) ListStocks(ctx context.Context) ([]entity.Stock, error) {
	return s.stockRepository.All(ctx)
}

func (s *stockUseCase) FindStock(ctx context.Context, code string) (*entity.Stock, error) {
	return s.stockRepository.FindOne(ctx, code)
}

func (s *stockUseCase) ListStocksWithPagination(ctx context.Context, limit, offset int64) ([]entity.Stock, int64, error) {
	return s.stockRepository.FindWithPagination(ctx, limit, offset)
}

func (s *stockUseCase) SearchStocks(ctx context.Context, query string) ([]entity.Stock, error) {
	return s.stockRepository.Search(ctx, query)
}
