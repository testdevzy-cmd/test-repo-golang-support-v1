package services

import (
	"fmt"

	"github.com/test-repo/test-repo-golang-support-v1/models"
	"github.com/test-repo/test-repo-golang-support-v1/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(id int, name, email string, age int) (*models.User, error) {
	user := models.NewUser(name, email, age, true)
	user.ID = id
	err := s.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetAllUsers() []*models.User {
	return s.repo.ListAll()
}

func (s *UserService) PrintUserReport(id int) string {
	user, err := s.GetUser(id)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return fmt.Sprintf("User Report: %s (%s) - Age: %d", user.FullName, user.EmailAddress, user.Age)
}
