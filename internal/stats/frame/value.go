package frame

import (
	"fmt"
)

type ValueInt int

var _ Value = ValueInt(0)

func (value ValueInt) String() string {
	return fmt.Sprintf("%d", value)
}

func (value ValueInt) Float() float32 {
	return float32(value)
}
func (value ValueInt) MarshalJSON() ([]byte, error) {
	return []byte(`"` + value.String() + `"`), nil
}

type ValueFloatDecimal float32

var _ Value = ValueFloatDecimal(0)

func (value ValueFloatDecimal) String() string {
	return fmt.Sprintf("%.2f", value)
}

func (value ValueFloatDecimal) Float() float32 {
	return float32(value)
}

func (value ValueFloatDecimal) MarshalJSON() ([]byte, error) {
	return []byte(`"` + value.String() + `"`), nil
}

type ValueFloatPercent float32

var _ Value = ValueFloatPercent(0)

func (value ValueFloatPercent) String() string {
	return fmt.Sprintf("%.2f%%", value)
}

func (value ValueFloatPercent) Float() float32 {
	return float32(value)
}

func (value ValueFloatPercent) MarshalJSON() ([]byte, error) {
	return []byte(`"` + value.String() + `"`), nil
}

type valueInvalid struct{}

var _ Value = InvalidValue

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
	return []byte(`"-1"`), nil
}

var InvalidValue = valueInvalid{}

type ValueSpecialRating float32

var _ Value = ValueSpecialRating(0)

func (value ValueSpecialRating) int() int {
	if value > 0 {
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

func (value ValueSpecialRating) MarshalJSON() ([]byte, error) {
	return []byte(`"` + value.String() + `"`), nil
}

type Value interface {
	String() string
	Float() float32
}
