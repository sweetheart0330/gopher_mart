package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher struct {
	pepper string
	cost   int
}

func NewPasswordHasher(pepper string, cost int) *PasswordHasher {
	return &PasswordHasher{pepper: pepper, cost: cost}
}

func (h *PasswordHasher) Hash(password string) (string, error) {
	pepperedPassword := h.pepper + password
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(pepperedPassword), h.cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedBytes), nil
}

func (h *PasswordHasher) Verify(password string, hashedPassword string) (bool, error) {
	pepperedPassword := h.pepper + password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pepperedPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}
