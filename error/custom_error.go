package error

import "fmt"

type customError struct {
	msg string
}

func (e customError) Error() string {
	return e.msg
}

func doSth() *customError {
	return nil
}

func handleError(err error) {
	if err != nil {
		fmt.Println("handleError: err != nil")
	} else {
		fmt.Println("handleError: err == nil")
	}
}
