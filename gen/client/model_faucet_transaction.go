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

// checks if the FaucetTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FaucetTransaction{}

// FaucetTransaction The faucet transaction
type FaucetTransaction struct {
	// The transaction hash of the transaction the faucet created.
	TransactionHash string `json:"transaction_hash"`
	// Link to the transaction on the blockchain explorer.
	TransactionLink string `json:"transaction_link"`
}

type _FaucetTransaction FaucetTransaction

// NewFaucetTransaction instantiates a new FaucetTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFaucetTransaction(transactionHash string, transactionLink string) *FaucetTransaction {
	this := FaucetTransaction{}
	this.TransactionHash = transactionHash
	this.TransactionLink = transactionLink
	return &this
}

// NewFaucetTransactionWithDefaults instantiates a new FaucetTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFaucetTransactionWithDefaults() *FaucetTransaction {
	this := FaucetTransaction{}
	return &this
}

// GetTransactionHash returns the TransactionHash field value
func (o *FaucetTransaction) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *FaucetTransaction) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *FaucetTransaction) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetTransactionLink returns the TransactionLink field value
func (o *FaucetTransaction) GetTransactionLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionLink
}

// GetTransactionLinkOk returns a tuple with the TransactionLink field value
// and a boolean to check if the value has been set.
func (o *FaucetTransaction) GetTransactionLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionLink, true
}

// SetTransactionLink sets field value
func (o *FaucetTransaction) SetTransactionLink(v string) {
	o.TransactionLink = v
}

func (o FaucetTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FaucetTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_hash"] = o.TransactionHash
	toSerialize["transaction_link"] = o.TransactionLink
	return toSerialize, nil
}

func (o *FaucetTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_hash",
		"transaction_link",
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

	varFaucetTransaction := _FaucetTransaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFaucetTransaction)

	if err != nil {
		return err
	}

	*o = FaucetTransaction(varFaucetTransaction)

	return err
}

type NullableFaucetTransaction struct {
	value *FaucetTransaction
	isSet bool
}

func (v NullableFaucetTransaction) Get() *FaucetTransaction {
	return v.value
}

func (v *NullableFaucetTransaction) Set(val *FaucetTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableFaucetTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableFaucetTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaucetTransaction(val *FaucetTransaction) *NullableFaucetTransaction {
	return &NullableFaucetTransaction{value: val, isSet: true}
}

func (v NullableFaucetTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaucetTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

