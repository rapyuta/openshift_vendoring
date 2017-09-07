// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	unversioned "github.com/openshift/kubernetes/pkg/api/unversioned"
	api_v1 "github.com/openshift/kubernetes/pkg/api/v1"
	conversion "github.com/openshift/kubernetes/pkg/conversion"
	runtime "github.com/openshift/kubernetes/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DockerImageReference, InType: reflect.TypeOf(&DockerImageReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Image, InType: reflect.TypeOf(&Image{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageImportSpec, InType: reflect.TypeOf(&ImageImportSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageImportStatus, InType: reflect.TypeOf(&ImageImportStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageLayer, InType: reflect.TypeOf(&ImageLayer{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageList, InType: reflect.TypeOf(&ImageList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageSignature, InType: reflect.TypeOf(&ImageSignature{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageStream, InType: reflect.TypeOf(&ImageStream{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageStreamImage, InType: reflect.TypeOf(&ImageStreamImage{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageStreamImport, InType: reflect.TypeOf(&ImageStreamImport{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageStreamImportSpec, InType: reflect.TypeOf(&ImageStreamImportSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageStreamImportStatus, InType: reflect.TypeOf(&ImageStreamImportStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageStreamList, InType: reflect.TypeOf(&ImageStreamList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageStreamMapping, InType: reflect.TypeOf(&ImageStreamMapping{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageStreamSpec, InType: reflect.TypeOf(&ImageStreamSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageStreamStatus, InType: reflect.TypeOf(&ImageStreamStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageStreamTag, InType: reflect.TypeOf(&ImageStreamTag{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageStreamTagList, InType: reflect.TypeOf(&ImageStreamTagList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_NamedTagEventList, InType: reflect.TypeOf(&NamedTagEventList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RepositoryImportSpec, InType: reflect.TypeOf(&RepositoryImportSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RepositoryImportStatus, InType: reflect.TypeOf(&RepositoryImportStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SignatureCondition, InType: reflect.TypeOf(&SignatureCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SignatureGenericEntity, InType: reflect.TypeOf(&SignatureGenericEntity{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SignatureIssuer, InType: reflect.TypeOf(&SignatureIssuer{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SignatureSubject, InType: reflect.TypeOf(&SignatureSubject{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_TagEvent, InType: reflect.TypeOf(&TagEvent{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_TagEventCondition, InType: reflect.TypeOf(&TagEventCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_TagImportPolicy, InType: reflect.TypeOf(&TagImportPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_TagReference, InType: reflect.TypeOf(&TagReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_TagReferencePolicy, InType: reflect.TypeOf(&TagReferencePolicy{})},
	)
}

func DeepCopy_v1_DockerImageReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DockerImageReference)
		out := out.(*DockerImageReference)
		out.Registry = in.Registry
		out.Namespace = in.Namespace
		out.Name = in.Name
		out.Tag = in.Tag
		out.ID = in.ID
		return nil
	}
}

func DeepCopy_v1_Image(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Image)
		out := out.(*Image)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.DockerImageReference = in.DockerImageReference
		if err := runtime.DeepCopy_runtime_RawExtension(&in.DockerImageMetadata, &out.DockerImageMetadata, c); err != nil {
			return err
		}
		out.DockerImageMetadataVersion = in.DockerImageMetadataVersion
		out.DockerImageManifest = in.DockerImageManifest
		if in.DockerImageLayers != nil {
			in, out := &in.DockerImageLayers, &out.DockerImageLayers
			*out = make([]ImageLayer, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.DockerImageLayers = nil
		}
		if in.Signatures != nil {
			in, out := &in.Signatures, &out.Signatures
			*out = make([]ImageSignature, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ImageSignature(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Signatures = nil
		}
		if in.DockerImageSignatures != nil {
			in, out := &in.DockerImageSignatures, &out.DockerImageSignatures
			*out = make([][]byte, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*[]byte)
				}
			}
		} else {
			out.DockerImageSignatures = nil
		}
		out.DockerImageManifestMediaType = in.DockerImageManifestMediaType
		out.DockerImageConfig = in.DockerImageConfig
		return nil
	}
}

func DeepCopy_v1_ImageImportSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageImportSpec)
		out := out.(*ImageImportSpec)
		out.From = in.From
		if in.To != nil {
			in, out := &in.To, &out.To
			*out = new(api_v1.LocalObjectReference)
			**out = **in
		} else {
			out.To = nil
		}
		out.ImportPolicy = in.ImportPolicy
		out.IncludeManifest = in.IncludeManifest
		return nil
	}
}

