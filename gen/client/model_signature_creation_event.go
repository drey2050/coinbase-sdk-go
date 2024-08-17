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

// checks if the SignatureCreationEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureCreationEvent{}

// SignatureCreationEvent An event representing a signature creation.
type SignatureCreationEvent struct {
	// The ID of the seed that the server-signer should create the signature for
	SeedId string `json:"seed_id"`
	// The ID of the wallet the signature is for
	WalletId string `json:"wallet_id"`
	// The ID of the user that the wallet belongs to
	WalletUserId string `json:"wallet_user_id"`
	// The ID of the address the transfer belongs to
	AddressId string `json:"address_id"`
	// The index of the address that the server-signer should sign with
	AddressIndex int32 `json:"address_index"`
	// The payload that the server-signer should sign
	SigningPayload string `json:"signing_payload"`
	TransactionType TransactionType `json:"transaction_type"`
	// The ID of the transaction that the server-signer should sign
	TransactionId string `json:"transaction_id"`
}

type _SignatureCreationEvent SignatureCreationEvent

// NewSignatureCreationEvent instantiates a new SignatureCreationEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureCreationEvent(seedId string, walletId string, walletUserId string, addressId string, addressIndex int32, signingPayload string, transactionType TransactionType, transactionId string) *SignatureCreationEvent {
	this := SignatureCreationEvent{}
	this.SeedId = seedId
	this.WalletId = walletId
	this.WalletUserId = walletUserId
	this.AddressId = addressId
	this.AddressIndex = addressIndex
	this.SigningPayload = signingPayload
	this.TransactionType = transactionType
	this.TransactionId = transactionId
	return &this
}

// NewSignatureCreationEventWithDefaults instantiates a new SignatureCreationEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureCreationEventWithDefaults() *SignatureCreationEvent {
	this := SignatureCreationEvent{}
	return &this
}

// GetSeedId returns the SeedId field value
func (o *SignatureCreationEvent) GetSeedId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeedId
}

// GetSeedIdOk returns a tuple with the SeedId field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEvent) GetSeedIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeedId, true
}

// SetSeedId sets field value
func (o *SignatureCreationEvent) SetSeedId(v string) {
	o.SeedId = v
}

// GetWalletId returns the WalletId field value
func (o *SignatureCreationEvent) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEvent) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *SignatureCreationEvent) SetWalletId(v string) {
	o.WalletId = v
}

// GetWalletUserId returns the WalletUserId field value
func (o *SignatureCreationEvent) GetWalletUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletUserId
}

// GetWalletUserIdOk returns a tuple with the WalletUserId field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEvent) GetWalletUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletUserId, true
}

// SetWalletUserId sets field value
func (o *SignatureCreationEvent) SetWalletUserId(v string) {
	o.WalletUserId = v
}

// GetAddressId returns the AddressId field value
func (o *SignatureCreationEvent) GetAddressId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEvent) GetAddressIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressId, true
}

// SetAddressId sets field value
func (o *SignatureCreationEvent) SetAddressId(v string) {
	o.AddressId = v
}

// GetAddressIndex returns the AddressIndex field value
func (o *SignatureCreationEvent) GetAddressIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AddressIndex
}

// GetAddressIndexOk returns a tuple with the AddressIndex field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEvent) GetAddressIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressIndex, true
}

// SetAddressIndex sets field value
func (o *SignatureCreationEvent) SetAddressIndex(v int32) {
	o.AddressIndex = v
}

// GetSigningPayload returns the SigningPayload field value
func (o *SignatureCreationEvent) GetSigningPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SigningPayload
}

// GetSigningPayloadOk returns a tuple with the SigningPayload field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEvent) GetSigningPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningPayload, true
}

// SetSigningPayload sets field value
func (o *SignatureCreationEvent) SetSigningPayload(v string) {
	o.SigningPayload = v
}

// GetTransactionType returns the TransactionType field value
func (o *SignatureCreationEvent) GetTransactionType() TransactionType {
	if o == nil {
		var ret TransactionType
		return ret
	}

	return o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEvent) GetTransactionTypeOk() (*TransactionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionType, true
}

// SetTransactionType sets field value
func (o *SignatureCreationEvent) SetTransactionType(v TransactionType) {
	o.TransactionType = v
}

// GetTransactionId returns the TransactionId field value
func (o *SignatureCreationEvent) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *SignatureCreationEvent) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *SignatureCreationEvent) SetTransactionId(v string) {
	o.TransactionId = v
}

func (o SignatureCreationEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureCreationEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["seed_id"] = o.SeedId
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["wallet_user_id"] = o.WalletUserId
	toSerialize["address_id"] = o.AddressId
	toSerialize["address_index"] = o.AddressIndex
	toSerialize["signing_payload"] = o.SigningPayload
	toSerialize["transaction_type"] = o.TransactionType
	toSerialize["transaction_id"] = o.TransactionId
	return toSerialize, nil
}

func (o *SignatureCreationEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"seed_id",
		"wallet_id",
		"wallet_user_id",
		"address_id",
		"address_index",
		"signing_payload",
		"transaction_type",
		"transaction_id",
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

	varSignatureCreationEvent := _SignatureCreationEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSignatureCreationEvent)

	if err != nil {
		return err
	}

	*o = SignatureCreationEvent(varSignatureCreationEvent)

	return err
}

type NullableSignatureCreationEvent struct {
	value *SignatureCreationEvent
	isSet bool
}

func (v NullableSignatureCreationEvent) Get() *SignatureCreationEvent {
	return v.value
}

func (v *NullableSignatureCreationEvent) Set(val *SignatureCreationEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureCreationEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureCreationEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureCreationEvent(val *SignatureCreationEvent) *NullableSignatureCreationEvent {
	return &NullableSignatureCreationEvent{value: val, isSet: true}
}

func (v NullableSignatureCreationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureCreationEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
