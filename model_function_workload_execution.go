/*
RIK API

RESTful public-facing API. The API is accessible through HTTP calls on specific URLs carrying JSON modeled data. 

API version: 0.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the FunctionWorkloadExecution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionWorkloadExecution{}

// FunctionWorkloadExecution struct for FunctionWorkloadExecution
type FunctionWorkloadExecution struct {
	Rootfs *string `json:"rootfs,omitempty"`
}

// NewFunctionWorkloadExecution instantiates a new FunctionWorkloadExecution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionWorkloadExecution() *FunctionWorkloadExecution {
	this := FunctionWorkloadExecution{}
	return &this
}

// NewFunctionWorkloadExecutionWithDefaults instantiates a new FunctionWorkloadExecution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionWorkloadExecutionWithDefaults() *FunctionWorkloadExecution {
	this := FunctionWorkloadExecution{}
	return &this
}

// GetRootfs returns the Rootfs field value if set, zero value otherwise.
func (o *FunctionWorkloadExecution) GetRootfs() string {
	if o == nil || IsNil(o.Rootfs) {
		var ret string
		return ret
	}
	return *o.Rootfs
}

// GetRootfsOk returns a tuple with the Rootfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionWorkloadExecution) GetRootfsOk() (*string, bool) {
	if o == nil || IsNil(o.Rootfs) {
		return nil, false
	}
	return o.Rootfs, true
}

// HasRootfs returns a boolean if a field has been set.
func (o *FunctionWorkloadExecution) HasRootfs() bool {
	if o != nil && !IsNil(o.Rootfs) {
		return true
	}

	return false
}

// SetRootfs gets a reference to the given string and assigns it to the Rootfs field.
func (o *FunctionWorkloadExecution) SetRootfs(v string) {
	o.Rootfs = &v
}

func (o FunctionWorkloadExecution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionWorkloadExecution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rootfs) {
		toSerialize["rootfs"] = o.Rootfs
	}
	return toSerialize, nil
}

type NullableFunctionWorkloadExecution struct {
	value *FunctionWorkloadExecution
	isSet bool
}

func (v NullableFunctionWorkloadExecution) Get() *FunctionWorkloadExecution {
	return v.value
}

func (v *NullableFunctionWorkloadExecution) Set(val *FunctionWorkloadExecution) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionWorkloadExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionWorkloadExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionWorkloadExecution(val *FunctionWorkloadExecution) *NullableFunctionWorkloadExecution {
	return &NullableFunctionWorkloadExecution{value: val, isSet: true}
}

func (v NullableFunctionWorkloadExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionWorkloadExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


