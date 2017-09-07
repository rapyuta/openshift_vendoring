package v1

import (
	kapi "github.com/openshift/kubernetes/pkg/api"
	"github.com/openshift/kubernetes/pkg/runtime"
)

func SetDefaults_ImagePolicyConfig(obj *ImagePolicyConfig) {
	if obj == nil {
		return
	}

	if len(obj.ResolveImages) == 0 {
		obj.ResolveImages = Attempt
	}

	for i := range obj.ExecutionRules {
		if len(obj.ExecutionRules[i].OnResources) == 0 {
			obj.ExecutionRules[i].OnResources = []GroupResource{{Resource: "pods", Group: kapi.GroupName}}
		}
	}

}

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return scheme.AddDefaultingFuncs(
		SetDefaults_ImagePolicyConfig,
	)
}
