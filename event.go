// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stainless-sdks/opencode-go/internal/apijson"
	"github.com/stainless-sdks/opencode-go/internal/requestconfig"
	"github.com/stainless-sdks/opencode-go/option"
	"github.com/stainless-sdks/opencode-go/packages/respjson"
	"github.com/stainless-sdks/opencode-go/packages/ssestream"
	"github.com/stainless-sdks/opencode-go/shared/constant"
)

// EventService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.Options = opts
	return
}

// Get events
func (r *EventService) ListStreaming(ctx context.Context, opts ...option.RequestOption) (stream *ssestream.Stream[EventListResponseUnion]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "event"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[EventListResponseUnion](ssestream.NewDecoder(raw), err)
}

// EventListResponseUnion contains all possible properties and values from
// [EventListResponseInstallationUpdated], [EventListResponseLspClientDiagnostics],
// [EventListResponseMessageUpdated], [EventListResponseMessageRemoved],
// [EventListResponseMessagePartUpdated], [EventListResponseMessagePartRemoved],
// [EventListResponsePermissionUpdated], [EventListResponsePermissionReplied],
// [EventListResponseFileEdited], [EventListResponseSessionUpdated],
// [EventListResponseSessionDeleted], [EventListResponseSessionIdle],
// [EventListResponseSessionError], [EventListResponseServerConnected].
//
// Use the [EventListResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EventListResponseUnion struct {
	// This field is a union of [EventListResponseInstallationUpdatedProperties],
	// [EventListResponseLspClientDiagnosticsProperties],
	// [EventListResponseMessageUpdatedProperties],
	// [EventListResponseMessageRemovedProperties],
	// [EventListResponseMessagePartUpdatedProperties],
	// [EventListResponseMessagePartRemovedProperties],
	// [EventListResponsePermissionUpdatedProperties],
	// [EventListResponsePermissionRepliedProperties],
	// [EventListResponseFileEditedProperties],
	// [EventListResponseSessionUpdatedProperties],
	// [EventListResponseSessionDeletedProperties],
	// [EventListResponseSessionIdleProperties],
	// [EventListResponseSessionErrorProperties], [any]
	Properties EventListResponseUnionProperties `json:"properties"`
	// Any of "installation.updated", "lsp.client.diagnostics", "message.updated",
	// "message.removed", "message.part.updated", "message.part.removed",
	// "permission.updated", "permission.replied", "file.edited", "session.updated",
	// "session.deleted", "session.idle", "session.error", "server.connected".
	Type string `json:"type"`
	JSON struct {
		Properties respjson.Field
		Type       respjson.Field
		raw        string
	} `json:"-"`
}

// anyEventListResponse is implemented by each variant of [EventListResponseUnion]
// to add type safety for the return type of [EventListResponseUnion.AsAny]
type anyEventListResponse interface {
	implEventListResponseUnion()
}

func (EventListResponseInstallationUpdated) implEventListResponseUnion()  {}
func (EventListResponseLspClientDiagnostics) implEventListResponseUnion() {}
func (EventListResponseMessageUpdated) implEventListResponseUnion()       {}
func (EventListResponseMessageRemoved) implEventListResponseUnion()       {}
func (EventListResponseMessagePartUpdated) implEventListResponseUnion()   {}
func (EventListResponseMessagePartRemoved) implEventListResponseUnion()   {}
func (EventListResponsePermissionUpdated) implEventListResponseUnion()    {}
func (EventListResponsePermissionReplied) implEventListResponseUnion()    {}
func (EventListResponseFileEdited) implEventListResponseUnion()           {}
func (EventListResponseSessionUpdated) implEventListResponseUnion()       {}
func (EventListResponseSessionDeleted) implEventListResponseUnion()       {}
func (EventListResponseSessionIdle) implEventListResponseUnion()          {}
func (EventListResponseSessionError) implEventListResponseUnion()         {}
func (EventListResponseServerConnected) implEventListResponseUnion()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := EventListResponseUnion.AsAny().(type) {
//	case opencode.EventListResponseInstallationUpdated:
//	case opencode.EventListResponseLspClientDiagnostics:
//	case opencode.EventListResponseMessageUpdated:
//	case opencode.EventListResponseMessageRemoved:
//	case opencode.EventListResponseMessagePartUpdated:
//	case opencode.EventListResponseMessagePartRemoved:
//	case opencode.EventListResponsePermissionUpdated:
//	case opencode.EventListResponsePermissionReplied:
//	case opencode.EventListResponseFileEdited:
//	case opencode.EventListResponseSessionUpdated:
//	case opencode.EventListResponseSessionDeleted:
//	case opencode.EventListResponseSessionIdle:
//	case opencode.EventListResponseSessionError:
//	case opencode.EventListResponseServerConnected:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u EventListResponseUnion) AsAny() anyEventListResponse {
	switch u.Type {
	case "installation.updated":
		return u.AsInstallationUpdated()
	case "lsp.client.diagnostics":
		return u.AsLspClientDiagnostics()
	case "message.updated":
		return u.AsMessageUpdated()
	case "message.removed":
		return u.AsMessageRemoved()
	case "message.part.updated":
		return u.AsMessagePartUpdated()
	case "message.part.removed":
		return u.AsMessagePartRemoved()
	case "permission.updated":
		return u.AsPermissionUpdated()
	case "permission.replied":
		return u.AsPermissionReplied()
	case "file.edited":
		return u.AsFileEdited()
	case "session.updated":
		return u.AsSessionUpdated()
	case "session.deleted":
		return u.AsSessionDeleted()
	case "session.idle":
		return u.AsSessionIdle()
	case "session.error":
		return u.AsSessionError()
	case "server.connected":
		return u.AsServerConnected()
	}
	return nil
}

