package main

import (
	"errors"
)

func justError() error {
	return errors.New("an error in main package")
}
