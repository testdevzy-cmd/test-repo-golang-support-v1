package models

import (
	"errors"
	"time"
)

type User struct {
	ID              int
	DisplayIdentity string
	EmailAddress    string
	Age             int
	IsActive        bool
	CreatedAt       time.Time
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
	u.DisplayIdentity = name
	u.EmailAddress = email
	return nil
}

func NewUser(name, email string, age int, isActive bool) *User {
	return &User{
		DisplayIdentity: name,
		EmailAddress:    email,
		Age:             age,
		IsActive:        isActive,
		CreatedAt:       time.Now(),
	}
}

func (u *User) GetDisplayName() string {
	if u.DisplayIdentity == "" {
		return "Anonymous"
	}
	return u.DisplayIdentity
}
