package repository

import (
	"errors"
	"fmt"
	"stage01-project-backend/entity"
	"stage01-project-backend/httperror"

	"gorm.io/gorm"
)

type UsersRepository interface {
	CreateUser(newUser *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
	FindUserReferral(referral string) (*entity.User, error)
}

type userRepositoryImp struct {
	db *gorm.DB
}

type UserRConfig struct {
	DB *gorm.DB
}

func NewUserRepository(cfg *UserRConfig) UsersRepository {
	return &userRepositoryImp{
		db: cfg.DB,
	}
}

func (r *userRepositoryImp) CreateUser(newUser *entity.User) error {
	if err := r.db.Create(newUser).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return httperror.ErrEmailAlreadyRegistered
		}

		return httperror.ErrCreateUser
	}

	return nil
}

func (r *userRepositoryImp) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, httperror.ErrUserNotFound
		}

		return nil, httperror.ErrFailedGetUserByEmail
	}

	return user, nil
}

func (r *userRepositoryImp) FindUserReferral(referral string) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.Where("referral = ?", referral).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("INI EKSEKUSI 1")
			return nil, httperror.ErrUserNotFound
		}

		fmt.Println("INI EKSEKUSI 2")
		return nil, err
	}

	fmt.Println("INI EKSEKUSI 3")
	return user, nil
}
