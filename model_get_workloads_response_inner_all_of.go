/*
RIK

RESTful public-facing API. The API is accessible through HTTP calls on specific URLs carrying JSON modeled data. 

API version: 0.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetWorkloadsResponseInnerAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWorkloadsResponseInnerAllOf{}

// GetWorkloadsResponseInnerAllOf struct for GetWorkloadsResponseInnerAllOf
type GetWorkloadsResponseInnerAllOf struct {
	Value *Workload `json:"value,omitempty"`
}

// NewGetWorkloadsResponseInnerAllOf instantiates a new GetWorkloadsResponseInnerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWorkloadsResponseInnerAllOf() *GetWorkloadsResponseInnerAllOf {
	this := GetWorkloadsResponseInnerAllOf{}
	return &this
}

// NewGetWorkloadsResponseInnerAllOfWithDefaults instantiates a new GetWorkloadsResponseInnerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWorkloadsResponseInnerAllOfWithDefaults() *GetWorkloadsResponseInnerAllOf {
	this := GetWorkloadsResponseInnerAllOf{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetWorkloadsResponseInnerAllOf) GetValue() Workload {
	if o == nil || IsNil(o.Value) {
		var ret Workload
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWorkloadsResponseInnerAllOf) GetValueOk() (*Workload, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetWorkloadsResponseInnerAllOf) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given Workload and assigns it to the Value field.
func (o *GetWorkloadsResponseInnerAllOf) SetValue(v Workload) {
	o.Value = &v
}

func (o GetWorkloadsResponseInnerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWorkloadsResponseInnerAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableGetWorkloadsResponseInnerAllOf struct {
	value *GetWorkloadsResponseInnerAllOf
	isSet bool
}

func (v NullableGetWorkloadsResponseInnerAllOf) Get() *GetWorkloadsResponseInnerAllOf {
	return v.value
}

func (v *NullableGetWorkloadsResponseInnerAllOf) Set(val *GetWorkloadsResponseInnerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWorkloadsResponseInnerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWorkloadsResponseInnerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWorkloadsResponseInnerAllOf(val *GetWorkloadsResponseInnerAllOf) *NullableGetWorkloadsResponseInnerAllOf {
	return &NullableGetWorkloadsResponseInnerAllOf{value: val, isSet: true}
}

func (v NullableGetWorkloadsResponseInnerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWorkloadsResponseInnerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

