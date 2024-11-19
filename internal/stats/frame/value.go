package frame

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
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
func (value *ValueInt) UnmarshalJSON(data []byte) error {
	if string(data) == valueInvalidEncoded {
		*value = ValueInt(InvalidValue.Float())
		return nil
	}
	v, err := strconv.Atoi(strings.ReplaceAll(string(data), `"`, ""))
	if err != nil {
		return errors.Wrap(err, "failed to parse int")
	}
	*value = ValueInt(v)
	return nil
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
func (value *ValueFloatDecimal) UnmarshalJSON(data []byte) error {
	if string(data) == valueInvalidEncoded {
		*value = ValueFloatDecimal(InvalidValue.Float())
		return nil
	}
	v, err := strconv.ParseFloat(strings.ReplaceAll(string(data), `"`, ""), 32)
	if err != nil {
		return errors.Wrap(err, "failed to parse float decimal")
	}
	*value = ValueFloatDecimal(v)
	return nil
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
func (value *ValueFloatPercent) UnmarshalJSON(data []byte) error {
	if string(data) == valueInvalidEncoded {
		*value = ValueFloatPercent(InvalidValue.Float())
		return nil
	}
	v, err := strconv.ParseFloat(strings.TrimSuffix(strings.ReplaceAll(string(data), `"`, ""), "%"), 32)
	if err != nil {
		return errors.Wrap(err, "failed to parse float percent")
	}
	*value = ValueFloatPercent(v)
	return nil
}

type valueInvalid struct{}

const valueInvalidEncoded = `"-"`

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
	return []byte(valueInvalidEncoded), nil
}
func (value *valueInvalid) UnmarshalJSON(data []byte) error {
	if string(data) == valueInvalidEncoded {
		value = &InvalidValue
		return nil
	}
	return errors.New("bad input for invalid value")
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
func (value *ValueSpecialRating) UnmarshalJSON(data []byte) error {
	if string(data) == valueInvalidEncoded {
		*value = ValueSpecialRating(InvalidValue.Float())
		return nil
	}
	v, err := strconv.ParseFloat(strings.ReplaceAll(string(data), `"`, ""), 32)
	if err != nil {
		return errors.Wrap(err, "failed to parse special rating")
	}
	*value = ValueSpecialRating(v)
	return nil
}

type Value interface {
	String() string
	Float() float32
}
