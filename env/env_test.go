package env

import (
	"os"

	"testing"
)


func TestMustBePresented(t *testing.T) {
	t.Log("Check with presented env variable TEST")
	{
		err := os.Setenv("TEST", "test-env-var")
		if err != nil {
			t.Errorf("Failed to set TEST env variable: %v", err)
		}

		MustBePresented("TEST")
		t.Log("Good. There is no error")
	}

	t.Log("Check with no presented env variable TEST")
	{
		err := os.Unsetenv("TEST")
		if err != nil {
			t.Errorf("Failed to unset TEST env variable: %v", err)
		}

		defer func() {
	        if r := recover(); r != nil {
	            t.Log("Recovered. Panic was occurred:", r)
	        }
	    }()

		MustBePresented("TEST")
	}
}
