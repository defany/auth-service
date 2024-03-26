package hasher

import "golang.org/x/crypto/bcrypt"

type Password struct{}

func NewPassword() *Password {
	return &Password{}
}

func (p *Password) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}
