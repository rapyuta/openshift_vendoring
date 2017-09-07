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
	scheme.AddTypeDefaultingFunc(&DeploymentConfig{}, func(obj interface{}) { SetObjectDefaults_DeploymentConfig(obj.(*DeploymentConfig)) })
	scheme.AddTypeDefaultingFunc(&DeploymentConfigList{}, func(obj interface{}) { SetObjectDefaults_DeploymentConfigList(obj.(*DeploymentConfigList)) })
	return nil
}

func SetObjectDefaults_DeploymentConfig(in *DeploymentConfig) {
	SetDefaults_DeploymentConfig(in)
	SetDefaults_DeploymentConfigSpec(&in.Spec)
	SetDefaults_DeploymentStrategy(&in.Spec.Strategy)
	if in.Spec.Strategy.CustomParams != nil {
		for i := range in.Spec.Strategy.CustomParams.Environment {
			a := &in.Spec.Strategy.CustomParams.Environment[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					api_v1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
	}
	if in.Spec.Strategy.RecreateParams != nil {
		SetDefaults_RecreateDeploymentStrategyParams(in.Spec.Strategy.RecreateParams)
		if in.Spec.Strategy.RecreateParams.Pre != nil {
			if in.Spec.Strategy.RecreateParams.Pre.ExecNewPod != nil {
				for i := range in.Spec.Strategy.RecreateParams.Pre.ExecNewPod.Env {
					a := &in.Spec.Strategy.RecreateParams.Pre.ExecNewPod.Env[i]
					if a.ValueFrom != nil {
						if a.ValueFrom.FieldRef != nil {
							api_v1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
						}
					}
				}
			}
		}
		if in.Spec.Strategy.RecreateParams.Mid != nil {
			if in.Spec.Strategy.RecreateParams.Mid.ExecNewPod != nil {
				for i := range in.Spec.Strategy.RecreateParams.Mid.ExecNewPod.Env {
					a := &in.Spec.Strategy.RecreateParams.Mid.ExecNewPod.Env[i]
					if a.ValueFrom != nil {
						if a.ValueFrom.FieldRef != nil {
							api_v1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
						}
					}
				}
			}
		}
		if in.Spec.Strategy.RecreateParams.Post != nil {
			if in.Spec.Strategy.RecreateParams.Post.ExecNewPod != nil {
				for i := range in.Spec.Strategy.RecreateParams.Post.ExecNewPod.Env {
					a := &in.Spec.Strategy.RecreateParams.Post.ExecNewPod.Env[i]
					if a.ValueFrom != nil {
						if a.ValueFrom.FieldRef != nil {
							api_v1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
						}
					}
				}
			}
		}
	}
	if in.Spec.Strategy.RollingParams != nil {
		SetDefaults_RollingDeploymentStrategyParams(in.Spec.Strategy.RollingParams)
		if in.Spec.Strategy.RollingParams.Pre != nil {
			if in.Spec.Strategy.RollingParams.Pre.ExecNewPod != nil {
				for i := range in.Spec.Strategy.RollingParams.Pre.ExecNewPod.Env {
					a := &in.Spec.Strategy.RollingParams.Pre.ExecNewPod.Env[i]
					if a.ValueFrom != nil {
						if a.ValueFrom.FieldRef != nil {
							api_v1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
						}
					}
				}
			}
		}
		if in.Spec.Strategy.RollingParams.Post != nil {
			if in.Spec.Strategy.RollingParams.Post.ExecNewPod != nil {
				for i := range in.Spec.Strategy.RollingParams.Post.ExecNewPod.Env {
					a := &in.Spec.Strategy.RollingParams.Post.ExecNewPod.Env[i]
					if a.ValueFrom != nil {
						if a.ValueFrom.FieldRef != nil {
							api_v1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
						}
					}
				}
			}
		}
	}
	api_v1.SetDefaults_ResourceList(&in.Spec.Strategy.Resources.Limits)
	api_v1.SetDefaults_ResourceList(&in.Spec.Strategy.Resources.Requests)
	if in.Spec.Template != nil {
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
	}
}

func SetObjectDefaults_DeploymentConfigList(in *DeploymentConfigList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_DeploymentConfig(a)
	}
}
