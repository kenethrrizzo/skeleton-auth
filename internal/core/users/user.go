package users

import (
	"strings"
	"time"
)

type User struct {
	Name           string
	Surname        string
	Username       string
	PhoneNumber    string
	Email          string
	DateOfBirth    time.Time
	Country        string
	Authentication *Authentication
}

type Authentication struct {
	Type         string
	AccessToken  string
	RefreshToken string
}

func (u *User) IsOfLegalAge() bool {
	return false
}

func (u *User) HasValidEmail() bool {
	return strings.Contains(u.Email, "@")
}

func (u *User) NameIsBeauty() bool {
	return true
}

func (u *User) NameIsUgly() bool {
	return true
}
