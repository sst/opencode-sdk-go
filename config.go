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
	"github.com/stainless-sdks/opencode-go/shared/constant"
)

// ConfigService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r ConfigService) {
	r = ConfigService{}
	r.Options = opts
	return
}

// Get config info
func (r *ConfigService) Get(ctx context.Context, opts ...option.RequestOption) (res *ConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all providers
func (r *ConfigService) ListProviders(ctx context.Context, opts ...option.RequestOption) (res *ConfigListProvidersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "config/providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ConfigGetResponse struct {
	// JSON schema reference for configuration validation
	Schema string `json:"$schema"`
	// Agent configuration, see https://opencode.ai/docs/agent
	Agent ConfigGetResponseAgent `json:"agent"`
	// @deprecated Use 'share' field instead. Share newly created sessions
	// automatically
	Autoshare bool `json:"autoshare"`
	// Automatically update to the latest version
	Autoupdate bool `json:"autoupdate"`
	// Command configuration, see https://opencode.ai/docs/commands
	Command map[string]ConfigGetResponseCommand `json:"command"`
	// Disable providers that are loaded automatically
	DisabledProviders []string                              `json:"disabled_providers"`
	Experimental      ConfigGetResponseExperimental         `json:"experimental"`
	Formatter         map[string]ConfigGetResponseFormatter `json:"formatter"`
	// Additional instruction files or patterns to include
	Instructions []string `json:"instructions"`
	// Custom keybind configurations
	Keybinds ConfigGetResponseKeybinds `json:"keybinds"`
	// @deprecated Always uses stretch layout.
	//
	// Any of "auto", "stretch".
	Layout ConfigGetResponseLayout              `json:"layout"`
	Lsp    map[string]ConfigGetResponseLspUnion `json:"lsp"`
	// MCP (Model Context Protocol) server configurations
	Mcp map[string]ConfigGetResponseMcpUnion `json:"mcp"`
	// @deprecated Use `agent` field instead.
	Mode ConfigGetResponseMode `json:"mode"`
	// Model to use in the format of provider/model, eg anthropic/claude-2
	Model      string                      `json:"model"`
	Permission ConfigGetResponsePermission `json:"permission"`
	Plugin     []string                    `json:"plugin"`
	// Custom provider configurations and model overrides
	Provider map[string]ConfigGetResponseProvider `json:"provider"`
	// Control sharing behavior:'manual' allows manual sharing via commands, 'auto'
	// enables automatic sharing, 'disabled' disables all sharing
	//
	// Any of "manual", "auto", "disabled".
	Share ConfigGetResponseShare `json:"share"`
	// Small model to use for tasks like title generation in the format of
	// provider/model
	SmallModel string `json:"small_model"`
	Snapshot   bool   `json:"snapshot"`
	// Theme name to use for the interface
	Theme string          `json:"theme"`
	Tools map[string]bool `json:"tools"`
	// TUI specific settings
	Tui ConfigGetResponseTui `json:"tui"`
	// Custom username to display in conversations instead of system username
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Schema            respjson.Field
		Agent             respjson.Field
		Autoshare         respjson.Field
		Autoupdate        respjson.Field
		Command           respjson.Field
		DisabledProviders respjson.Field
		Experimental      respjson.Field
		Formatter         respjson.Field
		Instructions      respjson.Field
		Keybinds          respjson.Field
		Layout            respjson.Field
		Lsp               respjson.Field
		Mcp               respjson.Field
		Mode              respjson.Field
		Model             respjson.Field
		Permission        respjson.Field
		Plugin            respjson.Field
		Provider          respjson.Field
		Share             respjson.Field
		SmallModel        respjson.Field
		Snapshot          respjson.Field
		Theme             respjson.Field
		Tools             respjson.Field
		Tui               respjson.Field
		Username          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Agent configuration, see https://opencode.ai/docs/agent
type ConfigGetResponseAgent struct {
	Build       ConfigGetResponseAgentBuild       `json:"build"`
	General     ConfigGetResponseAgentGeneral     `json:"general"`
	Plan        ConfigGetResponseAgentPlan        `json:"plan"`
	ExtraFields map[string]ConfigGetResponseAgent `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Build       respjson.Field
		General     respjson.Field
		Plan        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseAgent) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseAgentBuild struct {
	// Description of when to use the agent
	Description string `json:"description"`
	Disable     bool   `json:"disable"`
	// Any of "subagent", "primary", "all".
	Mode        ConfigGetResponseAgentBuildMode       `json:"mode"`
	Model       string                                `json:"model"`
	Permission  ConfigGetResponseAgentBuildPermission `json:"permission"`
	Prompt      string                                `json:"prompt"`
	Temperature float64                               `json:"temperature"`
	Tools       map[string]bool                       `json:"tools"`
	TopP        float64                               `json:"top_p"`
	ExtraFields map[string]any                        `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Disable     respjson.Field
		Mode        respjson.Field
		Model       respjson.Field
		Permission  respjson.Field
		Prompt      respjson.Field
		Temperature respjson.Field
		Tools       respjson.Field
		TopP        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseAgentBuild) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseAgentBuild) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseAgentBuildMode string

const (
	ConfigGetResponseAgentBuildModeSubagent ConfigGetResponseAgentBuildMode = "subagent"
	ConfigGetResponseAgentBuildModePrimary  ConfigGetResponseAgentBuildMode = "primary"
	ConfigGetResponseAgentBuildModeAll      ConfigGetResponseAgentBuildMode = "all"
)

type ConfigGetResponseAgentBuildPermission struct {
	Bash ConfigGetResponseAgentBuildPermissionBashUnion `json:"bash"`
	// Any of "ask", "allow", "deny".
	Edit ConfigGetResponseAgentBuildPermissionEdit `json:"edit"`
	// Any of "ask", "allow", "deny".
	Webfetch ConfigGetResponseAgentBuildPermissionWebfetch `json:"webfetch"`
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
func (r ConfigGetResponseAgentBuildPermission) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseAgentBuildPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseAgentBuildPermissionBashString string

const (
	ConfigGetResponseAgentBuildPermissionBashStringAsk   ConfigGetResponseAgentBuildPermissionBashString = "ask"
	ConfigGetResponseAgentBuildPermissionBashStringAllow ConfigGetResponseAgentBuildPermissionBashString = "allow"
	ConfigGetResponseAgentBuildPermissionBashStringDeny  ConfigGetResponseAgentBuildPermissionBashString = "deny"
)

type ConfigGetResponseAgentBuildPermissionBashMapItem string

const (
	ConfigGetResponseAgentBuildPermissionBashMapItemAsk   ConfigGetResponseAgentBuildPermissionBashMapItem = "ask"
	ConfigGetResponseAgentBuildPermissionBashMapItemAllow ConfigGetResponseAgentBuildPermissionBashMapItem = "allow"
	ConfigGetResponseAgentBuildPermissionBashMapItemDeny  ConfigGetResponseAgentBuildPermissionBashMapItem = "deny"
)

type ConfigGetResponseAgentBuildPermissionEdit string

const (
	ConfigGetResponseAgentBuildPermissionEditAsk   ConfigGetResponseAgentBuildPermissionEdit = "ask"
	ConfigGetResponseAgentBuildPermissionEditAllow ConfigGetResponseAgentBuildPermissionEdit = "allow"
	ConfigGetResponseAgentBuildPermissionEditDeny  ConfigGetResponseAgentBuildPermissionEdit = "deny"
)

type ConfigGetResponseAgentBuildPermissionWebfetch string

const (
	ConfigGetResponseAgentBuildPermissionWebfetchAsk   ConfigGetResponseAgentBuildPermissionWebfetch = "ask"
	ConfigGetResponseAgentBuildPermissionWebfetchAllow ConfigGetResponseAgentBuildPermissionWebfetch = "allow"
	ConfigGetResponseAgentBuildPermissionWebfetchDeny  ConfigGetResponseAgentBuildPermissionWebfetch = "deny"
)

type ConfigGetResponseAgentGeneral struct {
	// Description of when to use the agent
	Description string `json:"description"`
	Disable     bool   `json:"disable"`
	// Any of "subagent", "primary", "all".
	Mode        ConfigGetResponseAgentGeneralMode       `json:"mode"`
	Model       string                                  `json:"model"`
	Permission  ConfigGetResponseAgentGeneralPermission `json:"permission"`
	Prompt      string                                  `json:"prompt"`
	Temperature float64                                 `json:"temperature"`
	Tools       map[string]bool                         `json:"tools"`
	TopP        float64                                 `json:"top_p"`
	ExtraFields map[string]any                          `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Disable     respjson.Field
		Mode        respjson.Field
		Model       respjson.Field
		Permission  respjson.Field
		Prompt      respjson.Field
		Temperature respjson.Field
		Tools       respjson.Field
		TopP        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseAgentGeneral) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseAgentGeneral) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseAgentGeneralMode string

const (
	ConfigGetResponseAgentGeneralModeSubagent ConfigGetResponseAgentGeneralMode = "subagent"
	ConfigGetResponseAgentGeneralModePrimary  ConfigGetResponseAgentGeneralMode = "primary"
	ConfigGetResponseAgentGeneralModeAll      ConfigGetResponseAgentGeneralMode = "all"
)

type ConfigGetResponseAgentGeneralPermission struct {
	Bash ConfigGetResponseAgentGeneralPermissionBashUnion `json:"bash"`
	// Any of "ask", "allow", "deny".
	Edit ConfigGetResponseAgentGeneralPermissionEdit `json:"edit"`
	// Any of "ask", "allow", "deny".
	Webfetch ConfigGetResponseAgentGeneralPermissionWebfetch `json:"webfetch"`
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
func (r ConfigGetResponseAgentGeneralPermission) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseAgentGeneralPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseAgentGeneralPermissionBashString string

const (
	ConfigGetResponseAgentGeneralPermissionBashStringAsk   ConfigGetResponseAgentGeneralPermissionBashString = "ask"
	ConfigGetResponseAgentGeneralPermissionBashStringAllow ConfigGetResponseAgentGeneralPermissionBashString = "allow"
	ConfigGetResponseAgentGeneralPermissionBashStringDeny  ConfigGetResponseAgentGeneralPermissionBashString = "deny"
)

type ConfigGetResponseAgentGeneralPermissionBashMapItem string

const (
	ConfigGetResponseAgentGeneralPermissionBashMapItemAsk   ConfigGetResponseAgentGeneralPermissionBashMapItem = "ask"
	ConfigGetResponseAgentGeneralPermissionBashMapItemAllow ConfigGetResponseAgentGeneralPermissionBashMapItem = "allow"
	ConfigGetResponseAgentGeneralPermissionBashMapItemDeny  ConfigGetResponseAgentGeneralPermissionBashMapItem = "deny"
)

type ConfigGetResponseAgentGeneralPermissionEdit string

const (
	ConfigGetResponseAgentGeneralPermissionEditAsk   ConfigGetResponseAgentGeneralPermissionEdit = "ask"
	ConfigGetResponseAgentGeneralPermissionEditAllow ConfigGetResponseAgentGeneralPermissionEdit = "allow"
	ConfigGetResponseAgentGeneralPermissionEditDeny  ConfigGetResponseAgentGeneralPermissionEdit = "deny"
)

type ConfigGetResponseAgentGeneralPermissionWebfetch string

const (
	ConfigGetResponseAgentGeneralPermissionWebfetchAsk   ConfigGetResponseAgentGeneralPermissionWebfetch = "ask"
	ConfigGetResponseAgentGeneralPermissionWebfetchAllow ConfigGetResponseAgentGeneralPermissionWebfetch = "allow"
	ConfigGetResponseAgentGeneralPermissionWebfetchDeny  ConfigGetResponseAgentGeneralPermissionWebfetch = "deny"
)

type ConfigGetResponseAgentPlan struct {
	// Description of when to use the agent
	Description string `json:"description"`
	Disable     bool   `json:"disable"`
	// Any of "subagent", "primary", "all".
	Mode        ConfigGetResponseAgentPlanMode       `json:"mode"`
	Model       string                               `json:"model"`
	Permission  ConfigGetResponseAgentPlanPermission `json:"permission"`
	Prompt      string                               `json:"prompt"`
	Temperature float64                              `json:"temperature"`
	Tools       map[string]bool                      `json:"tools"`
	TopP        float64                              `json:"top_p"`
	ExtraFields map[string]any                       `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Disable     respjson.Field
		Mode        respjson.Field
		Model       respjson.Field
		Permission  respjson.Field
		Prompt      respjson.Field
		Temperature respjson.Field
		Tools       respjson.Field
		TopP        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseAgentPlan) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseAgentPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseAgentPlanMode string

const (
	ConfigGetResponseAgentPlanModeSubagent ConfigGetResponseAgentPlanMode = "subagent"
	ConfigGetResponseAgentPlanModePrimary  ConfigGetResponseAgentPlanMode = "primary"
	ConfigGetResponseAgentPlanModeAll      ConfigGetResponseAgentPlanMode = "all"
)

type ConfigGetResponseAgentPlanPermission struct {
	Bash ConfigGetResponseAgentPlanPermissionBashUnion `json:"bash"`
	// Any of "ask", "allow", "deny".
	Edit ConfigGetResponseAgentPlanPermissionEdit `json:"edit"`
	// Any of "ask", "allow", "deny".
	Webfetch ConfigGetResponseAgentPlanPermissionWebfetch `json:"webfetch"`
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
func (r ConfigGetResponseAgentPlanPermission) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseAgentPlanPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseAgentPlanPermissionBashString string

const (
	ConfigGetResponseAgentPlanPermissionBashStringAsk   ConfigGetResponseAgentPlanPermissionBashString = "ask"
	ConfigGetResponseAgentPlanPermissionBashStringAllow ConfigGetResponseAgentPlanPermissionBashString = "allow"
	ConfigGetResponseAgentPlanPermissionBashStringDeny  ConfigGetResponseAgentPlanPermissionBashString = "deny"
)

type ConfigGetResponseAgentPlanPermissionBashMapItem string

const (
	ConfigGetResponseAgentPlanPermissionBashMapItemAsk   ConfigGetResponseAgentPlanPermissionBashMapItem = "ask"
	ConfigGetResponseAgentPlanPermissionBashMapItemAllow ConfigGetResponseAgentPlanPermissionBashMapItem = "allow"
	ConfigGetResponseAgentPlanPermissionBashMapItemDeny  ConfigGetResponseAgentPlanPermissionBashMapItem = "deny"
)

type ConfigGetResponseAgentPlanPermissionEdit string

const (
	ConfigGetResponseAgentPlanPermissionEditAsk   ConfigGetResponseAgentPlanPermissionEdit = "ask"
	ConfigGetResponseAgentPlanPermissionEditAllow ConfigGetResponseAgentPlanPermissionEdit = "allow"
	ConfigGetResponseAgentPlanPermissionEditDeny  ConfigGetResponseAgentPlanPermissionEdit = "deny"
)

type ConfigGetResponseAgentPlanPermissionWebfetch string

const (
	ConfigGetResponseAgentPlanPermissionWebfetchAsk   ConfigGetResponseAgentPlanPermissionWebfetch = "ask"
	ConfigGetResponseAgentPlanPermissionWebfetchAllow ConfigGetResponseAgentPlanPermissionWebfetch = "allow"
	ConfigGetResponseAgentPlanPermissionWebfetchDeny  ConfigGetResponseAgentPlanPermissionWebfetch = "deny"
)

type ConfigGetResponseCommand struct {
	Template    string `json:"template,required"`
	Agent       string `json:"agent"`
	Description string `json:"description"`
	Model       string `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Template    respjson.Field
		Agent       respjson.Field
		Description respjson.Field
		Model       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseCommand) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseCommand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseExperimental struct {
	Hook ConfigGetResponseExperimentalHook `json:"hook"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hook        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseExperimental) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseExperimental) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseExperimentalHook struct {
	FileEdited       map[string][]ConfigGetResponseExperimentalHookFileEdited `json:"file_edited"`
	SessionCompleted []ConfigGetResponseExperimentalHookSessionCompleted      `json:"session_completed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileEdited       respjson.Field
		SessionCompleted respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseExperimentalHook) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseExperimentalHook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseExperimentalHookFileEdited struct {
	Command     []string          `json:"command,required"`
	Environment map[string]string `json:"environment"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command     respjson.Field
		Environment respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseExperimentalHookFileEdited) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseExperimentalHookFileEdited) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseExperimentalHookSessionCompleted struct {
	Command     []string          `json:"command,required"`
	Environment map[string]string `json:"environment"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command     respjson.Field
		Environment respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseExperimentalHookSessionCompleted) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseExperimentalHookSessionCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseFormatter struct {
	Command     []string          `json:"command"`
	Disabled    bool              `json:"disabled"`
	Environment map[string]string `json:"environment"`
	Extensions  []string          `json:"extensions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command     respjson.Field
		Disabled    respjson.Field
		Environment respjson.Field
		Extensions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseFormatter) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseFormatter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom keybind configurations
type ConfigGetResponseKeybinds struct {
	// Next agent
	AgentCycle string `json:"agent_cycle,required"`
	// Previous agent
	AgentCycleReverse string `json:"agent_cycle_reverse,required"`
	// List agents
	AgentList string `json:"agent_list,required"`
	// Exit the application
	AppExit string `json:"app_exit,required"`
	// Show help dialog
	AppHelp string `json:"app_help,required"`
	// Open external editor
	EditorOpen string `json:"editor_open,required"`
	// @deprecated Close file
	FileClose string `json:"file_close,required"`
	// @deprecated Split/unified diff
	FileDiffToggle string `json:"file_diff_toggle,required"`
	// @deprecated Currently not available. List files
	FileList string `json:"file_list,required"`
	// @deprecated Search file
	FileSearch string `json:"file_search,required"`
	// Clear input field
	InputClear string `json:"input_clear,required"`
	// Insert newline in input
	InputNewline string `json:"input_newline,required"`
	// Paste from clipboard
	InputPaste string `json:"input_paste,required"`
	// Submit input
	InputSubmit string `json:"input_submit,required"`
	// Leader key for keybind combinations
	Leader string `json:"leader,required"`
	// Copy message
	MessagesCopy string `json:"messages_copy,required"`
	// Navigate to first message
	MessagesFirst string `json:"messages_first,required"`
	// Scroll messages down by half page
	MessagesHalfPageDown string `json:"messages_half_page_down,required"`
	// Scroll messages up by half page
	MessagesHalfPageUp string `json:"messages_half_page_up,required"`
	// Navigate to last message
	MessagesLast string `json:"messages_last,required"`
	// @deprecated Toggle layout
	MessagesLayoutToggle string `json:"messages_layout_toggle,required"`
	// @deprecated Navigate to next message
	MessagesNext string `json:"messages_next,required"`
	// Scroll messages down by one page
	MessagesPageDown string `json:"messages_page_down,required"`
	// Scroll messages up by one page
	MessagesPageUp string `json:"messages_page_up,required"`
	// @deprecated Navigate to previous message
	MessagesPrevious string `json:"messages_previous,required"`
	// Redo message
	MessagesRedo string `json:"messages_redo,required"`
	// @deprecated use messages_undo. Revert message
	MessagesRevert string `json:"messages_revert,required"`
	// Undo message
	MessagesUndo string `json:"messages_undo,required"`
	// Next recent model
	ModelCycleRecent string `json:"model_cycle_recent,required"`
	// Previous recent model
	ModelCycleRecentReverse string `json:"model_cycle_recent_reverse,required"`
	// List available models
	ModelList string `json:"model_list,required"`
	// Create/update AGENTS.md
	ProjectInit string `json:"project_init,required"`
	// Cycle to next child session
	SessionChildCycle string `json:"session_child_cycle,required"`
	// Cycle to previous child session
	SessionChildCycleReverse string `json:"session_child_cycle_reverse,required"`
	// Compact the session
	SessionCompact string `json:"session_compact,required"`
	// Export session to editor
	SessionExport string `json:"session_export,required"`
	// Interrupt current session
	SessionInterrupt string `json:"session_interrupt,required"`
	// List all sessions
	SessionList string `json:"session_list,required"`
	// Create a new session
	SessionNew string `json:"session_new,required"`
	// Share current session
	SessionShare string `json:"session_share,required"`
	// Show session timeline
	SessionTimeline string `json:"session_timeline,required"`
	// Unshare current session
	SessionUnshare string `json:"session_unshare,required"`
	// @deprecated use agent_cycle. Next agent
	SwitchAgent string `json:"switch_agent,required"`
	// @deprecated use agent_cycle_reverse. Previous agent
	SwitchAgentReverse string `json:"switch_agent_reverse,required"`
	// @deprecated use agent_cycle. Next mode
	SwitchMode string `json:"switch_mode,required"`
	// @deprecated use agent_cycle_reverse. Previous mode
	SwitchModeReverse string `json:"switch_mode_reverse,required"`
	// List available themes
	ThemeList string `json:"theme_list,required"`
	// Toggle thinking blocks
	ThinkingBlocks string `json:"thinking_blocks,required"`
	// Toggle tool details
	ToolDetails string `json:"tool_details,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentCycle               respjson.Field
		AgentCycleReverse        respjson.Field
		AgentList                respjson.Field
		AppExit                  respjson.Field
		AppHelp                  respjson.Field
		EditorOpen               respjson.Field
		FileClose                respjson.Field
		FileDiffToggle           respjson.Field
		FileList                 respjson.Field
		FileSearch               respjson.Field
		InputClear               respjson.Field
		InputNewline             respjson.Field
		InputPaste               respjson.Field
		InputSubmit              respjson.Field
		Leader                   respjson.Field
		MessagesCopy             respjson.Field
		MessagesFirst            respjson.Field
		MessagesHalfPageDown     respjson.Field
		MessagesHalfPageUp       respjson.Field
		MessagesLast             respjson.Field
		MessagesLayoutToggle     respjson.Field
		MessagesNext             respjson.Field
		MessagesPageDown         respjson.Field
		MessagesPageUp           respjson.Field
		MessagesPrevious         respjson.Field
		MessagesRedo             respjson.Field
		MessagesRevert           respjson.Field
		MessagesUndo             respjson.Field
		ModelCycleRecent         respjson.Field
		ModelCycleRecentReverse  respjson.Field
		ModelList                respjson.Field
		ProjectInit              respjson.Field
		SessionChildCycle        respjson.Field
		SessionChildCycleReverse respjson.Field
		SessionCompact           respjson.Field
		SessionExport            respjson.Field
		SessionInterrupt         respjson.Field
		SessionList              respjson.Field
		SessionNew               respjson.Field
		SessionShare             respjson.Field
		SessionTimeline          respjson.Field
		SessionUnshare           respjson.Field
		SwitchAgent              respjson.Field
		SwitchAgentReverse       respjson.Field
		SwitchMode               respjson.Field
		SwitchModeReverse        respjson.Field
		ThemeList                respjson.Field
		ThinkingBlocks           respjson.Field
		ToolDetails              respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseKeybinds) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseKeybinds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// @deprecated Always uses stretch layout.
type ConfigGetResponseLayout string

const (
	ConfigGetResponseLayoutAuto    ConfigGetResponseLayout = "auto"
	ConfigGetResponseLayoutStretch ConfigGetResponseLayout = "stretch"
)

// ConfigGetResponseLspUnion contains all possible properties and values from
// [ConfigGetResponseLspDisabled], [ConfigGetResponseLspObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConfigGetResponseLspUnion struct {
	Disabled bool `json:"disabled"`
	// This field is from variant [ConfigGetResponseLspObject].
	Command []string `json:"command"`
	// This field is from variant [ConfigGetResponseLspObject].
	Env map[string]string `json:"env"`
	// This field is from variant [ConfigGetResponseLspObject].
	Extensions []string `json:"extensions"`
	// This field is from variant [ConfigGetResponseLspObject].
	Initialization map[string]any `json:"initialization"`
	JSON           struct {
		Disabled       respjson.Field
		Command        respjson.Field
		Env            respjson.Field
		Extensions     respjson.Field
		Initialization respjson.Field
		raw            string
	} `json:"-"`
}

func (u ConfigGetResponseLspUnion) AsConfigGetResponseLspDisabled() (v ConfigGetResponseLspDisabled) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConfigGetResponseLspUnion) AsConfigGetResponseLspObject() (v ConfigGetResponseLspObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConfigGetResponseLspUnion) RawJSON() string { return u.JSON.raw }

func (r *ConfigGetResponseLspUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseLspDisabled struct {
	Disabled bool `json:"disabled,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Disabled    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseLspDisabled) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseLspDisabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseLspObject struct {
	Command        []string          `json:"command,required"`
	Disabled       bool              `json:"disabled"`
	Env            map[string]string `json:"env"`
	Extensions     []string          `json:"extensions"`
	Initialization map[string]any    `json:"initialization"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command        respjson.Field
		Disabled       respjson.Field
		Env            respjson.Field
		Extensions     respjson.Field
		Initialization respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseLspObject) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseLspObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConfigGetResponseMcpUnion contains all possible properties and values from
// [ConfigGetResponseMcpLocal], [ConfigGetResponseMcpRemote].
//
// Use the [ConfigGetResponseMcpUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConfigGetResponseMcpUnion struct {
	// This field is from variant [ConfigGetResponseMcpLocal].
	Command []string `json:"command"`
	// Any of "local", "remote".
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
	// This field is from variant [ConfigGetResponseMcpLocal].
	Environment map[string]string `json:"environment"`
	// This field is from variant [ConfigGetResponseMcpRemote].
	URL string `json:"url"`
	// This field is from variant [ConfigGetResponseMcpRemote].
	Headers map[string]string `json:"headers"`
	JSON    struct {
		Command     respjson.Field
		Type        respjson.Field
		Enabled     respjson.Field
		Environment respjson.Field
		URL         respjson.Field
		Headers     respjson.Field
		raw         string
	} `json:"-"`
}

// anyConfigGetResponseMcp is implemented by each variant of
// [ConfigGetResponseMcpUnion] to add type safety for the return type of
// [ConfigGetResponseMcpUnion.AsAny]
type anyConfigGetResponseMcp interface {
	implConfigGetResponseMcpUnion()
}

func (ConfigGetResponseMcpLocal) implConfigGetResponseMcpUnion()  {}
func (ConfigGetResponseMcpRemote) implConfigGetResponseMcpUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConfigGetResponseMcpUnion.AsAny().(type) {
//	case opencode.ConfigGetResponseMcpLocal:
//	case opencode.ConfigGetResponseMcpRemote:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConfigGetResponseMcpUnion) AsAny() anyConfigGetResponseMcp {
	switch u.Type {
	case "local":
		return u.AsLocal()
	case "remote":
		return u.AsRemote()
	}
	return nil
}

func (u ConfigGetResponseMcpUnion) AsLocal() (v ConfigGetResponseMcpLocal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConfigGetResponseMcpUnion) AsRemote() (v ConfigGetResponseMcpRemote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConfigGetResponseMcpUnion) RawJSON() string { return u.JSON.raw }

func (r *ConfigGetResponseMcpUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseMcpLocal struct {
	// Command and arguments to run the MCP server
	Command []string `json:"command,required"`
	// Type of MCP server connection
	Type constant.Local `json:"type,required"`
	// Enable or disable the MCP server on startup
	Enabled bool `json:"enabled"`
	// Environment variables to set when running the MCP server
	Environment map[string]string `json:"environment"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command     respjson.Field
		Type        respjson.Field
		Enabled     respjson.Field
		Environment respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseMcpLocal) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseMcpLocal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseMcpRemote struct {
	// Type of MCP server connection
	Type constant.Remote `json:"type,required"`
	// URL of the remote MCP server
	URL string `json:"url,required"`
	// Enable or disable the MCP server on startup
	Enabled bool `json:"enabled"`
	// Headers to send with the request
	Headers map[string]string `json:"headers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		Enabled     respjson.Field
		Headers     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseMcpRemote) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseMcpRemote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// @deprecated Use `agent` field instead.
type ConfigGetResponseMode struct {
	Build       ConfigGetResponseModeBuild       `json:"build"`
	Plan        ConfigGetResponseModePlan        `json:"plan"`
	ExtraFields map[string]ConfigGetResponseMode `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Build       respjson.Field
		Plan        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseMode) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseMode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseModeBuild struct {
	// Description of when to use the agent
	Description string `json:"description"`
	Disable     bool   `json:"disable"`
	// Any of "subagent", "primary", "all".
	Mode        ConfigGetResponseModeBuildMode       `json:"mode"`
	Model       string                               `json:"model"`
	Permission  ConfigGetResponseModeBuildPermission `json:"permission"`
	Prompt      string                               `json:"prompt"`
	Temperature float64                              `json:"temperature"`
	Tools       map[string]bool                      `json:"tools"`
	TopP        float64                              `json:"top_p"`
	ExtraFields map[string]any                       `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Disable     respjson.Field
		Mode        respjson.Field
		Model       respjson.Field
		Permission  respjson.Field
		Prompt      respjson.Field
		Temperature respjson.Field
		Tools       respjson.Field
		TopP        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseModeBuild) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseModeBuild) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseModeBuildMode string

const (
	ConfigGetResponseModeBuildModeSubagent ConfigGetResponseModeBuildMode = "subagent"
	ConfigGetResponseModeBuildModePrimary  ConfigGetResponseModeBuildMode = "primary"
	ConfigGetResponseModeBuildModeAll      ConfigGetResponseModeBuildMode = "all"
)

type ConfigGetResponseModeBuildPermission struct {
	Bash ConfigGetResponseModeBuildPermissionBashUnion `json:"bash"`
	// Any of "ask", "allow", "deny".
	Edit ConfigGetResponseModeBuildPermissionEdit `json:"edit"`
	// Any of "ask", "allow", "deny".
	Webfetch ConfigGetResponseModeBuildPermissionWebfetch `json:"webfetch"`
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
func (r ConfigGetResponseModeBuildPermission) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseModeBuildPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseModeBuildPermissionBashString string

const (
	ConfigGetResponseModeBuildPermissionBashStringAsk   ConfigGetResponseModeBuildPermissionBashString = "ask"
	ConfigGetResponseModeBuildPermissionBashStringAllow ConfigGetResponseModeBuildPermissionBashString = "allow"
	ConfigGetResponseModeBuildPermissionBashStringDeny  ConfigGetResponseModeBuildPermissionBashString = "deny"
)

type ConfigGetResponseModeBuildPermissionBashMapItem string

const (
	ConfigGetResponseModeBuildPermissionBashMapItemAsk   ConfigGetResponseModeBuildPermissionBashMapItem = "ask"
	ConfigGetResponseModeBuildPermissionBashMapItemAllow ConfigGetResponseModeBuildPermissionBashMapItem = "allow"
	ConfigGetResponseModeBuildPermissionBashMapItemDeny  ConfigGetResponseModeBuildPermissionBashMapItem = "deny"
)

type ConfigGetResponseModeBuildPermissionEdit string

const (
	ConfigGetResponseModeBuildPermissionEditAsk   ConfigGetResponseModeBuildPermissionEdit = "ask"
	ConfigGetResponseModeBuildPermissionEditAllow ConfigGetResponseModeBuildPermissionEdit = "allow"
	ConfigGetResponseModeBuildPermissionEditDeny  ConfigGetResponseModeBuildPermissionEdit = "deny"
)

type ConfigGetResponseModeBuildPermissionWebfetch string

const (
	ConfigGetResponseModeBuildPermissionWebfetchAsk   ConfigGetResponseModeBuildPermissionWebfetch = "ask"
	ConfigGetResponseModeBuildPermissionWebfetchAllow ConfigGetResponseModeBuildPermissionWebfetch = "allow"
	ConfigGetResponseModeBuildPermissionWebfetchDeny  ConfigGetResponseModeBuildPermissionWebfetch = "deny"
)

type ConfigGetResponseModePlan struct {
	// Description of when to use the agent
	Description string `json:"description"`
	Disable     bool   `json:"disable"`
	// Any of "subagent", "primary", "all".
	Mode        ConfigGetResponseModePlanMode       `json:"mode"`
	Model       string                              `json:"model"`
	Permission  ConfigGetResponseModePlanPermission `json:"permission"`
	Prompt      string                              `json:"prompt"`
	Temperature float64                             `json:"temperature"`
	Tools       map[string]bool                     `json:"tools"`
	TopP        float64                             `json:"top_p"`
	ExtraFields map[string]any                      `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Disable     respjson.Field
		Mode        respjson.Field
		Model       respjson.Field
		Permission  respjson.Field
		Prompt      respjson.Field
		Temperature respjson.Field
		Tools       respjson.Field
		TopP        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseModePlan) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseModePlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseModePlanMode string

const (
	ConfigGetResponseModePlanModeSubagent ConfigGetResponseModePlanMode = "subagent"
	ConfigGetResponseModePlanModePrimary  ConfigGetResponseModePlanMode = "primary"
	ConfigGetResponseModePlanModeAll      ConfigGetResponseModePlanMode = "all"
)

type ConfigGetResponseModePlanPermission struct {
	Bash ConfigGetResponseModePlanPermissionBashUnion `json:"bash"`
	// Any of "ask", "allow", "deny".
	Edit ConfigGetResponseModePlanPermissionEdit `json:"edit"`
	// Any of "ask", "allow", "deny".
	Webfetch ConfigGetResponseModePlanPermissionWebfetch `json:"webfetch"`
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
func (r ConfigGetResponseModePlanPermission) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseModePlanPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseModePlanPermissionBashString string

const (
	ConfigGetResponseModePlanPermissionBashStringAsk   ConfigGetResponseModePlanPermissionBashString = "ask"
	ConfigGetResponseModePlanPermissionBashStringAllow ConfigGetResponseModePlanPermissionBashString = "allow"
	ConfigGetResponseModePlanPermissionBashStringDeny  ConfigGetResponseModePlanPermissionBashString = "deny"
)

type ConfigGetResponseModePlanPermissionBashMapItem string

const (
	ConfigGetResponseModePlanPermissionBashMapItemAsk   ConfigGetResponseModePlanPermissionBashMapItem = "ask"
	ConfigGetResponseModePlanPermissionBashMapItemAllow ConfigGetResponseModePlanPermissionBashMapItem = "allow"
	ConfigGetResponseModePlanPermissionBashMapItemDeny  ConfigGetResponseModePlanPermissionBashMapItem = "deny"
)

type ConfigGetResponseModePlanPermissionEdit string

const (
	ConfigGetResponseModePlanPermissionEditAsk   ConfigGetResponseModePlanPermissionEdit = "ask"
	ConfigGetResponseModePlanPermissionEditAllow ConfigGetResponseModePlanPermissionEdit = "allow"
	ConfigGetResponseModePlanPermissionEditDeny  ConfigGetResponseModePlanPermissionEdit = "deny"
)

type ConfigGetResponseModePlanPermissionWebfetch string

const (
	ConfigGetResponseModePlanPermissionWebfetchAsk   ConfigGetResponseModePlanPermissionWebfetch = "ask"
	ConfigGetResponseModePlanPermissionWebfetchAllow ConfigGetResponseModePlanPermissionWebfetch = "allow"
	ConfigGetResponseModePlanPermissionWebfetchDeny  ConfigGetResponseModePlanPermissionWebfetch = "deny"
)

type ConfigGetResponsePermission struct {
	Bash ConfigGetResponsePermissionBashUnion `json:"bash"`
	// Any of "ask", "allow", "deny".
	Edit ConfigGetResponsePermissionEdit `json:"edit"`
	// Any of "ask", "allow", "deny".
	Webfetch ConfigGetResponsePermissionWebfetch `json:"webfetch"`
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
func (r ConfigGetResponsePermission) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponsePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponsePermissionBashString string

const (
	ConfigGetResponsePermissionBashStringAsk   ConfigGetResponsePermissionBashString = "ask"
	ConfigGetResponsePermissionBashStringAllow ConfigGetResponsePermissionBashString = "allow"
	ConfigGetResponsePermissionBashStringDeny  ConfigGetResponsePermissionBashString = "deny"
)

type ConfigGetResponsePermissionBashMapItem string

const (
	ConfigGetResponsePermissionBashMapItemAsk   ConfigGetResponsePermissionBashMapItem = "ask"
	ConfigGetResponsePermissionBashMapItemAllow ConfigGetResponsePermissionBashMapItem = "allow"
	ConfigGetResponsePermissionBashMapItemDeny  ConfigGetResponsePermissionBashMapItem = "deny"
)

type ConfigGetResponsePermissionEdit string

const (
	ConfigGetResponsePermissionEditAsk   ConfigGetResponsePermissionEdit = "ask"
	ConfigGetResponsePermissionEditAllow ConfigGetResponsePermissionEdit = "allow"
	ConfigGetResponsePermissionEditDeny  ConfigGetResponsePermissionEdit = "deny"
)

type ConfigGetResponsePermissionWebfetch string

const (
	ConfigGetResponsePermissionWebfetchAsk   ConfigGetResponsePermissionWebfetch = "ask"
	ConfigGetResponsePermissionWebfetchAllow ConfigGetResponsePermissionWebfetch = "allow"
	ConfigGetResponsePermissionWebfetchDeny  ConfigGetResponsePermissionWebfetch = "deny"
)

type ConfigGetResponseProvider struct {
	ID      string                                    `json:"id"`
	API     string                                    `json:"api"`
	Env     []string                                  `json:"env"`
	Models  map[string]ConfigGetResponseProviderModel `json:"models"`
	Name    string                                    `json:"name"`
	Npm     string                                    `json:"npm"`
	Options ConfigGetResponseProviderOptions          `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		API         respjson.Field
		Env         respjson.Field
		Models      respjson.Field
		Name        respjson.Field
		Npm         respjson.Field
		Options     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseProvider) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseProviderModel struct {
	ID          string                              `json:"id"`
	Attachment  bool                                `json:"attachment"`
	Cost        ConfigGetResponseProviderModelCost  `json:"cost"`
	Limit       ConfigGetResponseProviderModelLimit `json:"limit"`
	Name        string                              `json:"name"`
	Options     map[string]any                      `json:"options"`
	Reasoning   bool                                `json:"reasoning"`
	ReleaseDate string                              `json:"release_date"`
	Temperature bool                                `json:"temperature"`
	ToolCall    bool                                `json:"tool_call"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attachment  respjson.Field
		Cost        respjson.Field
		Limit       respjson.Field
		Name        respjson.Field
		Options     respjson.Field
		Reasoning   respjson.Field
		ReleaseDate respjson.Field
		Temperature respjson.Field
		ToolCall    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseProviderModel) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseProviderModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseProviderModelCost struct {
	Input      float64 `json:"input,required"`
	Output     float64 `json:"output,required"`
	CacheRead  float64 `json:"cache_read"`
	CacheWrite float64 `json:"cache_write"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input       respjson.Field
		Output      respjson.Field
		CacheRead   respjson.Field
		CacheWrite  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseProviderModelCost) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseProviderModelCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseProviderModelLimit struct {
	Context float64 `json:"context,required"`
	Output  float64 `json:"output,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseProviderModelLimit) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseProviderModelLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigGetResponseProviderOptions struct {
	APIKey      string         `json:"apiKey"`
	BaseURL     string         `json:"baseURL"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		BaseURL     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseProviderOptions) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseProviderOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Control sharing behavior:'manual' allows manual sharing via commands, 'auto'
// enables automatic sharing, 'disabled' disables all sharing
type ConfigGetResponseShare string

const (
	ConfigGetResponseShareManual   ConfigGetResponseShare = "manual"
	ConfigGetResponseShareAuto     ConfigGetResponseShare = "auto"
	ConfigGetResponseShareDisabled ConfigGetResponseShare = "disabled"
)

// TUI specific settings
type ConfigGetResponseTui struct {
	// TUI scroll speed
	ScrollSpeed float64 `json:"scroll_speed,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScrollSpeed respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigGetResponseTui) RawJSON() string { return r.JSON.raw }
func (r *ConfigGetResponseTui) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigListProvidersResponse struct {
	Default   map[string]string                     `json:"default,required"`
	Providers []ConfigListProvidersResponseProvider `json:"providers,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Providers   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigListProvidersResponse) RawJSON() string { return r.JSON.raw }
func (r *ConfigListProvidersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigListProvidersResponseProvider struct {
	ID     string                                              `json:"id,required"`
	Env    []string                                            `json:"env,required"`
	Models map[string]ConfigListProvidersResponseProviderModel `json:"models,required"`
	Name   string                                              `json:"name,required"`
	API    string                                              `json:"api"`
	Npm    string                                              `json:"npm"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Env         respjson.Field
		Models      respjson.Field
		Name        respjson.Field
		API         respjson.Field
		Npm         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigListProvidersResponseProvider) RawJSON() string { return r.JSON.raw }
func (r *ConfigListProvidersResponseProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigListProvidersResponseProviderModel struct {
	ID          string                                        `json:"id,required"`
	Attachment  bool                                          `json:"attachment,required"`
	Cost        ConfigListProvidersResponseProviderModelCost  `json:"cost,required"`
	Limit       ConfigListProvidersResponseProviderModelLimit `json:"limit,required"`
	Name        string                                        `json:"name,required"`
	Options     map[string]any                                `json:"options,required"`
	Reasoning   bool                                          `json:"reasoning,required"`
	ReleaseDate string                                        `json:"release_date,required"`
	Temperature bool                                          `json:"temperature,required"`
	ToolCall    bool                                          `json:"tool_call,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attachment  respjson.Field
		Cost        respjson.Field
		Limit       respjson.Field
		Name        respjson.Field
		Options     respjson.Field
		Reasoning   respjson.Field
		ReleaseDate respjson.Field
		Temperature respjson.Field
		ToolCall    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigListProvidersResponseProviderModel) RawJSON() string { return r.JSON.raw }
func (r *ConfigListProvidersResponseProviderModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigListProvidersResponseProviderModelCost struct {
	Input      float64 `json:"input,required"`
	Output     float64 `json:"output,required"`
	CacheRead  float64 `json:"cache_read"`
	CacheWrite float64 `json:"cache_write"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input       respjson.Field
		Output      respjson.Field
		CacheRead   respjson.Field
		CacheWrite  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigListProvidersResponseProviderModelCost) RawJSON() string { return r.JSON.raw }
func (r *ConfigListProvidersResponseProviderModelCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigListProvidersResponseProviderModelLimit struct {
	Context float64 `json:"context,required"`
	Output  float64 `json:"output,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigListProvidersResponseProviderModelLimit) RawJSON() string { return r.JSON.raw }
func (r *ConfigListProvidersResponseProviderModelLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
