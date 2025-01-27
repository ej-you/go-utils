# go-utils

## Module with a set of different utils for Golang

### Submodules:

1. `env` funcs:

* MustBePresented(envVars ...string)

> Description
> <br>
> Input: env variables names
> <br>
> Output: -
> <br>
> Panic if at least one given env variable is not presented
> <br>
> Example: MustBePresented("VAR1", "VAR2")

2. `validator`

* Validate(dataToValidate validate.Validator) error

```go
// ! HINT !

type validate.Validator interface {
	IsValid(errors *validate.Errors)
}
```

> Description
> <br>
> Input: pointer to struct to validate
> <br>
> Output: \*echo.HTTPError
> <br>
> Validate given struct annd return \*echo.HTTPError if there are validate errors
> <br>
> Example: Validate(&dataStruct)
