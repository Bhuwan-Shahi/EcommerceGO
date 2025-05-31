package service

import (
	"ecommerceGO/internal/domain"
	"ecommerceGO/internal/dto"
	"ecommerceGO/internal/helper"
	"ecommerceGO/internal/repository"
	"errors"
	"log"
	"time"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) {
	// Hash the password
	hPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}

	// Create the user
	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hPassword,
		Phone:    input.Phone,
		UserType: "buyer",
	})
	if err != nil {
		log.Println("Error creating user:", err)
		return "", err
	}

	log.Println("User created:", user)

	// Generate a token for the user
	token, err := s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
	if err != nil {
		log.Println("Error generating token:", err)
		return "", err
	}

	return token, nil
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
	err = s.Auth.VerifyPassword(paswrod, user.Password)

	if err != nil {
		return "", err
	}

	//generate token
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)

}

func (s UserService) isVerifiedUser(id uint) bool {
	curentUser, err := s.Repo.FindUserById(id)
	return err == nil && curentUser.Verified
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) { //any is alias for empty interface

	//if usser already verified

	if s.isVerifiedUser(e.ID) {
		return 0, errors.New("user already verified")
	}

	//generate verification code

	code, err := s.Auth.GenerateCode()

	if err != nil {
		return 0, nil
	}

	//update user with verificatio code
	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	_, err = s.Repo.UpdateUser(e.ID, user)
	if err != nil {
		return 0, errors.New("unable to update verificaiton code")
	}

	//send SMS

	//return verificaiton code
	return code, nil

}

func (s UserService) VerifyCode(id uint, code int) error { //any is alias for empty interface

	//Perfomr DB operation and Business Logic
	if s.isVerifiedUser(id) {
		log.Print("verified")
		return errors.New("user already verified")
	}

	user, err := s.Repo.FindUserById(id)

	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("verification code doesnot match")
	}
	if !time.Now().Before(user.Expiry) {
		return errors.New("verirication code expired")

	}

	updateUsesr := domain.User{
		Verified: true,
	}
	_, err = s.Repo.UpdateUser(id, updateUsesr)
	if err != nil {
		return errors.New("unable to verify user")
	}

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
