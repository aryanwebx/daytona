/*
Daytona Server API

Daytona Server API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverapiclient

import (
	"encoding/json"
)

// checks if the ProviderTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderTarget{}

// ProviderTarget struct for ProviderTarget
type ProviderTarget struct {
	Name *string `json:"name,omitempty"`
	// JSON encoded map of options
	Options *string `json:"options,omitempty"`
	ProviderInfo *ProviderProviderInfo `json:"providerInfo,omitempty"`
}

// NewProviderTarget instantiates a new ProviderTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderTarget() *ProviderTarget {
	this := ProviderTarget{}
	return &this
}

// NewProviderTargetWithDefaults instantiates a new ProviderTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderTargetWithDefaults() *ProviderTarget {
	this := ProviderTarget{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProviderTarget) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderTarget) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProviderTarget) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProviderTarget) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ProviderTarget) GetOptions() string {
	if o == nil || IsNil(o.Options) {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderTarget) GetOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProviderTarget) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *ProviderTarget) SetOptions(v string) {
	o.Options = &v
}

// GetProviderInfo returns the ProviderInfo field value if set, zero value otherwise.
func (o *ProviderTarget) GetProviderInfo() ProviderProviderInfo {
	if o == nil || IsNil(o.ProviderInfo) {
		var ret ProviderProviderInfo
		return ret
	}
	return *o.ProviderInfo
}

// GetProviderInfoOk returns a tuple with the ProviderInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderTarget) GetProviderInfoOk() (*ProviderProviderInfo, bool) {
	if o == nil || IsNil(o.ProviderInfo) {
		return nil, false
	}
	return o.ProviderInfo, true
}

// HasProviderInfo returns a boolean if a field has been set.
func (o *ProviderTarget) HasProviderInfo() bool {
	if o != nil && !IsNil(o.ProviderInfo) {
		return true
	}

	return false
}

// SetProviderInfo gets a reference to the given ProviderProviderInfo and assigns it to the ProviderInfo field.
func (o *ProviderTarget) SetProviderInfo(v ProviderProviderInfo) {
	o.ProviderInfo = &v
}

func (o ProviderTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.ProviderInfo) {
		toSerialize["providerInfo"] = o.ProviderInfo
	}
	return toSerialize, nil
}

type NullableProviderTarget struct {
	value *ProviderTarget
	isSet bool
}

func (v NullableProviderTarget) Get() *ProviderTarget {
	return v.value
}

func (v *NullableProviderTarget) Set(val *ProviderTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderTarget(val *ProviderTarget) *NullableProviderTarget {
	return &NullableProviderTarget{value: val, isSet: true}
}

func (v NullableProviderTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

