package usecase

import (
	"context"
	"fmt"
	"go-stock/internal/config"
	"go-stock/internal/entity"
	"go-stock/internal/infrastructure/idx"
	"go-stock/internal/repository"
)

type FinancialReportUseCase interface {
	UpdateFinancialReport(ctx context.Context, period string, year string) error
	Find(ctx context.Context, stockCode, reportPeriod, reportYear string) (*entity.FinancialReport, error)
}

type financialReportUseCase struct {
	financialReportRepository repository.FinancialReportRepository
	idxClient                 idx.IdxClient
	cfg                       config.Config
}

func NewFinancialReportUseCase(cfg config.Config, idxClient idx.IdxClient, financialReportRepository repository.FinancialReportRepository) FinancialReportUseCase {
	return &financialReportUseCase{
		financialReportRepository: financialReportRepository,
		idxClient:                 idxClient,
		cfg:                       cfg,
	}
}

func (b *financialReportUseCase) UpdateFinancialReport(ctx context.Context, period string, year string) error {
	list, err := b.idxClient.GetFinancialReports(ctx, period, year)
	if err != nil {
		return err
	}

	var financialReports []entity.FinancialReport
	for _, financialReport := range list.Results {
		var attachments []entity.Attachment
		for _, attachment := range financialReport.Attachments {
			attachments = append(attachments, entity.Attachment{
				StockCode:    attachment.EmitenCode,
				StockName:    attachment.NamaEmiten,
				FileID:       attachment.FileID,
				FileModified: attachment.FileModified,
				FileName:     attachment.FileName,
				FilePath:     fmt.Sprintf("%s%s", b.cfg.GetService().IDXService.BaseURL, attachment.FilePath),
				FileSize:     attachment.FileSize,
				FileType:     attachment.FileType,
				ReportPeriod: attachment.ReportPeriod,
				ReportType:   attachment.ReportType,
				ReportYear:   attachment.ReportYear,
			})
		}

		financialReports = append(financialReports, entity.FinancialReport{
			StockCode:    financialReport.KodeEmiten,
			FileModified: financialReport.FileModified,
			ReportPeriod: financialReport.ReportPeriod,
			ReportYear:   financialReport.ReportYear,
			StockName:    financialReport.NamaEmiten,
			Attachment:   attachments,
		})
	}

	if len(financialReports) == 0 {
		return nil // no broker data to update
	}

	if err := b.financialReportRepository.BulkUpsert(ctx, financialReports); err != nil {
		return fmt.Errorf("bulk upsert failed: %w", err)
	}

	return nil
}

func (b *financialReportUseCase) Find(ctx context.Context, stockCode, reportPeriod, reportYear string) (*entity.FinancialReport, error) {
	return b.financialReportRepository.Find(ctx, stockCode, reportPeriod, reportYear)
}
