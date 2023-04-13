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

// checks if the ContainerWorkloadContainersInnerNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerWorkloadContainersInnerNetwork{}

// ContainerWorkloadContainersInnerNetwork struct for ContainerWorkloadContainersInnerNetwork
type ContainerWorkloadContainersInnerNetwork struct {
	Port *float32 `json:"port,omitempty"`
	TargetPort *float32 `json:"targetPort,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewContainerWorkloadContainersInnerNetwork instantiates a new ContainerWorkloadContainersInnerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerWorkloadContainersInnerNetwork() *ContainerWorkloadContainersInnerNetwork {
	this := ContainerWorkloadContainersInnerNetwork{}
	return &this
}

// NewContainerWorkloadContainersInnerNetworkWithDefaults instantiates a new ContainerWorkloadContainersInnerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerWorkloadContainersInnerNetworkWithDefaults() *ContainerWorkloadContainersInnerNetwork {
	this := ContainerWorkloadContainersInnerNetwork{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ContainerWorkloadContainersInnerNetwork) GetPort() float32 {
	if o == nil || IsNil(o.Port) {
		var ret float32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerWorkloadContainersInnerNetwork) GetPortOk() (*float32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ContainerWorkloadContainersInnerNetwork) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given float32 and assigns it to the Port field.
func (o *ContainerWorkloadContainersInnerNetwork) SetPort(v float32) {
	o.Port = &v
}

// GetTargetPort returns the TargetPort field value if set, zero value otherwise.
func (o *ContainerWorkloadContainersInnerNetwork) GetTargetPort() float32 {
	if o == nil || IsNil(o.TargetPort) {
		var ret float32
		return ret
	}
	return *o.TargetPort
}

// GetTargetPortOk returns a tuple with the TargetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerWorkloadContainersInnerNetwork) GetTargetPortOk() (*float32, bool) {
	if o == nil || IsNil(o.TargetPort) {
		return nil, false
	}
	return o.TargetPort, true
}

// HasTargetPort returns a boolean if a field has been set.
func (o *ContainerWorkloadContainersInnerNetwork) HasTargetPort() bool {
	if o != nil && !IsNil(o.TargetPort) {
		return true
	}

	return false
}

// SetTargetPort gets a reference to the given float32 and assigns it to the TargetPort field.
func (o *ContainerWorkloadContainersInnerNetwork) SetTargetPort(v float32) {
	o.TargetPort = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ContainerWorkloadContainersInnerNetwork) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerWorkloadContainersInnerNetwork) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ContainerWorkloadContainersInnerNetwork) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ContainerWorkloadContainersInnerNetwork) SetProtocol(v string) {
	o.Protocol = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContainerWorkloadContainersInnerNetwork) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerWorkloadContainersInnerNetwork) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContainerWorkloadContainersInnerNetwork) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ContainerWorkloadContainersInnerNetwork) SetType(v string) {
	o.Type = &v
}

func (o ContainerWorkloadContainersInnerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerWorkloadContainersInnerNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.TargetPort) {
		toSerialize["targetPort"] = o.TargetPort
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableContainerWorkloadContainersInnerNetwork struct {
	value *ContainerWorkloadContainersInnerNetwork
	isSet bool
}

func (v NullableContainerWorkloadContainersInnerNetwork) Get() *ContainerWorkloadContainersInnerNetwork {
	return v.value
}

func (v *NullableContainerWorkloadContainersInnerNetwork) Set(val *ContainerWorkloadContainersInnerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerWorkloadContainersInnerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerWorkloadContainersInnerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerWorkloadContainersInnerNetwork(val *ContainerWorkloadContainersInnerNetwork) *NullableContainerWorkloadContainersInnerNetwork {
	return &NullableContainerWorkloadContainersInnerNetwork{value: val, isSet: true}
}

func (v NullableContainerWorkloadContainersInnerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerWorkloadContainersInnerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

