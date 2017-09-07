// +build !ignore_autogenerated_openshift

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	api "github.com/openshift/origin/pkg/project/api"
	pkg_api "github.com/openshift/kubernetes/pkg/api"
	api_v1 "github.com/openshift/kubernetes/pkg/api/v1"
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
		Convert_v1_Project_To_api_Project,
		Convert_api_Project_To_v1_Project,
		Convert_v1_ProjectList_To_api_ProjectList,
		Convert_api_ProjectList_To_v1_ProjectList,
		Convert_v1_ProjectRequest_To_api_ProjectRequest,
		Convert_api_ProjectRequest_To_v1_ProjectRequest,
		Convert_v1_ProjectSpec_To_api_ProjectSpec,
		Convert_api_ProjectSpec_To_v1_ProjectSpec,
		Convert_v1_ProjectStatus_To_api_ProjectStatus,
		Convert_api_ProjectStatus_To_v1_ProjectStatus,
	)
}

func autoConvert_v1_Project_To_api_Project(in *Project, out *api.Project, s conversion.Scope) error {
	if err := api_v1.Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_ProjectSpec_To_api_ProjectSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ProjectStatus_To_api_ProjectStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_Project_To_api_Project(in *Project, out *api.Project, s conversion.Scope) error {
	return autoConvert_v1_Project_To_api_Project(in, out, s)
}

func autoConvert_api_Project_To_v1_Project(in *api.Project, out *Project, s conversion.Scope) error {
	if err := api_v1.Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_api_ProjectSpec_To_v1_ProjectSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_api_ProjectStatus_To_v1_ProjectStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_Project_To_v1_Project(in *api.Project, out *Project, s conversion.Scope) error {
	return autoConvert_api_Project_To_v1_Project(in, out, s)
}

func autoConvert_v1_ProjectList_To_api_ProjectList(in *ProjectList, out *api.ProjectList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.Project, len(*in))
		for i := range *in {
			if err := Convert_v1_Project_To_api_Project(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_ProjectList_To_api_ProjectList(in *ProjectList, out *api.ProjectList, s conversion.Scope) error {
	return autoConvert_v1_ProjectList_To_api_ProjectList(in, out, s)
}

func autoConvert_api_ProjectList_To_v1_ProjectList(in *api.ProjectList, out *ProjectList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			if err := Convert_api_Project_To_v1_Project(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_ProjectList_To_v1_ProjectList(in *api.ProjectList, out *ProjectList, s conversion.Scope) error {
	return autoConvert_api_ProjectList_To_v1_ProjectList(in, out, s)
}

func autoConvert_v1_ProjectRequest_To_api_ProjectRequest(in *ProjectRequest, out *api.ProjectRequest, s conversion.Scope) error {
	if err := api_v1.Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	return nil
}

func Convert_v1_ProjectRequest_To_api_ProjectRequest(in *ProjectRequest, out *api.ProjectRequest, s conversion.Scope) error {
	return autoConvert_v1_ProjectRequest_To_api_ProjectRequest(in, out, s)
}

func autoConvert_api_ProjectRequest_To_v1_ProjectRequest(in *api.ProjectRequest, out *ProjectRequest, s conversion.Scope) error {
	if err := api_v1.Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	return nil
}

func Convert_api_ProjectRequest_To_v1_ProjectRequest(in *api.ProjectRequest, out *ProjectRequest, s conversion.Scope) error {
	return autoConvert_api_ProjectRequest_To_v1_ProjectRequest(in, out, s)
}

func autoConvert_v1_ProjectSpec_To_api_ProjectSpec(in *ProjectSpec, out *api.ProjectSpec, s conversion.Scope) error {
	out.Finalizers = *(*[]pkg_api.FinalizerName)(unsafe.Pointer(&in.Finalizers))
	return nil
}

func Convert_v1_ProjectSpec_To_api_ProjectSpec(in *ProjectSpec, out *api.ProjectSpec, s conversion.Scope) error {
	return autoConvert_v1_ProjectSpec_To_api_ProjectSpec(in, out, s)
}

func autoConvert_api_ProjectSpec_To_v1_ProjectSpec(in *api.ProjectSpec, out *ProjectSpec, s conversion.Scope) error {
	out.Finalizers = *(*[]api_v1.FinalizerName)(unsafe.Pointer(&in.Finalizers))
	return nil
}

func Convert_api_ProjectSpec_To_v1_ProjectSpec(in *api.ProjectSpec, out *ProjectSpec, s conversion.Scope) error {
	return autoConvert_api_ProjectSpec_To_v1_ProjectSpec(in, out, s)
}

func autoConvert_v1_ProjectStatus_To_api_ProjectStatus(in *ProjectStatus, out *api.ProjectStatus, s conversion.Scope) error {
	out.Phase = pkg_api.NamespacePhase(in.Phase)
	return nil
}

func Convert_v1_ProjectStatus_To_api_ProjectStatus(in *ProjectStatus, out *api.ProjectStatus, s conversion.Scope) error {
	return autoConvert_v1_ProjectStatus_To_api_ProjectStatus(in, out, s)
}

func autoConvert_api_ProjectStatus_To_v1_ProjectStatus(in *api.ProjectStatus, out *ProjectStatus, s conversion.Scope) error {
	out.Phase = api_v1.NamespacePhase(in.Phase)
	return nil
}

func Convert_api_ProjectStatus_To_v1_ProjectStatus(in *api.ProjectStatus, out *ProjectStatus, s conversion.Scope) error {
	return autoConvert_api_ProjectStatus_To_v1_ProjectStatus(in, out, s)
}
