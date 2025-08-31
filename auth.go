// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/opencode-go/internal/apijson"
	"github.com/stainless-sdks/opencode-go/internal/requestconfig"
	"github.com/stainless-sdks/opencode-go/option"
	"github.com/stainless-sdks/opencode-go/packages/param"
	"github.com/stainless-sdks/opencode-go/shared/constant"
)

// AuthService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
	r.Options = opts
	return
}

// Set authentication credentials
func (r *AuthService) SetCredentials(ctx context.Context, id string, body AuthSetCredentialsParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("auth/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AuthSetCredentialsParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfOAuth *AuthSetCredentialsParamsBodyOAuth `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAPI *AuthSetCredentialsParamsBodyAPI `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWellknown *AuthSetCredentialsParamsBodyWellknown `json:",inline"`

	paramObj
}

func (u AuthSetCredentialsParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOAuth, u.OfAPI, u.OfWellknown)
}
func (r *AuthSetCredentialsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Access, Expires, Refresh, Type are required.
type AuthSetCredentialsParamsBodyOAuth struct {
	Access  string  `json:"access,required"`
	Expires float64 `json:"expires,required"`
	Refresh string  `json:"refresh,required"`
	// This field can be elided, and will marshal its zero value as "oauth".
	Type constant.OAuth `json:"type,required"`
	paramObj
}

func (r AuthSetCredentialsParamsBodyOAuth) MarshalJSON() (data []byte, err error) {
	type shadow AuthSetCredentialsParamsBodyOAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthSetCredentialsParamsBodyOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Type are required.
type AuthSetCredentialsParamsBodyAPI struct {
	Key string `json:"key,required"`
	// This field can be elided, and will marshal its zero value as "api".
	Type constant.API `json:"type,required"`
	paramObj
}

func (r AuthSetCredentialsParamsBodyAPI) MarshalJSON() (data []byte, err error) {
	type shadow AuthSetCredentialsParamsBodyAPI
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthSetCredentialsParamsBodyAPI) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Token, Key, Type are required.
type AuthSetCredentialsParamsBodyWellknown struct {
	Token string `json:"token,required"`
	Key   string `json:"key,required"`
	// This field can be elided, and will marshal its zero value as "wellknown".
	Type constant.Wellknown `json:"type,required"`
	paramObj
}

func (r AuthSetCredentialsParamsBodyWellknown) MarshalJSON() (data []byte, err error) {
	type shadow AuthSetCredentialsParamsBodyWellknown
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthSetCredentialsParamsBodyWellknown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
