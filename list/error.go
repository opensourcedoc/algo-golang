package list

import (
	"errors"
)

func indexOutOfRange() error {
	return errors.New("Index out of range")
}

func notFound() error {
	return errors.New("Not found")
}

func emptyList() error {
	return errors.New("Empty list")
}
