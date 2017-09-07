// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package api

import (
	pkg_api "github.com/openshift/kubernetes/pkg/api"
	unversioned "github.com/openshift/kubernetes/pkg/api/unversioned"
	conversion "github.com/openshift/kubernetes/pkg/conversion"
	runtime "github.com/openshift/kubernetes/pkg/runtime"
	sets "github.com/openshift/kubernetes/pkg/util/sets"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_Action, InType: reflect.TypeOf(&Action{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ClusterPolicy, InType: reflect.TypeOf(&ClusterPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ClusterPolicyBinding, InType: reflect.TypeOf(&ClusterPolicyBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ClusterPolicyBindingList, InType: reflect.TypeOf(&ClusterPolicyBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ClusterPolicyList, InType: reflect.TypeOf(&ClusterPolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ClusterRole, InType: reflect.TypeOf(&ClusterRole{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ClusterRoleBinding, InType: reflect.TypeOf(&ClusterRoleBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ClusterRoleBindingList, InType: reflect.TypeOf(&ClusterRoleBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ClusterRoleList, InType: reflect.TypeOf(&ClusterRoleList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_GroupRestriction, InType: reflect.TypeOf(&GroupRestriction{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_IsPersonalSubjectAccessReview, InType: reflect.TypeOf(&IsPersonalSubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_LocalResourceAccessReview, InType: reflect.TypeOf(&LocalResourceAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_LocalSubjectAccessReview, InType: reflect.TypeOf(&LocalSubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_Policy, InType: reflect.TypeOf(&Policy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_PolicyBinding, InType: reflect.TypeOf(&PolicyBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_PolicyBindingList, InType: reflect.TypeOf(&PolicyBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_PolicyList, InType: reflect.TypeOf(&PolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_PolicyRule, InType: reflect.TypeOf(&PolicyRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_PolicyRuleBuilder, InType: reflect.TypeOf(&PolicyRuleBuilder{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ResourceAccessReview, InType: reflect.TypeOf(&ResourceAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ResourceAccessReviewResponse, InType: reflect.TypeOf(&ResourceAccessReviewResponse{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_Role, InType: reflect.TypeOf(&Role{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RoleBinding, InType: reflect.TypeOf(&RoleBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RoleBindingList, InType: reflect.TypeOf(&RoleBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RoleBindingRestriction, InType: reflect.TypeOf(&RoleBindingRestriction{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RoleBindingRestrictionList, InType: reflect.TypeOf(&RoleBindingRestrictionList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RoleBindingRestrictionSpec, InType: reflect.TypeOf(&RoleBindingRestrictionSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RoleList, InType: reflect.TypeOf(&RoleList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_SelfSubjectRulesReview, InType: reflect.TypeOf(&SelfSubjectRulesReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_SelfSubjectRulesReviewSpec, InType: reflect.TypeOf(&SelfSubjectRulesReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ServiceAccountReference, InType: reflect.TypeOf(&ServiceAccountReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ServiceAccountRestriction, InType: reflect.TypeOf(&ServiceAccountRestriction{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_SubjectAccessReview, InType: reflect.TypeOf(&SubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_SubjectAccessReviewResponse, InType: reflect.TypeOf(&SubjectAccessReviewResponse{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_SubjectRulesReview, InType: reflect.TypeOf(&SubjectRulesReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_SubjectRulesReviewSpec, InType: reflect.TypeOf(&SubjectRulesReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_SubjectRulesReviewStatus, InType: reflect.TypeOf(&SubjectRulesReviewStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserRestriction, InType: reflect.TypeOf(&UserRestriction{})},
	)
}

func DeepCopy_api_Action(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Action)
		out := out.(*Action)
		out.Namespace = in.Namespace
		out.Verb = in.Verb
		out.Group = in.Group
		out.Version = in.Version
		out.Resource = in.Resource
		out.ResourceName = in.ResourceName
		out.Path = in.Path
		out.IsNonResourceURL = in.IsNonResourceURL
		if in.Content == nil {
			out.Content = nil
		} else if newVal, err := c.DeepCopy(&in.Content); err != nil {
			return err
		} else {
			out.Content = *newVal.(*runtime.Object)
		}
		return nil
	}
}

func DeepCopy_api_ClusterPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPolicy)
		out := out.(*ClusterPolicy)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.LastModified = in.LastModified.DeepCopy()
		if in.Roles != nil {
			in, out := &in.Roles, &out.Roles
			*out = make(ClusterRolesByName)
			for key, val := range *in {
				if newVal, err := c.DeepCopy(&val); err != nil {
					return err
				} else {
					(*out)[key] = *newVal.(**ClusterRole)
				}
			}
		} else {
			out.Roles = nil
		}
		return nil
	}
}

func DeepCopy_api_ClusterPolicyBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPolicyBinding)
		out := out.(*ClusterPolicyBinding)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.LastModified = in.LastModified.DeepCopy()
		out.PolicyRef = in.PolicyRef
		if in.RoleBindings != nil {
			in, out := &in.RoleBindings, &out.RoleBindings
			*out = make(ClusterRoleBindingsByName)
			for key, val := range *in {
				if newVal, err := c.DeepCopy(&val); err != nil {
					return err
				} else {
					(*out)[key] = *newVal.(**ClusterRoleBinding)
				}
			}
		} else {
			out.RoleBindings = nil
		}
		return nil
	}
}

func DeepCopy_api_ClusterPolicyBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPolicyBindingList)
		out := out.(*ClusterPolicyBindingList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterPolicyBinding, len(*in))
			for i := range *in {
				if err := DeepCopy_api_ClusterPolicyBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_ClusterPolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPolicyList)
		out := out.(*ClusterPolicyList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterPolicy, len(*in))
			for i := range *in {
				if err := DeepCopy_api_ClusterPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_ClusterRole(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRole)
		out := out.(*ClusterRole)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if err := DeepCopy_api_PolicyRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Rules = nil
		}
		return nil
	}
}

func DeepCopy_api_ClusterRoleBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRoleBinding)
		out := out.(*ClusterRoleBinding)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if in.Subjects != nil {
			in, out := &in.Subjects, &out.Subjects
			*out = make([]pkg_api.ObjectReference, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Subjects = nil
		}
		out.RoleRef = in.RoleRef
		return nil
	}
}

func DeepCopy_api_ClusterRoleBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRoleBindingList)
		out := out.(*ClusterRoleBindingList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterRoleBinding, len(*in))
			for i := range *in {
				if err := DeepCopy_api_ClusterRoleBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_ClusterRoleList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRoleList)
		out := out.(*ClusterRoleList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterRole, len(*in))
			for i := range *in {
				if err := DeepCopy_api_ClusterRole(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_GroupRestriction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupRestriction)
		out := out.(*GroupRestriction)
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Groups = nil
		}
		if in.Selectors != nil {
			in, out := &in.Selectors, &out.Selectors
			*out = make([]unversioned.LabelSelector, len(*in))
			for i := range *in {
				if err := unversioned.DeepCopy_unversioned_LabelSelector(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Selectors = nil
		}
		return nil
	}
}

func DeepCopy_api_IsPersonalSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IsPersonalSubjectAccessReview)
		out := out.(*IsPersonalSubjectAccessReview)
		out.TypeMeta = in.TypeMeta
		return nil
	}
}

func DeepCopy_api_LocalResourceAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LocalResourceAccessReview)
		out := out.(*LocalResourceAccessReview)
		out.TypeMeta = in.TypeMeta
		if err := DeepCopy_api_Action(&in.Action, &out.Action, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_api_LocalSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LocalSubjectAccessReview)
		out := out.(*LocalSubjectAccessReview)
		out.TypeMeta = in.TypeMeta
		if err := DeepCopy_api_Action(&in.Action, &out.Action, c); err != nil {
			return err
		}
		out.User = in.User
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Groups = nil
		}
		if in.Scopes != nil {
			in, out := &in.Scopes, &out.Scopes
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Scopes = nil
		}
		return nil
	}
}

func DeepCopy_api_Policy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Policy)
		out := out.(*Policy)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.LastModified = in.LastModified.DeepCopy()
		if in.Roles != nil {
			in, out := &in.Roles, &out.Roles
			*out = make(RolesByName)
			for key, val := range *in {
				if newVal, err := c.DeepCopy(&val); err != nil {
					return err
				} else {
					(*out)[key] = *newVal.(**Role)
				}
			}
		} else {
			out.Roles = nil
		}
		return nil
	}
}

func DeepCopy_api_PolicyBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyBinding)
		out := out.(*PolicyBinding)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.LastModified = in.LastModified.DeepCopy()
		out.PolicyRef = in.PolicyRef
		if in.RoleBindings != nil {
			in, out := &in.RoleBindings, &out.RoleBindings
			*out = make(RoleBindingsByName)
			for key, val := range *in {
				if newVal, err := c.DeepCopy(&val); err != nil {
					return err
				} else {
					(*out)[key] = *newVal.(**RoleBinding)
				}
			}
		} else {
			out.RoleBindings = nil
		}
		return nil
	}
}

func DeepCopy_api_PolicyBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyBindingList)
		out := out.(*PolicyBindingList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]PolicyBinding, len(*in))
			for i := range *in {
				if err := DeepCopy_api_PolicyBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_PolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyList)
		out := out.(*PolicyList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Policy, len(*in))
			for i := range *in {
				if err := DeepCopy_api_Policy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_PolicyRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyRule)
		out := out.(*PolicyRule)
		if in.Verbs != nil {
			in, out := &in.Verbs, &out.Verbs
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Verbs = nil
		}
		if in.AttributeRestrictions == nil {
			out.AttributeRestrictions = nil
		} else if newVal, err := c.DeepCopy(&in.AttributeRestrictions); err != nil {
			return err
		} else {
			out.AttributeRestrictions = *newVal.(*runtime.Object)
		}
		if in.APIGroups != nil {
			in, out := &in.APIGroups, &out.APIGroups
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.APIGroups = nil
		}
		if in.Resources != nil {
			in, out := &in.Resources, &out.Resources
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Resources = nil
		}
		if in.ResourceNames != nil {
			in, out := &in.ResourceNames, &out.ResourceNames
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.ResourceNames = nil
		}
		if in.NonResourceURLs != nil {
			in, out := &in.NonResourceURLs, &out.NonResourceURLs
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.NonResourceURLs = nil
		}
		return nil
	}
}

func DeepCopy_api_PolicyRuleBuilder(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyRuleBuilder)
		out := out.(*PolicyRuleBuilder)
		if err := DeepCopy_api_PolicyRule(&in.PolicyRule, &out.PolicyRule, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_api_ResourceAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResourceAccessReview)
		out := out.(*ResourceAccessReview)
		out.TypeMeta = in.TypeMeta
		if err := DeepCopy_api_Action(&in.Action, &out.Action, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_api_ResourceAccessReviewResponse(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResourceAccessReviewResponse)
		out := out.(*ResourceAccessReviewResponse)
		out.TypeMeta = in.TypeMeta
		out.Namespace = in.Namespace
		if in.Users != nil {
			in, out := &in.Users, &out.Users
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Users = nil
		}
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Groups = nil
		}
		out.EvaluationError = in.EvaluationError
		return nil
	}
}

func DeepCopy_api_Role(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Role)
		out := out.(*Role)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if err := DeepCopy_api_PolicyRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Rules = nil
		}
		return nil
	}
}

func DeepCopy_api_RoleBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBinding)
		out := out.(*RoleBinding)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if in.Subjects != nil {
			in, out := &in.Subjects, &out.Subjects
			*out = make([]pkg_api.ObjectReference, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Subjects = nil
		}
		out.RoleRef = in.RoleRef
		return nil
	}
}

func DeepCopy_api_RoleBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingList)
		out := out.(*RoleBindingList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]RoleBinding, len(*in))
			for i := range *in {
				if err := DeepCopy_api_RoleBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_RoleBindingRestriction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingRestriction)
		out := out.(*RoleBindingRestriction)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_api_RoleBindingRestrictionSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_api_RoleBindingRestrictionList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingRestrictionList)
		out := out.(*RoleBindingRestrictionList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]RoleBindingRestriction, len(*in))
			for i := range *in {
				if err := DeepCopy_api_RoleBindingRestriction(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_RoleBindingRestrictionSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingRestrictionSpec)
		out := out.(*RoleBindingRestrictionSpec)
		if in.UserRestriction != nil {
			in, out := &in.UserRestriction, &out.UserRestriction
			*out = new(UserRestriction)
			if err := DeepCopy_api_UserRestriction(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.UserRestriction = nil
		}
		if in.GroupRestriction != nil {
			in, out := &in.GroupRestriction, &out.GroupRestriction
			*out = new(GroupRestriction)
			if err := DeepCopy_api_GroupRestriction(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.GroupRestriction = nil
		}
		if in.ServiceAccountRestriction != nil {
			in, out := &in.ServiceAccountRestriction, &out.ServiceAccountRestriction
			*out = new(ServiceAccountRestriction)
			if err := DeepCopy_api_ServiceAccountRestriction(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.ServiceAccountRestriction = nil
		}
		return nil
	}
}

func DeepCopy_api_RoleList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleList)
		out := out.(*RoleList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Role, len(*in))
			for i := range *in {
				if err := DeepCopy_api_Role(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_SelfSubjectRulesReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SelfSubjectRulesReview)
		out := out.(*SelfSubjectRulesReview)
		out.TypeMeta = in.TypeMeta
		if err := DeepCopy_api_SelfSubjectRulesReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_api_SubjectRulesReviewStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_api_SelfSubjectRulesReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SelfSubjectRulesReviewSpec)
		out := out.(*SelfSubjectRulesReviewSpec)
		if in.Scopes != nil {
			in, out := &in.Scopes, &out.Scopes
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Scopes = nil
		}
		return nil
	}
}

func DeepCopy_api_ServiceAccountReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceAccountReference)
		out := out.(*ServiceAccountReference)
		out.Name = in.Name
		out.Namespace = in.Namespace
		return nil
	}
}

