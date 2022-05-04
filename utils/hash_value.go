package utils

import "golang.org/x/crypto/bcrypt"

func Generate(raw string) (string,error) {
	hashValue ,err := bcrypt.GenerateFromPassword([]byte(raw),10)
	if err != nil {
		return "",err
	}
	return string(hashValue),nil
}
func Verify(hash string,raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash),[]byte(raw))
}
