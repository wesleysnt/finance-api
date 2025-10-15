package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (pass string, err error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	pass = string(hashedPass)
	return
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil

}