func DeepCopy_api_ServiceAccountRestriction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceAccountRestriction)
		out := out.(*ServiceAccountRestriction)
		if in.ServiceAccounts != nil {
			in, out := &in.ServiceAccounts, &out.ServiceAccounts
			*out = make([]ServiceAccountReference, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.ServiceAccounts = nil
		}
		if in.Namespaces != nil {
			in, out := &in.Namespaces, &out.Namespaces
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Namespaces = nil
		}
		return nil
	}
}

func DeepCopy_api_SubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReview)
		out := out.(*SubjectAccessReview)
		out.TypeMeta = in.TypeMeta
		if err := DeepCopy_api_Action(&in.Action, &out.Action, c); err != nil {
			return err
		}
		out.User = in.User
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Groups = nil
		}
		if in.Scopes != nil {
			in, out := &in.Scopes, &out.Scopes
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Scopes = nil
		}
		return nil
	}
}

func DeepCopy_api_SubjectAccessReviewResponse(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReviewResponse)
		out := out.(*SubjectAccessReviewResponse)
		out.TypeMeta = in.TypeMeta
		out.Namespace = in.Namespace
		out.Allowed = in.Allowed
		out.Reason = in.Reason
		out.EvaluationError = in.EvaluationError
		return nil
	}
}

