package service

import (
	"ecommerceGO/internal/domain"
	"ecommerceGO/internal/dto"
	"log"
)

type UserService struct {
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	log.Println(input)
	return "this is my token as of now", nil

}

func (s UserService) findserByEmail(email string) (*domain.User, error) {

	//Perfomr DB operation and Business Logic
	return nil, nil

}

func (s UserService) Login(input any) (string, error) { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	return "", nil

}

func (s UserService) GetVerificationCode(e domain.User) (int, error) { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	return 0, nil

}

func (s UserService) VerifyCode(id uint, code int) error { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	return nil

}

func (s UserService) CreateProfile(id uint, input any) error { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	return nil

}

func (s UserService) GetProfile(id uint) (*domain.User, error) { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	return nil, nil

}

func (s UserService) UpdateProfile(id uint, input any) error { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	return nil

}
func (s UserService) BecomeSeller(id uint, input any) (string, error) { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	return "", nil

}

func (s UserService) FindCart(id uint) ([]interface{}, error) { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	return nil, nil

}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	return nil, nil

}

func (s UserService) CreateOrder(u domain.User) (int, error) { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	return 0, nil

}

func (s UserService) GetOrder(u domain.User) ([]interface{}, error) { //any is alias for empty interface

	//Perfomr DB operation and Business Logic

	return nil, nil

}

func (s UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	return nil, nil

}
