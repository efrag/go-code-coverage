package calculator

type Calculator struct {
	name string
}

func Sum(a, b uint8) uint16 {
	return uint16(a) + uint16(b)
}

func Multiply(a, b int8) int16 {
	res := int16(a) * int16(b)
	return res
}

func Abs(a int8) int8 {
	if a < 0 {
		return -a
	}

	return a
}
