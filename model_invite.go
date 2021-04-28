/*
 * The User API
 *
 * API to manage teams, members and tokens
 *
 * API version: 1.3.9 cumulative-hardin
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package userapi

import (
	"encoding/json"
)

// Invite Invite. Invites are used to share access to teams. You must be an administrator to generate invites for a team.
type Invite struct {
	Code *string `json:"code,omitempty"`
	// Time stamp when the invite was generated.
	CreatedAt *string `json:"createdAt,omitempty"`
}

// NewInvite instantiates a new Invite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvite() *Invite {
	this := Invite{}
	return &this
}

// NewInviteWithDefaults instantiates a new Invite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteWithDefaults() *Invite {
	this := Invite{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Invite) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Invite) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Invite) SetCode(v string) {
	o.Code = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Invite) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Invite) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Invite) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o Invite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInvite struct {
	value *Invite
	isSet bool
}

func (v NullableInvite) Get() *Invite {
	return v.value
}

func (v *NullableInvite) Set(val *Invite) {
	v.value = val
	v.isSet = true
}

func (v NullableInvite) IsSet() bool {
	return v.isSet
}

func (v *NullableInvite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvite(val *Invite) *NullableInvite {
	return &NullableInvite{value: val, isSet: true}
}

func (v NullableInvite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
