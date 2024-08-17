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

// checks if the FetchStakingRewards200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchStakingRewards200Response{}

// FetchStakingRewards200Response 
type FetchStakingRewards200Response struct {
	Data []StakingReward `json:"data"`
	// True if this list has another page of items after this one that can be fetched.
	HasMore bool `json:"has_more"`
	// The page token to be used to fetch the next page.
	NextPage string `json:"next_page"`
}

type _FetchStakingRewards200Response FetchStakingRewards200Response

// NewFetchStakingRewards200Response instantiates a new FetchStakingRewards200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchStakingRewards200Response(data []StakingReward, hasMore bool, nextPage string) *FetchStakingRewards200Response {
	this := FetchStakingRewards200Response{}
	this.Data = data
	this.HasMore = hasMore
	this.NextPage = nextPage
	return &this
}

// NewFetchStakingRewards200ResponseWithDefaults instantiates a new FetchStakingRewards200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchStakingRewards200ResponseWithDefaults() *FetchStakingRewards200Response {
	this := FetchStakingRewards200Response{}
	return &this
}

// GetData returns the Data field value
func (o *FetchStakingRewards200Response) GetData() []StakingReward {
	if o == nil {
		var ret []StakingReward
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FetchStakingRewards200Response) GetDataOk() ([]StakingReward, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *FetchStakingRewards200Response) SetData(v []StakingReward) {
	o.Data = v
}

// GetHasMore returns the HasMore field value
func (o *FetchStakingRewards200Response) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *FetchStakingRewards200Response) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *FetchStakingRewards200Response) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextPage returns the NextPage field value
func (o *FetchStakingRewards200Response) GetNextPage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value
// and a boolean to check if the value has been set.
func (o *FetchStakingRewards200Response) GetNextPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPage, true
}

// SetNextPage sets field value
func (o *FetchStakingRewards200Response) SetNextPage(v string) {
	o.NextPage = v
}

func (o FetchStakingRewards200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchStakingRewards200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["has_more"] = o.HasMore
	toSerialize["next_page"] = o.NextPage
	return toSerialize, nil
}

func (o *FetchStakingRewards200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"has_more",
		"next_page",
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

	varFetchStakingRewards200Response := _FetchStakingRewards200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFetchStakingRewards200Response)

	if err != nil {
		return err
	}

	*o = FetchStakingRewards200Response(varFetchStakingRewards200Response)

	return err
}

type NullableFetchStakingRewards200Response struct {
	value *FetchStakingRewards200Response
	isSet bool
}

func (v NullableFetchStakingRewards200Response) Get() *FetchStakingRewards200Response {
	return v.value
}

func (v *NullableFetchStakingRewards200Response) Set(val *FetchStakingRewards200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchStakingRewards200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchStakingRewards200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchStakingRewards200Response(val *FetchStakingRewards200Response) *NullableFetchStakingRewards200Response {
	return &NullableFetchStakingRewards200Response{value: val, isSet: true}
}

func (v NullableFetchStakingRewards200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchStakingRewards200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
