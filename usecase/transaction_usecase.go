package usecase

import (
	"stage01-project-backend/dto"
	"stage01-project-backend/entity"
	"stage01-project-backend/repository"
)

type TransactionUsecase interface {
	CreateNewTransaction(newTransactionDTO *dto.TransactionRequestDTO, userId uint64) (*dto.TransactionResponseDTO, error)
	FindUserTransactions(userId int) ([]*dto.TransactionResponseDTO, error)
	UpdateTransaction(updatedTransactionDTO *dto.EditTransactionRequestDTO) error
}

type transactionUsecaseImp struct {
	transactionRepository repository.TransactionRepository
}

type TransactionUConfig struct {
	TransactionRepository repository.TransactionRepository
}

func NewSubscriptionUsecase(cfg *TransactionUConfig) TransactionUsecase {
	return &transactionUsecaseImp{
		transactionRepository: cfg.TransactionRepository,
	}
}

func (u *transactionUsecaseImp) CreateNewTransaction(newTransactionDTO *dto.TransactionRequestDTO, userId uint64) (*dto.TransactionResponseDTO, error) {

	newTransaction := &entity.Transaction{
		UserId:         userId,
		Status:         newTransactionDTO.Status,
		Total:          newTransactionDTO.Total,
		PaymentDate:    newTransactionDTO.PaymentDate,
		SubscriptionId: newTransactionDTO.SubscriptionId,
		VoucherId:      &newTransactionDTO.VoucherId,
	}

	if newTransactionDTO.VoucherId == 0 {
		newTransaction.VoucherId = nil
	}

	transaction, err := u.transactionRepository.CreateNewTransaction(newTransaction)
	if err != nil {
		return nil, err
	}

	transactionResponse := &dto.TransactionResponseDTO{
		TransactionId: transaction.Id,
		Status:        transaction.Status,
		Total:         transaction.Total,
		Subscription:  transaction.Subscription.Name,
	}

	return transactionResponse, nil
}

func (u *transactionUsecaseImp) FindUserTransactions(userId int) ([]*dto.TransactionResponseDTO, error) {
	transactions, err := u.transactionRepository.FindTransactionsByUserId(userId)
	if err != nil {
		return nil, err
	}

	transactionsDTO := make([]*dto.TransactionResponseDTO, 0, len(transactions))
	for _, transaction := range transactions {
		transaction := &dto.TransactionResponseDTO{
			TransactionId: transaction.Id,
			Status:        transaction.Status,
			Total:         transaction.Total,
			Subscription:  transaction.Subscription.Name,
		}
		transactionsDTO = append(transactionsDTO, transaction)
	}

	return transactionsDTO, nil
}

func (u *transactionUsecaseImp) UpdateTransaction(updatedTransactionDTO *dto.EditTransactionRequestDTO) error {
	err := u.transactionRepository.UpdateTransaction(updatedTransactionDTO)
	if err != nil {
		return err
	}
	return nil
}
