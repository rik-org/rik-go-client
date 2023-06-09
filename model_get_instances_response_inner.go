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

// checks if the GetInstancesResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInstancesResponseInner{}

// GetInstancesResponseInner struct for GetInstancesResponseInner
type GetInstancesResponseInner struct {
	// Simple helper schema to define an UUID
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Value *Instance `json:"value,omitempty"`
}

// NewGetInstancesResponseInner instantiates a new GetInstancesResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInstancesResponseInner() *GetInstancesResponseInner {
	this := GetInstancesResponseInner{}
	return &this
}

// NewGetInstancesResponseInnerWithDefaults instantiates a new GetInstancesResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInstancesResponseInnerWithDefaults() *GetInstancesResponseInner {
	this := GetInstancesResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetInstancesResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstancesResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetInstancesResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetInstancesResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetInstancesResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstancesResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetInstancesResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetInstancesResponseInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetInstancesResponseInner) GetValue() Instance {
	if o == nil || IsNil(o.Value) {
		var ret Instance
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstancesResponseInner) GetValueOk() (*Instance, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetInstancesResponseInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given Instance and assigns it to the Value field.
func (o *GetInstancesResponseInner) SetValue(v Instance) {
	o.Value = &v
}

func (o GetInstancesResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInstancesResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableGetInstancesResponseInner struct {
	value *GetInstancesResponseInner
	isSet bool
}

func (v NullableGetInstancesResponseInner) Get() *GetInstancesResponseInner {
	return v.value
}

func (v *NullableGetInstancesResponseInner) Set(val *GetInstancesResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInstancesResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInstancesResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInstancesResponseInner(val *GetInstancesResponseInner) *NullableGetInstancesResponseInner {
	return &NullableGetInstancesResponseInner{value: val, isSet: true}
}

func (v NullableGetInstancesResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInstancesResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


