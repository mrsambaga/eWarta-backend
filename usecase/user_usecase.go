package usecase

import (
	"errors"
	"stage01-project-backend/dto"
	"stage01-project-backend/entity"
	"stage01-project-backend/httperror"
	"stage01-project-backend/repository"
	"stage01-project-backend/util"
)

type UsersUsecase interface {
	Register(newUserDTO *dto.RegisterRequestDTO) error
	Login(loginUserDTO *dto.LoginRequestDTO) (*dto.TokenResponse, error)
	GetProfile(id int) (*dto.UserRespondDTO, error)
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
	defaultRole := "user"

	if newUserDTO.Role == "" {
		newUserDTO.Role = defaultRole
	}

	existingUser, _ := u.usersRepository.GetUserByEmailRole(newUserDTO.Email, newUserDTO.Role)
	if existingUser != nil {
		return httperror.ErrEmailAlreadyRegistered
	}

	hashedPassword, err := util.HashPassword(newUserDTO.Password)
	if err != nil {
		return err
	}

	referralCode := util.GenerateReferral(newUserDTO.Name, newUserDTO.Email)

	if newUserDTO.RefReferral != "" {
		_, err = u.usersRepository.FindUserReferral(newUserDTO.RefReferral)
		if err != nil {
			if errors.Is(err, httperror.ErrUserNotFound) {
				return httperror.ErrInvalidReferral
			}
			return httperror.ErrInvalidReferral
		}
	}

	newUser := &entity.User{
		Name:        newUserDTO.Name,
		Email:       newUserDTO.Email,
		Password:    hashedPassword,
		Phone:       newUserDTO.Phone,
		Address:     newUserDTO.Address,
		Role:        newUserDTO.Role,
		Referral:    referralCode,
		RefReferral: newUserDTO.RefReferral,
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
		Role:     loginUserDTO.Role,
	}

	registeredUser, err := u.usersRepository.GetUserByEmailRole(loginUser.Email, loginUser.Role)
	if err != nil {
		if errors.Is(err, httperror.ErrUserNotFound) {
			return nil, httperror.ErrInvalidEmailPassword
		}

		return nil, err
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

func (u *usersUsecaseImp) GetProfile(id int) (*dto.UserRespondDTO, error) {
	user, err := u.usersRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	userDTO := &dto.UserRespondDTO{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Address:  user.Address,
		Quota:    user.Quota,
		Referral: user.Referral,
	}

	return userDTO, nil
}
