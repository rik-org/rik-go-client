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

// checks if the Workload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Workload{}

// Workload struct for Workload
type Workload struct {
	ApiVersion *string `json:"apiVersion,omitempty"`
	Kind *Kind `json:"kind,omitempty"`
	Name *string `json:"name,omitempty"`
	Spec *WorkloadSpec `json:"spec,omitempty"`
}

// NewWorkload instantiates a new Workload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkload() *Workload {
	this := Workload{}
	return &this
}

// NewWorkloadWithDefaults instantiates a new Workload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadWithDefaults() *Workload {
	this := Workload{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *Workload) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workload) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *Workload) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *Workload) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Workload) GetKind() Kind {
	if o == nil || IsNil(o.Kind) {
		var ret Kind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workload) GetKindOk() (*Kind, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Workload) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given Kind and assigns it to the Kind field.
func (o *Workload) SetKind(v Kind) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Workload) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Workload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Workload) SetName(v string) {
	o.Name = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *Workload) GetSpec() WorkloadSpec {
	if o == nil || IsNil(o.Spec) {
		var ret WorkloadSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workload) GetSpecOk() (*WorkloadSpec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *Workload) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given WorkloadSpec and assigns it to the Spec field.
func (o *Workload) SetSpec(v WorkloadSpec) {
	o.Spec = &v
}

func (o Workload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Workload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	return toSerialize, nil
}

type NullableWorkload struct {
	value *Workload
	isSet bool
}

func (v NullableWorkload) Get() *Workload {
	return v.value
}

func (v *NullableWorkload) Set(val *Workload) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkload) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkload(val *Workload) *NullableWorkload {
	return &NullableWorkload{value: val, isSet: true}
}

func (v NullableWorkload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


