// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/opencode-go/internal/apijson"
	"github.com/stainless-sdks/opencode-go/internal/requestconfig"
	"github.com/stainless-sdks/opencode-go/option"
	"github.com/stainless-sdks/opencode-go/packages/param"
	"github.com/stainless-sdks/opencode-go/packages/respjson"
	"github.com/stainless-sdks/opencode-go/shared/constant"
)

// SessionService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionService] method instead.
type SessionService struct {
	Options []option.RequestOption
	Share   SessionShareService
	Message SessionMessageService
}

// NewSessionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSessionService(opts ...option.RequestOption) (r SessionService) {
	r = SessionService{}
	r.Options = opts
	r.Share = NewSessionShareService(opts...)
	r.Message = NewSessionMessageService(opts...)
	return
}

// Create a new session
func (r *SessionService) New(ctx context.Context, body SessionNewParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = append(r.Options[:], opts...)
	path := "session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get session
func (r *SessionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update session properties
func (r *SessionService) Update(ctx context.Context, id string, body SessionUpdateParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all sessions
func (r *SessionService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Session, err error) {
	opts = append(r.Options[:], opts...)
	path := "session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a session and all its data
func (r *SessionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Abort a session
func (r *SessionService) Abort(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/abort", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Analyze the app and create an AGENTS.md file
func (r *SessionService) Analyze(ctx context.Context, id string, body SessionAnalyzeParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/init", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a session's children
func (r *SessionService) GetChildren(ctx context.Context, id string, opts ...option.RequestOption) (res *[]Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/children", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Respond to a permission request
func (r *SessionService) RespondToPermission(ctx context.Context, permissionID string, params SessionRespondToPermissionParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if permissionID == "" {
		err = errors.New("missing required permissionID parameter")
		return
	}
	path := fmt.Sprintf("session/%s/permissions/%s", params.ID, permissionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Restore all reverted messages
func (r *SessionService) RestoreReverted(ctx context.Context, id string, opts ...option.RequestOption) (res *Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/unrevert", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Revert a message
func (r *SessionService) Revert(ctx context.Context, id string, body SessionRevertParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/revert", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Run a shell command
func (r *SessionService) RunShell(ctx context.Context, id string, body SessionRunShellParams, opts ...option.RequestOption) (res *AssistantMessage, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/shell", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send a new command to a session
func (r *SessionService) SendCommand(ctx context.Context, id string, body SessionSendCommandParams, opts ...option.RequestOption) (res *SessionSendCommandResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/command", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Summarize the session
func (r *SessionService) Summarize(ctx context.Context, id string, body SessionSummarizeParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/summarize", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AssistantMessage struct {
	ID         string                     `json:"id,required"`
	Cost       float64                    `json:"cost,required"`
	Mode       string                     `json:"mode,required"`
	ModelID    string                     `json:"modelID,required"`
	Path       AssistantMessagePath       `json:"path,required"`
	ProviderID string                     `json:"providerID,required"`
	Role       constant.Assistant         `json:"role,required"`
	SessionID  string                     `json:"sessionID,required"`
	System     []string                   `json:"system,required"`
	Time       AssistantMessageTime       `json:"time,required"`
	Tokens     AssistantMessageTokens     `json:"tokens,required"`
	Error      AssistantMessageErrorUnion `json:"error"`
	Summary    bool                       `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cost        respjson.Field
		Mode        respjson.Field
		ModelID     respjson.Field
		Path        respjson.Field
		ProviderID  respjson.Field
		Role        respjson.Field
		SessionID   respjson.Field
		System      respjson.Field
		Time        respjson.Field
		Tokens      respjson.Field
		Error       respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantMessage) RawJSON() string { return r.JSON.raw }
func (r *AssistantMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantMessagePath struct {
	Cwd  string `json:"cwd,required"`
	Root string `json:"root,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cwd         respjson.Field
		Root        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantMessagePath) RawJSON() string { return r.JSON.raw }
func (r *AssistantMessagePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantMessageTime struct {
	Created   float64 `json:"created,required"`
	Completed float64 `json:"completed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created     respjson.Field
		Completed   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantMessageTime) RawJSON() string { return r.JSON.raw }
func (r *AssistantMessageTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantMessageTokens struct {
	Cache     AssistantMessageTokensCache `json:"cache,required"`
	Input     float64                     `json:"input,required"`
	Output    float64                     `json:"output,required"`
	Reasoning float64                     `json:"reasoning,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cache       respjson.Field
		Input       respjson.Field
		Output      respjson.Field
		Reasoning   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantMessageTokens) RawJSON() string { return r.JSON.raw }
func (r *AssistantMessageTokens) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantMessageTokensCache struct {
	Read  float64 `json:"read,required"`
	Write float64 `json:"write,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Read        respjson.Field
		Write       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantMessageTokensCache) RawJSON() string { return r.JSON.raw }
func (r *AssistantMessageTokensCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssistantMessageErrorUnion contains all possible properties and values from
// [ProviderAuthError], [UnknownError], [MessageOutputLengthError],
// [MessageAbortedError].
//
// Use the [AssistantMessageErrorUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AssistantMessageErrorUnion struct {
	// This field is a union of [ProviderAuthErrorData], [UnknownErrorData], [any],
	// [any]
	Data AssistantMessageErrorUnionData `json:"data"`
	// Any of "ProviderAuthError", "UnknownError", "MessageOutputLengthError",
	// "MessageAbortedError".
	Name string `json:"name"`
	JSON struct {
		Data respjson.Field
		Name respjson.Field
		raw  string
	} `json:"-"`
}

// anyAssistantMessageError is implemented by each variant of
// [AssistantMessageErrorUnion] to add type safety for the return type of
// [AssistantMessageErrorUnion.AsAny]
type anyAssistantMessageError interface {
	implAssistantMessageErrorUnion()
}

func (ProviderAuthError) implAssistantMessageErrorUnion()        {}
func (UnknownError) implAssistantMessageErrorUnion()             {}
func (MessageOutputLengthError) implAssistantMessageErrorUnion() {}
func (MessageAbortedError) implAssistantMessageErrorUnion()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AssistantMessageErrorUnion.AsAny().(type) {
//	case opencode.ProviderAuthError:
//	case opencode.UnknownError:
//	case opencode.MessageOutputLengthError:
//	case opencode.MessageAbortedError:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AssistantMessageErrorUnion) AsAny() anyAssistantMessageError {
	switch u.Name {
	case "ProviderAuthError":
		return u.AsProviderAuthError()
	case "UnknownError":
		return u.AsUnknownError()
	case "MessageOutputLengthError":
		return u.AsMessageOutputLengthError()
	case "MessageAbortedError":
		return u.AsMessageAbortedError()
	}
	return nil
}

func (u AssistantMessageErrorUnion) AsProviderAuthError() (v ProviderAuthError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantMessageErrorUnion) AsUnknownError() (v UnknownError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantMessageErrorUnion) AsMessageOutputLengthError() (v MessageOutputLengthError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantMessageErrorUnion) AsMessageAbortedError() (v MessageAbortedError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AssistantMessageErrorUnion) RawJSON() string { return u.JSON.raw }

func (r *AssistantMessageErrorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssistantMessageErrorUnionData is an implicit subunion of
// [AssistantMessageErrorUnion]. AssistantMessageErrorUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AssistantMessageErrorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMessageAbortedErrorData]
type AssistantMessageErrorUnionData struct {
	// This field will be present if the value is a [any] instead of an object.
	OfMessageAbortedErrorData any    `json:",inline"`
	Message                   string `json:"message"`
	// This field is from variant [ProviderAuthErrorData].
	ProviderID string `json:"providerID"`
	JSON       struct {
		OfMessageAbortedErrorData respjson.Field
		Message                   respjson.Field
		ProviderID                respjson.Field
		raw                       string
	} `json:"-"`
}

func (r *AssistantMessageErrorUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageAbortedError struct {
	Data any                          `json:"data,required"`
	Name constant.MessageAbortedError `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageAbortedError) RawJSON() string { return r.JSON.raw }
func (r *MessageAbortedError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageOutputLengthError struct {
	Data any                               `json:"data,required"`
	Name constant.MessageOutputLengthError `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageOutputLengthError) RawJSON() string { return r.JSON.raw }
func (r *MessageOutputLengthError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProviderAuthError struct {
	Data ProviderAuthErrorData      `json:"data,required"`
	Name constant.ProviderAuthError `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderAuthError) RawJSON() string { return r.JSON.raw }
func (r *ProviderAuthError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProviderAuthErrorData struct {
	Message    string `json:"message,required"`
	ProviderID string `json:"providerID,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ProviderID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderAuthErrorData) RawJSON() string { return r.JSON.raw }
func (r *ProviderAuthErrorData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Session struct {
	ID        string        `json:"id,required"`
	Directory string        `json:"directory,required"`
	ProjectID string        `json:"projectID,required"`
	Time      SessionTime   `json:"time,required"`
	Title     string        `json:"title,required"`
	Version   string        `json:"version,required"`
	ParentID  string        `json:"parentID"`
	Revert    SessionRevert `json:"revert"`
	Share     SessionShare  `json:"share"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Directory   respjson.Field
		ProjectID   respjson.Field
		Time        respjson.Field
		Title       respjson.Field
		Version     respjson.Field
		ParentID    respjson.Field
		Revert      respjson.Field
		Share       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Session) RawJSON() string { return r.JSON.raw }
func (r *Session) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionTime struct {
	Created float64 `json:"created,required"`
	Updated float64 `json:"updated,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created     respjson.Field
		Updated     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionTime) RawJSON() string { return r.JSON.raw }
func (r *SessionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionRevert struct {
	MessageID string `json:"messageID,required"`
	Diff      string `json:"diff"`
	PartID    string `json:"partID"`
	Snapshot  string `json:"snapshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		Diff        respjson.Field
		PartID      respjson.Field
		Snapshot    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionRevert) RawJSON() string { return r.JSON.raw }
func (r *SessionRevert) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionShare struct {
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionShare) RawJSON() string { return r.JSON.raw }
func (r *SessionShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnknownError struct {
	Data UnknownErrorData      `json:"data,required"`
	Name constant.UnknownError `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnknownError) RawJSON() string { return r.JSON.raw }
func (r *UnknownError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnknownErrorData struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnknownErrorData) RawJSON() string { return r.JSON.raw }
func (r *UnknownErrorData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionSendCommandResponse struct {
	Info  AssistantMessage `json:"info,required"`
	Parts []PartUnion      `json:"parts,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		Parts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionSendCommandResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionSendCommandResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionNewParams struct {
	ParentID param.Opt[string] `json:"parentID,omitzero"`
	Title    param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r SessionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionUpdateParams struct {
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r SessionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionAnalyzeParams struct {
	MessageID  string `json:"messageID,required"`
	ModelID    string `json:"modelID,required"`
	ProviderID string `json:"providerID,required"`
	paramObj
}

func (r SessionAnalyzeParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionAnalyzeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionAnalyzeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionRespondToPermissionParams struct {
	ID string `path:"id,required" json:"-"`
	// Any of "once", "always", "reject".
	Response SessionRespondToPermissionParamsResponse `json:"response,omitzero,required"`
	paramObj
}

func (r SessionRespondToPermissionParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionRespondToPermissionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionRespondToPermissionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionRespondToPermissionParamsResponse string

const (
	SessionRespondToPermissionParamsResponseOnce   SessionRespondToPermissionParamsResponse = "once"
	SessionRespondToPermissionParamsResponseAlways SessionRespondToPermissionParamsResponse = "always"
	SessionRespondToPermissionParamsResponseReject SessionRespondToPermissionParamsResponse = "reject"
)

type SessionRevertParams struct {
	MessageID string            `json:"messageID,required"`
	PartID    param.Opt[string] `json:"partID,omitzero"`
	paramObj
}

func (r SessionRevertParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionRevertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionRevertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionRunShellParams struct {
	Agent   string `json:"agent,required"`
	Command string `json:"command,required"`
	paramObj
}

func (r SessionRunShellParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionRunShellParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionRunShellParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionSendCommandParams struct {
	Arguments string            `json:"arguments,required"`
	Command   string            `json:"command,required"`
	Agent     param.Opt[string] `json:"agent,omitzero"`
	MessageID param.Opt[string] `json:"messageID,omitzero"`
	Model     param.Opt[string] `json:"model,omitzero"`
	paramObj
}

func (r SessionSendCommandParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionSendCommandParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionSendCommandParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionSummarizeParams struct {
	ModelID    string `json:"modelID,required"`
	ProviderID string `json:"providerID,required"`
	paramObj
}

func (r SessionSummarizeParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionSummarizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionSummarizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
