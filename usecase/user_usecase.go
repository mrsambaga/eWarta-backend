package usecase

import (
	"regexp"
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

	if newUserDTO.Role == "" {
		newUserDTO.Role = "user"
	}

	err := checkValidEmail(newUserDTO.Email)
	if err != nil {
		return err
	}

	err = checkValidPassword(newUserDTO.Password)
	if err != nil {
		return err
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

func checkValidEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return httperror.ErrInvalidEmailFormat
	}
	return nil
}

func checkValidPassword(password string) error {
	if len(password) < 8 {
		return httperror.ErrInvalidPasswordLength
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
