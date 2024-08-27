package utility

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	cost := 10
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func CheckPass(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
