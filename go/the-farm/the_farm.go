package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(calculator FodderCalculator, numCows int) (float64, error) {
	totalFodder, err := calculator.FodderAmount(numCows)
	if err != nil {
		return 0, err
	}

	factor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (totalFodder * factor) / float64(numCows), nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, numCows int) (float64, error) {
	if numCows > 0 {
		return DivideFood(calculator, numCows)
	}

	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	numCows int
	message string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.numCows, err.message)
}

func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{
			numCows: numCows,
			message: "there are no negative cows",
		}
	}

	if numCows == 0 {
		return &InvalidCowsError{
			numCows: numCows,
			message: "no cows don't need food",
		}
	}

	return nil
}
