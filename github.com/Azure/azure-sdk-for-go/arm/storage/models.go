package storage

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/openshift/github.com/Azure/go-autorest/autorest"
	"github.com/openshift/github.com/Azure/go-autorest/autorest/date"
)

// AccessTier enumerates the values for access tier.
type AccessTier string

const (
	// Cool specifies the cool state for access tier.
	Cool AccessTier = "Cool"
	// Hot specifies the hot state for access tier.
	Hot AccessTier = "Hot"
)

// AccountStatus enumerates the values for account status.
type AccountStatus string

const (
	// Available specifies the available state for account status.
	Available AccountStatus = "Available"
	// Unavailable specifies the unavailable state for account status.
	Unavailable AccountStatus = "Unavailable"
)

// KeyPermission enumerates the values for key permission.
type KeyPermission string

const (
	// FULL specifies the full state for key permission.
	FULL KeyPermission = "FULL"
	// READ specifies the read state for key permission.
	READ KeyPermission = "READ"
)

// Kind enumerates the values for kind.
type Kind string

const (
	// BlobStorage specifies the blob storage state for kind.
	BlobStorage Kind = "BlobStorage"
	// Storage specifies the storage state for kind.
	Storage Kind = "Storage"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating specifies the creating state for provisioning state.
	Creating ProvisioningState = "Creating"
	// ResolvingDNS specifies the resolving dns state for provisioning state.
	ResolvingDNS ProvisioningState = "ResolvingDNS"
	// Succeeded specifies the succeeded state for provisioning state.
	Succeeded ProvisioningState = "Succeeded"
)

// Reason enumerates the values for reason.
type Reason string

const (
	// AccountNameInvalid specifies the account name invalid state for reason.
	AccountNameInvalid Reason = "AccountNameInvalid"
	// AlreadyExists specifies the already exists state for reason.
	AlreadyExists Reason = "AlreadyExists"
)

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// PremiumLRS specifies the premium lrs state for sku name.
	PremiumLRS SkuName = "Premium_LRS"
	// StandardGRS specifies the standard grs state for sku name.
	StandardGRS SkuName = "Standard_GRS"
	// StandardLRS specifies the standard lrs state for sku name.
	StandardLRS SkuName = "Standard_LRS"
	// StandardRAGRS specifies the standard ragrs state for sku name.
	StandardRAGRS SkuName = "Standard_RAGRS"
	// StandardZRS specifies the standard zrs state for sku name.
	StandardZRS SkuName = "Standard_ZRS"
)

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Premium specifies the premium state for sku tier.
	Premium SkuTier = "Premium"
	// Standard specifies the standard state for sku tier.
	Standard SkuTier = "Standard"
)

// UsageUnit enumerates the values for usage unit.
type UsageUnit string

const (
	// Bytes specifies the bytes state for usage unit.
	Bytes UsageUnit = "Bytes"
	// BytesPerSecond specifies the bytes per second state for usage unit.
	BytesPerSecond UsageUnit = "BytesPerSecond"
	// Count specifies the count state for usage unit.
	Count UsageUnit = "Count"
	// CountsPerSecond specifies the counts per second state for usage unit.
	CountsPerSecond UsageUnit = "CountsPerSecond"
	// Percent specifies the percent state for usage unit.
	Percent UsageUnit = "Percent"
	// Seconds specifies the seconds state for usage unit.
	Seconds UsageUnit = "Seconds"
)

