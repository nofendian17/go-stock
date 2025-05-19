package cron

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-stock/internal/app"
	"go-stock/internal/shared/cron"
)

func Start(ctx context.Context, bootstrap app.Bootstrap) {
	// Load timezone
	timezone := bootstrap.GetConfig().GetApplication().Timezone
	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Printf("⚠️ Invalid timezone %q: defaulting to UTC", timezone)
		location = time.UTC
	}

	// Initialize cron client
	client := cron.NewCronClient(timezone)
	config := bootstrap.GetConfig().GetCronJob()

	// Helper function to register jobs
	registerJob := func(jobName, schedule string, jobFunc func()) {
		if schedule == "" {
			log.Printf("❌ Skipping %s job: no schedule configured", jobName)
			return
		}

		jobID, err := client.AddJob(schedule, jobFunc)
		if err != nil {
			log.Fatalf("❌ Failed to schedule %s job: %v", jobName, err)
		}

		log.Printf("📌 Scheduled %s job (ID: %d)", jobName, jobID)
	}

	// Register: UpdateSummaries
	registerJob("UpdateSummaries", config.UpdateStockSummaryList, func() {
		now := time.Now().In(location)
		date := now.Format("20060102")

		err := bootstrap.GetUsecase().StockSummaryUsecase.UpdateSummaries(ctx, date)
		if err != nil {
			log.Printf("❌ Failed to update stock summaries for %s: %v", date, err)
			return
		}
		log.Printf("✅ FindOne summary job executed for %s at %s", date, now.Format(time.RFC3339))
	})

	// Register: UpdateStock
	registerJob("UpdateStock", config.UpdateStockList, func() {
		err := bootstrap.GetUsecase().StockUsecase.UpdateStock(ctx)
		if err != nil {
			log.Printf("❌ Failed to update stock list: %v", err)
			return
		}
		log.Printf("✅ FindOne list updated at %s", time.Now().In(location).Format(time.RFC3339))
	})

	// Register: UpdateBroker
	registerJob("UpdateBroker", config.UpdateBrokerList, func() {
		err := bootstrap.GetUsecase().BrokerUsecase.UpdateBroker(ctx)
		if err != nil {
			log.Printf("❌ Failed to update broker list: %v", err)
			return
		}
		log.Printf("✅ Broker list updated at %s", time.Now().In(location).Format(time.RFC3339))
	})

	// Register: UpdateFinancialReport
	registerJob("UpdateFinancialReport", config.UpdateFinancialReport, func() {
		now := time.Now().In(location)

		year := now.Year()
		month := int(now.Month())

		var period string

		switch {
		case month >= 4 && month <= 6:
			period = "TW1"
		case month >= 7 && month <= 9:
			period = "TW2"
		case month >= 10 && month <= 12:
			period = "TW3"
		case month >= 1 && month <= 3:
			period = "Audit"
			// Adjust year back since Audit refers to previous year
			year -= 1
		}

		log.Printf("Update financial report for year %d and period %s", year, period)

		err := bootstrap.GetUsecase().FinancialReportUseCase.UpdateFinancialReport(ctx, period, fmt.Sprintf("%d", year))
		if err != nil {
			log.Printf("❌ Failed to financial report: %v", err)
			return
		}
		log.Printf("✅ Financial report updated at %s", time.Now().In(location).Format(time.RFC3339))
	})

	// Start scheduler
	client.Start()
	log.Println("🚀 Cron scheduler started.")

	// Wait for shutdown
	<-ctx.Done()
	log.Println("🛑 Shutting down cron scheduler...")
	client.Stop()
	log.Println("✅ Cron scheduler shut down cleanly.")
}
