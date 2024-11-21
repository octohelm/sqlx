/*
Package model GENERATED BY gengo:enum 
DON'T EDIT THIS FILE
*/
package model

import (
	bytes "bytes"
	driver "database/sql/driver"
	errors "errors"
	fmt "fmt"

	enumeration "github.com/octohelm/enumeration/pkg/enumeration"
	pkgscanner "github.com/octohelm/enumeration/pkg/scanner"
)

var InvalidGender = errors.New("invalid Gender")

func (Gender) EnumValues() []any {
	return []any{
		GENDER__MALE, GENDER__FEMALE,
	}
}
func (v Gender) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v *Gender) UnmarshalText(data []byte) error {
	vv, err := ParseGenderFromString(string(bytes.ToUpper(data)))
	if err != nil {
		return err
	}
	*v = vv
	return nil
}

func ParseGenderFromString(s string) (Gender, error) {
	switch s {
	case "MALE":
		return GENDER__MALE, nil
	case "FEMALE":
		return GENDER__FEMALE, nil

	default:
		var i Gender
		_, err := fmt.Sscanf(s, "UNKNOWN_%d", &i)
		if err == nil {
			return i, nil
		}
		return GENDER_UNKNOWN, InvalidGender
	}
}

func (v Gender) IsZero() bool {
	return v == GENDER_UNKNOWN
}

func (v Gender) String() string {
	switch v {
	case GENDER__MALE:
		return "MALE"
	case GENDER__FEMALE:
		return "FEMALE"

	case GENDER_UNKNOWN:
		return "UNKNOWN"
	default:
		return fmt.Sprintf("UNKNOWN_%d", v)
	}
}

func ParseGenderLabelString(label string) (Gender, error) {
	switch label {
	case "男":
		return GENDER__MALE, nil
	case "女":
		return GENDER__FEMALE, nil

	default:
		return GENDER_UNKNOWN, InvalidGender
	}
}

func (v Gender) Label() string {
	switch v {
	case GENDER__MALE:
		return "男"
	case GENDER__FEMALE:
		return "女"

	default:
		return fmt.Sprint(v)
	}
}

func (v Gender) Value() (driver.Value, error) {
	offset := 0
	if o, ok := any(v).(enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *Gender) Scan(src any) error {
	offset := 0
	if o, ok := any(v).(enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}

	i, err := pkgscanner.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = Gender(i)
	return nil
}

var InvalidGenderExt = errors.New("invalid GenderExt")

func (GenderExt) EnumValues() []any {
	return []any{
		GENDER_EXT__FEMALE, GENDER_EXT__MALE,
	}
}
func ParseGenderExtLabelString(label string) (GenderExt, error) {
	switch label {
	case "女":
		return GENDER_EXT__FEMALE, nil
	case "男":
		return GENDER_EXT__MALE, nil

	default:
		return "", InvalidGenderExt
	}
}

func (v GenderExt) Label() string {
	switch v {
	case GENDER_EXT__FEMALE:
		return "女"
	case GENDER_EXT__MALE:
		return "男"

	default:
		return fmt.Sprint(v)
	}
}
