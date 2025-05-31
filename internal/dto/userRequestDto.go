package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignup struct {
	UserLogin
	Phone string
}

type VerificationCodeInput struct {
	Code int `json:"code"`
}
