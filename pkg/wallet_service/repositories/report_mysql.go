package repositories

import (
	"Expense_Manager/pkg/wallet_service/models"
	"github.com/uptrace/bun"
	"time"
)

type MySQLReportRepo interface {
	GetDailyReport(walletID int, startDate time.Time, endDate time.Time) ([]models.WalletReport, error)
	GetWeeklyReport(walletID int, startDate time.Time, endDate time.Time) ([]models.WalletReport, error)
	GetMonthlyReport(walletID int, startDate time.Time, endDate time.Time) ([]models.WalletReport, error)
	GetAnnuallyReport(walletID int, startDate time.Time, endDate time.Time) ([]models.WalletReport, error)
}

type MySQLReportRepoImpl struct {
	db *bun.DB
}

func (r *MySQLReportRepoImpl) GetDailyReport(walletID int, startDate time.Time, endDate time.Time) ([]models.WalletReport, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MySQLReportRepoImpl) GetWeeklyReport(walletID int, startDate time.Time, endDate time.Time) ([]models.WalletReport, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MySQLReportRepoImpl) GetMonthlyReport(walletID int, startDate time.Time, endDate time.Time) ([]models.WalletReport, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MySQLReportRepoImpl) GetAnnuallyReport(walletID int, startDate time.Time, endDate time.Time) ([]models.WalletReport, error) {
	//TODO implement me
	panic("implement me")
}