// Account is the storage account.
type Account struct {
	autorest.Response  `json:"-"`
	ID                 *string             `json:"id,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Type               *string             `json:"type,omitempty"`
	Location           *string             `json:"location,omitempty"`
	Tags               *map[string]*string `json:"tags,omitempty"`
	Sku                *Sku                `json:"sku,omitempty"`
	Kind               Kind                `json:"kind,omitempty"`
	*AccountProperties `json:"properties,omitempty"`
}

// AccountCheckNameAvailabilityParameters is
type AccountCheckNameAvailabilityParameters struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// AccountCreateParameters is the parameters used when creating a storage
// account.
type AccountCreateParameters struct {
	Sku                                *Sku                `json:"sku,omitempty"`
	Kind                               Kind                `json:"kind,omitempty"`
	Location                           *string             `json:"location,omitempty"`
	Tags                               *map[string]*string `json:"tags,omitempty"`
	*AccountPropertiesCreateParameters `json:"properties,omitempty"`
}

// AccountKey is an access key for the storage account.
type AccountKey struct {
	KeyName     *string       `json:"keyName,omitempty"`
	Value       *string       `json:"value,omitempty"`
	Permissions KeyPermission `json:"permissions,omitempty"`
}

// AccountListKeysResult is the response from the ListKeys operation.
type AccountListKeysResult struct {
	autorest.Response `json:"-"`
	Keys              *[]AccountKey `json:"keys,omitempty"`
}

// AccountListResult is the response from the List Storage Accounts operation.
type AccountListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Account `json:"value,omitempty"`
}

// AccountProperties is
type AccountProperties struct {
	ProvisioningState   ProvisioningState `json:"provisioningState,omitempty"`
	PrimaryEndpoints    *Endpoints        `json:"primaryEndpoints,omitempty"`
	PrimaryLocation     *string           `json:"primaryLocation,omitempty"`
	StatusOfPrimary     AccountStatus     `json:"statusOfPrimary,omitempty"`
	LastGeoFailoverTime *date.Time        `json:"lastGeoFailoverTime,omitempty"`
	SecondaryLocation   *string           `json:"secondaryLocation,omitempty"`
	StatusOfSecondary   AccountStatus     `json:"statusOfSecondary,omitempty"`
	CreationTime        *date.Time        `json:"creationTime,omitempty"`
	CustomDomain        *CustomDomain     `json:"customDomain,omitempty"`
	SecondaryEndpoints  *Endpoints        `json:"secondaryEndpoints,omitempty"`
	Encryption          *Encryption       `json:"encryption,omitempty"`
	AccessTier          AccessTier        `json:"accessTier,omitempty"`
}

// AccountPropertiesCreateParameters is
type AccountPropertiesCreateParameters struct {
	CustomDomain *CustomDomain `json:"customDomain,omitempty"`
	Encryption   *Encryption   `json:"encryption,omitempty"`
	AccessTier   AccessTier    `json:"accessTier,omitempty"`
}

// AccountPropertiesUpdateParameters is
type AccountPropertiesUpdateParameters struct {
	CustomDomain *CustomDomain `json:"customDomain,omitempty"`
	Encryption   *Encryption   `json:"encryption,omitempty"`
	AccessTier   AccessTier    `json:"accessTier,omitempty"`
}

// AccountRegenerateKeyParameters is
type AccountRegenerateKeyParameters struct {
	KeyName *string `json:"keyName,omitempty"`
}

// AccountUpdateParameters is the parameters that can be provided when
// updating the storage account properties.
type AccountUpdateParameters struct {
	Sku                                *Sku                `json:"sku,omitempty"`
	Tags                               *map[string]*string `json:"tags,omitempty"`
	*AccountPropertiesUpdateParameters `json:"properties,omitempty"`
}

// CheckNameAvailabilityResult is the CheckNameAvailability operation response.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	NameAvailable     *bool   `json:"nameAvailable,omitempty"`
	Reason            Reason  `json:"reason,omitempty"`
	Message           *string `json:"message,omitempty"`
}

// CustomDomain is the custom domain assigned to this storage account. This
// can be set via Update.
type CustomDomain struct {
	Name         *string `json:"name,omitempty"`
	UseSubDomain *bool   `json:"useSubDomain,omitempty"`
}

// Encryption is the encryption settings on the storage account.
type Encryption struct {
	Services  *EncryptionServices `json:"services,omitempty"`
	KeySource *string             `json:"keySource,omitempty"`
}

// EncryptionService is a service that allows server-side encryption to be
// used.
type EncryptionService struct {
	Enabled         *bool      `json:"enabled,omitempty"`
	LastEnabledTime *date.Time `json:"lastEnabledTime,omitempty"`
}

// EncryptionServices is a list of services that support encryption.
type EncryptionServices struct {
	Blob *EncryptionService `json:"blob,omitempty"`
}

// Endpoints is the URIs that are used to perform a retrieval of a public
// blob, queue, or table object.
type Endpoints struct {
	Blob  *string `json:"blob,omitempty"`
	Queue *string `json:"queue,omitempty"`
	Table *string `json:"table,omitempty"`
	File  *string `json:"file,omitempty"`
}

// Resource is
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// Sku is the SKU of the storage account.
type Sku struct {
	Name SkuName `json:"name,omitempty"`
	Tier SkuTier `json:"tier,omitempty"`
}

// Usage is describes Storage Resource Usage.
type Usage struct {
	Unit         UsageUnit  `json:"unit,omitempty"`
	CurrentValue *int32     `json:"currentValue,omitempty"`
	Limit        *int32     `json:"limit,omitempty"`
	Name         *UsageName `json:"name,omitempty"`
}

// UsageListResult is the response from the List Usages operation.
type UsageListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Usage `json:"value,omitempty"`
}

// UsageName is the usage names that can be used; currently limited to
// StorageAccount.
type UsageName struct {
	Value          *string `json:"value,omitempty"`
	LocalizedValue *string `json:"localizedValue,omitempty"`
}
