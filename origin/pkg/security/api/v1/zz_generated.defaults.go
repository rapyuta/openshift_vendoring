// +build !ignore_autogenerated_openshift

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1

import (
	api_v1 "github.com/openshift/kubernetes/pkg/api/v1"
	runtime "github.com/openshift/kubernetes/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&PodSecurityPolicyReview{}, func(obj interface{}) { SetObjectDefaults_PodSecurityPolicyReview(obj.(*PodSecurityPolicyReview)) })
	scheme.AddTypeDefaultingFunc(&PodSecurityPolicySelfSubjectReview{}, func(obj interface{}) {
		SetObjectDefaults_PodSecurityPolicySelfSubjectReview(obj.(*PodSecurityPolicySelfSubjectReview))
	})
	scheme.AddTypeDefaultingFunc(&PodSecurityPolicySubjectReview{}, func(obj interface{}) {
		SetObjectDefaults_PodSecurityPolicySubjectReview(obj.(*PodSecurityPolicySubjectReview))
	})
	return nil
}

func SetObjectDefaults_PodSecurityPolicyReview(in *PodSecurityPolicyReview) {
	api_v1.SetDefaults_PodSpec(&in.Spec.Template.Spec)
	for i := range in.Spec.Template.Spec.Volumes {
		a := &in.Spec.Template.Spec.Volumes[i]
		api_v1.SetDefaults_Volume(a)
		if a.VolumeSource.Secret != nil {
			api_v1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			api_v1.SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
		}
		if a.VolumeSource.RBD != nil {
			api_v1.SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
		}
		if a.VolumeSource.DownwardAPI != nil {
			api_v1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			api_v1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			api_v1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Metadata != nil {
			api_v1.SetDefaults_DeprecatedDownwardAPIVolumeSource(a.VolumeSource.Metadata)
			for j := range a.VolumeSource.Metadata.Items {
				b := &a.VolumeSource.Metadata.Items[j]
				if b.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.InitContainers {
		a := &in.Spec.Template.Spec.InitContainers[i]
		api_v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			api_v1.SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		api_v1.SetDefaults_ResourceList(&a.Resources.Limits)
		api_v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			api_v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			api_v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.Containers {
		a := &in.Spec.Template.Spec.Containers[i]
		api_v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			api_v1.SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		api_v1.SetDefaults_ResourceList(&a.Resources.Limits)
		api_v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			api_v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			api_v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Status.AllowedServiceAccounts {
		a := &in.Status.AllowedServiceAccounts[i]
		api_v1.SetDefaults_PodSpec(&a.PodSecurityPolicySubjectReviewStatus.Template.Spec)
		for j := range a.PodSecurityPolicySubjectReviewStatus.Template.Spec.Volumes {
			b := &a.PodSecurityPolicySubjectReviewStatus.Template.Spec.Volumes[j]
			api_v1.SetDefaults_Volume(b)
			if b.VolumeSource.Secret != nil {
				api_v1.SetDefaults_SecretVolumeSource(b.VolumeSource.Secret)
			}
			if b.VolumeSource.ISCSI != nil {
				api_v1.SetDefaults_ISCSIVolumeSource(b.VolumeSource.ISCSI)
			}
			if b.VolumeSource.RBD != nil {
				api_v1.SetDefaults_RBDVolumeSource(b.VolumeSource.RBD)
			}
			if b.VolumeSource.DownwardAPI != nil {
				api_v1.SetDefaults_DownwardAPIVolumeSource(b.VolumeSource.DownwardAPI)
				for k := range b.VolumeSource.DownwardAPI.Items {
					c := &b.VolumeSource.DownwardAPI.Items[k]
					if c.FieldRef != nil {
						api_v1.SetDefaults_ObjectFieldSelector(c.FieldRef)
					}
				}
			}
			if b.VolumeSource.ConfigMap != nil {
				api_v1.SetDefaults_ConfigMapVolumeSource(b.VolumeSource.ConfigMap)
			}
			if b.VolumeSource.AzureDisk != nil {
				api_v1.SetDefaults_AzureDiskVolumeSource(b.VolumeSource.AzureDisk)
			}
			if b.VolumeSource.Metadata != nil {
				api_v1.SetDefaults_DeprecatedDownwardAPIVolumeSource(b.VolumeSource.Metadata)
				for k := range b.VolumeSource.Metadata.Items {
					c := &b.VolumeSource.Metadata.Items[k]
					if c.FieldRef != nil {
						api_v1.SetDefaults_ObjectFieldSelector(c.FieldRef)
					}
				}
			}
		}
		for j := range a.PodSecurityPolicySubjectReviewStatus.Template.Spec.InitContainers {
			b := &a.PodSecurityPolicySubjectReviewStatus.Template.Spec.InitContainers[j]
			api_v1.SetDefaults_Container(b)
			for k := range b.Ports {
				c := &b.Ports[k]
				api_v1.SetDefaults_ContainerPort(c)
			}
			for k := range b.Env {
				c := &b.Env[k]
				if c.ValueFrom != nil {
					if c.ValueFrom.FieldRef != nil {
						api_v1.SetDefaults_ObjectFieldSelector(c.ValueFrom.FieldRef)
					}
				}
			}
			api_v1.SetDefaults_ResourceList(&b.Resources.Limits)
			api_v1.SetDefaults_ResourceList(&b.Resources.Requests)
			if b.LivenessProbe != nil {
				api_v1.SetDefaults_Probe(b.LivenessProbe)
				if b.LivenessProbe.Handler.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(b.LivenessProbe.Handler.HTTPGet)
				}
			}
			if b.ReadinessProbe != nil {
				api_v1.SetDefaults_Probe(b.ReadinessProbe)
				if b.ReadinessProbe.Handler.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(b.ReadinessProbe.Handler.HTTPGet)
				}
			}
			if b.Lifecycle != nil {
				if b.Lifecycle.PostStart != nil {
					if b.Lifecycle.PostStart.HTTPGet != nil {
						api_v1.SetDefaults_HTTPGetAction(b.Lifecycle.PostStart.HTTPGet)
					}
				}
				if b.Lifecycle.PreStop != nil {
					if b.Lifecycle.PreStop.HTTPGet != nil {
						api_v1.SetDefaults_HTTPGetAction(b.Lifecycle.PreStop.HTTPGet)
					}
				}
			}
		}
		for j := range a.PodSecurityPolicySubjectReviewStatus.Template.Spec.Containers {
			b := &a.PodSecurityPolicySubjectReviewStatus.Template.Spec.Containers[j]
			api_v1.SetDefaults_Container(b)
			for k := range b.Ports {
				c := &b.Ports[k]
				api_v1.SetDefaults_ContainerPort(c)
			}
			for k := range b.Env {
				c := &b.Env[k]
				if c.ValueFrom != nil {
					if c.ValueFrom.FieldRef != nil {
						api_v1.SetDefaults_ObjectFieldSelector(c.ValueFrom.FieldRef)
					}
				}
			}
			api_v1.SetDefaults_ResourceList(&b.Resources.Limits)
			api_v1.SetDefaults_ResourceList(&b.Resources.Requests)
			if b.LivenessProbe != nil {
				api_v1.SetDefaults_Probe(b.LivenessProbe)
				if b.LivenessProbe.Handler.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(b.LivenessProbe.Handler.HTTPGet)
				}
			}
			if b.ReadinessProbe != nil {
				api_v1.SetDefaults_Probe(b.ReadinessProbe)
				if b.ReadinessProbe.Handler.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(b.ReadinessProbe.Handler.HTTPGet)
				}
			}
			if b.Lifecycle != nil {
				if b.Lifecycle.PostStart != nil {
					if b.Lifecycle.PostStart.HTTPGet != nil {
						api_v1.SetDefaults_HTTPGetAction(b.Lifecycle.PostStart.HTTPGet)
					}
				}
				if b.Lifecycle.PreStop != nil {
					if b.Lifecycle.PreStop.HTTPGet != nil {
						api_v1.SetDefaults_HTTPGetAction(b.Lifecycle.PreStop.HTTPGet)
					}
				}
			}
		}
	}
}

func SetObjectDefaults_PodSecurityPolicySelfSubjectReview(in *PodSecurityPolicySelfSubjectReview) {
	api_v1.SetDefaults_PodSpec(&in.Spec.Template.Spec)
	for i := range in.Spec.Template.Spec.Volumes {
		a := &in.Spec.Template.Spec.Volumes[i]
		api_v1.SetDefaults_Volume(a)
		if a.VolumeSource.Secret != nil {
			api_v1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			api_v1.SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
		}
		if a.VolumeSource.RBD != nil {
			api_v1.SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
		}
		if a.VolumeSource.DownwardAPI != nil {
			api_v1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			api_v1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			api_v1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Metadata != nil {
			api_v1.SetDefaults_DeprecatedDownwardAPIVolumeSource(a.VolumeSource.Metadata)
			for j := range a.VolumeSource.Metadata.Items {
				b := &a.VolumeSource.Metadata.Items[j]
				if b.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.InitContainers {
		a := &in.Spec.Template.Spec.InitContainers[i]
		api_v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			api_v1.SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		api_v1.SetDefaults_ResourceList(&a.Resources.Limits)
		api_v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			api_v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			api_v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.Containers {
		a := &in.Spec.Template.Spec.Containers[i]
		api_v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			api_v1.SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		api_v1.SetDefaults_ResourceList(&a.Resources.Limits)
		api_v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			api_v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			api_v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	api_v1.SetDefaults_PodSpec(&in.Status.Template.Spec)
	for i := range in.Status.Template.Spec.Volumes {
		a := &in.Status.Template.Spec.Volumes[i]
		api_v1.SetDefaults_Volume(a)
		if a.VolumeSource.Secret != nil {
			api_v1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			api_v1.SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
		}
		if a.VolumeSource.RBD != nil {
			api_v1.SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
		}
		if a.VolumeSource.DownwardAPI != nil {
			api_v1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			api_v1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			api_v1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Metadata != nil {
			api_v1.SetDefaults_DeprecatedDownwardAPIVolumeSource(a.VolumeSource.Metadata)
			for j := range a.VolumeSource.Metadata.Items {
				b := &a.VolumeSource.Metadata.Items[j]
				if b.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
	}
	for i := range in.Status.Template.Spec.InitContainers {
		a := &in.Status.Template.Spec.InitContainers[i]
		api_v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			api_v1.SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		api_v1.SetDefaults_ResourceList(&a.Resources.Limits)
		api_v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			api_v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			api_v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Status.Template.Spec.Containers {
		a := &in.Status.Template.Spec.Containers[i]
		api_v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			api_v1.SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		api_v1.SetDefaults_ResourceList(&a.Resources.Limits)
		api_v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			api_v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			api_v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
}

func SetObjectDefaults_PodSecurityPolicySubjectReview(in *PodSecurityPolicySubjectReview) {
	api_v1.SetDefaults_PodSpec(&in.Spec.Template.Spec)
	for i := range in.Spec.Template.Spec.Volumes {
		a := &in.Spec.Template.Spec.Volumes[i]
		api_v1.SetDefaults_Volume(a)
		if a.VolumeSource.Secret != nil {
			api_v1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			api_v1.SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
		}
		if a.VolumeSource.RBD != nil {
			api_v1.SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
		}
		if a.VolumeSource.DownwardAPI != nil {
			api_v1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			api_v1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			api_v1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Metadata != nil {
			api_v1.SetDefaults_DeprecatedDownwardAPIVolumeSource(a.VolumeSource.Metadata)
			for j := range a.VolumeSource.Metadata.Items {
				b := &a.VolumeSource.Metadata.Items[j]
				if b.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.InitContainers {
		a := &in.Spec.Template.Spec.InitContainers[i]
		api_v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			api_v1.SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		api_v1.SetDefaults_ResourceList(&a.Resources.Limits)
		api_v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			api_v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			api_v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.Containers {
		a := &in.Spec.Template.Spec.Containers[i]
		api_v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			api_v1.SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		api_v1.SetDefaults_ResourceList(&a.Resources.Limits)
		api_v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			api_v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			api_v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	api_v1.SetDefaults_PodSpec(&in.Status.Template.Spec)
	for i := range in.Status.Template.Spec.Volumes {
		a := &in.Status.Template.Spec.Volumes[i]
		api_v1.SetDefaults_Volume(a)
		if a.VolumeSource.Secret != nil {
			api_v1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			api_v1.SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
		}
		if a.VolumeSource.RBD != nil {
			api_v1.SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
		}
		if a.VolumeSource.DownwardAPI != nil {
			api_v1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			api_v1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			api_v1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Metadata != nil {
			api_v1.SetDefaults_DeprecatedDownwardAPIVolumeSource(a.VolumeSource.Metadata)
			for j := range a.VolumeSource.Metadata.Items {
				b := &a.VolumeSource.Metadata.Items[j]
				if b.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
	}
	for i := range in.Status.Template.Spec.InitContainers {
		a := &in.Status.Template.Spec.InitContainers[i]
		api_v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			api_v1.SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		api_v1.SetDefaults_ResourceList(&a.Resources.Limits)
		api_v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			api_v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			api_v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Status.Template.Spec.Containers {
		a := &in.Status.Template.Spec.Containers[i]
		api_v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			api_v1.SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		api_v1.SetDefaults_ResourceList(&a.Resources.Limits)
		api_v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			api_v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			api_v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				api_v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					api_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
}
