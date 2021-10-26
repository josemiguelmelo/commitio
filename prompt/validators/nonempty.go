package validators

import (
	"errors"
)

func NonEmpty(input string) error {
	if len(input) <= 0 {
		return errors.New("Input should not be empty")
	}
	return nil
}
