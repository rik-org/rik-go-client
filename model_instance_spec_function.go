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

// checks if the InstanceSpecFunction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceSpecFunction{}

// InstanceSpecFunction struct for InstanceSpecFunction
type InstanceSpecFunction struct {
	Execution *FunctionExecution `json:"execution,omitempty"`
	Exposure *InstanceSpecFunctionAllOfExposure `json:"exposure,omitempty"`
}

// NewInstanceSpecFunction instantiates a new InstanceSpecFunction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceSpecFunction() *InstanceSpecFunction {
	this := InstanceSpecFunction{}
	return &this
}

// NewInstanceSpecFunctionWithDefaults instantiates a new InstanceSpecFunction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceSpecFunctionWithDefaults() *InstanceSpecFunction {
	this := InstanceSpecFunction{}
	return &this
}

// GetExecution returns the Execution field value if set, zero value otherwise.
func (o *InstanceSpecFunction) GetExecution() FunctionExecution {
	if o == nil || IsNil(o.Execution) {
		var ret FunctionExecution
		return ret
	}
	return *o.Execution
}

// GetExecutionOk returns a tuple with the Execution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceSpecFunction) GetExecutionOk() (*FunctionExecution, bool) {
	if o == nil || IsNil(o.Execution) {
		return nil, false
	}
	return o.Execution, true
}

// HasExecution returns a boolean if a field has been set.
func (o *InstanceSpecFunction) HasExecution() bool {
	if o != nil && !IsNil(o.Execution) {
		return true
	}

	return false
}

// SetExecution gets a reference to the given FunctionExecution and assigns it to the Execution field.
func (o *InstanceSpecFunction) SetExecution(v FunctionExecution) {
	o.Execution = &v
}

// GetExposure returns the Exposure field value if set, zero value otherwise.
func (o *InstanceSpecFunction) GetExposure() InstanceSpecFunctionAllOfExposure {
	if o == nil || IsNil(o.Exposure) {
		var ret InstanceSpecFunctionAllOfExposure
		return ret
	}
	return *o.Exposure
}

// GetExposureOk returns a tuple with the Exposure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceSpecFunction) GetExposureOk() (*InstanceSpecFunctionAllOfExposure, bool) {
	if o == nil || IsNil(o.Exposure) {
		return nil, false
	}
	return o.Exposure, true
}

// HasExposure returns a boolean if a field has been set.
func (o *InstanceSpecFunction) HasExposure() bool {
	if o != nil && !IsNil(o.Exposure) {
		return true
	}

	return false
}

// SetExposure gets a reference to the given InstanceSpecFunctionAllOfExposure and assigns it to the Exposure field.
func (o *InstanceSpecFunction) SetExposure(v InstanceSpecFunctionAllOfExposure) {
	o.Exposure = &v
}

func (o InstanceSpecFunction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceSpecFunction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Execution) {
		toSerialize["execution"] = o.Execution
	}
	if !IsNil(o.Exposure) {
		toSerialize["exposure"] = o.Exposure
	}
	return toSerialize, nil
}

type NullableInstanceSpecFunction struct {
	value *InstanceSpecFunction
	isSet bool
}

func (v NullableInstanceSpecFunction) Get() *InstanceSpecFunction {
	return v.value
}

func (v *NullableInstanceSpecFunction) Set(val *InstanceSpecFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceSpecFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceSpecFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceSpecFunction(val *InstanceSpecFunction) *NullableInstanceSpecFunction {
	return &NullableInstanceSpecFunction{value: val, isSet: true}
}

func (v NullableInstanceSpecFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceSpecFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


