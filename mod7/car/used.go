package car

import "github.com/efrag/go-code-coverage/mod7/age"

type Car struct {
	manufacturedYear int
}

func AverageAge(cars []Car) (float32, error) {
	var years float32

	for i := 0; i < len(cars); i++ {
		y, e := age.YearsSince(cars[i].manufacturedYear)
		if e != nil {
			return years, e
		}
		years += float32(y)
	}

	return years/float32(len(cars)), nil
}
