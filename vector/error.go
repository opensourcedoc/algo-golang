package vector

import "errors"

func negativeSize() {
	panic("Size should be larger than zero")
}

func indexOutOfRange() {
	panic("Index Out Of Range")
}

func unequalLength() {
	panic("Unequal Leagth")
}

func unequalType() error {
	return errors.New("Unequal Type")
}

func unknownType() error {
	return errors.New("Unknown Type. Consider AddBy method.")
}
