// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/opencode-go/internal/apijson"
	"github.com/stainless-sdks/opencode-go/internal/apiquery"
	"github.com/stainless-sdks/opencode-go/internal/requestconfig"
	"github.com/stainless-sdks/opencode-go/option"
	"github.com/stainless-sdks/opencode-go/packages/param"
	"github.com/stainless-sdks/opencode-go/packages/respjson"
)

// FindService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFindService] method instead.
type FindService struct {
	Options []option.RequestOption
}

// NewFindService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFindService(opts ...option.RequestOption) (r FindService) {
	r = FindService{}
	r.Options = opts
	return
}

// Find text in files
func (r *FindService) Get(ctx context.Context, query FindGetParams, opts ...option.RequestOption) (res *[]FindGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "find"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Find files
func (r *FindService) GetFile(ctx context.Context, query FindGetFileParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	path := "find/file"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Find workspace symbols
func (r *FindService) GetSymbol(ctx context.Context, query FindGetSymbolParams, opts ...option.RequestOption) (res *[]FindGetSymbolResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "find/symbol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Range struct {
	End   RangeEnd   `json:"end,required"`
	Start RangeStart `json:"start,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Range) RawJSON() string { return r.JSON.raw }
func (r *Range) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Range to a RangeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RangeParam.Overrides()
func (r Range) ToParam() RangeParam {
	return param.Override[RangeParam](json.RawMessage(r.RawJSON()))
}

type RangeEnd struct {
	Character float64 `json:"character,required"`
	Line      float64 `json:"line,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Character   respjson.Field
		Line        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RangeEnd) RawJSON() string { return r.JSON.raw }
func (r *RangeEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RangeStart struct {
	Character float64 `json:"character,required"`
	Line      float64 `json:"line,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Character   respjson.Field
		Line        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RangeStart) RawJSON() string { return r.JSON.raw }
func (r *RangeStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties End, Start are required.
type RangeParam struct {
	End   RangeEndParam   `json:"end,omitzero,required"`
	Start RangeStartParam `json:"start,omitzero,required"`
	paramObj
}

func (r RangeParam) MarshalJSON() (data []byte, err error) {
	type shadow RangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Character, Line are required.
type RangeEndParam struct {
	Character float64 `json:"character,required"`
	Line      float64 `json:"line,required"`
	paramObj
}

func (r RangeEndParam) MarshalJSON() (data []byte, err error) {
	type shadow RangeEndParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RangeEndParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Character, Line are required.
type RangeStartParam struct {
	Character float64 `json:"character,required"`
	Line      float64 `json:"line,required"`
	paramObj
}

func (r RangeStartParam) MarshalJSON() (data []byte, err error) {
	type shadow RangeStartParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RangeStartParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FindGetResponse struct {
	AbsoluteOffset float64                   `json:"absolute_offset,required"`
	LineNumber     float64                   `json:"line_number,required"`
	Lines          FindGetResponseLines      `json:"lines,required"`
	Path           FindGetResponsePath       `json:"path,required"`
	Submatches     []FindGetResponseSubmatch `json:"submatches,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AbsoluteOffset respjson.Field
		LineNumber     respjson.Field
		Lines          respjson.Field
		Path           respjson.Field
		Submatches     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FindGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FindGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FindGetResponseLines struct {
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FindGetResponseLines) RawJSON() string { return r.JSON.raw }
func (r *FindGetResponseLines) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FindGetResponsePath struct {
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FindGetResponsePath) RawJSON() string { return r.JSON.raw }
func (r *FindGetResponsePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FindGetResponseSubmatch struct {
	End   float64                      `json:"end,required"`
	Match FindGetResponseSubmatchMatch `json:"match,required"`
	Start float64                      `json:"start,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Match       respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FindGetResponseSubmatch) RawJSON() string { return r.JSON.raw }
func (r *FindGetResponseSubmatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FindGetResponseSubmatchMatch struct {
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FindGetResponseSubmatchMatch) RawJSON() string { return r.JSON.raw }
func (r *FindGetResponseSubmatchMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FindGetSymbolResponse struct {
	Kind     float64                       `json:"kind,required"`
	Location FindGetSymbolResponseLocation `json:"location,required"`
	Name     string                        `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Location    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FindGetSymbolResponse) RawJSON() string { return r.JSON.raw }
func (r *FindGetSymbolResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FindGetSymbolResponseLocation struct {
	Range Range  `json:"range,required"`
	Uri   string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Range       respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FindGetSymbolResponseLocation) RawJSON() string { return r.JSON.raw }
func (r *FindGetSymbolResponseLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FindGetParams struct {
	Pattern string `query:"pattern,required" json:"-"`
	paramObj
}

// URLQuery serializes [FindGetParams]'s query parameters as `url.Values`.
func (r FindGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FindGetFileParams struct {
	Query string `query:"query,required" json:"-"`
	paramObj
}

// URLQuery serializes [FindGetFileParams]'s query parameters as `url.Values`.
func (r FindGetFileParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FindGetSymbolParams struct {
	Query string `query:"query,required" json:"-"`
	paramObj
}

// URLQuery serializes [FindGetSymbolParams]'s query parameters as `url.Values`.
func (r FindGetSymbolParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
