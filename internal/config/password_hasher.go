package config

import (
	"github.com/fkrhykal/upside-api/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHasher struct {}

func (b *BcryptPasswordHasher) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func (b *BcryptPasswordHasher) Verify(hashedPassword string, rawPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
}

func NewBcryptPasswordHasher() util.PasswordHasher {
	return &BcryptPasswordHasher{}
}