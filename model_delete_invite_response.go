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

// DeleteInviteResponse struct for DeleteInviteResponse
type DeleteInviteResponse struct {
	Invite *Invite `json:"invite,omitempty"`
}

// NewDeleteInviteResponse instantiates a new DeleteInviteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteInviteResponse() *DeleteInviteResponse {
	this := DeleteInviteResponse{}
	return &this
}

// NewDeleteInviteResponseWithDefaults instantiates a new DeleteInviteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteInviteResponseWithDefaults() *DeleteInviteResponse {
	this := DeleteInviteResponse{}
	return &this
}

// GetInvite returns the Invite field value if set, zero value otherwise.
func (o *DeleteInviteResponse) GetInvite() Invite {
	if o == nil || o.Invite == nil {
		var ret Invite
		return ret
	}
	return *o.Invite
}

// GetInviteOk returns a tuple with the Invite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteInviteResponse) GetInviteOk() (*Invite, bool) {
	if o == nil || o.Invite == nil {
		return nil, false
	}
	return o.Invite, true
}

// HasInvite returns a boolean if a field has been set.
func (o *DeleteInviteResponse) HasInvite() bool {
	if o != nil && o.Invite != nil {
		return true
	}

	return false
}

// SetInvite gets a reference to the given Invite and assigns it to the Invite field.
func (o *DeleteInviteResponse) SetInvite(v Invite) {
	o.Invite = &v
}

func (o DeleteInviteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Invite != nil {
		toSerialize["invite"] = o.Invite
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteInviteResponse struct {
	value *DeleteInviteResponse
	isSet bool
}

func (v NullableDeleteInviteResponse) Get() *DeleteInviteResponse {
	return v.value
}

func (v *NullableDeleteInviteResponse) Set(val *DeleteInviteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteInviteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteInviteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteInviteResponse(val *DeleteInviteResponse) *NullableDeleteInviteResponse {
	return &NullableDeleteInviteResponse{value: val, isSet: true}
}

func (v NullableDeleteInviteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteInviteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
