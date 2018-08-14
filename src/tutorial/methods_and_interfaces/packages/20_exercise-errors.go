package methods_and_interfaces

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negativ number: %g", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}
	return 0, nil
}

func ErrorExercise20() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
