package remotetokenreview

import (
	"errors"

	"github.com/openshift/kubernetes/pkg/apis/authentication"
	"github.com/openshift/kubernetes/pkg/auth/user"
	internalauthentication "github.com/openshift/kubernetes/pkg/client/clientset_generated/internalclientset/typed/authentication/internalversion"
)

type Authenticator struct {
	authenticationClient internalauthentication.TokenReviewsGetter
}

// NewAuthenticator authenticates by doing a tokenreview
func NewAuthenticator(authenticationClient internalauthentication.TokenReviewsGetter) (*Authenticator, error) {
	return &Authenticator{
		authenticationClient: authenticationClient,
	}, nil
}

func (a *Authenticator) AuthenticateToken(value string) (user.Info, bool, error) {
	if len(value) == 0 {
		return nil, false, nil
	}
	tokenReview := &authentication.TokenReview{}
	tokenReview.Spec.Token = value

	response, err := a.authenticationClient.TokenReviews().Create(tokenReview)
	if err != nil {
		return nil, false, err
	}

	if len(response.Status.Error) > 0 {
		return nil, false, errors.New(response.Status.Error)
	}
	if !response.Status.Authenticated {
		return nil, false, nil
	}

	userInfo := &user.DefaultInfo{
		Name:   response.Status.User.Username,
		UID:    response.Status.User.UID,
		Groups: response.Status.User.Groups,
		Extra:  map[string][]string{},
	}
	for k, v := range response.Status.User.Extra {
		userInfo.Extra[k] = v
	}

	return userInfo, true, nil
}
