/*
RIK API

RESTful public-facing API. The API is accessible through HTTP calls on specific URLs carrying JSON modeled data. 

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the InstanceDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceDefinition{}

// InstanceDefinition struct for InstanceDefinition
type InstanceDefinition struct {
	Name string `json:"name"`
	WorkloadId string `json:"workload_id"`
	Replicas *int32 `json:"replicas,omitempty"`
}

// NewInstanceDefinition instantiates a new InstanceDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceDefinition(name string, workloadId string) *InstanceDefinition {
	this := InstanceDefinition{}
	this.Name = name
	this.WorkloadId = workloadId
	return &this
}

// NewInstanceDefinitionWithDefaults instantiates a new InstanceDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceDefinitionWithDefaults() *InstanceDefinition {
	this := InstanceDefinition{}
	return &this
}

// GetName returns the Name field value
func (o *InstanceDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InstanceDefinition) SetName(v string) {
	o.Name = v
}

// GetWorkloadId returns the WorkloadId field value
func (o *InstanceDefinition) GetWorkloadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkloadId
}

// GetWorkloadIdOk returns a tuple with the WorkloadId field value
// and a boolean to check if the value has been set.
func (o *InstanceDefinition) GetWorkloadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkloadId, true
}

// SetWorkloadId sets field value
func (o *InstanceDefinition) SetWorkloadId(v string) {
	o.WorkloadId = v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *InstanceDefinition) GetReplicas() int32 {
	if o == nil || IsNil(o.Replicas) {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceDefinition) GetReplicasOk() (*int32, bool) {
	if o == nil || IsNil(o.Replicas) {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *InstanceDefinition) HasReplicas() bool {
	if o != nil && !IsNil(o.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *InstanceDefinition) SetReplicas(v int32) {
	o.Replicas = &v
}

func (o InstanceDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["workload_id"] = o.WorkloadId
	if !IsNil(o.Replicas) {
		toSerialize["replicas"] = o.Replicas
	}
	return toSerialize, nil
}

type NullableInstanceDefinition struct {
	value *InstanceDefinition
	isSet bool
}

func (v NullableInstanceDefinition) Get() *InstanceDefinition {
	return v.value
}

func (v *NullableInstanceDefinition) Set(val *InstanceDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceDefinition(val *InstanceDefinition) *NullableInstanceDefinition {
	return &NullableInstanceDefinition{value: val, isSet: true}
}

func (v NullableInstanceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


