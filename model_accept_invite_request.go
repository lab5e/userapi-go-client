/*
 * The User API
 *
 * API to manage teams, members and tokens
 *
 * API version: 1.3.11 lucky-fremont
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package userapi

import (
	"encoding/json"
)

// AcceptInviteRequest struct for AcceptInviteRequest
type AcceptInviteRequest struct {
	// The invite code to use.
	Code *string `json:"code,omitempty"`
}

// NewAcceptInviteRequest instantiates a new AcceptInviteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptInviteRequest() *AcceptInviteRequest {
	this := AcceptInviteRequest{}
	return &this
}

// NewAcceptInviteRequestWithDefaults instantiates a new AcceptInviteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptInviteRequestWithDefaults() *AcceptInviteRequest {
	this := AcceptInviteRequest{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AcceptInviteRequest) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptInviteRequest) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AcceptInviteRequest) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AcceptInviteRequest) SetCode(v string) {
	o.Code = &v
}

func (o AcceptInviteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableAcceptInviteRequest struct {
	value *AcceptInviteRequest
	isSet bool
}

func (v NullableAcceptInviteRequest) Get() *AcceptInviteRequest {
	return v.value
}

func (v *NullableAcceptInviteRequest) Set(val *AcceptInviteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptInviteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptInviteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptInviteRequest(val *AcceptInviteRequest) *NullableAcceptInviteRequest {
	return &NullableAcceptInviteRequest{value: val, isSet: true}
}

func (v NullableAcceptInviteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptInviteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


