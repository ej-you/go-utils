package validator

import (
	"maps"
	"testing"

    "github.com/google/uuid"
)

// структура для проверки валидации
type ValidateData struct {
	ID 					uuid.UUID `json:"id" myvalid:"required" validate:"required"`
	Email 				string `json:"email" myvalid:"email" validate:"email"`
	Login 				string `json:"login" myvalid:"required|max:10" validate:"required,max=10"`
	Password 			string `json:"password" myvalid:"required|min:5" validate:"required,min=5"`
	Weight 				float64 `json:"weight" myvalid:"required" validate:"required"`
	Sale 				float64 `json:"sale" myvalid:"min:0|max:100" validate:"min=0,max=100"`
	Age 				int `json:"age" myvalid:"required|min:2" validate:"required,min=2"`
	HandFingersAmount 	int `json:"handFingersAmount" myvalid:"max:20" validate:"max=20"`
}


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

var validator = New()


func TestNoOneErrorCases(t *testing.T) {
	t.Log("Check full valid data")
	{
		err := validator.Validate(&validData)
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

		err := validator.Validate(&modifiedData)
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

		err := validator.Validate(&modifiedData)
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
		// modifiedData.ID = [16]byte{}
		modifiedData.ID = uuid.Nil

		err := validator.Validate(&modifiedData)
		if err == nil || err.Error() != "Key: 'ValidateData.ID' Error:Field validation for 'ID' failed on the 'required' tag" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with empty ID")
		}
	}

	t.Log("Check empty email")
	{
		modifiedData := validData
		modifiedData.Email = ""

		err := validator.Validate(&modifiedData)
		if err == nil || err.Error() != "Key: 'ValidateData.Email' Error:Field validation for 'Email' failed on the 'email' tag" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with incorrect email")
		}
	}

	t.Log("Check empty weight")
	{
		modifiedData := validData
		modifiedData.Weight = 0.0

		err := validator.Validate(&modifiedData)
		if err == nil || err.Error() != "Key: 'ValidateData.Weight' Error:Field validation for 'Weight' failed on the 'required' tag" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with empty weight")
		}
	}

	t.Log("Check empty age")
	{
		modifiedData := validData
		modifiedData.Age = 0

		err := validator.Validate(&modifiedData)
		if err == nil || err.Error() != "Key: 'ValidateData.Age' Error:Field validation for 'Age' failed on the 'required' tag" {
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

		err := validator.Validate(&modifiedData)
		if err == nil || err.Error() != "Key: 'ValidateData.Email' Error:Field validation for 'Email' failed on the 'email' tag" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with invalid email")
		}
	}

	t.Log("Check invalid login")
	{
		modifiedData := validData
		modifiedData.Login = "qwerty0123456789"

		err := validator.Validate(&modifiedData)
		if err == nil || err.Error() != "Key: 'ValidateData.Login' Error:Field validation for 'Login' failed on the 'max' tag" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with invalid login")
		}
	}

	t.Log("Check invalid password")
	{
		modifiedData := validData
		modifiedData.Password = "123"

		err := validator.Validate(&modifiedData)
		if err == nil || err.Error() != "Key: 'ValidateData.Password' Error:Field validation for 'Password' failed on the 'min' tag" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with invalid password")
		}
	}

	t.Log("Check invalid sale")
	{
		modifiedData := validData
		modifiedData.Sale = 120.0

		err := validator.Validate(&modifiedData)
		if err == nil || err.Error() != "Key: 'ValidateData.Sale' Error:Field validation for 'Sale' failed on the 'max' tag" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with invalid sale")
		}
	}

	t.Log("Check invalid age")
	{
		modifiedData := validData
		modifiedData.Age = -50

		err := validator.Validate(&modifiedData)
		if err == nil || err.Error() != "Key: 'ValidateData.Age' Error:Field validation for 'Age' failed on the 'min' tag" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with invalid age")
		}
	}

	t.Log("Check invalid handFingersAmount")
	{
		modifiedData := validData
		modifiedData.HandFingersAmount = 100

		err := validator.Validate(&modifiedData)
		if err == nil || err.Error() != "Key: 'ValidateData.HandFingersAmount' Error:Field validation for 'HandFingersAmount' failed on the 'max' tag" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Log("OK. Got 1 error with invalid handFingersAmount")
		}
	}
}


func TestGetTranslatedMap(t *testing.T) {
	t.Log("Check invalid login and empty age")
	{
		modifiedData := validData
		modifiedData.Login = "qwerty0123456789"
		modifiedData.Age = 0

		err := validator.Validate(&modifiedData)
		if err == nil { // NOT err
			t.Errorf("Unexpected error: %v", err)
			return
		}

		expectedMap := map[string]string{
			"fieldAge": "Age is a required field",
			"fieldLogin": "Login must be a maximum of 10 characters in length",
		}
		errMap, ok := validator.GetMapFromValidationError(err)
		if !ok {
			t.Errorf("Unexpected type of error interface (NOT validatorModule.ValidationErrors)")
		}

		if !maps.Equal(errMap, expectedMap) {
			t.Errorf("Unexpected map: %v", errMap)
		} else {
			t.Log("OK. Got expected errors map")
		}
	}
}


func TestGetTranslatedString(t *testing.T) {
	t.Log("Check invalid login and empty age")
	{
		modifiedData := validData
		modifiedData.Login = "qwerty0123456789"
		modifiedData.Age = 0

		err := validator.Validate(&modifiedData)
		if err == nil { // NOT err
			t.Errorf("Unexpected error: %v", err)
			return
		}

		expectedString := "fieldLogin: Login must be a maximum of 10 characters in length" + " | " + "fieldAge: Age is a required field"
		errString, ok := validator.GetStringFromValidationError(err)
		if !ok {
			t.Errorf("Unexpected type of error interface (NOT validatorModule.ValidationErrors)")
		}

		if errString != expectedString {
			t.Errorf("Unexpected string: %v", errString)
		} else {
			t.Log("OK. Got expected errors string")
		}
	}
}