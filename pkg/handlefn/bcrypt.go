package handlefn

import (
	"errors"

	"github.com/yumosx/got/pkg/errx"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	config int
}

func NewBcrypt(config int) *Bcrypt {
	return &Bcrypt{config: config}
}

func (b *Bcrypt) HashSecret(args string) errx.Option[string] {
	hPassword, err := bcrypt.GenerateFromPassword([]byte(args), bcrypt.DefaultCost)
	if err != nil {
		return errx.Err[string](err)
	}
	return errx.Ok(string(hPassword))
}

func (b *Bcrypt) Authenticate(hashedPassword, plaintextPassword string) errx.Option[bool] {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintextPassword))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return errx.Ok(false)
	}
	if err != nil {
		return errx.Err[bool](err)
	}
	return errx.Ok(true)
}
