/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
Contact: yuga.cohler@coinbase.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AddressList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressList{}

// AddressList 
type AddressList struct {
	Data []Address `json:"data"`
	// True if this list has another page of items after this one that can be fetched.
	HasMore bool `json:"has_more"`
	// The page token to be used to fetch the next page.
	NextPage string `json:"next_page"`
	// The total number of addresses for the wallet.
	TotalCount int32 `json:"total_count"`
}

type _AddressList AddressList

// NewAddressList instantiates a new AddressList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressList(data []Address, hasMore bool, nextPage string, totalCount int32) *AddressList {
	this := AddressList{}
	this.Data = data
	this.HasMore = hasMore
	this.NextPage = nextPage
	this.TotalCount = totalCount
	return &this
}

// NewAddressListWithDefaults instantiates a new AddressList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressListWithDefaults() *AddressList {
	this := AddressList{}
	return &this
}

// GetData returns the Data field value
func (o *AddressList) GetData() []Address {
	if o == nil {
		var ret []Address
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressList) GetDataOk() ([]Address, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AddressList) SetData(v []Address) {
	o.Data = v
}

// GetHasMore returns the HasMore field value
func (o *AddressList) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *AddressList) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *AddressList) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextPage returns the NextPage field value
func (o *AddressList) GetNextPage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value
// and a boolean to check if the value has been set.
func (o *AddressList) GetNextPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPage, true
}

// SetNextPage sets field value
func (o *AddressList) SetNextPage(v string) {
	o.NextPage = v
}

// GetTotalCount returns the TotalCount field value
func (o *AddressList) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *AddressList) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *AddressList) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o AddressList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["has_more"] = o.HasMore
	toSerialize["next_page"] = o.NextPage
	toSerialize["total_count"] = o.TotalCount
	return toSerialize, nil
}

func (o *AddressList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"has_more",
		"next_page",
		"total_count",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddressList := _AddressList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressList)

	if err != nil {
		return err
	}

	*o = AddressList(varAddressList)

	return err
}

type NullableAddressList struct {
	value *AddressList
	isSet bool
}

func (v NullableAddressList) Get() *AddressList {
	return v.value
}

func (v *NullableAddressList) Set(val *AddressList) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressList) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressList(val *AddressList) *NullableAddressList {
	return &NullableAddressList{value: val, isSet: true}
}

func (v NullableAddressList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