func DeepCopy_v1_ImageImportStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageImportStatus)
		out := out.(*ImageImportStatus)
		if err := unversioned.DeepCopy_unversioned_Status(&in.Status, &out.Status, c); err != nil {
			return err
		}
		if in.Image != nil {
			in, out := &in.Image, &out.Image
			*out = new(Image)
			if err := DeepCopy_v1_Image(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Image = nil
		}
		out.Tag = in.Tag
		return nil
	}
}

func DeepCopy_v1_ImageLayer(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageLayer)
		out := out.(*ImageLayer)
		out.Name = in.Name
		out.LayerSize = in.LayerSize
		out.MediaType = in.MediaType
		return nil
	}
}

func DeepCopy_v1_ImageList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageList)
		out := out.(*ImageList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Image, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_Image(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1_ImageSignature(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageSignature)
		out := out.(*ImageSignature)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.Type = in.Type
		if in.Content != nil {
			in, out := &in.Content, &out.Content
			*out = make([]byte, len(*in))
			copy(*out, *in)
		} else {
			out.Content = nil
		}
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]SignatureCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_SignatureCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Conditions = nil
		}
		out.ImageIdentity = in.ImageIdentity
		if in.SignedClaims != nil {
			in, out := &in.SignedClaims, &out.SignedClaims
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.SignedClaims = nil
		}
		if in.Created != nil {
			in, out := &in.Created, &out.Created
			*out = new(unversioned.Time)
			**out = (*in).DeepCopy()
		} else {
			out.Created = nil
		}
		if in.IssuedBy != nil {
			in, out := &in.IssuedBy, &out.IssuedBy
			*out = new(SignatureIssuer)
			**out = **in
		} else {
			out.IssuedBy = nil
		}
		if in.IssuedTo != nil {
			in, out := &in.IssuedTo, &out.IssuedTo
			*out = new(SignatureSubject)
			**out = **in
		} else {
			out.IssuedTo = nil
		}
		return nil
	}
}

func DeepCopy_v1_ImageStream(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageStream)
		out := out.(*ImageStream)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_ImageStreamSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_ImageStreamStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_ImageStreamImage(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageStreamImage)
		out := out.(*ImageStreamImage)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_Image(&in.Image, &out.Image, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_ImageStreamImport(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageStreamImport)
		out := out.(*ImageStreamImport)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_ImageStreamImportSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_ImageStreamImportStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_ImageStreamImportSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageStreamImportSpec)
		out := out.(*ImageStreamImportSpec)
		out.Import = in.Import
		if in.Repository != nil {
			in, out := &in.Repository, &out.Repository
			*out = new(RepositoryImportSpec)
			**out = **in
		} else {
			out.Repository = nil
		}
		if in.Images != nil {
			in, out := &in.Images, &out.Images
			*out = make([]ImageImportSpec, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ImageImportSpec(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Images = nil
		}
		return nil
	}
}

func DeepCopy_v1_ImageStreamImportStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageStreamImportStatus)
		out := out.(*ImageStreamImportStatus)
		if in.Import != nil {
			in, out := &in.Import, &out.Import
			*out = new(ImageStream)
			if err := DeepCopy_v1_ImageStream(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Import = nil
		}
		if in.Repository != nil {
			in, out := &in.Repository, &out.Repository
			*out = new(RepositoryImportStatus)
			if err := DeepCopy_v1_RepositoryImportStatus(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Repository = nil
		}
		if in.Images != nil {
			in, out := &in.Images, &out.Images
			*out = make([]ImageImportStatus, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ImageImportStatus(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Images = nil
		}
		return nil
	}
}

func DeepCopy_v1_ImageStreamList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageStreamList)
		out := out.(*ImageStreamList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ImageStream, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ImageStream(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1_ImageStreamMapping(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageStreamMapping)
		out := out.(*ImageStreamMapping)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_Image(&in.Image, &out.Image, c); err != nil {
			return err
		}
		out.Tag = in.Tag
		return nil
	}
}

func DeepCopy_v1_ImageStreamSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageStreamSpec)
		out := out.(*ImageStreamSpec)
		out.DockerImageRepository = in.DockerImageRepository
		if in.Tags != nil {
			in, out := &in.Tags, &out.Tags
			*out = make([]TagReference, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_TagReference(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Tags = nil
		}
		return nil
	}
}

func DeepCopy_v1_ImageStreamStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageStreamStatus)
		out := out.(*ImageStreamStatus)
		out.DockerImageRepository = in.DockerImageRepository
		if in.Tags != nil {
			in, out := &in.Tags, &out.Tags
			*out = make([]NamedTagEventList, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_NamedTagEventList(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Tags = nil
		}
		return nil
	}
}

