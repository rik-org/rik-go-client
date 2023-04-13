/*
RIK API

RESTful public-facing API. The API is accessible through HTTP calls on specific URLs carrying JSON modeled data. 

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ContainerWorkloadContainersInnerEnvInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerWorkloadContainersInnerEnvInner{}

// ContainerWorkloadContainersInnerEnvInner struct for ContainerWorkloadContainersInnerEnvInner
type ContainerWorkloadContainersInnerEnvInner struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewContainerWorkloadContainersInnerEnvInner instantiates a new ContainerWorkloadContainersInnerEnvInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerWorkloadContainersInnerEnvInner() *ContainerWorkloadContainersInnerEnvInner {
	this := ContainerWorkloadContainersInnerEnvInner{}
	return &this
}

// NewContainerWorkloadContainersInnerEnvInnerWithDefaults instantiates a new ContainerWorkloadContainersInnerEnvInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerWorkloadContainersInnerEnvInnerWithDefaults() *ContainerWorkloadContainersInnerEnvInner {
	this := ContainerWorkloadContainersInnerEnvInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerWorkloadContainersInnerEnvInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerWorkloadContainersInnerEnvInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerWorkloadContainersInnerEnvInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerWorkloadContainersInnerEnvInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ContainerWorkloadContainersInnerEnvInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerWorkloadContainersInnerEnvInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ContainerWorkloadContainersInnerEnvInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ContainerWorkloadContainersInnerEnvInner) SetValue(v string) {
	o.Value = &v
}

func (o ContainerWorkloadContainersInnerEnvInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerWorkloadContainersInnerEnvInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableContainerWorkloadContainersInnerEnvInner struct {
	value *ContainerWorkloadContainersInnerEnvInner
	isSet bool
}

func (v NullableContainerWorkloadContainersInnerEnvInner) Get() *ContainerWorkloadContainersInnerEnvInner {
	return v.value
}

func (v *NullableContainerWorkloadContainersInnerEnvInner) Set(val *ContainerWorkloadContainersInnerEnvInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerWorkloadContainersInnerEnvInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerWorkloadContainersInnerEnvInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerWorkloadContainersInnerEnvInner(val *ContainerWorkloadContainersInnerEnvInner) *NullableContainerWorkloadContainersInnerEnvInner {
	return &NullableContainerWorkloadContainersInnerEnvInner{value: val, isSet: true}
}

func (v NullableContainerWorkloadContainersInnerEnvInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerWorkloadContainersInnerEnvInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


