package validator

import (
    "github.com/google/uuid"
	validate "github.com/gobuffalo/validate/v3"

	"testing"
)

// структура для проверки валидации
type ValidateData struct {
	ID 					uuid.UUID `json:"id" myvalid:"required"`
	Email 				string `json:"email" myvalid:"email"`
	Login 				string `json:"login" myvalid:"required|max:10"`
	Password 			string `json:"password" myvalid:"required|min:5"`
	Weight 				float64 `json:"weight" myvalid:"required"`
	Sale 				float64 `json:"sale" myvalid:"min:0|max:100"`
	Age 				int `json:"age" myvalid:"required|min:0"`
	HandFingersAmount 	int `json:"handFingersAmount" myvalid:"max:20"`
}
// обязательный метод для валидации
func (self ValidateData) IsValid(errors *validate.Errors) {}


var validData = ValidateData{
	ID: uuid.New(),
	Email: "example@gmail.com",
	Login: "user1",
	Password: "qwerty",
	Weight: 65.7,
	Sale: 15.5,
	Age: 20,
	HandFingersAmount: 10,
}


func TestNoOneErrorCases(t *testing.T) {
	t.Log("Check full valid data")
	{
		err := Validate(&validData)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. No one error")
		}
	}

	t.Log("Check empty sale")
	{
		modifiedData := validData
		modifiedData.Sale = 0.0

		err := Validate(&modifiedData)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. No one error")
		}
	}

	t.Log("Check empty hand fingers amount")
	{
		modifiedData := validData
		modifiedData.HandFingersAmount = 0

		err := Validate(&modifiedData)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. No one error")
		}
	}
}


func TestOneErrorEmptyValuesCases(t *testing.T) {
	t.Log("Check empty ID")
	{
		modifiedData := validData
		modifiedData.ID = [16]byte{}

		err := Validate(&modifiedData)
		if err == nil || err.Error() != "code=400, message=map[validate_id:id field must not be blank]" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with empty ID")
		}
	}

	t.Log("Check empty email")
	{
		modifiedData := validData
		modifiedData.Email = ""

		err := Validate(&modifiedData)
		if err == nil || err.Error() != "code=400, message=map[validate_email:Email is not in the right format]" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with incorrect email")
		}
	}

	t.Log("Check empty weight")
	{
		modifiedData := validData
		modifiedData.Weight = 0.0

		err := Validate(&modifiedData)
		if err == nil || err.Error() != "code=400, message=map[validate_weight:weight field must not be blank]" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with empty weight")
		}
	}

	t.Log("Check empty age")
	{
		modifiedData := validData
		modifiedData.Age = 0

		err := Validate(&modifiedData)
		if err == nil || err.Error() != "code=400, message=map[validate_age:age field must not be blank]" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with empty age")
		}
	}
}


func TestOneErrorInvalidValuesCases(t *testing.T) {
	t.Log("Check invalid email")
	{
		modifiedData := validData
		modifiedData.Email = "invalid_email"

		err := Validate(&modifiedData)
		if err == nil || err.Error() != "code=400, message=map[validate_email:Email is not in the right format]" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with invalid email")
		}
	}

	t.Log("Check invalid login")
	{
		modifiedData := validData
		modifiedData.Login = "qwerty0123456789"

		err := Validate(&modifiedData)
		if err == nil || err.Error() != "code=400, message=map[validate_login:login field must contain less than 10 symbols]" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with invalid login")
		}
	}

	t.Log("Check invalid password")
	{
		modifiedData := validData
		modifiedData.Password = "123"

		err := Validate(&modifiedData)
		if err == nil || err.Error() != "code=400, message=map[validate_password:password field must contain at least 5 symbols]" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with invalid password")
		}
	}

	t.Log("Check invalid sale")
	{
		modifiedData := validData
		modifiedData.Sale = 120.0

		err := Validate(&modifiedData)
		if err == nil || err.Error() != "code=400, message=map[validate_sale:sale field must be less than or equal to 100.000000]" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with invalid sale")
		}
	}

	t.Log("Check invalid age")
	{
		modifiedData := validData
		modifiedData.Age = -50

		err := Validate(&modifiedData)
		if err == nil || err.Error() != "code=400, message=map[validate_age:age field must be greater than or equal to 0]" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with invalid age")
		}
	}

	t.Log("Check invalid handFingersAmount")
	{
		modifiedData := validData
		modifiedData.HandFingersAmount = 100

		err := Validate(&modifiedData)
		if err == nil || err.Error() != "code=400, message=map[validate_handFingersAmount:handFingersAmount field must be less than or equal to 20]" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with invalid handFingersAmount")
		}
	}
}


func TestManyErrorsCases(t *testing.T) {
	t.Log("Check invalid login and empty age")
	{
		modifiedData := validData
		modifiedData.Login = "qwerty0123456789"
		modifiedData.Age = 0

		err := Validate(&modifiedData)
		if err == nil || err.Error() != "code=400, message=map[validate_age:age field must not be blank validate_login:login field must contain less than 10 symbols]" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 2 errors with invalid login and empty age")
		}
	}
}
