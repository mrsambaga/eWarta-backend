package usecase

import (
	"stage01-project-backend/dto"
	"stage01-project-backend/entity"
	"stage01-project-backend/httperror"
	"stage01-project-backend/repository"
	"stage01-project-backend/util"
)

type UsersUsecase interface {
	Register(newUserDTO *dto.RegisterRequestDTO) error
	Login(loginUserDTO *dto.LoginRequestDTO) (*dto.TokenResponse, error)
}

type usersUsecaseImp struct {
	usersRepository repository.UsersRepository
}

type UsersUsecaseConfig struct {
	UsersRepository repository.UsersRepository
}

func NewUsersUsecase(cfg *UsersUsecaseConfig) UsersUsecase {
	return &usersUsecaseImp{
		usersRepository: cfg.UsersRepository,
	}
}

func (u *usersUsecaseImp) Register(newUserDTO *dto.RegisterRequestDTO) error {
	defaultQuota := 0
	defaultRole := "user"

	if newUserDTO.Role == "" {
		newUserDTO.Role = defaultRole
	}

	hashedPassword, err := util.HashPassword(newUserDTO.Password)
	if err != nil {
		return err
	}

	newUser := &entity.User{
		Name:     newUserDTO.Name,
		Email:    newUserDTO.Email,
		Password: hashedPassword,
		Phone:    newUserDTO.Phone,
		Address:  newUserDTO.Address,
		Role:     newUserDTO.Role,
		Quota:    defaultQuota,
	}

	err = u.usersRepository.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *usersUsecaseImp) Login(loginUserDTO *dto.LoginRequestDTO) (*dto.TokenResponse, error) {
	loginUser := &entity.User{
		Email:    loginUserDTO.Email,
		Password: loginUserDTO.Password,
	}

	registeredUser, err := u.usersRepository.GetUserByEmail(loginUser.Email)
	if err != nil {
		return nil, httperror.ErrInvalidEmailPassword
	}

	ok := util.ComparePassword(registeredUser.Password, loginUser.Password)
	if !ok {
		return nil, httperror.ErrInvalidEmailPassword
	}

	loginUser.UserId = registeredUser.UserId
	token, err := util.GenerateAccessToken(loginUser)
	if err != nil {
		return nil, err
	}

	return token, nil
}
