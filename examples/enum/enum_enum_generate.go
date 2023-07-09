// Code generated by github.com/yoas0bi/micro-toolkit
// DO NOT EDIT

package enum_test

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	// Status_name name map
	Status_name = map[Status]string{
		Running: "running",
	}

	// Status_value value map
	Status_value = map[string]Status{
		"running": Running,
	}
)

// ParseStatusFromString Parse Status from string
func ParseStatusFromString(str string) (Status, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := Status_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown Status: %s", str)
	}

	return Status(v), nil
}

// Equal type compare
func (t Status) Equal(target Status) bool {
	return t == target
}

// IsIn todo
func (t Status) IsIn(targets ...Status) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// String stringer
func (t Status) String() string {
	v, ok := Status_name[t]
	if !ok {
		return "unknown"
	}

	return v
}

// MarshalJSON todo
func (t Status) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *Status) UnmarshalJSON(b []byte) error {
	ins, err := ParseStatusFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

var (
	// Enum_name name map
	Enum_name = map[Enum]string{
		E1: "e1",
	}

	// Enum_value value map
	Enum_value = map[string]Enum{
		"e1": E1,
	}
)

// ParseEnumFromString Parse Enum from string
func ParseEnumFromString(str string) (Enum, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := Enum_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown Enum: %s", str)
	}

	return Enum(v), nil
}

// Equal type compare
func (t Enum) Equal(target Enum) bool {
	return t == target
}

// IsIn todo
func (t Enum) IsIn(targets ...Enum) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// String stringer
func (t Enum) String() string {
	v, ok := Enum_name[t]
	if !ok {
		return "unknown"
	}

	return v
}

// MarshalJSON todo
func (t Enum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *Enum) UnmarshalJSON(b []byte) error {
	ins, err := ParseEnumFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