func DeepCopy_v1_ImageStreamTag(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageStreamTag)
		out := out.(*ImageStreamTag)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if in.Tag != nil {
			in, out := &in.Tag, &out.Tag
			*out = new(TagReference)
			if err := DeepCopy_v1_TagReference(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Tag = nil
		}
		out.Generation = in.Generation
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]TagEventCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_TagEventCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Conditions = nil
		}
		if err := DeepCopy_v1_Image(&in.Image, &out.Image, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_ImageStreamTagList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageStreamTagList)
		out := out.(*ImageStreamTagList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ImageStreamTag, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ImageStreamTag(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1_NamedTagEventList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NamedTagEventList)
		out := out.(*NamedTagEventList)
		out.Tag = in.Tag
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]TagEvent, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_TagEvent(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]TagEventCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_TagEventCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Conditions = nil
		}
		return nil
	}
}

func DeepCopy_v1_RepositoryImportSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RepositoryImportSpec)
		out := out.(*RepositoryImportSpec)
		out.From = in.From
		out.ImportPolicy = in.ImportPolicy
		out.IncludeManifest = in.IncludeManifest
		return nil
	}
}

func DeepCopy_v1_RepositoryImportStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RepositoryImportStatus)
		out := out.(*RepositoryImportStatus)
		if err := unversioned.DeepCopy_unversioned_Status(&in.Status, &out.Status, c); err != nil {
			return err
		}
		if in.Images != nil {
			in, out := &in.Images, &out.Images
			*out = make([]ImageImportStatus, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ImageImportStatus(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Images = nil
		}
		if in.AdditionalTags != nil {
			in, out := &in.AdditionalTags, &out.AdditionalTags
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.AdditionalTags = nil
		}
		return nil
	}
}

func DeepCopy_v1_SignatureCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SignatureCondition)
		out := out.(*SignatureCondition)
		out.Type = in.Type
		out.Status = in.Status
		out.LastProbeTime = in.LastProbeTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		out.Reason = in.Reason
		out.Message = in.Message
		return nil
	}
}

func DeepCopy_v1_SignatureGenericEntity(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SignatureGenericEntity)
		out := out.(*SignatureGenericEntity)
		out.Organization = in.Organization
		out.CommonName = in.CommonName
		return nil
	}
}

func DeepCopy_v1_SignatureIssuer(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SignatureIssuer)
		out := out.(*SignatureIssuer)
		out.SignatureGenericEntity = in.SignatureGenericEntity
		return nil
	}
}

func DeepCopy_v1_SignatureSubject(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SignatureSubject)
		out := out.(*SignatureSubject)
		out.SignatureGenericEntity = in.SignatureGenericEntity
		out.PublicKeyID = in.PublicKeyID
		return nil
	}
}

func DeepCopy_v1_TagEvent(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TagEvent)
		out := out.(*TagEvent)
		out.Created = in.Created.DeepCopy()
		out.DockerImageReference = in.DockerImageReference
		out.Image = in.Image
		out.Generation = in.Generation
		return nil
	}
}

func DeepCopy_v1_TagEventCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TagEventCondition)
		out := out.(*TagEventCondition)
		out.Type = in.Type
		out.Status = in.Status
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		out.Reason = in.Reason
		out.Message = in.Message
		out.Generation = in.Generation
		return nil
	}
}

func DeepCopy_v1_TagImportPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TagImportPolicy)
		out := out.(*TagImportPolicy)
		out.Insecure = in.Insecure
		out.Scheduled = in.Scheduled
		return nil
	}
}

func DeepCopy_v1_TagReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TagReference)
		out := out.(*TagReference)
		out.Name = in.Name
		if in.Annotations != nil {
			in, out := &in.Annotations, &out.Annotations
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Annotations = nil
		}
		if in.From != nil {
			in, out := &in.From, &out.From
			*out = new(api_v1.ObjectReference)
			**out = **in
		} else {
			out.From = nil
		}
		out.Reference = in.Reference
		if in.Generation != nil {
			in, out := &in.Generation, &out.Generation
			*out = new(int64)
			**out = **in
		} else {
			out.Generation = nil
		}
		out.ImportPolicy = in.ImportPolicy
		out.ReferencePolicy = in.ReferencePolicy
		return nil
	}
}

func DeepCopy_v1_TagReferencePolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TagReferencePolicy)
		out := out.(*TagReferencePolicy)
		out.Type = in.Type
		return nil
	}
}
