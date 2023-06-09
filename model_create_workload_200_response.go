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

// CreateWorkload200Response - struct for CreateWorkload200Response
type CreateWorkload200Response struct {
	CreateWorkloadResponse *CreateWorkloadResponse
}

// CreateWorkloadResponseAsCreateWorkload200Response is a convenience function that returns CreateWorkloadResponse wrapped in CreateWorkload200Response
func CreateWorkloadResponseAsCreateWorkload200Response(v *CreateWorkloadResponse) CreateWorkload200Response {
	return CreateWorkload200Response{
		CreateWorkloadResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateWorkload200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateWorkloadResponse
	err = newStrictDecoder(data).Decode(&dst.CreateWorkloadResponse)
	if err == nil {
		jsonCreateWorkloadResponse, _ := json.Marshal(dst.CreateWorkloadResponse)
		if string(jsonCreateWorkloadResponse) == "{}" { // empty struct
			dst.CreateWorkloadResponse = nil
		} else {
			match++
		}
	} else {
		dst.CreateWorkloadResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateWorkloadResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateWorkload200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateWorkload200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateWorkload200Response) MarshalJSON() ([]byte, error) {
	if src.CreateWorkloadResponse != nil {
		return json.Marshal(&src.CreateWorkloadResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateWorkload200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CreateWorkloadResponse != nil {
		return obj.CreateWorkloadResponse
	}

	// all schemas are nil
	return nil
}

type NullableCreateWorkload200Response struct {
	value *CreateWorkload200Response
	isSet bool
}

func (v NullableCreateWorkload200Response) Get() *CreateWorkload200Response {
	return v.value
}

func (v *NullableCreateWorkload200Response) Set(val *CreateWorkload200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkload200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkload200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkload200Response(val *CreateWorkload200Response) *NullableCreateWorkload200Response {
	return &NullableCreateWorkload200Response{value: val, isSet: true}
}

func (v NullableCreateWorkload200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkload200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


