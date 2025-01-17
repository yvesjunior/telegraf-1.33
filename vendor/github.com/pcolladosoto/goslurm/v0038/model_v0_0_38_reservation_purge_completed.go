/*
Slurm Rest API RO

API to access Slurm. Only GET requests are implemented.

API version: 0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v0038

import (
	"encoding/json"
)

// checks if the V0038ReservationPurgeCompleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038ReservationPurgeCompleted{}

// V0038ReservationPurgeCompleted If PURGE_COMP flag is set the amount of seconds this reservation will sit idle until it is revoked
type V0038ReservationPurgeCompleted struct {
	// amount of seconds this reservation will sit idle until it is revoked
	Time *int32 `json:"time,omitempty"`
}

// NewV0038ReservationPurgeCompleted instantiates a new V0038ReservationPurgeCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038ReservationPurgeCompleted() *V0038ReservationPurgeCompleted {
	this := V0038ReservationPurgeCompleted{}
	return &this
}

// NewV0038ReservationPurgeCompletedWithDefaults instantiates a new V0038ReservationPurgeCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038ReservationPurgeCompletedWithDefaults() *V0038ReservationPurgeCompleted {
	this := V0038ReservationPurgeCompleted{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V0038ReservationPurgeCompleted) GetTime() int32 {
	if o == nil || IsNil(o.Time) {
		var ret int32
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038ReservationPurgeCompleted) GetTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V0038ReservationPurgeCompleted) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int32 and assigns it to the Time field.
func (o *V0038ReservationPurgeCompleted) SetTime(v int32) {
	o.Time = &v
}

func (o V0038ReservationPurgeCompleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038ReservationPurgeCompleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableV0038ReservationPurgeCompleted struct {
	value *V0038ReservationPurgeCompleted
	isSet bool
}

func (v NullableV0038ReservationPurgeCompleted) Get() *V0038ReservationPurgeCompleted {
	return v.value
}

func (v *NullableV0038ReservationPurgeCompleted) Set(val *V0038ReservationPurgeCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038ReservationPurgeCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038ReservationPurgeCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038ReservationPurgeCompleted(val *V0038ReservationPurgeCompleted) *NullableV0038ReservationPurgeCompleted {
	return &NullableV0038ReservationPurgeCompleted{value: val, isSet: true}
}

func (v NullableV0038ReservationPurgeCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038ReservationPurgeCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


