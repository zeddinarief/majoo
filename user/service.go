package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Service yang berfungsi untuk menangani user
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	SaveFoto(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
	GetAllUsers() ([]User, error)
	// UpdateUser(input FormUpdateUserInput) (User, error)
}

type service struct {
	repository Repository
}

// NewService init
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Username = input.Username
	user.Password = input.Password
	user.NamaLengkap = input.NamaLengkap

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) SaveFoto(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.Foto = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

// func (s *service) UpdateUser(input FormUpdateUserInput) (User, error) {
// 	user, err := s.repository.FindByID(input.ID)
// 	if err != nil {
// 		return user, err
// 	}

// 	user.Username = input.Username
// 	user.Password = input.Password
// 	user.NamaLengkap = input.NamaLengkap

// 	updatedUser, err := s.repository.Update(user)
// 	if err != nil {
// 		return updatedUser, err
// 	}

// 	return updatedUser, nil
// }
