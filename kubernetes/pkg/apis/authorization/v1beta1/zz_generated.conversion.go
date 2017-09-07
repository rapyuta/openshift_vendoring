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

package v1beta1

import (
	authorization "github.com/openshift/kubernetes/pkg/apis/authorization"
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
		Convert_v1beta1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview,
		Convert_authorization_LocalSubjectAccessReview_To_v1beta1_LocalSubjectAccessReview,
		Convert_v1beta1_NonResourceAttributes_To_authorization_NonResourceAttributes,
		Convert_authorization_NonResourceAttributes_To_v1beta1_NonResourceAttributes,
		Convert_v1beta1_ResourceAttributes_To_authorization_ResourceAttributes,
		Convert_authorization_ResourceAttributes_To_v1beta1_ResourceAttributes,
		Convert_v1beta1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview,
		Convert_authorization_SelfSubjectAccessReview_To_v1beta1_SelfSubjectAccessReview,
		Convert_v1beta1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec,
		Convert_authorization_SelfSubjectAccessReviewSpec_To_v1beta1_SelfSubjectAccessReviewSpec,
		Convert_v1beta1_SubjectAccessReview_To_authorization_SubjectAccessReview,
		Convert_authorization_SubjectAccessReview_To_v1beta1_SubjectAccessReview,
		Convert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec,
		Convert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec,
		Convert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus,
		Convert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus,
	)
}

func autoConvert_v1beta1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in *LocalSubjectAccessReview, out *authorization.LocalSubjectAccessReview, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1beta1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in *LocalSubjectAccessReview, out *authorization.LocalSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_v1beta1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in, out, s)
}

