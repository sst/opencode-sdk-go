// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/stainless-sdks/opencode-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Agent string                    // Always "agent"
type API string                      // Always "api"
type Assistant string                // Always "assistant"
type Completed string                // Always "completed"
type Error string                    // Always "error"
type File string                     // Always "file"
type FileEdited string               // Always "file.edited"
type InstallationUpdated string      // Always "installation.updated"
type Local string                    // Always "local"
type LspClientDiagnostics string     // Always "lsp.client.diagnostics"
type MessagePartRemoved string       // Always "message.part.removed"
type MessagePartUpdated string       // Always "message.part.updated"
type MessageRemoved string           // Always "message.removed"
type MessageUpdated string           // Always "message.updated"
type MessageAbortedError string      // Always "MessageAbortedError"
type MessageOutputLengthError string // Always "MessageOutputLengthError"
type OAuth string                    // Always "oauth"
type Patch string                    // Always "patch"
type Pending string                  // Always "pending"
type PermissionReplied string        // Always "permission.replied"
type PermissionUpdated string        // Always "permission.updated"
type ProviderAuthError string        // Always "ProviderAuthError"
type Reasoning string                // Always "reasoning"
type Remote string                   // Always "remote"
type Running string                  // Always "running"
type ServerConnected string          // Always "server.connected"
type SessionDeleted string           // Always "session.deleted"
type SessionError string             // Always "session.error"
type SessionIdle string              // Always "session.idle"
type SessionUpdated string           // Always "session.updated"
type Snapshot string                 // Always "snapshot"
type StepFinish string               // Always "step-finish"
type StepStart string                // Always "step-start"
type Symbol string                   // Always "symbol"
type Text string                     // Always "text"
type Tool string                     // Always "tool"
type UnknownError string             // Always "UnknownError"
type User string                     // Always "user"
type Wellknown string                // Always "wellknown"

func (c Agent) Default() Agent                               { return "agent" }
func (c API) Default() API                                   { return "api" }
func (c Assistant) Default() Assistant                       { return "assistant" }
func (c Completed) Default() Completed                       { return "completed" }
func (c Error) Default() Error                               { return "error" }
func (c File) Default() File                                 { return "file" }
func (c FileEdited) Default() FileEdited                     { return "file.edited" }
func (c InstallationUpdated) Default() InstallationUpdated   { return "installation.updated" }
func (c Local) Default() Local                               { return "local" }
func (c LspClientDiagnostics) Default() LspClientDiagnostics { return "lsp.client.diagnostics" }
func (c MessagePartRemoved) Default() MessagePartRemoved     { return "message.part.removed" }
func (c MessagePartUpdated) Default() MessagePartUpdated     { return "message.part.updated" }
func (c MessageRemoved) Default() MessageRemoved             { return "message.removed" }
func (c MessageUpdated) Default() MessageUpdated             { return "message.updated" }
func (c MessageAbortedError) Default() MessageAbortedError   { return "MessageAbortedError" }
func (c MessageOutputLengthError) Default() MessageOutputLengthError {
	return "MessageOutputLengthError"
}
func (c OAuth) Default() OAuth                         { return "oauth" }
func (c Patch) Default() Patch                         { return "patch" }
func (c Pending) Default() Pending                     { return "pending" }
func (c PermissionReplied) Default() PermissionReplied { return "permission.replied" }
func (c PermissionUpdated) Default() PermissionUpdated { return "permission.updated" }
func (c ProviderAuthError) Default() ProviderAuthError { return "ProviderAuthError" }
func (c Reasoning) Default() Reasoning                 { return "reasoning" }
func (c Remote) Default() Remote                       { return "remote" }
func (c Running) Default() Running                     { return "running" }
func (c ServerConnected) Default() ServerConnected     { return "server.connected" }
func (c SessionDeleted) Default() SessionDeleted       { return "session.deleted" }
func (c SessionError) Default() SessionError           { return "session.error" }
func (c SessionIdle) Default() SessionIdle             { return "session.idle" }
func (c SessionUpdated) Default() SessionUpdated       { return "session.updated" }
func (c Snapshot) Default() Snapshot                   { return "snapshot" }
func (c StepFinish) Default() StepFinish               { return "step-finish" }
func (c StepStart) Default() StepStart                 { return "step-start" }
func (c Symbol) Default() Symbol                       { return "symbol" }
func (c Text) Default() Text                           { return "text" }
func (c Tool) Default() Tool                           { return "tool" }
func (c UnknownError) Default() UnknownError           { return "UnknownError" }
func (c User) Default() User                           { return "user" }
func (c Wellknown) Default() Wellknown                 { return "wellknown" }

func (c Agent) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c API) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c Assistant) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c Completed) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c Error) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c File) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c FileEdited) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c InstallationUpdated) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Local) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c LspClientDiagnostics) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c MessagePartRemoved) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c MessagePartUpdated) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c MessageRemoved) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c MessageUpdated) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c MessageAbortedError) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c MessageOutputLengthError) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c OAuth) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c Patch) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c Pending) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c PermissionReplied) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c PermissionUpdated) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c ProviderAuthError) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Reasoning) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c Remote) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Running) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c ServerConnected) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c SessionDeleted) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c SessionError) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c SessionIdle) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c SessionUpdated) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Snapshot) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c StepFinish) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c StepStart) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c Symbol) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c Tool) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c UnknownError) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c User) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c Wellknown) MarshalJSON() ([]byte, error)                { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
