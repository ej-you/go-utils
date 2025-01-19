package env

import (
	"fmt"
	"os"
)


// panic if at least one given env variable is not presented
func MustBePresented(envVars ...string) {
	var isPresented bool

	for _, envVar := range envVars {
		_, isPresented = os.LookupEnv(envVar)
		if !isPresented {
			panic(fmt.Errorf("env variable %s is not presented", envVar))
		}
	}
}
