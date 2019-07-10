package crypto

import "golang.org/x/crypto/bcrypt"

// Hash returns the bcrypt hash of a password.
func Hash(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(b), err
}

// Check ensures that password matches the hashed password.
func Check(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
