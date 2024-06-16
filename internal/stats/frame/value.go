package frame

import "fmt"

type ValueInt int

func (value ValueInt) String() string {
	return fmt.Sprintf("%d", value)
}

func (value ValueInt) Float() float32 {
	return float32(value)
}

type ValueFloatDecimal float32

func (value ValueFloatDecimal) String() string {
	return fmt.Sprintf("%.2f", value)
}

func (value ValueFloatDecimal) Float() float32 {
	return float32(value)
}

type ValueFloatPercent float32

func (value ValueFloatPercent) String() string {
	return fmt.Sprintf("%.2f%%", value)
}

func (value ValueFloatPercent) Float() float32 {
	return float32(value)
}

type valueInvalid struct{}

func (value valueInvalid) String() string {
	return "-"
}

func (value valueInvalid) Float() float32 {
	return -1
}

func (value valueInvalid) Equals(compareTo Value) bool {
	return compareTo.Float() == value.Float()
}

func (value valueInvalid) MarshalJSON() ([]byte, error) {
	return []byte("-1"), nil
}

var InvalidValue = valueInvalid{}

type ValueSpecialRating float32

func (value ValueSpecialRating) int() int {
	if value >= 0 {
		return int((value * 10) + 3000)
	}
	return int(InvalidValue.Float())
}

func (value ValueSpecialRating) String() string {
	if value > 1 {
		return fmt.Sprintf("%d", value.int())
	}
	return InvalidValue.String()
}

func (value ValueSpecialRating) Float() float32 {
	return float32(value.int())
}

type Value interface {
	String() string
	Float() float32
}
