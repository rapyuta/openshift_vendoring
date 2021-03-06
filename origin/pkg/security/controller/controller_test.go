package controller

import (
	"fmt"
	"strings"
	"testing"

	kapi "github.com/openshift/kubernetes/pkg/api"
	"github.com/openshift/kubernetes/pkg/api/errors"
	"github.com/openshift/kubernetes/pkg/client/clientset_generated/internalclientset/fake"
	"github.com/openshift/kubernetes/pkg/client/testing/core"
	"github.com/openshift/kubernetes/pkg/runtime"

	"github.com/openshift/origin/pkg/security"
	"github.com/openshift/origin/pkg/security/mcs"
	"github.com/openshift/origin/pkg/security/uid"
	"github.com/openshift/origin/pkg/security/uidallocator"
)

func TestController(t *testing.T) {
	var action core.Action
	client := &fake.Clientset{}
	client.AddReactor("*", "*", func(a core.Action) (handled bool, ret runtime.Object, err error) {
		action = a
		return true, (*kapi.Namespace)(nil), nil
	})

	uidr, _ := uid.NewRange(10, 20, 2)
	mcsr, _ := mcs.NewRange("s0:", 10, 2)
	uida := uidallocator.NewInMemory(uidr)
	c := Allocation{
		uid:    uida,
		mcs:    DefaultMCSAllocation(uidr, mcsr, 5),
		client: client.Core().Namespaces(),
	}

	err := c.Next(&kapi.Namespace{ObjectMeta: kapi.ObjectMeta{Name: "test"}})
	if err != nil {
		t.Fatal(err)
	}

	got := action.(core.CreateAction).GetObject().(*kapi.Namespace)
	if got.Annotations[security.UIDRangeAnnotation] != "10/2" {
		t.Errorf("unexpected uid annotation: %#v", got)
	}
	if got.Annotations[security.SupplementalGroupsAnnotation] != "10/2" {
		t.Errorf("unexpected supplemental group annotation: %#v", got)
	}
	if got.Annotations[security.MCSAnnotation] != "s0:c1,c0" {
		t.Errorf("unexpected mcs annotation: %#v", got)
	}
	if !uida.Has(uid.Block{Start: 10, End: 11}) {
		t.Errorf("did not allocate uid: %#v", uida)
	}
}

func TestControllerError(t *testing.T) {
	testCases := map[string]struct {
		err     func() error
		errFn   func(err error) bool
		reactFn core.ReactionFunc
		actions int
	}{
		"not found": {
			err:     func() error { return errors.NewNotFound(kapi.Resource("Namespace"), "test") },
			errFn:   func(err error) bool { return err == nil },
			actions: 1,
		},
		"unknown": {
			err:     func() error { return fmt.Errorf("unknown") },
			errFn:   func(err error) bool { return err.Error() == "unknown" },
			actions: 1,
		},
		"conflict": {
			actions: 4,
			reactFn: func(a core.Action) (bool, runtime.Object, error) {
				if a.Matches("get", "namespaces") {
					return true, &kapi.Namespace{ObjectMeta: kapi.ObjectMeta{Name: "test"}}, nil
				}
				return true, (*kapi.Namespace)(nil), errors.NewConflict(kapi.Resource("namespace"), "test", fmt.Errorf("test conflict"))
			},
			errFn: func(err error) bool {
				return err != nil && strings.Contains(err.Error(), "unable to allocate security info")
			},
		},
	}

	for s, testCase := range testCases {
		client := &fake.Clientset{}

		if testCase.reactFn == nil {
			testCase.reactFn = func(a core.Action) (bool, runtime.Object, error) {
				return true, (*kapi.Namespace)(nil), testCase.err()
			}
		}

		client.AddReactor("*", "*", testCase.reactFn)

		uidr, _ := uid.NewRange(10, 19, 2)
		mcsr, _ := mcs.NewRange("s0:", 10, 2)
		uida := uidallocator.NewInMemory(uidr)
		c := Allocation{
			uid:    uida,
			mcs:    DefaultMCSAllocation(uidr, mcsr, 5),
			client: client.Core().Namespaces(),
		}

		err := c.Next(&kapi.Namespace{ObjectMeta: kapi.ObjectMeta{Name: "test"}})
		if !testCase.errFn(err) {
			t.Errorf("%s: unexpected error: %v", s, err)
		}

		if len(client.Actions()) != testCase.actions {
			t.Errorf("%s: expected %d actions: %v", s, testCase.actions, client.Actions())
		}
		if uida.Free() != 5 {
			t.Errorf("%s: should not have allocated uid: %d/%d", s, uida.Free(), uidr.Size())
		}
	}
}
