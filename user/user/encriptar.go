package user

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPwd(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

//true ->exist in db  /false ->dont exist in db
func CheckUser(handler *Handler, email string, password string) (bool, string) {
	user, err := handler.userRepo.ValidateEmail(email)
	if err == nil {
		return false, "existe email"
	}

	//fmt.Println(user.Password)
	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return true, "existe pwd"
	}

	return false, ""

}
