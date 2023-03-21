package libs

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPass), nil 
}

func CheckPassword(bodyPass, dbPass string) bool {
	
	err := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(bodyPass))
	return err != nil 
	
}