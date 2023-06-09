/*
RIK

RESTful public-facing API. The API is accessible through HTTP calls on specific URLs carrying JSON modeled data. 

API version: 0.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// Kind The kind of a workload.
type Kind string

// List of Kind
const (
	KIND_POD Kind = "Pod"
	KIND_FUNCTION Kind = "Function"
)

// All allowed values of Kind enum
var AllowedKindEnumValues = []Kind{
	"Pod",
	"Function",
}

func (v *Kind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Kind(value)
	for _, existing := range AllowedKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Kind", value)
}

// NewKindFromValue returns a pointer to a valid Kind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKindFromValue(v string) (*Kind, error) {
	ev := Kind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Kind: valid values are %v", v, AllowedKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Kind) IsValid() bool {
	for _, existing := range AllowedKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Kind value
func (v Kind) Ptr() *Kind {
	return &v
}

type NullableKind struct {
	value *Kind
	isSet bool
}

func (v NullableKind) Get() *Kind {
	return v.value
}

func (v *NullableKind) Set(val *Kind) {
	v.value = val
	v.isSet = true
}

func (v NullableKind) IsSet() bool {
	return v.isSet
}

func (v *NullableKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKind(val *Kind) *NullableKind {
	return &NullableKind{value: val, isSet: true}
}

func (v NullableKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