func (u EventListResponseUnion) AsInstallationUpdated() (v EventListResponseInstallationUpdated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsLspClientDiagnostics() (v EventListResponseLspClientDiagnostics) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsMessageUpdated() (v EventListResponseMessageUpdated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsMessageRemoved() (v EventListResponseMessageRemoved) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsMessagePartUpdated() (v EventListResponseMessagePartUpdated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsMessagePartRemoved() (v EventListResponseMessagePartRemoved) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsPermissionUpdated() (v EventListResponsePermissionUpdated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsPermissionReplied() (v EventListResponsePermissionReplied) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsFileEdited() (v EventListResponseFileEdited) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsSessionUpdated() (v EventListResponseSessionUpdated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsSessionDeleted() (v EventListResponseSessionDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsSessionIdle() (v EventListResponseSessionIdle) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsSessionError() (v EventListResponseSessionError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseUnion) AsServerConnected() (v EventListResponseServerConnected) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *EventListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventListResponseUnionProperties is an implicit subunion of
// [EventListResponseUnion]. EventListResponseUnionProperties provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [EventListResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfEventListResponseServerConnectedProperties]
type EventListResponseUnionProperties struct {
	// This field will be present if the value is a [any] instead of an object.
	OfEventListResponseServerConnectedProperties any `json:",inline"`
	// This field is from variant [EventListResponseInstallationUpdatedProperties].
	Version string `json:"version"`
	// This field is from variant [EventListResponseLspClientDiagnosticsProperties].
	Path string `json:"path"`
	// This field is from variant [EventListResponseLspClientDiagnosticsProperties].
	ServerID string `json:"serverID"`
	// This field is a union of [MessageUnion], [Session]
	Info      EventListResponseUnionPropertiesInfo `json:"info"`
	MessageID string                               `json:"messageID"`
	SessionID string                               `json:"sessionID"`
	// This field is from variant [EventListResponseMessagePartUpdatedProperties].
	Part PartUnion `json:"part"`
	// This field is from variant [EventListResponseMessagePartRemovedProperties].
	PartID string `json:"partID"`
	// This field is from variant [EventListResponsePermissionUpdatedProperties].
	ID string `json:"id"`
	// This field is from variant [EventListResponsePermissionUpdatedProperties].
	Metadata map[string]any `json:"metadata"`
	// This field is from variant [EventListResponsePermissionUpdatedProperties].
	Time EventListResponsePermissionUpdatedPropertiesTime `json:"time"`
	// This field is from variant [EventListResponsePermissionUpdatedProperties].
	Title string `json:"title"`
	// This field is from variant [EventListResponsePermissionUpdatedProperties].
	Type string `json:"type"`
	// This field is from variant [EventListResponsePermissionUpdatedProperties].
	CallID string `json:"callID"`
	// This field is from variant [EventListResponsePermissionUpdatedProperties].
	Pattern string `json:"pattern"`
	// This field is from variant [EventListResponsePermissionRepliedProperties].
	PermissionID string `json:"permissionID"`
	// This field is from variant [EventListResponsePermissionRepliedProperties].
	Response string `json:"response"`
	// This field is from variant [EventListResponseFileEditedProperties].
	File string `json:"file"`
	// This field is from variant [EventListResponseSessionErrorProperties].
	Error EventListResponseSessionErrorPropertiesErrorUnion `json:"error"`
	JSON  struct {
		OfEventListResponseServerConnectedProperties respjson.Field
		Version                                      respjson.Field
		Path                                         respjson.Field
		ServerID                                     respjson.Field
		Info                                         respjson.Field
		MessageID                                    respjson.Field
		SessionID                                    respjson.Field
		Part                                         respjson.Field
		PartID                                       respjson.Field
		ID                                           respjson.Field
		Metadata                                     respjson.Field
		Time                                         respjson.Field
		Title                                        respjson.Field
		Type                                         respjson.Field
		CallID                                       respjson.Field
		Pattern                                      respjson.Field
		PermissionID                                 respjson.Field
		Response                                     respjson.Field
		File                                         respjson.Field
		Error                                        respjson.Field
		raw                                          string
	} `json:"-"`
}

func (r *EventListResponseUnionProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventListResponseUnionPropertiesInfo is an implicit subunion of
// [EventListResponseUnion]. EventListResponseUnionPropertiesInfo provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [EventListResponseUnion].
type EventListResponseUnionPropertiesInfo struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	SessionID string `json:"sessionID"`
	// This field is a union of [MessageUserTime], [AssistantMessageTime],
	// [SessionTime]
	Time EventListResponseUnionPropertiesInfoTime `json:"time"`
	// This field is from variant [MessageUnion].
	Cost float64 `json:"cost"`
	// This field is from variant [MessageUnion].
	Mode string `json:"mode"`
	// This field is from variant [MessageUnion].
	ModelID string `json:"modelID"`
	// This field is from variant [MessageUnion].
	Path AssistantMessagePath `json:"path"`
	// This field is from variant [MessageUnion].
	ProviderID string `json:"providerID"`
	// This field is from variant [MessageUnion].
	System []string `json:"system"`
	// This field is from variant [MessageUnion].
	Tokens AssistantMessageTokens `json:"tokens"`
	// This field is from variant [MessageUnion].
	Error AssistantMessageErrorUnion `json:"error"`
	// This field is from variant [MessageUnion].
	Summary bool `json:"summary"`
	// This field is from variant [Session].
	Directory string `json:"directory"`
	// This field is from variant [Session].
	ProjectID string `json:"projectID"`
	// This field is from variant [Session].
	Title string `json:"title"`
	// This field is from variant [Session].
	Version string `json:"version"`
	// This field is from variant [Session].
	ParentID string `json:"parentID"`
	// This field is from variant [Session].
	Revert SessionRevert `json:"revert"`
	// This field is from variant [Session].
	Share SessionShare `json:"share"`
	JSON  struct {
		ID         respjson.Field
		Role       respjson.Field
		SessionID  respjson.Field
		Time       respjson.Field
		Cost       respjson.Field
		Mode       respjson.Field
		ModelID    respjson.Field
		Path       respjson.Field
		ProviderID respjson.Field
		System     respjson.Field
		Tokens     respjson.Field
		Error      respjson.Field
		Summary    respjson.Field
		Directory  respjson.Field
		ProjectID  respjson.Field
		Title      respjson.Field
		Version    respjson.Field
		ParentID   respjson.Field
		Revert     respjson.Field
		Share      respjson.Field
		raw        string
	} `json:"-"`
}

func (r *EventListResponseUnionPropertiesInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventListResponseUnionPropertiesInfoTime is an implicit subunion of
// [EventListResponseUnion]. EventListResponseUnionPropertiesInfoTime provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [EventListResponseUnion].
type EventListResponseUnionPropertiesInfoTime struct {
	Created float64 `json:"created"`
	// This field is from variant [AssistantMessageTime].
	Completed float64 `json:"completed"`
	// This field is from variant [SessionTime].
	Updated float64 `json:"updated"`
	JSON    struct {
		Created   respjson.Field
		Completed respjson.Field
		Updated   respjson.Field
		raw       string
	} `json:"-"`
}

func (r *EventListResponseUnionPropertiesInfoTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseInstallationUpdated struct {
	Properties EventListResponseInstallationUpdatedProperties `json:"properties,required"`
	Type       constant.InstallationUpdated                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseInstallationUpdated) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseInstallationUpdated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseInstallationUpdatedProperties struct {
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseInstallationUpdatedProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseInstallationUpdatedProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseLspClientDiagnostics struct {
	Properties EventListResponseLspClientDiagnosticsProperties `json:"properties,required"`
	Type       constant.LspClientDiagnostics                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseLspClientDiagnostics) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseLspClientDiagnostics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseLspClientDiagnosticsProperties struct {
	Path     string `json:"path,required"`
	ServerID string `json:"serverID,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		ServerID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseLspClientDiagnosticsProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseLspClientDiagnosticsProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseMessageUpdated struct {
	Properties EventListResponseMessageUpdatedProperties `json:"properties,required"`
	Type       constant.MessageUpdated                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseMessageUpdated) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseMessageUpdated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseMessageUpdatedProperties struct {
	Info MessageUnion `json:"info,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseMessageUpdatedProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseMessageUpdatedProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseMessageRemoved struct {
	Properties EventListResponseMessageRemovedProperties `json:"properties,required"`
	Type       constant.MessageRemoved                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseMessageRemoved) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseMessageRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseMessageRemovedProperties struct {
	MessageID string `json:"messageID,required"`
	SessionID string `json:"sessionID,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseMessageRemovedProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseMessageRemovedProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseMessagePartUpdated struct {
	Properties EventListResponseMessagePartUpdatedProperties `json:"properties,required"`
	Type       constant.MessagePartUpdated                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseMessagePartUpdated) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseMessagePartUpdated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseMessagePartUpdatedProperties struct {
	Part PartUnion `json:"part,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Part        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseMessagePartUpdatedProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseMessagePartUpdatedProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseMessagePartRemoved struct {
	Properties EventListResponseMessagePartRemovedProperties `json:"properties,required"`
	Type       constant.MessagePartRemoved                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseMessagePartRemoved) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseMessagePartRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseMessagePartRemovedProperties struct {
	MessageID string `json:"messageID,required"`
	PartID    string `json:"partID,required"`
	SessionID string `json:"sessionID,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		PartID      respjson.Field
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseMessagePartRemovedProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseMessagePartRemovedProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponsePermissionUpdated struct {
	Properties EventListResponsePermissionUpdatedProperties `json:"properties,required"`
	Type       constant.PermissionUpdated                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponsePermissionUpdated) RawJSON() string { return r.JSON.raw }
func (r *EventListResponsePermissionUpdated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponsePermissionUpdatedProperties struct {
	ID        string                                           `json:"id,required"`
	MessageID string                                           `json:"messageID,required"`
	Metadata  map[string]any                                   `json:"metadata,required"`
	SessionID string                                           `json:"sessionID,required"`
	Time      EventListResponsePermissionUpdatedPropertiesTime `json:"time,required"`
	Title     string                                           `json:"title,required"`
	Type      string                                           `json:"type,required"`
	CallID    string                                           `json:"callID"`
	Pattern   string                                           `json:"pattern"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		MessageID   respjson.Field
		Metadata    respjson.Field
		SessionID   respjson.Field
		Time        respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		CallID      respjson.Field
		Pattern     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponsePermissionUpdatedProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponsePermissionUpdatedProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponsePermissionUpdatedPropertiesTime struct {
	Created float64 `json:"created,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponsePermissionUpdatedPropertiesTime) RawJSON() string { return r.JSON.raw }
func (r *EventListResponsePermissionUpdatedPropertiesTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponsePermissionReplied struct {
	Properties EventListResponsePermissionRepliedProperties `json:"properties,required"`
	Type       constant.PermissionReplied                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponsePermissionReplied) RawJSON() string { return r.JSON.raw }
func (r *EventListResponsePermissionReplied) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponsePermissionRepliedProperties struct {
	PermissionID string `json:"permissionID,required"`
	Response     string `json:"response,required"`
	SessionID    string `json:"sessionID,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PermissionID respjson.Field
		Response     respjson.Field
		SessionID    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponsePermissionRepliedProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponsePermissionRepliedProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseFileEdited struct {
	Properties EventListResponseFileEditedProperties `json:"properties,required"`
	Type       constant.FileEdited                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseFileEdited) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseFileEdited) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseFileEditedProperties struct {
	File string `json:"file,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		File        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseFileEditedProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseFileEditedProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseSessionUpdated struct {
	Properties EventListResponseSessionUpdatedProperties `json:"properties,required"`
	Type       constant.SessionUpdated                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseSessionUpdated) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseSessionUpdated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseSessionUpdatedProperties struct {
	Info Session `json:"info,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseSessionUpdatedProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseSessionUpdatedProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseSessionDeleted struct {
	Properties EventListResponseSessionDeletedProperties `json:"properties,required"`
	Type       constant.SessionDeleted                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseSessionDeleted) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseSessionDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseSessionDeletedProperties struct {
	Info Session `json:"info,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseSessionDeletedProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseSessionDeletedProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseSessionIdle struct {
	Properties EventListResponseSessionIdleProperties `json:"properties,required"`
	Type       constant.SessionIdle                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseSessionIdle) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseSessionIdle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseSessionIdleProperties struct {
	SessionID string `json:"sessionID,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseSessionIdleProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseSessionIdleProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseSessionError struct {
	Properties EventListResponseSessionErrorProperties `json:"properties,required"`
	Type       constant.SessionError                   `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseSessionError) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseSessionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseSessionErrorProperties struct {
	Error     EventListResponseSessionErrorPropertiesErrorUnion `json:"error"`
	SessionID string                                            `json:"sessionID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseSessionErrorProperties) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseSessionErrorProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventListResponseSessionErrorPropertiesErrorUnion contains all possible
// properties and values from [ProviderAuthError], [UnknownError],
// [MessageOutputLengthError], [MessageAbortedError].
//
// Use the [EventListResponseSessionErrorPropertiesErrorUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EventListResponseSessionErrorPropertiesErrorUnion struct {
	// This field is a union of [ProviderAuthErrorData], [UnknownErrorData], [any],
	// [any]
	Data EventListResponseSessionErrorPropertiesErrorUnionData `json:"data"`
	// Any of "ProviderAuthError", "UnknownError", "MessageOutputLengthError",
	// "MessageAbortedError".
	Name string `json:"name"`
	JSON struct {
		Data respjson.Field
		Name respjson.Field
		raw  string
	} `json:"-"`
}

// anyEventListResponseSessionErrorPropertiesError is implemented by each variant
// of [EventListResponseSessionErrorPropertiesErrorUnion] to add type safety for
// the return type of [EventListResponseSessionErrorPropertiesErrorUnion.AsAny]
type anyEventListResponseSessionErrorPropertiesError interface {
	implEventListResponseSessionErrorPropertiesErrorUnion()
}

func (ProviderAuthError) implEventListResponseSessionErrorPropertiesErrorUnion()        {}
func (UnknownError) implEventListResponseSessionErrorPropertiesErrorUnion()             {}
func (MessageOutputLengthError) implEventListResponseSessionErrorPropertiesErrorUnion() {}
func (MessageAbortedError) implEventListResponseSessionErrorPropertiesErrorUnion()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := EventListResponseSessionErrorPropertiesErrorUnion.AsAny().(type) {
//	case opencode.ProviderAuthError:
//	case opencode.UnknownError:
//	case opencode.MessageOutputLengthError:
//	case opencode.MessageAbortedError:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u EventListResponseSessionErrorPropertiesErrorUnion) AsAny() anyEventListResponseSessionErrorPropertiesError {
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

func (u EventListResponseSessionErrorPropertiesErrorUnion) AsProviderAuthError() (v ProviderAuthError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseSessionErrorPropertiesErrorUnion) AsUnknownError() (v UnknownError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseSessionErrorPropertiesErrorUnion) AsMessageOutputLengthError() (v MessageOutputLengthError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventListResponseSessionErrorPropertiesErrorUnion) AsMessageAbortedError() (v MessageAbortedError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventListResponseSessionErrorPropertiesErrorUnion) RawJSON() string { return u.JSON.raw }

func (r *EventListResponseSessionErrorPropertiesErrorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventListResponseSessionErrorPropertiesErrorUnionData is an implicit subunion of
// [EventListResponseSessionErrorPropertiesErrorUnion].
// EventListResponseSessionErrorPropertiesErrorUnionData provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [EventListResponseSessionErrorPropertiesErrorUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMessageAbortedErrorData]
type EventListResponseSessionErrorPropertiesErrorUnionData struct {
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

func (r *EventListResponseSessionErrorPropertiesErrorUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListResponseServerConnected struct {
	Properties any                      `json:"properties,required"`
	Type       constant.ServerConnected `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponseServerConnected) RawJSON() string { return r.JSON.raw }
func (r *EventListResponseServerConnected) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
