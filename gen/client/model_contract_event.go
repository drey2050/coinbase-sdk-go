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
	"time"
	"bytes"
	"fmt"
)

// checks if the ContractEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractEvent{}

// ContractEvent Represents a single decoded event emitted by a smart contract
type ContractEvent struct {
	// The name of the blockchain network
	NetworkId string `json:"network_id"`
	// The name of the blockchain project or protocol
	ProtocolName string `json:"protocol_name"`
	// The name of the specific contract within the project
	ContractName string `json:"contract_name"`
	// The name of the event emitted by the contract
	EventName string `json:"event_name"`
	// The signature of the event, including parameter types
	Sig string `json:"sig"`
	// The first four bytes of the Keccak hash of the event signature
	FourBytes string `json:"four_bytes"`
	// The EVM address of the smart contract
	ContractAddress string `json:"contract_address"`
	// The timestamp of the block in which the event was emitted
	BlockTime time.Time `json:"block_time"`
	// The block number in which the event was emitted
	BlockHeight int32 `json:"block_height"`
	// The transaction hash in which the event was emitted
	TxHash string `json:"tx_hash"`
	// The index of the transaction within the block
	TxIndex int32 `json:"tx_index"`
	// The index of the event within the transaction
	EventIndex int32 `json:"event_index"`
	// The event data in a stringified format
	Data string `json:"data"`
}

type _ContractEvent ContractEvent

// NewContractEvent instantiates a new ContractEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractEvent(networkId string, protocolName string, contractName string, eventName string, sig string, fourBytes string, contractAddress string, blockTime time.Time, blockHeight int32, txHash string, txIndex int32, eventIndex int32, data string) *ContractEvent {
	this := ContractEvent{}
	this.NetworkId = networkId
	this.ProtocolName = protocolName
	this.ContractName = contractName
	this.EventName = eventName
	this.Sig = sig
	this.FourBytes = fourBytes
	this.ContractAddress = contractAddress
	this.BlockTime = blockTime
	this.BlockHeight = blockHeight
	this.TxHash = txHash
	this.TxIndex = txIndex
	this.EventIndex = eventIndex
	this.Data = data
	return &this
}

// NewContractEventWithDefaults instantiates a new ContractEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractEventWithDefaults() *ContractEvent {
	this := ContractEvent{}
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *ContractEvent) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *ContractEvent) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetProtocolName returns the ProtocolName field value
func (o *ContractEvent) GetProtocolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProtocolName
}

// GetProtocolNameOk returns a tuple with the ProtocolName field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetProtocolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtocolName, true
}

// SetProtocolName sets field value
func (o *ContractEvent) SetProtocolName(v string) {
	o.ProtocolName = v
}

// GetContractName returns the ContractName field value
func (o *ContractEvent) GetContractName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetContractNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractName, true
}

// SetContractName sets field value
func (o *ContractEvent) SetContractName(v string) {
	o.ContractName = v
}

// GetEventName returns the EventName field value
func (o *ContractEvent) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value
func (o *ContractEvent) SetEventName(v string) {
	o.EventName = v
}

// GetSig returns the Sig field value
func (o *ContractEvent) GetSig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sig
}

// GetSigOk returns a tuple with the Sig field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetSigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sig, true
}

// SetSig sets field value
func (o *ContractEvent) SetSig(v string) {
	o.Sig = v
}

// GetFourBytes returns the FourBytes field value
func (o *ContractEvent) GetFourBytes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FourBytes
}

// GetFourBytesOk returns a tuple with the FourBytes field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetFourBytesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FourBytes, true
}

// SetFourBytes sets field value
func (o *ContractEvent) SetFourBytes(v string) {
	o.FourBytes = v
}

// GetContractAddress returns the ContractAddress field value
func (o *ContractEvent) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *ContractEvent) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetBlockTime returns the BlockTime field value
func (o *ContractEvent) GetBlockTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetBlockTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTime, true
}

// SetBlockTime sets field value
func (o *ContractEvent) SetBlockTime(v time.Time) {
	o.BlockTime = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *ContractEvent) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *ContractEvent) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetTxHash returns the TxHash field value
func (o *ContractEvent) GetTxHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetTxHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxHash, true
}

// SetTxHash sets field value
func (o *ContractEvent) SetTxHash(v string) {
	o.TxHash = v
}

// GetTxIndex returns the TxIndex field value
func (o *ContractEvent) GetTxIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetTxIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxIndex, true
}

// SetTxIndex sets field value
func (o *ContractEvent) SetTxIndex(v int32) {
	o.TxIndex = v
}

// GetEventIndex returns the EventIndex field value
func (o *ContractEvent) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *ContractEvent) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetData returns the Data field value
func (o *ContractEvent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ContractEvent) SetData(v string) {
	o.Data = v
}

func (o ContractEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_id"] = o.NetworkId
	toSerialize["protocol_name"] = o.ProtocolName
	toSerialize["contract_name"] = o.ContractName
	toSerialize["event_name"] = o.EventName
	toSerialize["sig"] = o.Sig
	toSerialize["four_bytes"] = o.FourBytes
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["block_time"] = o.BlockTime
	toSerialize["block_height"] = o.BlockHeight
	toSerialize["tx_hash"] = o.TxHash
	toSerialize["tx_index"] = o.TxIndex
	toSerialize["event_index"] = o.EventIndex
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ContractEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_id",
		"protocol_name",
		"contract_name",
		"event_name",
		"sig",
		"four_bytes",
		"contract_address",
		"block_time",
		"block_height",
		"tx_hash",
		"tx_index",
		"event_index",
		"data",
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

	varContractEvent := _ContractEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractEvent)

	if err != nil {
		return err
	}

	*o = ContractEvent(varContractEvent)

	return err
}

type NullableContractEvent struct {
	value *ContractEvent
	isSet bool
}

func (v NullableContractEvent) Get() *ContractEvent {
	return v.value
}

func (v *NullableContractEvent) Set(val *ContractEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableContractEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableContractEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractEvent(val *ContractEvent) *NullableContractEvent {
	return &NullableContractEvent{value: val, isSet: true}
}

func (v NullableContractEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
