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

// AgentService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options []option.RequestOption
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r AgentService) {
	r = AgentService{}
	r.Options = opts
	return
}

// List all agents
func (r *AgentService) List(ctx context.Context, opts ...option.RequestOption) (res *[]AgentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "agent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AgentListResponse struct {
	BuiltIn bool `json:"builtIn,required"`
	// Any of "subagent", "primary", "all".
	Mode        AgentListResponseMode       `json:"mode,required"`
	Name        string                      `json:"name,required"`
	Options     map[string]any              `json:"options,required"`
	Permission  AgentListResponsePermission `json:"permission,required"`
	Tools       map[string]bool             `json:"tools,required"`
	Description string                      `json:"description"`
	Model       AgentListResponseModel      `json:"model"`
	Prompt      string                      `json:"prompt"`
	Temperature float64                     `json:"temperature"`
	TopP        float64                     `json:"topP"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BuiltIn     respjson.Field
		Mode        respjson.Field
		Name        respjson.Field
		Options     respjson.Field
		Permission  respjson.Field
		Tools       respjson.Field
		Description respjson.Field
		Model       respjson.Field
		Prompt      respjson.Field
		Temperature respjson.Field
		TopP        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListResponseMode string

const (
	AgentListResponseModeSubagent AgentListResponseMode = "subagent"
	AgentListResponseModePrimary  AgentListResponseMode = "primary"
	AgentListResponseModeAll      AgentListResponseMode = "all"
)

type AgentListResponsePermission struct {
	Bash map[string]AgentListResponsePermissionBash `json:"bash,required"`
	// Any of "ask", "allow", "deny".
	Edit AgentListResponsePermissionEdit `json:"edit,required"`
	// Any of "ask", "allow", "deny".
	Webfetch AgentListResponsePermissionWebfetch `json:"webfetch"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bash        respjson.Field
		Edit        respjson.Field
		Webfetch    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentListResponsePermission) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponsePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListResponsePermissionBash string

const (
	AgentListResponsePermissionBashAsk   AgentListResponsePermissionBash = "ask"
	AgentListResponsePermissionBashAllow AgentListResponsePermissionBash = "allow"
	AgentListResponsePermissionBashDeny  AgentListResponsePermissionBash = "deny"
)

type AgentListResponsePermissionEdit string

const (
	AgentListResponsePermissionEditAsk   AgentListResponsePermissionEdit = "ask"
	AgentListResponsePermissionEditAllow AgentListResponsePermissionEdit = "allow"
	AgentListResponsePermissionEditDeny  AgentListResponsePermissionEdit = "deny"
)

type AgentListResponsePermissionWebfetch string

const (
	AgentListResponsePermissionWebfetchAsk   AgentListResponsePermissionWebfetch = "ask"
	AgentListResponsePermissionWebfetchAllow AgentListResponsePermissionWebfetch = "allow"
	AgentListResponsePermissionWebfetchDeny  AgentListResponsePermissionWebfetch = "deny"
)

type AgentListResponseModel struct {
	ModelID    string `json:"modelID,required"`
	ProviderID string `json:"providerID,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ModelID     respjson.Field
		ProviderID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentListResponseModel) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
