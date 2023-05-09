package repository

import (
	"fmt"
	"stage01-project-backend/entity"
	"stage01-project-backend/httperror"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateNewTransaction(newInvoice *entity.Invoice, subscriptionId uint64) error
}

type transactionRepoImp struct {
	db *gorm.DB
}

type SubscriptionRepoConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(cfg *SubscriptionRepoConfig) TransactionRepository {
	return &transactionRepoImp{
		db: cfg.DB,
	}
}

func (r *transactionRepoImp) CreateNewTransaction(newInvoice *entity.Invoice, transactionId uint64) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Create(newInvoice).Error; err != nil {
			return err
		}

		fmt.Println("INVOICE ID : ", newInvoice.Id)
		newTransaction := &entity.Transaction{
			InvoiceId:      newInvoice.Id,
			SubscriptionId: transactionId,
		}

		if err := tx.Create(newTransaction).Error; err != nil {
			return httperror.ErrCreateTransaction
		}

		fmt.Println("EKSEKUSI 1")
		return nil
	}); err != nil {
		return err
	}

	fmt.Println("EKSEKUSI 2")
	return nil
}
