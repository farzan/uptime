package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Get(userId int) (User, error)
	GetByEmail(email string) (User, error)
	Create(email string, password string) User
	EmailExists(email string) bool
	Add(user User)
}

func NewUserService() UserServiceInterface {
	return &userService{}
}

type userService struct {

}

func (s *userService) Get(userId int) (User, error) {
	return GetUserRepository().Get(userId)
}

func (s *userService) GetByEmail(email string) (User, error) {
	return GetUserRepository().GetByEmail(email)
}

func (s *userService) Create(email string, password string) User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	user := User{
		Email: email,
		Password: string(hashedPassword),
		Name: email,
	}

	return user
}

func (s *userService) Add(user User) {
	GetUserRepository().Add(user)
}

func (s *userService) EmailExists(email string) bool {
	return GetUserRepository().EmailExists(email)
}
