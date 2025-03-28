package datatypes

import (
	"github.com/go-json-experiment/json"
)

type Bool int

const (
	BOOL_UNKNOWN Bool = iota
	BOOL_TRUE         // true
	BOOL_FALSE        // false
)

var _ interface {
	json.Unmarshaler
	json.Marshaler
} = (*Bool)(nil)

func (Bool) OpenAPISchemaType() []string {
	return []string{"boolean"}
}

func (v Bool) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}

func (v *Bool) UnmarshalJSON(data []byte) (err error) {
	return v.UnmarshalText(data)
}

func (v Bool) MarshalText() ([]byte, error) {
	switch v {
	case BOOL_FALSE:
		return []byte("false"), nil
	case BOOL_TRUE:
		return []byte("true"), nil
	default:
		return []byte("null"), nil
	}
}

func (v *Bool) UnmarshalText(data []byte) (err error) {
	switch string(data) {
	case "false":
		*v = BOOL_FALSE
	case "true":
		*v = BOOL_TRUE
	}
	return
}
