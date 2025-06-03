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

type SellerInput struct {
	FirstName         string `"first_name": "min"`
	LastName          string `"last_name": "shahi"`
	PhoneNumber       string `"phone_number": "9765225479"`
	BankAccountNumber uint   `"bankAccountNumber": 123456789`
	SwiftCode         string `"swiftCode":"DB120029"`
	PaymentType       string `"paymentType": "regular"`
}



