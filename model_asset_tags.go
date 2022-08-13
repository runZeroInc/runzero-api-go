/*
 * Rumble API
 *
 * Rumble Network Discovery API
 *
 * API version: 1.0.4
 * Contact: support@rumble.run
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// AssetTags struct for AssetTags
type AssetTags struct {
	Tags map[string]string `json:"tags"`
}

// NewAssetTags instantiates a new AssetTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTags(tags map[string]string, ) *AssetTags {
	this := AssetTags{}
	this.Tags = tags
	return &this
}

// NewAssetTagsWithDefaults instantiates a new AssetTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTagsWithDefaults() *AssetTags {
	this := AssetTags{}
	return &this
}

// GetTags returns the Tags field value
func (o *AssetTags) GetTags() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *AssetTags) GetTagsOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *AssetTags) SetTags(v map[string]string) {
	o.Tags = v
}

func (o AssetTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableAssetTags struct {
	value *AssetTags
	isSet bool
}

func (v NullableAssetTags) Get() *AssetTags {
	return v.value
}

func (v *NullableAssetTags) Set(val *AssetTags) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTags) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTags(val *AssetTags) *NullableAssetTags {
	return &NullableAssetTags{value: val, isSet: true}
}

func (v NullableAssetTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
