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

// checks if the DeleteInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteInstanceRequest{}

// DeleteInstanceRequest struct for DeleteInstanceRequest
type DeleteInstanceRequest struct {
	Id *string `json:"id,omitempty"`
}

// NewDeleteInstanceRequest instantiates a new DeleteInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteInstanceRequest() *DeleteInstanceRequest {
	this := DeleteInstanceRequest{}
	return &this
}

// NewDeleteInstanceRequestWithDefaults instantiates a new DeleteInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteInstanceRequestWithDefaults() *DeleteInstanceRequest {
	this := DeleteInstanceRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteInstanceRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteInstanceRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteInstanceRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeleteInstanceRequest) SetId(v string) {
	o.Id = &v
}

func (o DeleteInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableDeleteInstanceRequest struct {
	value *DeleteInstanceRequest
	isSet bool
}

func (v NullableDeleteInstanceRequest) Get() *DeleteInstanceRequest {
	return v.value
}

func (v *NullableDeleteInstanceRequest) Set(val *DeleteInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteInstanceRequest(val *DeleteInstanceRequest) *NullableDeleteInstanceRequest {
	return &NullableDeleteInstanceRequest{value: val, isSet: true}
}

func (v NullableDeleteInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


