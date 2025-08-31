// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/opencode-go/internal/apijson"
	"github.com/stainless-sdks/opencode-go/internal/apiquery"
	"github.com/stainless-sdks/opencode-go/internal/requestconfig"
	"github.com/stainless-sdks/opencode-go/option"
	"github.com/stainless-sdks/opencode-go/packages/respjson"
)

// FileService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.Options = opts
	return
}

// List files and directories
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *[]FileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "file"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get file status
func (r *FileService) GetStatus(ctx context.Context, opts ...option.RequestOption) (res *[]FileGetStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "file/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Read a file
func (r *FileService) Read(ctx context.Context, query FileReadParams, opts ...option.RequestOption) (res *FileReadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "file/content"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FileListResponse struct {
	Ignored bool   `json:"ignored,required"`
	Name    string `json:"name,required"`
	Path    string `json:"path,required"`
	// Any of "file", "directory".
	Type FileListResponseType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ignored     respjson.Field
		Name        respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileListResponse) RawJSON() string { return r.JSON.raw }
func (r *FileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListResponseType string

const (
	FileListResponseTypeFile      FileListResponseType = "file"
	FileListResponseTypeDirectory FileListResponseType = "directory"
)

type FileGetStatusResponse struct {
	Added   int64  `json:"added,required"`
	Path    string `json:"path,required"`
	Removed int64  `json:"removed,required"`
	// Any of "added", "deleted", "modified".
	Status FileGetStatusResponseStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Path        respjson.Field
		Removed     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *FileGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileGetStatusResponseStatus string

const (
	FileGetStatusResponseStatusAdded    FileGetStatusResponseStatus = "added"
	FileGetStatusResponseStatusDeleted  FileGetStatusResponseStatus = "deleted"
	FileGetStatusResponseStatusModified FileGetStatusResponseStatus = "modified"
)

type FileReadResponse struct {
	Content string `json:"content,required"`
	// Any of "raw", "patch".
	Type FileReadResponseType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileReadResponse) RawJSON() string { return r.JSON.raw }
func (r *FileReadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileReadResponseType string

const (
	FileReadResponseTypeRaw   FileReadResponseType = "raw"
	FileReadResponseTypePatch FileReadResponseType = "patch"
)

type FileListParams struct {
	Path string `query:"path,required" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileReadParams struct {
	Path string `query:"path,required" json:"-"`
	paramObj
}

// URLQuery serializes [FileReadParams]'s query parameters as `url.Values`.
func (r FileReadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
