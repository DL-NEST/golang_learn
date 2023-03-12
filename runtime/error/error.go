package main

import (
	"errors"
	"fmt"
)

const (
	CommonError = "commonError"
)

var CommonErrorVar = errors.New(CommonError)

func main() {
	err := errors.New(CommonError)

	if err.Error() == CommonError {
		fmt.Println("commonError")
	}

	if errors.As(err, &CommonErrorVar) {
		fmt.Println("commonError")
	}
}
