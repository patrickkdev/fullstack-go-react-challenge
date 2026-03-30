package domain

import "github.com/google/uuid"

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash []byte
	SessionToken string
}

func NewUser(name, email string, passwordHash []byte) User {
	return User{Name: name, Email: email, PasswordHash: passwordHash, SessionToken: uuid.NewString()}
}

func (u *User) UpdateSessionToken() {
	u.SessionToken = uuid.NewString()
}

func (u *User) ClearSessionToken() {
	u.SessionToken = ""
}
