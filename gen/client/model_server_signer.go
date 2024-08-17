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

// checks if the ServerSigner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerSigner{}

// ServerSigner A Server-Signer assigned to sign transactions in a wallet.
type ServerSigner struct {
	// The ID of the server-signer
	ServerSignerId string `json:"server_signer_id"`
	// The IDs of the wallets that the server-signer can sign for
	Wallets []string `json:"wallets,omitempty"`
	// Whether the Server-Signer uses MPC.
	IsMpc bool `json:"is_mpc"`
}

type _ServerSigner ServerSigner

// NewServerSigner instantiates a new ServerSigner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerSigner(serverSignerId string, isMpc bool) *ServerSigner {
	this := ServerSigner{}
	this.ServerSignerId = serverSignerId
	this.IsMpc = isMpc
	return &this
}

// NewServerSignerWithDefaults instantiates a new ServerSigner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerSignerWithDefaults() *ServerSigner {
	this := ServerSigner{}
	return &this
}

// GetServerSignerId returns the ServerSignerId field value
func (o *ServerSigner) GetServerSignerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerSignerId
}

// GetServerSignerIdOk returns a tuple with the ServerSignerId field value
// and a boolean to check if the value has been set.
func (o *ServerSigner) GetServerSignerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerSignerId, true
}

// SetServerSignerId sets field value
func (o *ServerSigner) SetServerSignerId(v string) {
	o.ServerSignerId = v
}

// GetWallets returns the Wallets field value if set, zero value otherwise.
func (o *ServerSigner) GetWallets() []string {
	if o == nil || IsNil(o.Wallets) {
		var ret []string
		return ret
	}
	return o.Wallets
}

// GetWalletsOk returns a tuple with the Wallets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerSigner) GetWalletsOk() ([]string, bool) {
	if o == nil || IsNil(o.Wallets) {
		return nil, false
	}
	return o.Wallets, true
}

// HasWallets returns a boolean if a field has been set.
func (o *ServerSigner) HasWallets() bool {
	if o != nil && !IsNil(o.Wallets) {
		return true
	}

	return false
}

// SetWallets gets a reference to the given []string and assigns it to the Wallets field.
func (o *ServerSigner) SetWallets(v []string) {
	o.Wallets = v
}

// GetIsMpc returns the IsMpc field value
func (o *ServerSigner) GetIsMpc() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMpc
}

// GetIsMpcOk returns a tuple with the IsMpc field value
// and a boolean to check if the value has been set.
func (o *ServerSigner) GetIsMpcOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMpc, true
}

// SetIsMpc sets field value
func (o *ServerSigner) SetIsMpc(v bool) {
	o.IsMpc = v
}

func (o ServerSigner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerSigner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["server_signer_id"] = o.ServerSignerId
	if !IsNil(o.Wallets) {
		toSerialize["wallets"] = o.Wallets
	}
	toSerialize["is_mpc"] = o.IsMpc
	return toSerialize, nil
}

func (o *ServerSigner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"server_signer_id",
		"is_mpc",
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

	varServerSigner := _ServerSigner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServerSigner)

	if err != nil {
		return err
	}

	*o = ServerSigner(varServerSigner)

	return err
}

type NullableServerSigner struct {
	value *ServerSigner
	isSet bool
}

func (v NullableServerSigner) Get() *ServerSigner {
	return v.value
}

func (v *NullableServerSigner) Set(val *ServerSigner) {
	v.value = val
	v.isSet = true
}

func (v NullableServerSigner) IsSet() bool {
	return v.isSet
}

func (v *NullableServerSigner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerSigner(val *ServerSigner) *NullableServerSigner {
	return &NullableServerSigner{value: val, isSet: true}
}

func (v NullableServerSigner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerSigner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