func autoConvert_authorization_LocalSubjectAccessReview_To_v1beta1_LocalSubjectAccessReview(in *authorization.LocalSubjectAccessReview, out *LocalSubjectAccessReview, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_authorization_LocalSubjectAccessReview_To_v1beta1_LocalSubjectAccessReview(in *authorization.LocalSubjectAccessReview, out *LocalSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_authorization_LocalSubjectAccessReview_To_v1beta1_LocalSubjectAccessReview(in, out, s)
}

func autoConvert_v1beta1_NonResourceAttributes_To_authorization_NonResourceAttributes(in *NonResourceAttributes, out *authorization.NonResourceAttributes, s conversion.Scope) error {
	out.Path = in.Path
	out.Verb = in.Verb
	return nil
}

func Convert_v1beta1_NonResourceAttributes_To_authorization_NonResourceAttributes(in *NonResourceAttributes, out *authorization.NonResourceAttributes, s conversion.Scope) error {
	return autoConvert_v1beta1_NonResourceAttributes_To_authorization_NonResourceAttributes(in, out, s)
}

func autoConvert_authorization_NonResourceAttributes_To_v1beta1_NonResourceAttributes(in *authorization.NonResourceAttributes, out *NonResourceAttributes, s conversion.Scope) error {
	out.Path = in.Path
	out.Verb = in.Verb
	return nil
}

func Convert_authorization_NonResourceAttributes_To_v1beta1_NonResourceAttributes(in *authorization.NonResourceAttributes, out *NonResourceAttributes, s conversion.Scope) error {
	return autoConvert_authorization_NonResourceAttributes_To_v1beta1_NonResourceAttributes(in, out, s)
}

func autoConvert_v1beta1_ResourceAttributes_To_authorization_ResourceAttributes(in *ResourceAttributes, out *authorization.ResourceAttributes, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Verb = in.Verb
	out.Group = in.Group
	out.Version = in.Version
	out.Resource = in.Resource
	out.Subresource = in.Subresource
	out.Name = in.Name
	return nil
}

func Convert_v1beta1_ResourceAttributes_To_authorization_ResourceAttributes(in *ResourceAttributes, out *authorization.ResourceAttributes, s conversion.Scope) error {
	return autoConvert_v1beta1_ResourceAttributes_To_authorization_ResourceAttributes(in, out, s)
}

func autoConvert_authorization_ResourceAttributes_To_v1beta1_ResourceAttributes(in *authorization.ResourceAttributes, out *ResourceAttributes, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Verb = in.Verb
	out.Group = in.Group
	out.Version = in.Version
	out.Resource = in.Resource
	out.Subresource = in.Subresource
	out.Name = in.Name
	return nil
}

func Convert_authorization_ResourceAttributes_To_v1beta1_ResourceAttributes(in *authorization.ResourceAttributes, out *ResourceAttributes, s conversion.Scope) error {
	return autoConvert_authorization_ResourceAttributes_To_v1beta1_ResourceAttributes(in, out, s)
}

func autoConvert_v1beta1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(in *SelfSubjectAccessReview, out *authorization.SelfSubjectAccessReview, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v1beta1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1beta1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(in *SelfSubjectAccessReview, out *authorization.SelfSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_v1beta1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(in, out, s)
}

func autoConvert_authorization_SelfSubjectAccessReview_To_v1beta1_SelfSubjectAccessReview(in *authorization.SelfSubjectAccessReview, out *SelfSubjectAccessReview, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_authorization_SelfSubjectAccessReviewSpec_To_v1beta1_SelfSubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_authorization_SelfSubjectAccessReview_To_v1beta1_SelfSubjectAccessReview(in *authorization.SelfSubjectAccessReview, out *SelfSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_authorization_SelfSubjectAccessReview_To_v1beta1_SelfSubjectAccessReview(in, out, s)
}

func autoConvert_v1beta1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(in *SelfSubjectAccessReviewSpec, out *authorization.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*authorization.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*authorization.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	return nil
}

func Convert_v1beta1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(in *SelfSubjectAccessReviewSpec, out *authorization.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(in, out, s)
}

func autoConvert_authorization_SelfSubjectAccessReviewSpec_To_v1beta1_SelfSubjectAccessReviewSpec(in *authorization.SelfSubjectAccessReviewSpec, out *SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	return nil
}

func Convert_authorization_SelfSubjectAccessReviewSpec_To_v1beta1_SelfSubjectAccessReviewSpec(in *authorization.SelfSubjectAccessReviewSpec, out *SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_authorization_SelfSubjectAccessReviewSpec_To_v1beta1_SelfSubjectAccessReviewSpec(in, out, s)
}

func autoConvert_v1beta1_SubjectAccessReview_To_authorization_SubjectAccessReview(in *SubjectAccessReview, out *authorization.SubjectAccessReview, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1beta1_SubjectAccessReview_To_authorization_SubjectAccessReview(in *SubjectAccessReview, out *authorization.SubjectAccessReview, s conversion.Scope) error {
	return autoConvert_v1beta1_SubjectAccessReview_To_authorization_SubjectAccessReview(in, out, s)
}

func autoConvert_authorization_SubjectAccessReview_To_v1beta1_SubjectAccessReview(in *authorization.SubjectAccessReview, out *SubjectAccessReview, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_authorization_SubjectAccessReview_To_v1beta1_SubjectAccessReview(in *authorization.SubjectAccessReview, out *SubjectAccessReview, s conversion.Scope) error {
	return autoConvert_authorization_SubjectAccessReview_To_v1beta1_SubjectAccessReview(in, out, s)
}

func autoConvert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(in *SubjectAccessReviewSpec, out *authorization.SubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*authorization.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*authorization.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]authorization.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

func Convert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(in *SubjectAccessReviewSpec, out *authorization.SubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(in, out, s)
}

func autoConvert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec(in *authorization.SubjectAccessReviewSpec, out *SubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

func Convert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec(in *authorization.SubjectAccessReviewSpec, out *SubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec(in, out, s)
}

func autoConvert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(in *SubjectAccessReviewStatus, out *authorization.SubjectAccessReviewStatus, s conversion.Scope) error {
	out.Allowed = in.Allowed
	out.Reason = in.Reason
	out.EvaluationError = in.EvaluationError
	return nil
}

func Convert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(in *SubjectAccessReviewStatus, out *authorization.SubjectAccessReviewStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(in, out, s)
}

func autoConvert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(in *authorization.SubjectAccessReviewStatus, out *SubjectAccessReviewStatus, s conversion.Scope) error {
	out.Allowed = in.Allowed
	out.Reason = in.Reason
	out.EvaluationError = in.EvaluationError
	return nil
}

func Convert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(in *authorization.SubjectAccessReviewStatus, out *SubjectAccessReviewStatus, s conversion.Scope) error {
	return autoConvert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(in, out, s)
}
