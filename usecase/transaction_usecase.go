package usecase

import (
	"fmt"
	"stage01-project-backend/dto"
	"stage01-project-backend/entity"
	"stage01-project-backend/repository"
)

type TransactionUsecase interface {
	CreateNewTransaction(newTransactionDTO *dto.TransactionRequestDTO, userId uint64) error
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

func (u *transactionUsecaseImp) CreateNewTransaction(newTransactionDTO *dto.TransactionRequestDTO, userId uint64) error {

	newInvoice := &entity.Invoice{
		UserId:      userId,
		Status:      newTransactionDTO.Status,
		Total:       newTransactionDTO.Total,
		PaymentDate: newTransactionDTO.PaymentDate,
	}

	if newTransactionDTO.VoucherId != 0 {
		newInvoice.VoucherId = &newTransactionDTO.VoucherId
	}

	fmt.Println("Voucher ID : ", newTransactionDTO.VoucherId)
	fmt.Println("Voucher ID Transaction : ", newInvoice.VoucherId)

	err := u.transactionRepository.CreateNewTransaction(newInvoice, newTransactionDTO.SubscriptionId)
	if err != nil {
		return err
	}

	return nil
}
