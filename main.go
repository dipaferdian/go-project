package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return e.Field + ": " + e.Message
}


func validateUser() error {
    return &ValidationError{
        Field:   "email",
        Message: "invalid email format",
    }
}

func main() {
    err := validateUser()

    var validationErr *ValidationError


    if errors.As(err, &validationErr) {
        fmt.Println(validationErr.Field)
        fmt.Println(validationErr.Message)
    }
}