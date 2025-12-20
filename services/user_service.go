package services

import (
	"errors"
	"fmt"

	"github.com/test-repo/test-repo-golang-support-v1/models"
)

type UserService struct {
	users map[int]*models.User
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[int]*models.User),
	}
}

func (s *UserService) CreateUser(id int, name, email string, age int) (*models.User, error) {
	if _, exists := s.users[id]; exists {
		return nil, errors.New("user already exists")
	}
	user := models.NewUser(name, email, age, true)
	user.ID = id
	s.users[id] = user
	return user, nil
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	user, exists := s.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) GetAllUsers() []*models.User {
	result := make([]*models.User, 0, len(s.users))
	for _, user := range s.users {
		result = append(result, user)
	}
	return result
}

func (s *UserService) PrintUserReport(id int) string {
	user, err := s.GetUser(id)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return fmt.Sprintf("User Report: %s (%s) - Age: %d", user.FullName, user.EmailAddress, user.Age)
}
