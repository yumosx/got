package handlefn

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	config int
}

func NewBcrypt(config int) *Bcrypt {
	return &Bcrypt{config: config}
}

func (b *Bcrypt) HashSecret(args string) (string, error) {
	hPassword, err := bcrypt.GenerateFromPassword([]byte(args), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hPassword), nil
}

func (b *Bcrypt) Authenticate(hashedPassword, plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintextPassword))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
