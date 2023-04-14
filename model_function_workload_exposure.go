/*
RIK API

RESTful public-facing API. The API is accessible through HTTP calls on specific URLs carrying JSON modeled data. 

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the FunctionWorkloadExposure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionWorkloadExposure{}

// FunctionWorkloadExposure struct for FunctionWorkloadExposure
type FunctionWorkloadExposure struct {
	Port *int32 `json:"port,omitempty"`
	TargetPort *int32 `json:"targetPort,omitempty"`
	// The type of the exposure
	Type *string `json:"type,omitempty"`
}

// NewFunctionWorkloadExposure instantiates a new FunctionWorkloadExposure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionWorkloadExposure() *FunctionWorkloadExposure {
	this := FunctionWorkloadExposure{}
	return &this
}

// NewFunctionWorkloadExposureWithDefaults instantiates a new FunctionWorkloadExposure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionWorkloadExposureWithDefaults() *FunctionWorkloadExposure {
	this := FunctionWorkloadExposure{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *FunctionWorkloadExposure) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionWorkloadExposure) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *FunctionWorkloadExposure) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *FunctionWorkloadExposure) SetPort(v int32) {
	o.Port = &v
}

// GetTargetPort returns the TargetPort field value if set, zero value otherwise.
func (o *FunctionWorkloadExposure) GetTargetPort() int32 {
	if o == nil || IsNil(o.TargetPort) {
		var ret int32
		return ret
	}
	return *o.TargetPort
}

// GetTargetPortOk returns a tuple with the TargetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionWorkloadExposure) GetTargetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetPort) {
		return nil, false
	}
	return o.TargetPort, true
}

// HasTargetPort returns a boolean if a field has been set.
func (o *FunctionWorkloadExposure) HasTargetPort() bool {
	if o != nil && !IsNil(o.TargetPort) {
		return true
	}

	return false
}

// SetTargetPort gets a reference to the given int32 and assigns it to the TargetPort field.
func (o *FunctionWorkloadExposure) SetTargetPort(v int32) {
	o.TargetPort = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FunctionWorkloadExposure) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionWorkloadExposure) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FunctionWorkloadExposure) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FunctionWorkloadExposure) SetType(v string) {
	o.Type = &v
}

func (o FunctionWorkloadExposure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionWorkloadExposure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.TargetPort) {
		toSerialize["targetPort"] = o.TargetPort
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableFunctionWorkloadExposure struct {
	value *FunctionWorkloadExposure
	isSet bool
}

func (v NullableFunctionWorkloadExposure) Get() *FunctionWorkloadExposure {
	return v.value
}

func (v *NullableFunctionWorkloadExposure) Set(val *FunctionWorkloadExposure) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionWorkloadExposure) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionWorkloadExposure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionWorkloadExposure(val *FunctionWorkloadExposure) *NullableFunctionWorkloadExposure {
	return &NullableFunctionWorkloadExposure{value: val, isSet: true}
}

func (v NullableFunctionWorkloadExposure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionWorkloadExposure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


