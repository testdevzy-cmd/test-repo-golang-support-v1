package models

import (
	"errors"
	"time"
)

type User struct {
	ID           int
	Name         string
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

func (u *User) UpdateProfile(name, email string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if email == "" {
		return errors.New("email cannot be empty")
	}
	u.Name = name
	u.EmailAddress = email
	return nil
}

func NewUser(name, email string, age int, isActive bool) *User {
	return &User{
		Name:         name,
		EmailAddress: email,
		Age:          age,
		IsActive:     isActive,
		CreatedAt:    time.Now(),
	}
}

func (u *User) GetDisplayName() string {
	if u.Name == "" {
		return "Anonymous"
	}
	return u.Name
}
