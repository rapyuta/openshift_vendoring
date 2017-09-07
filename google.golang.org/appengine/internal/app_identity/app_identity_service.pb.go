// Code generated by protoc-gen-go.
// source: google.golang.org/appengine/internal/app_identity/app_identity_service.proto
// DO NOT EDIT!

/*
Package app_identity is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/appengine/internal/app_identity/app_identity_service.proto

It has these top-level messages:
	AppIdentityServiceError
	SignForAppRequest
	SignForAppResponse
	GetPublicCertificateForAppRequest
	PublicCertificate
	GetPublicCertificateForAppResponse
	GetServiceAccountNameRequest
	GetServiceAccountNameResponse
	GetAccessTokenRequest
	GetAccessTokenResponse
	GetDefaultGcsBucketNameRequest
	GetDefaultGcsBucketNameResponse
*/
package app_identity

import proto "github.com/openshift/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AppIdentityServiceError_ErrorCode int32

const (
	AppIdentityServiceError_SUCCESS           AppIdentityServiceError_ErrorCode = 0
	AppIdentityServiceError_UNKNOWN_SCOPE     AppIdentityServiceError_ErrorCode = 9
	AppIdentityServiceError_BLOB_TOO_LARGE    AppIdentityServiceError_ErrorCode = 1000
	AppIdentityServiceError_DEADLINE_EXCEEDED AppIdentityServiceError_ErrorCode = 1001
	AppIdentityServiceError_NOT_A_VALID_APP   AppIdentityServiceError_ErrorCode = 1002
	AppIdentityServiceError_UNKNOWN_ERROR     AppIdentityServiceError_ErrorCode = 1003
	AppIdentityServiceError_NOT_ALLOWED       AppIdentityServiceError_ErrorCode = 1005
	AppIdentityServiceError_NOT_IMPLEMENTED   AppIdentityServiceError_ErrorCode = 1006
)

var AppIdentityServiceError_ErrorCode_name = map[int32]string{
	0:    "SUCCESS",
	9:    "UNKNOWN_SCOPE",
	1000: "BLOB_TOO_LARGE",
	1001: "DEADLINE_EXCEEDED",
	1002: "NOT_A_VALID_APP",
	1003: "UNKNOWN_ERROR",
	1005: "NOT_ALLOWED",
	1006: "NOT_IMPLEMENTED",
}
var AppIdentityServiceError_ErrorCode_value = map[string]int32{
	"SUCCESS":           0,
	"UNKNOWN_SCOPE":     9,
	"BLOB_TOO_LARGE":    1000,
	"DEADLINE_EXCEEDED": 1001,
	"NOT_A_VALID_APP":   1002,
	"UNKNOWN_ERROR":     1003,
	"NOT_ALLOWED":       1005,
	"NOT_IMPLEMENTED":   1006,
}

func (x AppIdentityServiceError_ErrorCode) Enum() *AppIdentityServiceError_ErrorCode {
	p := new(AppIdentityServiceError_ErrorCode)
	*p = x
	return p
}
func (x AppIdentityServiceError_ErrorCode) String() string {
	return proto.EnumName(AppIdentityServiceError_ErrorCode_name, int32(x))
}
func (x *AppIdentityServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AppIdentityServiceError_ErrorCode_value, data, "AppIdentityServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = AppIdentityServiceError_ErrorCode(value)
	return nil
}

type AppIdentityServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *AppIdentityServiceError) Reset()         { *m = AppIdentityServiceError{} }
func (m *AppIdentityServiceError) String() string { return proto.CompactTextString(m) }
func (*AppIdentityServiceError) ProtoMessage()    {}

type SignForAppRequest struct {
	BytesToSign      []byte `protobuf:"bytes,1,opt,name=bytes_to_sign" json:"bytes_to_sign,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SignForAppRequest) Reset()         { *m = SignForAppRequest{} }
func (m *SignForAppRequest) String() string { return proto.CompactTextString(m) }
func (*SignForAppRequest) ProtoMessage()    {}

func (m *SignForAppRequest) GetBytesToSign() []byte {
	if m != nil {
		return m.BytesToSign
	}
	return nil
}

type SignForAppResponse struct {
	KeyName          *string `protobuf:"bytes,1,opt,name=key_name" json:"key_name,omitempty"`
	SignatureBytes   []byte  `protobuf:"bytes,2,opt,name=signature_bytes" json:"signature_bytes,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SignForAppResponse) Reset()         { *m = SignForAppResponse{} }
func (m *SignForAppResponse) String() string { return proto.CompactTextString(m) }
func (*SignForAppResponse) ProtoMessage()    {}

func (m *SignForAppResponse) GetKeyName() string {
	if m != nil && m.KeyName != nil {
		return *m.KeyName
	}
	return ""
}

func (m *SignForAppResponse) GetSignatureBytes() []byte {
	if m != nil {
		return m.SignatureBytes
	}
	return nil
}

type GetPublicCertificateForAppRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetPublicCertificateForAppRequest) Reset()         { *m = GetPublicCertificateForAppRequest{} }
func (m *GetPublicCertificateForAppRequest) String() string { return proto.CompactTextString(m) }
func (*GetPublicCertificateForAppRequest) ProtoMessage()    {}