func DeepCopy_api_SubjectRulesReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectRulesReview)
		out := out.(*SubjectRulesReview)
		out.TypeMeta = in.TypeMeta
		if err := DeepCopy_api_SubjectRulesReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_api_SubjectRulesReviewStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_api_SubjectRulesReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectRulesReviewSpec)
		out := out.(*SubjectRulesReviewSpec)
		out.User = in.User
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Groups = nil
		}
		if in.Scopes != nil {
			in, out := &in.Scopes, &out.Scopes
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Scopes = nil
		}
		return nil
	}
}

func DeepCopy_api_SubjectRulesReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectRulesReviewStatus)
		out := out.(*SubjectRulesReviewStatus)
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if err := DeepCopy_api_PolicyRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Rules = nil
		}
		out.EvaluationError = in.EvaluationError
		return nil
	}
}

func DeepCopy_api_UserRestriction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserRestriction)
		out := out.(*UserRestriction)
		if in.Users != nil {
			in, out := &in.Users, &out.Users
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Users = nil
		}
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Groups = nil
		}
		if in.Selectors != nil {
			in, out := &in.Selectors, &out.Selectors
			*out = make([]unversioned.LabelSelector, len(*in))
			for i := range *in {
				if err := unversioned.DeepCopy_unversioned_LabelSelector(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Selectors = nil
		}
		return nil
	}
}
