package pgt

import (
	"fmt"

	bat "github.com/robert-zaremba/go-bat"
)

// convertToString converts interface{} to string
func convertToString(source interface{}) (string, error) {
	switch t := source.(type) {
	default:
		return "", fmt.Errorf("Unable to parse %T as string. ", source)
	case []byte:
		return bat.UnsafeByteArrayToStr(t), nil
	case string:
		return t, nil
	}
}

// convertToBytes converts interface{} to []byte
func convertToBytes(source interface{}) ([]byte, error) {
	switch t := source.(type) {
	default:
		return nil, fmt.Errorf("Unable to parse %T as string. ", source)
	case []byte:
		return t, nil
	case string:
		return bat.UnsafeStrToByteArray(t), nil
	}
}
