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

// checks if the CreateWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWebhookRequest{}

// CreateWebhookRequest struct for CreateWebhookRequest
type CreateWebhookRequest struct {
	// The ID of the blockchain network
	NetworkId string `json:"network_id"`
	EventType *WebhookEventType `json:"event_type,omitempty"`
	// Webhook will monitor all events that matches any one of the event filters.
	EventFilters []WebhookEventFilter `json:"event_filters,omitempty"`
	// The URL to which the notifications will be sent
	NotificationUri string `json:"notification_uri"`
}

type _CreateWebhookRequest CreateWebhookRequest

// NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebhookRequest(networkId string, notificationUri string) *CreateWebhookRequest {
	this := CreateWebhookRequest{}
	this.NetworkId = networkId
	this.NotificationUri = notificationUri
	return &this
}

// NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest {
	this := CreateWebhookRequest{}
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *CreateWebhookRequest) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *CreateWebhookRequest) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetEventType() WebhookEventType {
	if o == nil || IsNil(o.EventType) {
		var ret WebhookEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetEventTypeOk() (*WebhookEventType, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given WebhookEventType and assigns it to the EventType field.
func (o *CreateWebhookRequest) SetEventType(v WebhookEventType) {
	o.EventType = &v
}

// GetEventFilters returns the EventFilters field value if set, zero value otherwise.
func (o *CreateWebhookRequest) GetEventFilters() []WebhookEventFilter {
	if o == nil || IsNil(o.EventFilters) {
		var ret []WebhookEventFilter
		return ret
	}
	return o.EventFilters
}

// GetEventFiltersOk returns a tuple with the EventFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetEventFiltersOk() ([]WebhookEventFilter, bool) {
	if o == nil || IsNil(o.EventFilters) {
		return nil, false
	}
	return o.EventFilters, true
}

// HasEventFilters returns a boolean if a field has been set.
func (o *CreateWebhookRequest) HasEventFilters() bool {
	if o != nil && !IsNil(o.EventFilters) {
		return true
	}

	return false
}

// SetEventFilters gets a reference to the given []WebhookEventFilter and assigns it to the EventFilters field.
func (o *CreateWebhookRequest) SetEventFilters(v []WebhookEventFilter) {
	o.EventFilters = v
}

// GetNotificationUri returns the NotificationUri field value
func (o *CreateWebhookRequest) GetNotificationUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequest) GetNotificationUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationUri, true
}

// SetNotificationUri sets field value
func (o *CreateWebhookRequest) SetNotificationUri(v string) {
	o.NotificationUri = v
}

func (o CreateWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_id"] = o.NetworkId
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if !IsNil(o.EventFilters) {
		toSerialize["event_filters"] = o.EventFilters
	}
	toSerialize["notification_uri"] = o.NotificationUri
	return toSerialize, nil
}

func (o *CreateWebhookRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_id",
		"notification_uri",
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

	varCreateWebhookRequest := _CreateWebhookRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWebhookRequest)

	if err != nil {
		return err
	}

	*o = CreateWebhookRequest(varCreateWebhookRequest)

	return err
}

type NullableCreateWebhookRequest struct {
	value *CreateWebhookRequest
	isSet bool
}

func (v NullableCreateWebhookRequest) Get() *CreateWebhookRequest {
	return v.value
}

func (v *NullableCreateWebhookRequest) Set(val *CreateWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebhookRequest(val *CreateWebhookRequest) *NullableCreateWebhookRequest {
	return &NullableCreateWebhookRequest{value: val, isSet: true}
}

func (v NullableCreateWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
