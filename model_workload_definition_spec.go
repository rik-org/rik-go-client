/*
RIK API

RESTful public-facing API. The API is accessible through HTTP calls on specific URLs carrying JSON modeled data. 

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// WorkloadDefinitionSpec - struct for WorkloadDefinitionSpec
type WorkloadDefinitionSpec struct {
	ContainerWorkload *ContainerWorkload
	FunctionWorkload *FunctionWorkload
}

// ContainerWorkloadAsWorkloadDefinitionSpec is a convenience function that returns ContainerWorkload wrapped in WorkloadDefinitionSpec
func ContainerWorkloadAsWorkloadDefinitionSpec(v *ContainerWorkload) WorkloadDefinitionSpec {
	return WorkloadDefinitionSpec{
		ContainerWorkload: v,
	}
}

// FunctionWorkloadAsWorkloadDefinitionSpec is a convenience function that returns FunctionWorkload wrapped in WorkloadDefinitionSpec
func FunctionWorkloadAsWorkloadDefinitionSpec(v *FunctionWorkload) WorkloadDefinitionSpec {
	return WorkloadDefinitionSpec{
		FunctionWorkload: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *WorkloadDefinitionSpec) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContainerWorkload
	err = newStrictDecoder(data).Decode(&dst.ContainerWorkload)
	if err == nil {
		jsonContainerWorkload, _ := json.Marshal(dst.ContainerWorkload)
		if string(jsonContainerWorkload) == "{}" { // empty struct
			dst.ContainerWorkload = nil
		} else {
			match++
		}
	} else {
		dst.ContainerWorkload = nil
	}

	// try to unmarshal data into FunctionWorkload
	err = newStrictDecoder(data).Decode(&dst.FunctionWorkload)
	if err == nil {
		jsonFunctionWorkload, _ := json.Marshal(dst.FunctionWorkload)
		if string(jsonFunctionWorkload) == "{}" { // empty struct
			dst.FunctionWorkload = nil
		} else {
			match++
		}
	} else {
		dst.FunctionWorkload = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ContainerWorkload = nil
		dst.FunctionWorkload = nil

		return fmt.Errorf("data matches more than one schema in oneOf(WorkloadDefinitionSpec)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(WorkloadDefinitionSpec)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WorkloadDefinitionSpec) MarshalJSON() ([]byte, error) {
	if src.ContainerWorkload != nil {
		return json.Marshal(&src.ContainerWorkload)
	}

	if src.FunctionWorkload != nil {
		return json.Marshal(&src.FunctionWorkload)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WorkloadDefinitionSpec) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ContainerWorkload != nil {
		return obj.ContainerWorkload
	}

	if obj.FunctionWorkload != nil {
		return obj.FunctionWorkload
	}

	// all schemas are nil
	return nil
}

type NullableWorkloadDefinitionSpec struct {
	value *WorkloadDefinitionSpec
	isSet bool
}

func (v NullableWorkloadDefinitionSpec) Get() *WorkloadDefinitionSpec {
	return v.value
}

func (v *NullableWorkloadDefinitionSpec) Set(val *WorkloadDefinitionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadDefinitionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadDefinitionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadDefinitionSpec(val *WorkloadDefinitionSpec) *NullableWorkloadDefinitionSpec {
	return &NullableWorkloadDefinitionSpec{value: val, isSet: true}
}

func (v NullableWorkloadDefinitionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadDefinitionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


