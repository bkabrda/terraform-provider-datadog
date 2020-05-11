/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// LogsIndex Object describing a Datadog Log index.
type LogsIndex struct {
	// The number of log events you can send in this index per day before you are rate-limited.
	DailyLimit *int64 `json:"daily_limit,omitempty"`
	// An array of exclusion objects. The logs are tested against the query of each filter, following the order of the array. Only the first matching active exclusion matters, others (if any) are ignored.
	ExclusionFilters *[]LogsExclusion `json:"exclusion_filters,omitempty"`
	Filter           LogsFilter       `json:"filter"`
	// A boolean stating if the index is rate limited, meaning more logs than the daily limit have been sent. Rate limit is reset every-day at 2pm UTC.
	IsRateLimited *bool `json:"is_rate_limited,omitempty"`
	// The name of the index.
	Name *string `json:"name,omitempty"`
	// The number of days before logs are deleted from this index.
	NumRetentionDays *int64 `json:"num_retention_days,omitempty"`
}

// NewLogsIndex instantiates a new LogsIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsIndex(filter LogsFilter) *LogsIndex {
	this := LogsIndex{}
	this.Filter = filter
	return &this
}

// NewLogsIndexWithDefaults instantiates a new LogsIndex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsIndexWithDefaults() *LogsIndex {
	this := LogsIndex{}
	return &this
}

// GetDailyLimit returns the DailyLimit field value if set, zero value otherwise.
func (o *LogsIndex) GetDailyLimit() int64 {
	if o == nil || o.DailyLimit == nil {
		var ret int64
		return ret
	}
	return *o.DailyLimit
}

// GetDailyLimitOk returns a tuple with the DailyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndex) GetDailyLimitOk() (*int64, bool) {
	if o == nil || o.DailyLimit == nil {
		return nil, false
	}
	return o.DailyLimit, true
}

// HasDailyLimit returns a boolean if a field has been set.
func (o *LogsIndex) HasDailyLimit() bool {
	if o != nil && o.DailyLimit != nil {
		return true
	}

	return false
}

// SetDailyLimit gets a reference to the given int64 and assigns it to the DailyLimit field.
func (o *LogsIndex) SetDailyLimit(v int64) {
	o.DailyLimit = &v
}

// GetExclusionFilters returns the ExclusionFilters field value if set, zero value otherwise.
func (o *LogsIndex) GetExclusionFilters() []LogsExclusion {
	if o == nil || o.ExclusionFilters == nil {
		var ret []LogsExclusion
		return ret
	}
	return *o.ExclusionFilters
}

// GetExclusionFiltersOk returns a tuple with the ExclusionFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndex) GetExclusionFiltersOk() (*[]LogsExclusion, bool) {
	if o == nil || o.ExclusionFilters == nil {
		return nil, false
	}
	return o.ExclusionFilters, true
}

// HasExclusionFilters returns a boolean if a field has been set.
func (o *LogsIndex) HasExclusionFilters() bool {
	if o != nil && o.ExclusionFilters != nil {
		return true
	}

	return false
}

// SetExclusionFilters gets a reference to the given []LogsExclusion and assigns it to the ExclusionFilters field.
func (o *LogsIndex) SetExclusionFilters(v []LogsExclusion) {
	o.ExclusionFilters = &v
}

// GetFilter returns the Filter field value
func (o *LogsIndex) GetFilter() LogsFilter {
	if o == nil {
		var ret LogsFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *LogsIndex) GetFilterOk() (*LogsFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *LogsIndex) SetFilter(v LogsFilter) {
	o.Filter = v
}

// GetIsRateLimited returns the IsRateLimited field value if set, zero value otherwise.
func (o *LogsIndex) GetIsRateLimited() bool {
	if o == nil || o.IsRateLimited == nil {
		var ret bool
		return ret
	}
	return *o.IsRateLimited
}

// GetIsRateLimitedOk returns a tuple with the IsRateLimited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndex) GetIsRateLimitedOk() (*bool, bool) {
	if o == nil || o.IsRateLimited == nil {
		return nil, false
	}
	return o.IsRateLimited, true
}

// HasIsRateLimited returns a boolean if a field has been set.
func (o *LogsIndex) HasIsRateLimited() bool {
	if o != nil && o.IsRateLimited != nil {
		return true
	}

	return false
}

// SetIsRateLimited gets a reference to the given bool and assigns it to the IsRateLimited field.
func (o *LogsIndex) SetIsRateLimited(v bool) {
	o.IsRateLimited = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsIndex) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndex) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsIndex) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsIndex) SetName(v string) {
	o.Name = &v
}

// GetNumRetentionDays returns the NumRetentionDays field value if set, zero value otherwise.
func (o *LogsIndex) GetNumRetentionDays() int64 {
	if o == nil || o.NumRetentionDays == nil {
		var ret int64
		return ret
	}
	return *o.NumRetentionDays
}

// GetNumRetentionDaysOk returns a tuple with the NumRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndex) GetNumRetentionDaysOk() (*int64, bool) {
	if o == nil || o.NumRetentionDays == nil {
		return nil, false
	}
	return o.NumRetentionDays, true
}

// HasNumRetentionDays returns a boolean if a field has been set.
func (o *LogsIndex) HasNumRetentionDays() bool {
	if o != nil && o.NumRetentionDays != nil {
		return true
	}

	return false
}

// SetNumRetentionDays gets a reference to the given int64 and assigns it to the NumRetentionDays field.
func (o *LogsIndex) SetNumRetentionDays(v int64) {
	o.NumRetentionDays = &v
}

func (o LogsIndex) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DailyLimit != nil {
		toSerialize["daily_limit"] = o.DailyLimit
	}
	if o.ExclusionFilters != nil {
		toSerialize["exclusion_filters"] = o.ExclusionFilters
	}
	if true {
		toSerialize["filter"] = o.Filter
	}
	if o.IsRateLimited != nil {
		toSerialize["is_rate_limited"] = o.IsRateLimited
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NumRetentionDays != nil {
		toSerialize["num_retention_days"] = o.NumRetentionDays
	}
	return json.Marshal(toSerialize)
}

type NullableLogsIndex struct {
	value *LogsIndex
	isSet bool
}

func (v NullableLogsIndex) Get() *LogsIndex {
	return v.value
}

func (v *NullableLogsIndex) Set(val *LogsIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsIndex(val *LogsIndex) *NullableLogsIndex {
	return &NullableLogsIndex{value: val, isSet: true}
}

func (v NullableLogsIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}