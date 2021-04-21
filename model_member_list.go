/*
 * The User API
 *
 * API to manage teams, members and tokens
 *
 * API version: 1.3.7 frequent-amara
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package userapi

import (
	"encoding/json"
)

// MemberList struct for MemberList
type MemberList struct {
	Members *[]Member `json:"members,omitempty"`
}

// NewMemberList instantiates a new MemberList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberList() *MemberList {
	this := MemberList{}
	return &this
}

// NewMemberListWithDefaults instantiates a new MemberList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberListWithDefaults() *MemberList {
	this := MemberList{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *MemberList) GetMembers() []Member {
	if o == nil || o.Members == nil {
		var ret []Member
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberList) GetMembersOk() (*[]Member, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *MemberList) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []Member and assigns it to the Members field.
func (o *MemberList) SetMembers(v []Member) {
	o.Members = &v
}

func (o MemberList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	return json.Marshal(toSerialize)
}

type NullableMemberList struct {
	value *MemberList
	isSet bool
}

func (v NullableMemberList) Get() *MemberList {
	return v.value
}

func (v *NullableMemberList) Set(val *MemberList) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberList) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberList(val *MemberList) *NullableMemberList {
	return &NullableMemberList{value: val, isSet: true}
}

func (v NullableMemberList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
