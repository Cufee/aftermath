package encoding

import (
	"bytes"
	"encoding/gob"
)

func EncodeGob(data any) ([]byte, error) {
	if data == nil {
		return []byte{}, nil
	}
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(data)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func DecodeGob(data []byte, target any) error {
	err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(target)
	if err != nil {
		return err
	}
	return nil
}
