package domain

import (
	debug "UptimeMonitor/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryInterface interface {
	Get(userId int) (User, error)
	GetByEmail(email string) (User, error)
	Add(user User)
	EmailExists(email string) bool
}

func NewUserRepositoryFake() UserRepositoryInterface {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123"), 8)
	users := []User{
		{1, "farzan.b@gmail.com", string(hashedPassword), "Farzan"},
		{2, "f.behzadian@dunro.com", string(hashedPassword), "f.behzadian@dunro.com"},
	}

	return &userRepositoryFake{
		users: users,
		index: len(users),
	}
}

type userRepositoryFake struct {
	users []User
	index int
}

func (r *userRepositoryFake) Get(userId int) (User, error) {
	for _, user := range r.users {
		if user.Id == userId {
			return user, nil
		}
	}

	return User{}, errorString("User not found")
}

func (r *userRepositoryFake) GetByEmail(email string) (User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return User{}, errorString("User not found")
}

func (r *userRepositoryFake) Add(user User) {
	r.index++
	user.Id = r.index
	r.users = append(r.users, user)
	debug.Printf("User added; user:%v\n", user)
}

func (r *userRepositoryFake) EmailExists(email string) bool {
	for _, user := range r.users {
		if user.Email == email {
			return true
		}
	}

	return false
}
