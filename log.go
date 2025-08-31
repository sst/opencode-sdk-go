// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/opencode-go/internal/apijson"
	"github.com/stainless-sdks/opencode-go/internal/requestconfig"
	"github.com/stainless-sdks/opencode-go/option"
	"github.com/stainless-sdks/opencode-go/packages/param"
)

// LogService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogService] method instead.
type LogService struct {
	Options []option.RequestOption
}

// NewLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogService(opts ...option.RequestOption) (r LogService) {
	r = LogService{}
	r.Options = opts
	return
}

// Write a log entry to the server logs
func (r *LogService) Write(ctx context.Context, body LogWriteParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	path := "log"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LogWriteParams struct {
	// Log level
	//
	// Any of "debug", "info", "error", "warn".
	Level LogWriteParamsLevel `json:"level,omitzero,required"`
	// Log message
	Message string `json:"message,required"`
	// Service name for the log entry
	Service string `json:"service,required"`
	// Additional metadata for the log entry
	Extra map[string]any `json:"extra,omitzero"`
	paramObj
}

func (r LogWriteParams) MarshalJSON() (data []byte, err error) {
	type shadow LogWriteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogWriteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Log level
type LogWriteParamsLevel string

const (
	LogWriteParamsLevelDebug LogWriteParamsLevel = "debug"
	LogWriteParamsLevelInfo  LogWriteParamsLevel = "info"
	LogWriteParamsLevelError LogWriteParamsLevel = "error"
	LogWriteParamsLevelWarn  LogWriteParamsLevel = "warn"
)
