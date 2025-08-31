// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/opencode-go/internal/apijson"
	"github.com/stainless-sdks/opencode-go/internal/requestconfig"
	"github.com/stainless-sdks/opencode-go/option"
	"github.com/stainless-sdks/opencode-go/packages/respjson"
)

// CommandService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommandService] method instead.
type CommandService struct {
	Options []option.RequestOption
}

// NewCommandService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCommandService(opts ...option.RequestOption) (r CommandService) {
	r = CommandService{}
	r.Options = opts
	return
}

// List all commands
func (r *CommandService) List(ctx context.Context, opts ...option.RequestOption) (res *[]CommandListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "command"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CommandListResponse struct {
	Name        string `json:"name,required"`
	Template    string `json:"template,required"`
	Agent       string `json:"agent"`
	Description string `json:"description"`
	Model       string `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Template    respjson.Field
		Agent       respjson.Field
		Description respjson.Field
		Model       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommandListResponse) RawJSON() string { return r.JSON.raw }
func (r *CommandListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
