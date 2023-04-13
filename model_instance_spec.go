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

// checks if the InstanceSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceSpec{}

// InstanceSpec struct for InstanceSpec
type InstanceSpec struct {
	Containers []ContainerWorkload `json:"containers,omitempty"`
	Function *InstanceSpecFunction `json:"function,omitempty"`
}

// NewInstanceSpec instantiates a new InstanceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceSpec() *InstanceSpec {
	this := InstanceSpec{}
	return &this
}

// NewInstanceSpecWithDefaults instantiates a new InstanceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceSpecWithDefaults() *InstanceSpec {
	this := InstanceSpec{}
	return &this
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *InstanceSpec) GetContainers() []ContainerWorkload {
	if o == nil || IsNil(o.Containers) {
		var ret []ContainerWorkload
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceSpec) GetContainersOk() ([]ContainerWorkload, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *InstanceSpec) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []ContainerWorkload and assigns it to the Containers field.
func (o *InstanceSpec) SetContainers(v []ContainerWorkload) {
	o.Containers = v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *InstanceSpec) GetFunction() InstanceSpecFunction {
	if o == nil || IsNil(o.Function) {
		var ret InstanceSpecFunction
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceSpec) GetFunctionOk() (*InstanceSpecFunction, bool) {
	if o == nil || IsNil(o.Function) {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *InstanceSpec) HasFunction() bool {
	if o != nil && !IsNil(o.Function) {
		return true
	}

	return false
}

// SetFunction gets a reference to the given InstanceSpecFunction and assigns it to the Function field.
func (o *InstanceSpec) SetFunction(v InstanceSpecFunction) {
	o.Function = &v
}

func (o InstanceSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	if !IsNil(o.Function) {
		toSerialize["function"] = o.Function
	}
	return toSerialize, nil
}

type NullableInstanceSpec struct {
	value *InstanceSpec
	isSet bool
}

func (v NullableInstanceSpec) Get() *InstanceSpec {
	return v.value
}

func (v *NullableInstanceSpec) Set(val *InstanceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceSpec(val *InstanceSpec) *NullableInstanceSpec {
	return &NullableInstanceSpec{value: val, isSet: true}
}

func (v NullableInstanceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


