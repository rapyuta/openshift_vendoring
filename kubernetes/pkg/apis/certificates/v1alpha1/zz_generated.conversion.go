// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	certificates "github.com/openshift/kubernetes/pkg/apis/certificates"
	conversion "github.com/openshift/kubernetes/pkg/conversion"
	runtime "github.com/openshift/kubernetes/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_CertificateSigningRequest_To_certificates_CertificateSigningRequest,
		Convert_certificates_CertificateSigningRequest_To_v1alpha1_CertificateSigningRequest,
		Convert_v1alpha1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition,
		Convert_certificates_CertificateSigningRequestCondition_To_v1alpha1_CertificateSigningRequestCondition,
		Convert_v1alpha1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList,
		Convert_certificates_CertificateSigningRequestList_To_v1alpha1_CertificateSigningRequestList,
		Convert_v1alpha1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec,
		Convert_certificates_CertificateSigningRequestSpec_To_v1alpha1_CertificateSigningRequestSpec,
		Convert_v1alpha1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus,
		Convert_certificates_CertificateSigningRequestStatus_To_v1alpha1_CertificateSigningRequestStatus,
	)
}

func autoConvert_v1alpha1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(in *CertificateSigningRequest, out *certificates.CertificateSigningRequest, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v1alpha1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(in *CertificateSigningRequest, out *certificates.CertificateSigningRequest, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequest_To_v1alpha1_CertificateSigningRequest(in *certificates.CertificateSigningRequest, out *CertificateSigningRequest, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_certificates_CertificateSigningRequestSpec_To_v1alpha1_CertificateSigningRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_certificates_CertificateSigningRequestStatus_To_v1alpha1_CertificateSigningRequestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_certificates_CertificateSigningRequest_To_v1alpha1_CertificateSigningRequest(in *certificates.CertificateSigningRequest, out *CertificateSigningRequest, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequest_To_v1alpha1_CertificateSigningRequest(in, out, s)
}

func autoConvert_v1alpha1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(in *CertificateSigningRequestCondition, out *certificates.CertificateSigningRequestCondition, s conversion.Scope) error {
	out.Type = certificates.RequestConditionType(in.Type)
	out.Reason = in.Reason
	out.Message = in.Message
	out.LastUpdateTime = in.LastUpdateTime
	return nil
}

func Convert_v1alpha1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(in *CertificateSigningRequestCondition, out *certificates.CertificateSigningRequestCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequestCondition_To_v1alpha1_CertificateSigningRequestCondition(in *certificates.CertificateSigningRequestCondition, out *CertificateSigningRequestCondition, s conversion.Scope) error {
	out.Type = RequestConditionType(in.Type)
	out.Reason = in.Reason
	out.Message = in.Message
	out.LastUpdateTime = in.LastUpdateTime
	return nil
}

func Convert_certificates_CertificateSigningRequestCondition_To_v1alpha1_CertificateSigningRequestCondition(in *certificates.CertificateSigningRequestCondition, out *CertificateSigningRequestCondition, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequestCondition_To_v1alpha1_CertificateSigningRequestCondition(in, out, s)
}

func autoConvert_v1alpha1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(in *CertificateSigningRequestList, out *certificates.CertificateSigningRequestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]certificates.CertificateSigningRequest)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1alpha1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(in *CertificateSigningRequestList, out *certificates.CertificateSigningRequestList, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequestList_To_v1alpha1_CertificateSigningRequestList(in *certificates.CertificateSigningRequestList, out *CertificateSigningRequestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]CertificateSigningRequest)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_certificates_CertificateSigningRequestList_To_v1alpha1_CertificateSigningRequestList(in *certificates.CertificateSigningRequestList, out *CertificateSigningRequestList, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequestList_To_v1alpha1_CertificateSigningRequestList(in, out, s)
}

func autoConvert_v1alpha1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(in *CertificateSigningRequestSpec, out *certificates.CertificateSigningRequestSpec, s conversion.Scope) error {
	out.Request = *(*[]byte)(unsafe.Pointer(&in.Request))
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

func Convert_v1alpha1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(in *CertificateSigningRequestSpec, out *certificates.CertificateSigningRequestSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequestSpec_To_v1alpha1_CertificateSigningRequestSpec(in *certificates.CertificateSigningRequestSpec, out *CertificateSigningRequestSpec, s conversion.Scope) error {
	out.Request = *(*[]byte)(unsafe.Pointer(&in.Request))
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

func Convert_certificates_CertificateSigningRequestSpec_To_v1alpha1_CertificateSigningRequestSpec(in *certificates.CertificateSigningRequestSpec, out *CertificateSigningRequestSpec, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequestSpec_To_v1alpha1_CertificateSigningRequestSpec(in, out, s)
}

func autoConvert_v1alpha1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(in *CertificateSigningRequestStatus, out *certificates.CertificateSigningRequestStatus, s conversion.Scope) error {
	out.Conditions = *(*[]certificates.CertificateSigningRequestCondition)(unsafe.Pointer(&in.Conditions))
	out.Certificate = *(*[]byte)(unsafe.Pointer(&in.Certificate))
	return nil
}

func Convert_v1alpha1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(in *CertificateSigningRequestStatus, out *certificates.CertificateSigningRequestStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequestStatus_To_v1alpha1_CertificateSigningRequestStatus(in *certificates.CertificateSigningRequestStatus, out *CertificateSigningRequestStatus, s conversion.Scope) error {
	out.Conditions = *(*[]CertificateSigningRequestCondition)(unsafe.Pointer(&in.Conditions))
	out.Certificate = *(*[]byte)(unsafe.Pointer(&in.Certificate))
	return nil
}

func Convert_certificates_CertificateSigningRequestStatus_To_v1alpha1_CertificateSigningRequestStatus(in *certificates.CertificateSigningRequestStatus, out *CertificateSigningRequestStatus, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequestStatus_To_v1alpha1_CertificateSigningRequestStatus(in, out, s)
}
