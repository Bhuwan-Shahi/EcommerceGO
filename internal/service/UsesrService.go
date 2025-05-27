package service

import (
	"ecommerceGO/internal/domain"
	"ecommerceGO/internal/dto"
	"ecommerceGO/internal/repository"
	"errors"
	"fmt"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) { //any is alias for empty interface

	user, _ := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})

	//Generate token
	log.Println(user)
	userInfo := fmt.Sprint("%v,%v,%v", user.ID, user.Email, user.UserType)

	//Perfomr DB operation and Business Logic
	log.Println(input)
	return userInfo, nil

}

func (s UserService) findserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUser(email)

	//Perform DB operation and Business Logic
	return &user, err

}

func (s UserService) Login(email, paswrod string) (string, error) { //any is alias for empty interface

	//Perfomr DB operation and Business Logic

	user, err := s.findserByEmail(email)

	if err != nil {
		return "", errors.New("user does not exits with provided email")
	}

	//Compare password and generate token
	return user.Email, nil

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
