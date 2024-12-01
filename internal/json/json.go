package json

import (
	"io"

	"github.com/goccy/go-json"
)

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func NewDecoder(reader io.Reader) *json.Decoder {
	return json.NewDecoder(reader)
}

func NewEncoder(writer io.Writer) *json.Encoder {
	return json.NewEncoder(writer)
}
