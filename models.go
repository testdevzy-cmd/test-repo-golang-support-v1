package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID           int
	FullName     string
	EmailAddress string
	Age          int
	IsActive     bool
	CreatedAt    time.Time
}

type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

func (a Address) String() string {
	return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.Country, a.ZipCode)
}

func (a Address) HasZipCode() bool {
	if len(a.ZipCode) > 0 {
		return true
	}
	return false
}

type Profile struct {
	User
	Address Address
	Bio     string
}

type Validator interface {
	Validate() error
}

type Repository interface {
	Save(data interface{}) error
	FindByID(id int) (interface{}, error)
	Delete(id int) error
}

type UserRepository struct {
	data map[int]*User
}

func NewUser(name, email string, age int, isActive bool) *User {
	user := &User{
		FullName:     name,
		EmailAddress: email,
		Age:          age,
		IsActive:     isActive,
		CreatedAt:    time.Now(),
	}
	return user
}

func (u *User) UpdateEmail(newEmail string) error {
	if newEmail == "" {
		return errors.New("email cannot be empty")
	}
	u.EmailAddress = newEmail
	return nil
}

func (u *User) IsAdult() bool {
	if u.Age >= 18 {
		return true
	}
	return false
}

func (u User) GetDisplayName() string {
	displayName := u.FullName
	if displayName == "" {
		displayName = "Anonymous"
	}
	return displayName
}

func (u *User) Activate() error {
	if u.IsActive {
		return errors.New("user is already active")
	}
	u.IsActive = true
	return nil
}

func (u *User) Deactivate() error {
	if !u.IsActive {
		return errors.New("user is already inactive")
	}
	u.IsActive = false
	return nil
}

func (r *UserRepository) Save(user *User) error {
	if r.data == nil {
		r.data = make(map[int]*User)
	}
	r.data[user.ID] = user
	return nil
}

func (r *UserRepository) FindByID(id int) (*User, error) {
	user, exists := r.data[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) Delete(id int) error {
	if r.data == nil {
		return errors.New("repository not initialized")
	}
	delete(r.data, id)
	return nil
}

func (p *Profile) Validate() error {
	if p.FullName == "" {
		return errors.New("name is required")
	}
	if p.EmailAddress == "" {
		return errors.New("email is required")
	}
	return nil
}

func (u *User) UpdateProfile(name, email string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if email == "" {
		return errors.New("email cannot be empty")
	}
	u.FullName = name
	u.EmailAddress = email
	return nil
}