type PublicCertificate struct {
	KeyName            *string `protobuf:"bytes,1,opt,name=key_name" json:"key_name,omitempty"`
	X509CertificatePem *string `protobuf:"bytes,2,opt,name=x509_certificate_pem" json:"x509_certificate_pem,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *PublicCertificate) Reset()         { *m = PublicCertificate{} }
func (m *PublicCertificate) String() string { return proto.CompactTextString(m) }
func (*PublicCertificate) ProtoMessage()    {}

func (m *PublicCertificate) GetKeyName() string {
	if m != nil && m.KeyName != nil {
		return *m.KeyName
	}
	return ""
}

func (m *PublicCertificate) GetX509CertificatePem() string {
	if m != nil && m.X509CertificatePem != nil {
		return *m.X509CertificatePem
	}
	return ""
}

type GetPublicCertificateForAppResponse struct {
	PublicCertificateList      []*PublicCertificate `protobuf:"bytes,1,rep,name=public_certificate_list" json:"public_certificate_list,omitempty"`
	MaxClientCacheTimeInSecond *int64               `protobuf:"varint,2,opt,name=max_client_cache_time_in_second" json:"max_client_cache_time_in_second,omitempty"`
	XXX_unrecognized           []byte               `json:"-"`
}

func (m *GetPublicCertificateForAppResponse) Reset()         { *m = GetPublicCertificateForAppResponse{} }
func (m *GetPublicCertificateForAppResponse) String() string { return proto.CompactTextString(m) }
func (*GetPublicCertificateForAppResponse) ProtoMessage()    {}

func (m *GetPublicCertificateForAppResponse) GetPublicCertificateList() []*PublicCertificate {
	if m != nil {
		return m.PublicCertificateList
	}
	return nil
}

func (m *GetPublicCertificateForAppResponse) GetMaxClientCacheTimeInSecond() int64 {
	if m != nil && m.MaxClientCacheTimeInSecond != nil {
		return *m.MaxClientCacheTimeInSecond
	}
	return 0
}

type GetServiceAccountNameRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetServiceAccountNameRequest) Reset()         { *m = GetServiceAccountNameRequest{} }
func (m *GetServiceAccountNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetServiceAccountNameRequest) ProtoMessage()    {}

type GetServiceAccountNameResponse struct {
	ServiceAccountName *string `protobuf:"bytes,1,opt,name=service_account_name" json:"service_account_name,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *GetServiceAccountNameResponse) Reset()         { *m = GetServiceAccountNameResponse{} }
func (m *GetServiceAccountNameResponse) String() string { return proto.CompactTextString(m) }
func (*GetServiceAccountNameResponse) ProtoMessage()    {}

func (m *GetServiceAccountNameResponse) GetServiceAccountName() string {
	if m != nil && m.ServiceAccountName != nil {
		return *m.ServiceAccountName
	}
	return ""
}

type GetAccessTokenRequest struct {
	Scope              []string `protobuf:"bytes,1,rep,name=scope" json:"scope,omitempty"`
	ServiceAccountId   *int64   `protobuf:"varint,2,opt,name=service_account_id" json:"service_account_id,omitempty"`
	ServiceAccountName *string  `protobuf:"bytes,3,opt,name=service_account_name" json:"service_account_name,omitempty"`
	XXX_unrecognized   []byte   `json:"-"`
}

func (m *GetAccessTokenRequest) Reset()         { *m = GetAccessTokenRequest{} }
func (m *GetAccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccessTokenRequest) ProtoMessage()    {}

func (m *GetAccessTokenRequest) GetScope() []string {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *GetAccessTokenRequest) GetServiceAccountId() int64 {
	if m != nil && m.ServiceAccountId != nil {
		return *m.ServiceAccountId
	}
	return 0
}

func (m *GetAccessTokenRequest) GetServiceAccountName() string {
	if m != nil && m.ServiceAccountName != nil {
		return *m.ServiceAccountName
	}
	return ""
}

type GetAccessTokenResponse struct {
	AccessToken      *string `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
	ExpirationTime   *int64  `protobuf:"varint,2,opt,name=expiration_time" json:"expiration_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetAccessTokenResponse) Reset()         { *m = GetAccessTokenResponse{} }
func (m *GetAccessTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccessTokenResponse) ProtoMessage()    {}

func (m *GetAccessTokenResponse) GetAccessToken() string {
	if m != nil && m.AccessToken != nil {
		return *m.AccessToken
	}
	return ""
}

func (m *GetAccessTokenResponse) GetExpirationTime() int64 {
	if m != nil && m.ExpirationTime != nil {
		return *m.ExpirationTime
	}
	return 0
}

type GetDefaultGcsBucketNameRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetDefaultGcsBucketNameRequest) Reset()         { *m = GetDefaultGcsBucketNameRequest{} }
func (m *GetDefaultGcsBucketNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetDefaultGcsBucketNameRequest) ProtoMessage()    {}

type GetDefaultGcsBucketNameResponse struct {
	DefaultGcsBucketName *string `protobuf:"bytes,1,opt,name=default_gcs_bucket_name" json:"default_gcs_bucket_name,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *GetDefaultGcsBucketNameResponse) Reset()         { *m = GetDefaultGcsBucketNameResponse{} }
func (m *GetDefaultGcsBucketNameResponse) String() string { return proto.CompactTextString(m) }
func (*GetDefaultGcsBucketNameResponse) ProtoMessage()    {}

func (m *GetDefaultGcsBucketNameResponse) GetDefaultGcsBucketName() string {
	if m != nil && m.DefaultGcsBucketName != nil {
		return *m.DefaultGcsBucketName
	}
	return ""
}

func init() {
}
