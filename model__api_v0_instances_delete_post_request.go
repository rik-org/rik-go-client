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

// checks if the ApiV0InstancesDeletePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV0InstancesDeletePostRequest{}

// ApiV0InstancesDeletePostRequest struct for ApiV0InstancesDeletePostRequest
type ApiV0InstancesDeletePostRequest struct {
	Id *string `json:"id,omitempty"`
}

// NewApiV0InstancesDeletePostRequest instantiates a new ApiV0InstancesDeletePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV0InstancesDeletePostRequest() *ApiV0InstancesDeletePostRequest {
	this := ApiV0InstancesDeletePostRequest{}
	return &this
}

// NewApiV0InstancesDeletePostRequestWithDefaults instantiates a new ApiV0InstancesDeletePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV0InstancesDeletePostRequestWithDefaults() *ApiV0InstancesDeletePostRequest {
	this := ApiV0InstancesDeletePostRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiV0InstancesDeletePostRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV0InstancesDeletePostRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiV0InstancesDeletePostRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiV0InstancesDeletePostRequest) SetId(v string) {
	o.Id = &v
}

func (o ApiV0InstancesDeletePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV0InstancesDeletePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableApiV0InstancesDeletePostRequest struct {
	value *ApiV0InstancesDeletePostRequest
	isSet bool
}

func (v NullableApiV0InstancesDeletePostRequest) Get() *ApiV0InstancesDeletePostRequest {
	return v.value
}

func (v *NullableApiV0InstancesDeletePostRequest) Set(val *ApiV0InstancesDeletePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV0InstancesDeletePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV0InstancesDeletePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV0InstancesDeletePostRequest(val *ApiV0InstancesDeletePostRequest) *NullableApiV0InstancesDeletePostRequest {
	return &NullableApiV0InstancesDeletePostRequest{value: val, isSet: true}
}

func (v NullableApiV0InstancesDeletePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV0InstancesDeletePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


