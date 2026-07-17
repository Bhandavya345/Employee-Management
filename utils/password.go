package utils

import "golang.org/x/crypto/bcrypt"

// Hash Password
func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(bytes), err
}

// Compare Password
func CheckPassword(password, hashedPassword string) error {

	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}
