package hasher

type PasswordMock struct {
	passToReturn string
}

func NewPasswordMock(passToReturn string) *PasswordMock {
	return &PasswordMock{
		passToReturn: passToReturn,
	}
}

func (p *PasswordMock) GenerateFromPassword(password []byte) ([]byte, error) {
	return []byte(p.passToReturn), nil
}
