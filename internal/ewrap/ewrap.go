package ewrap

import "fmt"

func Error(err error) {
	// return fmt.Errorf("error: %w", err)
	fmt.Errorf("error: %w", err)
}
