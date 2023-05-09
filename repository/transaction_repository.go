package repository

import (
	"errors"
	"stage01-project-backend/dto"
	"stage01-project-backend/entity"
	"stage01-project-backend/httperror"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateNewTransaction(newTransaction *entity.Transaction) (*entity.Transaction, error)
	FindTransactionsByUserId(userId int) ([]*entity.Transaction, error)
	UpdateTransaction(updatedTransactionDTO *dto.EditTransactionRequestDTO) error
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

func (r *transactionRepoImp) CreateNewTransaction(newTransaction *entity.Transaction) (*entity.Transaction, error) {
	if err := r.db.Create(newTransaction).Error; err != nil {
		return nil, err
	}

	return newTransaction, nil
}

func (r *transactionRepoImp) FindTransactionsByUserId(userId int) ([]*entity.Transaction, error) {
	transactions := []*entity.Transaction{}

	if err := r.db.Preload("Subscription").Where("user_id = ?", userId).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *transactionRepoImp) UpdateTransaction(updatedTransactionDTO *dto.EditTransactionRequestDTO) error {
	existingTransaction := &entity.Transaction{}

	if err := r.db.Transaction(func(tx *gorm.DB) error {

		// Get existing transaction
		if err := tx.Where("id = ?", updatedTransactionDTO.TransactionId).First(existingTransaction).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("transaction not found")
			}

			return err
		}

		// Compare payment amount & bill amount
		if existingTransaction.Total != updatedTransactionDTO.Total {
			return errors.New("payment amount not equal to subscription bill")
		}

		// Update
		existingTransaction.Status = "paid"
		existingTransaction.PaymentDate = time.Now()

		if err := tx.Save(existingTransaction).Error; err != nil {
			return httperror.ErrUpdateTransactions
		}

		// Create new user subscription record
		newUserSubscription := &entity.UserSubscription{
			UserId:         int(existingTransaction.UserId),
			SubscriptionId: int(existingTransaction.SubscriptionId),
			DateStart:      time.Now(),
			DateEnd:        time.Now().AddDate(0, 0, 30),
		}

		if err := tx.Create(newUserSubscription).Error; err != nil {
			return errors.New("failed to create new user subscription record")
		}

		// Update user quota
		user := &entity.User{}
		if err := tx.Where("user_id = ?", existingTransaction.UserId).First(user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return httperror.ErrUserNotFound
			}

			return err
		}

		subscription := &entity.Subscription{}
		if err := tx.Where("subscription_id = ?", existingTransaction.SubscriptionId).First(subscription).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("subscription not found")
			}

			return err
		}

		user.Quota += subscription.Quota
		if err := tx.Save(user).Error; err != nil {
			return errors.New("failed to update user")
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
