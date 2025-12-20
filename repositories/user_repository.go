package repositories

import (
	"errors"

	"github.com/test-repo/test-repo-golang-support-v1/models"
)

type UserRepository struct {
	data map[int]*models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		data: make(map[int]*models.User),
	}
}

func (r *UserRepository) Save(u *models.User) error {
	if u == nil {
		return errors.New("cannot save nil user")
	}
	r.data[u.ID] = u
	return nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	u, ok := r.data[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (r *UserRepository) ListAll() []*models.User {
	users := make([]*models.User, 0, len(r.data))
	for _, u := range r.data {
		users = append(users, u)
	}
	return users
}
