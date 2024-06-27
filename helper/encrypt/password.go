package encrypt

import "golang.org/x/crypto/bcrypt"

type HashInterface interface {
	Compare(hashed, input string) error
	HashPassword(string) (string, error)
}

type hash struct{}

func New() HashInterface {
	return &hash{}
}

func (h *hash) Compare(hashed, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
}

func (h *hash) HashPassword(password string) (string, error) {
	const salt = 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
