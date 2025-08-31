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

// SessionMessageService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionMessageService] method instead.
type SessionMessageService struct {
	Options []option.RequestOption
}

// NewSessionMessageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSessionMessageService(opts ...option.RequestOption) (r SessionMessageService) {
	r = SessionMessageService{}
	r.Options = opts
	return
}

// Create and send a new message to a session
func (r *SessionMessageService) New(ctx context.Context, id string, body SessionMessageNewParams, opts ...option.RequestOption) (res *SessionMessageNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/message", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a message from a session
func (r *SessionMessageService) Get(ctx context.Context, messageID string, query SessionMessageGetParams, opts ...option.RequestOption) (res *SessionMessageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return
	}
	path := fmt.Sprintf("session/%s/message/%s", query.ID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List messages for a session
func (r *SessionMessageService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *[]SessionMessageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/message", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// FilePartSourceUnion contains all possible properties and values from
// [FilePartSourceFile], [FilePartSourceSymbol].
//
// Use the [FilePartSourceUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FilePartSourceUnion struct {
	Path string `json:"path"`
	// This field is from variant [FilePartSourceFile].
	Text FilePartSourceText `json:"text"`
	// Any of "file", "symbol".
	Type string `json:"type"`
	// This field is from variant [FilePartSourceSymbol].
	Kind int64 `json:"kind"`
	// This field is from variant [FilePartSourceSymbol].
	Name string `json:"name"`
	// This field is from variant [FilePartSourceSymbol].
	Range Range `json:"range"`
	JSON  struct {
		Path  respjson.Field
		Text  respjson.Field
		Type  respjson.Field
		Kind  respjson.Field
		Name  respjson.Field
		Range respjson.Field
		raw   string
	} `json:"-"`
}

// anyFilePartSource is implemented by each variant of [FilePartSourceUnion] to add
// type safety for the return type of [FilePartSourceUnion.AsAny]
type anyFilePartSource interface {
	implFilePartSourceUnion()
}

func (FilePartSourceFile) implFilePartSourceUnion()   {}
func (FilePartSourceSymbol) implFilePartSourceUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := FilePartSourceUnion.AsAny().(type) {
//	case opencode.FilePartSourceFile:
//	case opencode.FilePartSourceSymbol:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FilePartSourceUnion) AsAny() anyFilePartSource {
	switch u.Type {
	case "file":
		return u.AsFile()
	case "symbol":
		return u.AsSymbol()
	}
	return nil
}

func (u FilePartSourceUnion) AsFile() (v FilePartSourceFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FilePartSourceUnion) AsSymbol() (v FilePartSourceSymbol) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FilePartSourceUnion) RawJSON() string { return u.JSON.raw }

func (r *FilePartSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FilePartSourceUnion to a FilePartSourceUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FilePartSourceUnionParam.Overrides()
func (r FilePartSourceUnion) ToParam() FilePartSourceUnionParam {
	return param.Override[FilePartSourceUnionParam](json.RawMessage(r.RawJSON()))
}

type FilePartSourceFile struct {
	Path string             `json:"path,required"`
	Text FilePartSourceText `json:"text,required"`
	Type constant.File      `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilePartSourceFile) RawJSON() string { return r.JSON.raw }
func (r *FilePartSourceFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilePartSourceSymbol struct {
	Kind  int64              `json:"kind,required"`
	Name  string             `json:"name,required"`
	Path  string             `json:"path,required"`
	Range Range              `json:"range,required"`
	Text  FilePartSourceText `json:"text,required"`
	Type  constant.Symbol    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Name        respjson.Field
		Path        respjson.Field
		Range       respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilePartSourceSymbol) RawJSON() string { return r.JSON.raw }
func (r *FilePartSourceSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func FilePartSourceParamOfFile(path string, text FilePartSourceTextParam) FilePartSourceUnionParam {
	var file FilePartSourceFileParam
	file.Path = path
	file.Text = text
	return FilePartSourceUnionParam{OfFile: &file}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FilePartSourceUnionParam struct {
	OfFile   *FilePartSourceFileParam   `json:",omitzero,inline"`
	OfSymbol *FilePartSourceSymbolParam `json:",omitzero,inline"`
	paramUnion
}

func (u FilePartSourceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFile, u.OfSymbol)
}
func (u *FilePartSourceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FilePartSourceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFile) {
		return u.OfFile
	} else if !param.IsOmitted(u.OfSymbol) {
		return u.OfSymbol
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilePartSourceUnionParam) GetKind() *int64 {
	if vt := u.OfSymbol; vt != nil {
		return &vt.Kind
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilePartSourceUnionParam) GetName() *string {
	if vt := u.OfSymbol; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilePartSourceUnionParam) GetRange() *RangeParam {
	if vt := u.OfSymbol; vt != nil {
		return &vt.Range
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilePartSourceUnionParam) GetPath() *string {
	if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Path)
	} else if vt := u.OfSymbol; vt != nil {
		return (*string)(&vt.Path)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilePartSourceUnionParam) GetType() *string {
	if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSymbol; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's Text property, if present.
func (u FilePartSourceUnionParam) GetText() *FilePartSourceTextParam {
	if vt := u.OfFile; vt != nil {
		return &vt.Text
	} else if vt := u.OfSymbol; vt != nil {
		return &vt.Text
	}
	return nil
}

func init() {
	apijson.RegisterUnion[FilePartSourceUnionParam](
		"type",
		apijson.Discriminator[FilePartSourceFileParam]("file"),
		apijson.Discriminator[FilePartSourceSymbolParam]("symbol"),
	)
}

// The properties Path, Text, Type are required.
type FilePartSourceFileParam struct {
	Path string                  `json:"path,required"`
	Text FilePartSourceTextParam `json:"text,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "file".
	Type constant.File `json:"type,required"`
	paramObj
}

func (r FilePartSourceFileParam) MarshalJSON() (data []byte, err error) {
	type shadow FilePartSourceFileParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilePartSourceFileParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Kind, Name, Path, Range, Text, Type are required.
type FilePartSourceSymbolParam struct {
	Kind  int64                   `json:"kind,required"`
	Name  string                  `json:"name,required"`
	Path  string                  `json:"path,required"`
	Range RangeParam              `json:"range,omitzero,required"`
	Text  FilePartSourceTextParam `json:"text,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "symbol".
	Type constant.Symbol `json:"type,required"`
	paramObj
}

func (r FilePartSourceSymbolParam) MarshalJSON() (data []byte, err error) {
	type shadow FilePartSourceSymbolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilePartSourceSymbolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilePartSourceText struct {
	End   int64  `json:"end,required"`
	Start int64  `json:"start,required"`
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilePartSourceText) RawJSON() string { return r.JSON.raw }
func (r *FilePartSourceText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FilePartSourceText to a FilePartSourceTextParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FilePartSourceTextParam.Overrides()
func (r FilePartSourceText) ToParam() FilePartSourceTextParam {
	return param.Override[FilePartSourceTextParam](json.RawMessage(r.RawJSON()))
}

// The properties End, Start, Value are required.
type FilePartSourceTextParam struct {
	End   int64  `json:"end,required"`
	Start int64  `json:"start,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r FilePartSourceTextParam) MarshalJSON() (data []byte, err error) {
	type shadow FilePartSourceTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilePartSourceTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageUnion contains all possible properties and values from [MessageUser],
// [AssistantMessage].
//
// Use the [MessageUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessageUnion struct {
	ID string `json:"id"`
	// Any of "user", "assistant".
	Role      string `json:"role"`
	SessionID string `json:"sessionID"`
	// This field is a union of [MessageUserTime], [AssistantMessageTime]
	Time MessageUnionTime `json:"time"`
	// This field is from variant [AssistantMessage].
	Cost float64 `json:"cost"`
	// This field is from variant [AssistantMessage].
	Mode string `json:"mode"`
	// This field is from variant [AssistantMessage].
	ModelID string `json:"modelID"`
	// This field is from variant [AssistantMessage].
	Path AssistantMessagePath `json:"path"`
	// This field is from variant [AssistantMessage].
	ProviderID string `json:"providerID"`
	// This field is from variant [AssistantMessage].
	System []string `json:"system"`
	// This field is from variant [AssistantMessage].
	Tokens AssistantMessageTokens `json:"tokens"`
	// This field is from variant [AssistantMessage].
	Error AssistantMessageErrorUnion `json:"error"`
	// This field is from variant [AssistantMessage].
	Summary bool `json:"summary"`
	JSON    struct {
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
		raw        string
	} `json:"-"`
}

// anyMessage is implemented by each variant of [MessageUnion] to add type safety
// for the return type of [MessageUnion.AsAny]
type anyMessage interface {
	implMessageUnion()
}

func (MessageUser) implMessageUnion()      {}
func (AssistantMessage) implMessageUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MessageUnion.AsAny().(type) {
//	case opencode.MessageUser:
//	case opencode.AssistantMessage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MessageUnion) AsAny() anyMessage {
	switch u.Role {
	case "user":
		return u.AsUser()
	case "assistant":
		return u.AsAssistant()
	}
	return nil
}

func (u MessageUnion) AsUser() (v MessageUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageUnion) AsAssistant() (v AssistantMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageUnion) RawJSON() string { return u.JSON.raw }

func (r *MessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageUnionTime is an implicit subunion of [MessageUnion]. MessageUnionTime
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageUnion].
type MessageUnionTime struct {
	Created float64 `json:"created"`
	// This field is from variant [AssistantMessageTime].
	Completed float64 `json:"completed"`
	JSON      struct {
		Created   respjson.Field
		Completed respjson.Field
		raw       string
	} `json:"-"`
}

func (r *MessageUnionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageUser struct {
	ID        string          `json:"id,required"`
	Role      constant.User   `json:"role,required"`
	SessionID string          `json:"sessionID,required"`
	Time      MessageUserTime `json:"time,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Role        respjson.Field
		SessionID   respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageUser) RawJSON() string { return r.JSON.raw }
func (r *MessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageUserTime struct {
	Created float64 `json:"created,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageUserTime) RawJSON() string { return r.JSON.raw }
func (r *MessageUserTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PartUnion contains all possible properties and values from [PartText],
// [PartReasoning], [PartFile], [PartTool], [PartStepStart], [PartStepFinish],
// [PartSnapshot], [PartPatch], [PartAgent].
//
// Use the [PartUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PartUnion struct {
	ID        string `json:"id"`
	MessageID string `json:"messageID"`
	SessionID string `json:"sessionID"`
	Text      string `json:"text"`
	// Any of "text", "reasoning", "file", "tool", "step-start", "step-finish",
	// "snapshot", "patch", "agent".
	Type string `json:"type"`
	// This field is from variant [PartText].
	Synthetic bool `json:"synthetic"`
	// This field is a union of [PartTextTime], [PartReasoningTime]
	Time PartUnionTime `json:"time"`
	// This field is from variant [PartReasoning].
	Metadata map[string]any `json:"metadata"`
	// This field is from variant [PartFile].
	Mime string `json:"mime"`
	// This field is from variant [PartFile].
	URL string `json:"url"`
	// This field is from variant [PartFile].
	Filename string `json:"filename"`
	// This field is a union of [FilePartSourceUnion], [PartAgentSource]
	Source PartUnionSource `json:"source"`
	// This field is from variant [PartTool].
	CallID string `json:"callID"`
	// This field is from variant [PartTool].
	State PartToolStateUnion `json:"state"`
	// This field is from variant [PartTool].
	Tool string `json:"tool"`
	// This field is from variant [PartStepFinish].
	Cost float64 `json:"cost"`
	// This field is from variant [PartStepFinish].
	Tokens PartStepFinishTokens `json:"tokens"`
	// This field is from variant [PartSnapshot].
	Snapshot string `json:"snapshot"`
	// This field is from variant [PartPatch].
	Files []string `json:"files"`
	// This field is from variant [PartPatch].
	Hash string `json:"hash"`
	// This field is from variant [PartAgent].
	Name string `json:"name"`
	JSON struct {
		ID        respjson.Field
		MessageID respjson.Field
		SessionID respjson.Field
		Text      respjson.Field
		Type      respjson.Field
		Synthetic respjson.Field
		Time      respjson.Field
		Metadata  respjson.Field
		Mime      respjson.Field
		URL       respjson.Field
		Filename  respjson.Field
		Source    respjson.Field
		CallID    respjson.Field
		State     respjson.Field
		Tool      respjson.Field
		Cost      respjson.Field
		Tokens    respjson.Field
		Snapshot  respjson.Field
		Files     respjson.Field
		Hash      respjson.Field
		Name      respjson.Field
		raw       string
	} `json:"-"`
}

// anyPart is implemented by each variant of [PartUnion] to add type safety for the
// return type of [PartUnion.AsAny]
type anyPart interface {
	implPartUnion()
}

func (PartText) implPartUnion()       {}
func (PartReasoning) implPartUnion()  {}
func (PartFile) implPartUnion()       {}
func (PartTool) implPartUnion()       {}
func (PartStepStart) implPartUnion()  {}
func (PartStepFinish) implPartUnion() {}
func (PartSnapshot) implPartUnion()   {}
func (PartPatch) implPartUnion()      {}
func (PartAgent) implPartUnion()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PartUnion.AsAny().(type) {
//	case opencode.PartText:
//	case opencode.PartReasoning:
//	case opencode.PartFile:
//	case opencode.PartTool:
//	case opencode.PartStepStart:
//	case opencode.PartStepFinish:
//	case opencode.PartSnapshot:
//	case opencode.PartPatch:
//	case opencode.PartAgent:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PartUnion) AsAny() anyPart {
	switch u.Type {
	case "text":
		return u.AsText()
	case "reasoning":
		return u.AsReasoning()
	case "file":
		return u.AsFile()
	case "tool":
		return u.AsTool()
	case "step-start":
		return u.AsStepStart()
	case "step-finish":
		return u.AsStepFinish()
	case "snapshot":
		return u.AsSnapshot()
	case "patch":
		return u.AsPatch()
	case "agent":
		return u.AsAgent()
	}
	return nil
}

func (u PartUnion) AsText() (v PartText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartUnion) AsReasoning() (v PartReasoning) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartUnion) AsFile() (v PartFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartUnion) AsTool() (v PartTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartUnion) AsStepStart() (v PartStepStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartUnion) AsStepFinish() (v PartStepFinish) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartUnion) AsSnapshot() (v PartSnapshot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartUnion) AsPatch() (v PartPatch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartUnion) AsAgent() (v PartAgent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PartUnion) RawJSON() string { return u.JSON.raw }

func (r *PartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PartUnionTime is an implicit subunion of [PartUnion]. PartUnionTime provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the [PartUnion].
type PartUnionTime struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
	JSON  struct {
		Start respjson.Field
		End   respjson.Field
		raw   string
	} `json:"-"`
}

func (r *PartUnionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PartUnionSource is an implicit subunion of [PartUnion]. PartUnionSource provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the [PartUnion].
type PartUnionSource struct {
	Path string `json:"path"`
	// This field is from variant [FilePartSourceUnion].
	Text FilePartSourceText `json:"text"`
	Type string             `json:"type"`
	// This field is from variant [FilePartSourceUnion].
	Kind int64 `json:"kind"`
	// This field is from variant [FilePartSourceUnion].
	Name string `json:"name"`
	// This field is from variant [FilePartSourceUnion].
	Range Range `json:"range"`
	// This field is from variant [PartAgentSource].
	End int64 `json:"end"`
	// This field is from variant [PartAgentSource].
	Start int64 `json:"start"`
	// This field is from variant [PartAgentSource].
	Value string `json:"value"`
	JSON  struct {
		Path  respjson.Field
		Text  respjson.Field
		Type  respjson.Field
		Kind  respjson.Field
		Name  respjson.Field
		Range respjson.Field
		End   respjson.Field
		Start respjson.Field
		Value respjson.Field
		raw   string
	} `json:"-"`
}

func (r *PartUnionSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartText struct {
	ID        string        `json:"id,required"`
	MessageID string        `json:"messageID,required"`
	SessionID string        `json:"sessionID,required"`
	Text      string        `json:"text,required"`
	Type      constant.Text `json:"type,required"`
	Synthetic bool          `json:"synthetic"`
	Time      PartTextTime  `json:"time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		MessageID   respjson.Field
		SessionID   respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		Synthetic   respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartText) RawJSON() string { return r.JSON.raw }
func (r *PartText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartTextTime struct {
	Start float64 `json:"start,required"`
	End   float64 `json:"end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Start       respjson.Field
		End         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartTextTime) RawJSON() string { return r.JSON.raw }
func (r *PartTextTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartReasoning struct {
	ID        string             `json:"id,required"`
	MessageID string             `json:"messageID,required"`
	SessionID string             `json:"sessionID,required"`
	Text      string             `json:"text,required"`
	Time      PartReasoningTime  `json:"time,required"`
	Type      constant.Reasoning `json:"type,required"`
	Metadata  map[string]any     `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		MessageID   respjson.Field
		SessionID   respjson.Field
		Text        respjson.Field
		Time        respjson.Field
		Type        respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartReasoning) RawJSON() string { return r.JSON.raw }
func (r *PartReasoning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartReasoningTime struct {
	Start float64 `json:"start,required"`
	End   float64 `json:"end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Start       respjson.Field
		End         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartReasoningTime) RawJSON() string { return r.JSON.raw }
func (r *PartReasoningTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartFile struct {
	ID        string              `json:"id,required"`
	MessageID string              `json:"messageID,required"`
	Mime      string              `json:"mime,required"`
	SessionID string              `json:"sessionID,required"`
	Type      constant.File       `json:"type,required"`
	URL       string              `json:"url,required"`
	Filename  string              `json:"filename"`
	Source    FilePartSourceUnion `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		MessageID   respjson.Field
		Mime        respjson.Field
		SessionID   respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		Filename    respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartFile) RawJSON() string { return r.JSON.raw }
func (r *PartFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartTool struct {
	ID        string             `json:"id,required"`
	CallID    string             `json:"callID,required"`
	MessageID string             `json:"messageID,required"`
	SessionID string             `json:"sessionID,required"`
	State     PartToolStateUnion `json:"state,required"`
	Tool      string             `json:"tool,required"`
	Type      constant.Tool      `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CallID      respjson.Field
		MessageID   respjson.Field
		SessionID   respjson.Field
		State       respjson.Field
		Tool        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartTool) RawJSON() string { return r.JSON.raw }
func (r *PartTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PartToolStateUnion contains all possible properties and values from
// [PartToolStatePending], [PartToolStateRunning], [PartToolStateCompleted],
// [PartToolStateError].
//
// Use the [PartToolStateUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PartToolStateUnion struct {
	// Any of "pending", "running", "completed", "error".
	Status string `json:"status"`
	// This field is a union of [PartToolStateRunningTime],
	// [PartToolStateCompletedTime], [PartToolStateErrorTime]
	Time     PartToolStateUnionTime `json:"time"`
	Input    any                    `json:"input"`
	Metadata any                    `json:"metadata"`
	Title    string                 `json:"title"`
	// This field is from variant [PartToolStateCompleted].
	Output string `json:"output"`
	// This field is from variant [PartToolStateError].
	Error string `json:"error"`
	JSON  struct {
		Status   respjson.Field
		Time     respjson.Field
		Input    respjson.Field
		Metadata respjson.Field
		Title    respjson.Field
		Output   respjson.Field
		Error    respjson.Field
		raw      string
	} `json:"-"`
}

// anyPartToolState is implemented by each variant of [PartToolStateUnion] to add
// type safety for the return type of [PartToolStateUnion.AsAny]
type anyPartToolState interface {
	implPartToolStateUnion()
}

func (PartToolStatePending) implPartToolStateUnion()   {}
func (PartToolStateRunning) implPartToolStateUnion()   {}
func (PartToolStateCompleted) implPartToolStateUnion() {}
func (PartToolStateError) implPartToolStateUnion()     {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PartToolStateUnion.AsAny().(type) {
//	case opencode.PartToolStatePending:
//	case opencode.PartToolStateRunning:
//	case opencode.PartToolStateCompleted:
//	case opencode.PartToolStateError:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PartToolStateUnion) AsAny() anyPartToolState {
	switch u.Status {
	case "pending":
		return u.AsPending()
	case "running":
		return u.AsRunning()
	case "completed":
		return u.AsCompleted()
	case "error":
		return u.AsError()
	}
	return nil
}

func (u PartToolStateUnion) AsPending() (v PartToolStatePending) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartToolStateUnion) AsRunning() (v PartToolStateRunning) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartToolStateUnion) AsCompleted() (v PartToolStateCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartToolStateUnion) AsError() (v PartToolStateError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PartToolStateUnion) RawJSON() string { return u.JSON.raw }

func (r *PartToolStateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PartToolStateUnionTime is an implicit subunion of [PartToolStateUnion].
// PartToolStateUnionTime provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [PartToolStateUnion].
type PartToolStateUnionTime struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
	JSON  struct {
		Start respjson.Field
		End   respjson.Field
		raw   string
	} `json:"-"`
}

func (r *PartToolStateUnionTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartToolStatePending struct {
	Status constant.Pending `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartToolStatePending) RawJSON() string { return r.JSON.raw }
func (r *PartToolStatePending) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartToolStateRunning struct {
	Status   constant.Running         `json:"status,required"`
	Time     PartToolStateRunningTime `json:"time,required"`
	Input    any                      `json:"input"`
	Metadata map[string]any           `json:"metadata"`
	Title    string                   `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Time        respjson.Field
		Input       respjson.Field
		Metadata    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartToolStateRunning) RawJSON() string { return r.JSON.raw }
func (r *PartToolStateRunning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartToolStateRunningTime struct {
	Start float64 `json:"start,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartToolStateRunningTime) RawJSON() string { return r.JSON.raw }
func (r *PartToolStateRunningTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartToolStateCompleted struct {
	Input    map[string]any             `json:"input,required"`
	Metadata map[string]any             `json:"metadata,required"`
	Output   string                     `json:"output,required"`
	Status   constant.Completed         `json:"status,required"`
	Time     PartToolStateCompletedTime `json:"time,required"`
	Title    string                     `json:"title,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input       respjson.Field
		Metadata    respjson.Field
		Output      respjson.Field
		Status      respjson.Field
		Time        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartToolStateCompleted) RawJSON() string { return r.JSON.raw }
func (r *PartToolStateCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartToolStateCompletedTime struct {
	End   float64 `json:"end,required"`
	Start float64 `json:"start,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartToolStateCompletedTime) RawJSON() string { return r.JSON.raw }
func (r *PartToolStateCompletedTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartToolStateError struct {
	Error    string                 `json:"error,required"`
	Input    map[string]any         `json:"input,required"`
	Status   constant.Error         `json:"status,required"`
	Time     PartToolStateErrorTime `json:"time,required"`
	Metadata map[string]any         `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Input       respjson.Field
		Status      respjson.Field
		Time        respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartToolStateError) RawJSON() string { return r.JSON.raw }
func (r *PartToolStateError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartToolStateErrorTime struct {
	End   float64 `json:"end,required"`
	Start float64 `json:"start,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartToolStateErrorTime) RawJSON() string { return r.JSON.raw }
func (r *PartToolStateErrorTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartStepStart struct {
	ID        string             `json:"id,required"`
	MessageID string             `json:"messageID,required"`
	SessionID string             `json:"sessionID,required"`
	Type      constant.StepStart `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		MessageID   respjson.Field
		SessionID   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartStepStart) RawJSON() string { return r.JSON.raw }
func (r *PartStepStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartStepFinish struct {
	ID        string               `json:"id,required"`
	Cost      float64              `json:"cost,required"`
	MessageID string               `json:"messageID,required"`
	SessionID string               `json:"sessionID,required"`
	Tokens    PartStepFinishTokens `json:"tokens,required"`
	Type      constant.StepFinish  `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cost        respjson.Field
		MessageID   respjson.Field
		SessionID   respjson.Field
		Tokens      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartStepFinish) RawJSON() string { return r.JSON.raw }
func (r *PartStepFinish) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartStepFinishTokens struct {
	Cache     PartStepFinishTokensCache `json:"cache,required"`
	Input     float64                   `json:"input,required"`
	Output    float64                   `json:"output,required"`
	Reasoning float64                   `json:"reasoning,required"`
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
func (r PartStepFinishTokens) RawJSON() string { return r.JSON.raw }
func (r *PartStepFinishTokens) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartStepFinishTokensCache struct {
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
func (r PartStepFinishTokensCache) RawJSON() string { return r.JSON.raw }
func (r *PartStepFinishTokensCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartSnapshot struct {
	ID        string            `json:"id,required"`
	MessageID string            `json:"messageID,required"`
	SessionID string            `json:"sessionID,required"`
	Snapshot  string            `json:"snapshot,required"`
	Type      constant.Snapshot `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		MessageID   respjson.Field
		SessionID   respjson.Field
		Snapshot    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartSnapshot) RawJSON() string { return r.JSON.raw }
func (r *PartSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartPatch struct {
	ID        string         `json:"id,required"`
	Files     []string       `json:"files,required"`
	Hash      string         `json:"hash,required"`
	MessageID string         `json:"messageID,required"`
	SessionID string         `json:"sessionID,required"`
	Type      constant.Patch `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Files       respjson.Field
		Hash        respjson.Field
		MessageID   respjson.Field
		SessionID   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartPatch) RawJSON() string { return r.JSON.raw }
func (r *PartPatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartAgent struct {
	ID        string          `json:"id,required"`
	MessageID string          `json:"messageID,required"`
	Name      string          `json:"name,required"`
	SessionID string          `json:"sessionID,required"`
	Type      constant.Agent  `json:"type,required"`
	Source    PartAgentSource `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		MessageID   respjson.Field
		Name        respjson.Field
		SessionID   respjson.Field
		Type        respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartAgent) RawJSON() string { return r.JSON.raw }
func (r *PartAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartAgentSource struct {
	End   int64  `json:"end,required"`
	Start int64  `json:"start,required"`
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartAgentSource) RawJSON() string { return r.JSON.raw }
func (r *PartAgentSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionMessageNewResponse struct {
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
func (r SessionMessageNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionMessageNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionMessageGetResponse struct {
	Info  MessageUnion `json:"info,required"`
	Parts []PartUnion  `json:"parts,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		Parts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionMessageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionMessageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionMessageListResponse struct {
	Info  MessageUnion `json:"info,required"`
	Parts []PartUnion  `json:"parts,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		Parts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionMessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionMessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionMessageNewParams struct {
	ModelID    string                             `json:"modelID,required"`
	Parts      []SessionMessageNewParamsPartUnion `json:"parts,omitzero,required"`
	ProviderID string                             `json:"providerID,required"`
	Agent      param.Opt[string]                  `json:"agent,omitzero"`
	MessageID  param.Opt[string]                  `json:"messageID,omitzero"`
	System     param.Opt[string]                  `json:"system,omitzero"`
	Tools      map[string]bool                    `json:"tools,omitzero"`
	paramObj
}

func (r SessionMessageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionMessageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionMessageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SessionMessageNewParamsPartUnion struct {
	OfText  *SessionMessageNewParamsPartText  `json:",omitzero,inline"`
	OfFile  *SessionMessageNewParamsPartFile  `json:",omitzero,inline"`
	OfAgent *SessionMessageNewParamsPartAgent `json:",omitzero,inline"`
	paramUnion
}

func (u SessionMessageNewParamsPartUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfFile, u.OfAgent)
}
func (u *SessionMessageNewParamsPartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SessionMessageNewParamsPartUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfFile) {
		return u.OfFile
	} else if !param.IsOmitted(u.OfAgent) {
		return u.OfAgent
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionMessageNewParamsPartUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionMessageNewParamsPartUnion) GetSynthetic() *bool {
	if vt := u.OfText; vt != nil && vt.Synthetic.Valid() {
		return &vt.Synthetic.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionMessageNewParamsPartUnion) GetTime() *SessionMessageNewParamsPartTextTime {
	if vt := u.OfText; vt != nil {
		return &vt.Time
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionMessageNewParamsPartUnion) GetMime() *string {
	if vt := u.OfFile; vt != nil {
		return &vt.Mime
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionMessageNewParamsPartUnion) GetURL() *string {
	if vt := u.OfFile; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionMessageNewParamsPartUnion) GetFilename() *string {
	if vt := u.OfFile; vt != nil && vt.Filename.Valid() {
		return &vt.Filename.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionMessageNewParamsPartUnion) GetName() *string {
	if vt := u.OfAgent; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionMessageNewParamsPartUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAgent; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionMessageNewParamsPartUnion) GetID() *string {
	if vt := u.OfText; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfFile; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfAgent; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u SessionMessageNewParamsPartUnion) GetSource() (res sessionMessageNewParamsPartUnionSource) {
	if vt := u.OfFile; vt != nil {
		res.any = vt.Source.asAny()
	} else if vt := u.OfAgent; vt != nil {
		res.any = &vt.Source
	}
	return
}

// Can have the runtime types [*FilePartSourceFileParam],
// [*FilePartSourceSymbolParam], [*SessionMessageNewParamsPartAgentSource]
type sessionMessageNewParamsPartUnionSource struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *opencode.FilePartSourceFileParam:
//	case *opencode.FilePartSourceSymbolParam:
//	case *opencode.SessionMessageNewParamsPartAgentSource:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u sessionMessageNewParamsPartUnionSource) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u sessionMessageNewParamsPartUnionSource) GetKind() *int64 {
	switch vt := u.any.(type) {
	case *FilePartSourceUnionParam:
		return vt.GetKind()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u sessionMessageNewParamsPartUnionSource) GetName() *string {
	switch vt := u.any.(type) {
	case *FilePartSourceUnionParam:
		return vt.GetName()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u sessionMessageNewParamsPartUnionSource) GetRange() *RangeParam {
	switch vt := u.any.(type) {
	case *FilePartSourceUnionParam:
		return vt.GetRange()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u sessionMessageNewParamsPartUnionSource) GetEnd() *int64 {
	switch vt := u.any.(type) {
	case *SessionMessageNewParamsPartAgentSource:
		return &vt.End
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u sessionMessageNewParamsPartUnionSource) GetStart() *int64 {
	switch vt := u.any.(type) {
	case *SessionMessageNewParamsPartAgentSource:
		return &vt.Start
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u sessionMessageNewParamsPartUnionSource) GetValue() *string {
	switch vt := u.any.(type) {
	case *SessionMessageNewParamsPartAgentSource:
		return &vt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u sessionMessageNewParamsPartUnionSource) GetPath() *string {
	switch vt := u.any.(type) {
	case *FilePartSourceUnionParam:
		return vt.GetPath()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u sessionMessageNewParamsPartUnionSource) GetType() *string {
	switch vt := u.any.(type) {
	case *FilePartSourceUnionParam:
		return vt.GetType()
	}
	return nil
}

// Returns a pointer to the underlying variant's Text property, if present.
func (u sessionMessageNewParamsPartUnionSource) GetText() *FilePartSourceTextParam {
	switch vt := u.any.(type) {
	case *FilePartSourceUnionParam:
		return vt.GetText()
	}
	return nil
}

func init() {
	apijson.RegisterUnion[SessionMessageNewParamsPartUnion](
		"type",
		apijson.Discriminator[SessionMessageNewParamsPartText]("text"),
		apijson.Discriminator[SessionMessageNewParamsPartFile]("file"),
		apijson.Discriminator[SessionMessageNewParamsPartAgent]("agent"),
	)
}

// The properties Text, Type are required.
type SessionMessageNewParamsPartText struct {
	Text      string                              `json:"text,required"`
	ID        param.Opt[string]                   `json:"id,omitzero"`
	Synthetic param.Opt[bool]                     `json:"synthetic,omitzero"`
	Time      SessionMessageNewParamsPartTextTime `json:"time,omitzero"`
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r SessionMessageNewParamsPartText) MarshalJSON() (data []byte, err error) {
	type shadow SessionMessageNewParamsPartText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionMessageNewParamsPartText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Start is required.
type SessionMessageNewParamsPartTextTime struct {
	Start float64            `json:"start,required"`
	End   param.Opt[float64] `json:"end,omitzero"`
	paramObj
}

func (r SessionMessageNewParamsPartTextTime) MarshalJSON() (data []byte, err error) {
	type shadow SessionMessageNewParamsPartTextTime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionMessageNewParamsPartTextTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Mime, Type, URL are required.
type SessionMessageNewParamsPartFile struct {
	Mime     string                   `json:"mime,required"`
	URL      string                   `json:"url,required"`
	ID       param.Opt[string]        `json:"id,omitzero"`
	Filename param.Opt[string]        `json:"filename,omitzero"`
	Source   FilePartSourceUnionParam `json:"source,omitzero"`
	// This field can be elided, and will marshal its zero value as "file".
	Type constant.File `json:"type,required"`
	paramObj
}

func (r SessionMessageNewParamsPartFile) MarshalJSON() (data []byte, err error) {
	type shadow SessionMessageNewParamsPartFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionMessageNewParamsPartFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type SessionMessageNewParamsPartAgent struct {
	Name   string                                 `json:"name,required"`
	ID     param.Opt[string]                      `json:"id,omitzero"`
	Source SessionMessageNewParamsPartAgentSource `json:"source,omitzero"`
	// This field can be elided, and will marshal its zero value as "agent".
	Type constant.Agent `json:"type,required"`
	paramObj
}

func (r SessionMessageNewParamsPartAgent) MarshalJSON() (data []byte, err error) {
	type shadow SessionMessageNewParamsPartAgent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionMessageNewParamsPartAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties End, Start, Value are required.
type SessionMessageNewParamsPartAgentSource struct {
	End   int64  `json:"end,required"`
	Start int64  `json:"start,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r SessionMessageNewParamsPartAgentSource) MarshalJSON() (data []byte, err error) {
	type shadow SessionMessageNewParamsPartAgentSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionMessageNewParamsPartAgentSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionMessageGetParams struct {
	// Session ID
	ID string `path:"id,required" json:"-"`
	paramObj
}
