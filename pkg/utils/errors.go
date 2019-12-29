package utils

import (
	"fmt"
)

// PanicWithError Create panic if there is any error
func PanicWithError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
