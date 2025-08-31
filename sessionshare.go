// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/opencode-go/internal/requestconfig"
	"github.com/stainless-sdks/opencode-go/option"
)

// SessionShareService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionShareService] method instead.
type SessionShareService struct {
	Options []option.RequestOption
}

// NewSessionShareService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSessionShareService(opts ...option.RequestOption) (r SessionShareService) {
	r = SessionShareService{}
	r.Options = opts
	return
}

// Share a session
func (r *SessionShareService) New(ctx context.Context, id string, opts ...option.RequestOption) (res *Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/share", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Unshare the session
func (r *SessionShareService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/share", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}
