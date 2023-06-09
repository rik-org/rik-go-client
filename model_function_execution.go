/*
RIK

RESTful public-facing API. The API is accessible through HTTP calls on specific URLs carrying JSON modeled data. 

API version: 0.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the FunctionExecution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionExecution{}

// FunctionExecution struct for FunctionExecution
type FunctionExecution struct {
	Rootfs *string `json:"rootfs,omitempty"`
}

// NewFunctionExecution instantiates a new FunctionExecution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionExecution() *FunctionExecution {
	this := FunctionExecution{}
	return &this
}

// NewFunctionExecutionWithDefaults instantiates a new FunctionExecution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionExecutionWithDefaults() *FunctionExecution {
	this := FunctionExecution{}
	return &this
}

// GetRootfs returns the Rootfs field value if set, zero value otherwise.
func (o *FunctionExecution) GetRootfs() string {
	if o == nil || IsNil(o.Rootfs) {
		var ret string
		return ret
	}
	return *o.Rootfs
}

// GetRootfsOk returns a tuple with the Rootfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionExecution) GetRootfsOk() (*string, bool) {
	if o == nil || IsNil(o.Rootfs) {
		return nil, false
	}
	return o.Rootfs, true
}

// HasRootfs returns a boolean if a field has been set.
func (o *FunctionExecution) HasRootfs() bool {
	if o != nil && !IsNil(o.Rootfs) {
		return true
	}

	return false
}

// SetRootfs gets a reference to the given string and assigns it to the Rootfs field.
func (o *FunctionExecution) SetRootfs(v string) {
	o.Rootfs = &v
}

func (o FunctionExecution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionExecution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rootfs) {
		toSerialize["rootfs"] = o.Rootfs
	}
	return toSerialize, nil
}

type NullableFunctionExecution struct {
	value *FunctionExecution
	isSet bool
}

func (v NullableFunctionExecution) Get() *FunctionExecution {
	return v.value
}

func (v *NullableFunctionExecution) Set(val *FunctionExecution) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionExecution(val *FunctionExecution) *NullableFunctionExecution {
	return &NullableFunctionExecution{value: val, isSet: true}
}

func (v NullableFunctionExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


