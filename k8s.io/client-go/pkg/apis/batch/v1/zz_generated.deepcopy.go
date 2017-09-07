// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	unversioned "github.com/openshift/k8s.io/client-go/pkg/api/unversioned"
	api_v1 "github.com/openshift/k8s.io/client-go/pkg/api/v1"
	conversion "github.com/openshift/k8s.io/client-go/pkg/conversion"
	runtime "github.com/openshift/k8s.io/client-go/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Job, InType: reflect.TypeOf(&Job{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_JobCondition, InType: reflect.TypeOf(&JobCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_JobList, InType: reflect.TypeOf(&JobList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_JobSpec, InType: reflect.TypeOf(&JobSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_JobStatus, InType: reflect.TypeOf(&JobStatus{})},
	)
}

func DeepCopy_v1_Job(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Job)
		out := out.(*Job)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_JobSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_JobStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_JobCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobCondition)
		out := out.(*JobCondition)
		out.Type = in.Type
		out.Status = in.Status
		out.LastProbeTime = in.LastProbeTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		out.Reason = in.Reason
		out.Message = in.Message
		return nil
	}
}

func DeepCopy_v1_JobList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobList)
		out := out.(*JobList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Job, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_Job(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1_JobSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobSpec)
		out := out.(*JobSpec)
		if in.Parallelism != nil {
			in, out := &in.Parallelism, &out.Parallelism
			*out = new(int32)
			**out = **in
		} else {
			out.Parallelism = nil
		}
		if in.Completions != nil {
			in, out := &in.Completions, &out.Completions
			*out = new(int32)
			**out = **in
		} else {
			out.Completions = nil
		}
		if in.ActiveDeadlineSeconds != nil {
			in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
			*out = new(int64)
			**out = **in
		} else {
			out.ActiveDeadlineSeconds = nil
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = new(unversioned.LabelSelector)
			if err := unversioned.DeepCopy_unversioned_LabelSelector(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Selector = nil
		}
		if in.ManualSelector != nil {
			in, out := &in.ManualSelector, &out.ManualSelector
			*out = new(bool)
			**out = **in
		} else {
			out.ManualSelector = nil
		}
		if err := api_v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_JobStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobStatus)
		out := out.(*JobStatus)
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]JobCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_JobCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Conditions = nil
		}
		if in.StartTime != nil {
			in, out := &in.StartTime, &out.StartTime
			*out = new(unversioned.Time)
			**out = (*in).DeepCopy()
		} else {
			out.StartTime = nil
		}
		if in.CompletionTime != nil {
			in, out := &in.CompletionTime, &out.CompletionTime
			*out = new(unversioned.Time)
			**out = (*in).DeepCopy()
		} else {
			out.CompletionTime = nil
		}
		out.Active = in.Active
		out.Succeeded = in.Succeeded
		out.Failed = in.Failed
		return nil
	}
}
