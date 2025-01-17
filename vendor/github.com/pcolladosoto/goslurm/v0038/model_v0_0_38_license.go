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

// checks if the V0038License type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038License{}

// V0038License struct for V0038License
type V0038License struct {
	// name of license
	LicenseName *string `json:"LicenseName,omitempty"`
	// total number of licenses
	Total *int32 `json:"Total,omitempty"`
	// number of licenses in use
	Used *int32 `json:"Used,omitempty"`
	// number of licenses available
	Free *int32 `json:"Free,omitempty"`
	// number of licenses reserved
	Reserved *int32 `json:"Reserved,omitempty"`
	// license is remote
	Remote *bool `json:"Remote,omitempty"`
}

// NewV0038License instantiates a new V0038License object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038License() *V0038License {
	this := V0038License{}
	return &this
}

// NewV0038LicenseWithDefaults instantiates a new V0038License object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038LicenseWithDefaults() *V0038License {
	this := V0038License{}
	return &this
}

// GetLicenseName returns the LicenseName field value if set, zero value otherwise.
func (o *V0038License) GetLicenseName() string {
	if o == nil || IsNil(o.LicenseName) {
		var ret string
		return ret
	}
	return *o.LicenseName
}

// GetLicenseNameOk returns a tuple with the LicenseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038License) GetLicenseNameOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseName) {
		return nil, false
	}
	return o.LicenseName, true
}

// HasLicenseName returns a boolean if a field has been set.
func (o *V0038License) HasLicenseName() bool {
	if o != nil && !IsNil(o.LicenseName) {
		return true
	}

	return false
}

// SetLicenseName gets a reference to the given string and assigns it to the LicenseName field.
func (o *V0038License) SetLicenseName(v string) {
	o.LicenseName = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0038License) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038License) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0038License) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *V0038License) SetTotal(v int32) {
	o.Total = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *V0038License) GetUsed() int32 {
	if o == nil || IsNil(o.Used) {
		var ret int32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038License) GetUsedOk() (*int32, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *V0038License) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int32 and assigns it to the Used field.
func (o *V0038License) SetUsed(v int32) {
	o.Used = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *V0038License) GetFree() int32 {
	if o == nil || IsNil(o.Free) {
		var ret int32
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038License) GetFreeOk() (*int32, bool) {
	if o == nil || IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *V0038License) HasFree() bool {
	if o != nil && !IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given int32 and assigns it to the Free field.
func (o *V0038License) SetFree(v int32) {
	o.Free = &v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *V0038License) GetReserved() int32 {
	if o == nil || IsNil(o.Reserved) {
		var ret int32
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038License) GetReservedOk() (*int32, bool) {
	if o == nil || IsNil(o.Reserved) {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *V0038License) HasReserved() bool {
	if o != nil && !IsNil(o.Reserved) {
		return true
	}

	return false
}

// SetReserved gets a reference to the given int32 and assigns it to the Reserved field.
func (o *V0038License) SetReserved(v int32) {
	o.Reserved = &v
}

// GetRemote returns the Remote field value if set, zero value otherwise.
func (o *V0038License) GetRemote() bool {
	if o == nil || IsNil(o.Remote) {
		var ret bool
		return ret
	}
	return *o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038License) GetRemoteOk() (*bool, bool) {
	if o == nil || IsNil(o.Remote) {
		return nil, false
	}
	return o.Remote, true
}

// HasRemote returns a boolean if a field has been set.
func (o *V0038License) HasRemote() bool {
	if o != nil && !IsNil(o.Remote) {
		return true
	}

	return false
}

// SetRemote gets a reference to the given bool and assigns it to the Remote field.
func (o *V0038License) SetRemote(v bool) {
	o.Remote = &v
}

func (o V0038License) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038License) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LicenseName) {
		toSerialize["LicenseName"] = o.LicenseName
	}
	if !IsNil(o.Total) {
		toSerialize["Total"] = o.Total
	}
	if !IsNil(o.Used) {
		toSerialize["Used"] = o.Used
	}
	if !IsNil(o.Free) {
		toSerialize["Free"] = o.Free
	}
	if !IsNil(o.Reserved) {
		toSerialize["Reserved"] = o.Reserved
	}
	if !IsNil(o.Remote) {
		toSerialize["Remote"] = o.Remote
	}
	return toSerialize, nil
}

type NullableV0038License struct {
	value *V0038License
	isSet bool
}

func (v NullableV0038License) Get() *V0038License {
	return v.value
}

func (v *NullableV0038License) Set(val *V0038License) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038License) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038License) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038License(val *V0038License) *NullableV0038License {
	return &NullableV0038License{value: val, isSet: true}
}

func (v NullableV0038License) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038License) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


