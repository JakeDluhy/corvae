package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(unhashed_pass string) (string, error) {
	byteUnhashed := []byte(unhashed_pass)
	defer clear(byteUnhashed)
	byteHashed, err := bcrypt.GenerateFromPassword(byteUnhashed, bcrypt.DefaultCost)
	return string(byteHashed), err
}

func ComparePasswords(hashedPassword, unhashedPassword string) error {
	byteUnhashed := []byte(unhashedPassword)
	byteHashed := []byte(hashedPassword)
	return bcrypt.CompareHashAndPassword(byteHashed, byteUnhashed)
}

func clear(b []byte) {
    for i := 0; i < len(b); i++ {
        b[i] = 0;
    }
}