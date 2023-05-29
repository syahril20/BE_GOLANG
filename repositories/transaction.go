package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	FindTransactionId(Id int) (models.Transaction, error)
	GetTripId(Id int) (models.TripResponse, error)
	DeleteTransaction(Id int, Transaction models.Transaction) (models.Transaction, error)
	CreateTransaction(Transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(Id int, Transaction models.Transaction) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindTransaction() ([]models.Transaction, error) {
	var Tansactions []models.Transaction
	err := r.db.Preload("Trip.Country").Find(&Tansactions).Error

	return Tansactions, err
}

func (r *repositories) FindTransactionId(Id int) (models.Transaction, error) {
	var Tansactions models.Transaction
	err := r.db.Preload("Trip.Country").First(&Tansactions, Id).Error

	return Tansactions, err
}
func (r *repositories) GetTripId(Id int) (models.TripResponse, error) {
	var Tansactions models.TripResponse
	err := r.db.Preload("Country").First(&Tansactions, Id).Error

	return Tansactions, err
}

func (r *repositories) DeleteTransaction(Id int, Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&Transaction).Error

	return Transaction, err
}

func (r *repositories) CreateTransaction(Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Trip").Create(&Transaction).Error

	return Transaction, err
}

func (r *repositories) UpdateTransaction(Id int, Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&Transaction).Error

	return Transaction, err
}
