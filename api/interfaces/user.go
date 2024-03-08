package interfaces

import (
	"errors"
	models "gomboc/api/models"
)

func GetUserById(id string) (user models.UserModel, e error) {
	return user, errors.New("function 'GetUser' not implemented yet")
}

func GetAllUsers(page int, limit int) (user models.UserModel, e error) {
	return user, errors.New("function 'GetAllUsers' not implemented yet")
}

func CreateUser() (user models.UserModel, e error) {
	return user, errors.New("function 'CreateUser' not implemented yet")
}

func UpdateUser() (user models.UserModel, e error) {
	return user, errors.New("function 'UpdateUser' not implemented yet")
}

func DeleteUser() (user models.UserModel, e error) {
	return user, errors.New("function 'DeleteUser' not implemented yet")
}

func GetUserInfo(id string) (user models.UserModel, e error) {
	return user, errors.New("function 'GetUserInfo' not implemented yet")
}
