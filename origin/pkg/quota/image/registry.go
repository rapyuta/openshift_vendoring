// Package image implements evaluators of usage for imagestreams and images. They are supposed
// to be passed to resource quota controller and origin resource quota admission plugin.
package image

import (
	"github.com/openshift/kubernetes/pkg/api/unversioned"
	"github.com/openshift/kubernetes/pkg/quota"
	"github.com/openshift/kubernetes/pkg/quota/generic"

	osclient "github.com/openshift/origin/pkg/client"
	"github.com/openshift/origin/pkg/controller/shared"
)

// NewImageQuotaRegistry returns a registry for quota evaluation of OpenShift resources related to images in
// internal registry. It evaluates only image streams and related virtual resources that can cause a creation
// of new image stream objects.
func NewImageQuotaRegistry(isInformer shared.ImageStreamInformer, osClient osclient.Interface) quota.Registry {
	imageStream := NewImageStreamEvaluator(isInformer.Lister())
	imageStreamTag := NewImageStreamTagEvaluator(isInformer.Lister(), osClient)
	imageStreamImport := NewImageStreamImportEvaluator(isInformer.Lister())
	return &generic.GenericRegistry{
		InternalEvaluators: map[unversioned.GroupKind]quota.Evaluator{
			imageStream.GroupKind():       imageStream,
			imageStreamTag.GroupKind():    imageStreamTag,
			imageStreamImport.GroupKind(): imageStreamImport,
		},
	}
}
