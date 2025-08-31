# Project

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#ProjectListResponse">ProjectListResponse</a>

Methods:

- <code title="get /project">client.Project.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#ProjectListResponse">ProjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Event

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#EventListResponseUnion">EventListResponseUnion</a>

Methods:

- <code title="get /event">client.Event.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#EventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#EventListResponseUnion">EventListResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Config

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#ConfigGetResponse">ConfigGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#ConfigListProvidersResponse">ConfigListProvidersResponse</a>

Methods:

- <code title="get /config">client.Config.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#ConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#ConfigGetResponse">ConfigGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /config/providers">client.Config.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#ConfigService.ListProviders">ListProviders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#ConfigListProvidersResponse">ConfigListProvidersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Session

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#AssistantMessage">AssistantMessage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#MessageAbortedError">MessageAbortedError</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#MessageOutputLengthError">MessageOutputLengthError</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#ProviderAuthError">ProviderAuthError</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#Session">Session</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#UnknownError">UnknownError</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionSendCommandResponse">SessionSendCommandResponse</a>

Methods:

- <code title="post /session">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionNewParams">SessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /session/{id}">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /session/{id}">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionUpdateParams">SessionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /session">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /session/{id}">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/abort">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.Abort">Abort</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/init">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.Analyze">Analyze</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionAnalyzeParams">SessionAnalyzeParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /session/{id}/children">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.GetChildren">GetChildren</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/permissions/{permissionID}">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.RespondToPermission">RespondToPermission</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, permissionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionRespondToPermissionParams">SessionRespondToPermissionParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/unrevert">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.RestoreReverted">RestoreReverted</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/revert">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.Revert">Revert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionRevertParams">SessionRevertParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/shell">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.RunShell">RunShell</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionRunShellParams">SessionRunShellParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#AssistantMessage">AssistantMessage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/command">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.SendCommand">SendCommand</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionSendCommandParams">SessionSendCommandParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionSendCommandResponse">SessionSendCommandResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/summarize">client.Session.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionService.Summarize">Summarize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionSummarizeParams">SessionSummarizeParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Share

Methods:

- <code title="post /session/{id}/share">client.Session.Share.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionShareService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /session/{id}/share">client.Session.Share.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionShareService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Message

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FilePartSourceUnionParam">FilePartSourceUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FilePartSourceTextParam">FilePartSourceTextParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FilePartSourceUnion">FilePartSourceUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FilePartSourceText">FilePartSourceText</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#MessageUnion">MessageUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#PartUnion">PartUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionMessageNewResponse">SessionMessageNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionMessageGetResponse">SessionMessageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionMessageListResponse">SessionMessageListResponse</a>

Methods:

- <code title="post /session/{id}/message">client.Session.Message.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionMessageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionMessageNewParams">SessionMessageNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionMessageNewResponse">SessionMessageNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /session/{id}/message/{messageID}">client.Session.Message.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionMessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionMessageGetParams">SessionMessageGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionMessageGetResponse">SessionMessageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /session/{id}/message">client.Session.Message.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionMessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#SessionMessageListResponse">SessionMessageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Command

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#CommandListResponse">CommandListResponse</a>

Methods:

- <code title="get /command">client.Command.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#CommandService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#CommandListResponse">CommandListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Find

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#RangeParam">RangeParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#Range">Range</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FindGetResponse">FindGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FindGetSymbolResponse">FindGetSymbolResponse</a>

Methods:

- <code title="get /find">client.Find.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FindService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FindGetParams">FindGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FindGetResponse">FindGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /find/file">client.Find.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FindService.GetFile">GetFile</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FindGetFileParams">FindGetFileParams</a>) ([]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /find/symbol">client.Find.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FindService.GetSymbol">GetSymbol</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FindGetSymbolParams">FindGetSymbolParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FindGetSymbolResponse">FindGetSymbolResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# File

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FileListResponse">FileListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FileGetStatusResponse">FileGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FileReadResponse">FileReadResponse</a>

Methods:

- <code title="get /file">client.File.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FileListParams">FileListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FileListResponse">FileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /file/status">client.File.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FileService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FileGetStatusResponse">FileGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /file/content">client.File.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FileService.Read">Read</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FileReadParams">FileReadParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#FileReadResponse">FileReadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Log

Methods:

- <code title="post /log">client.Log.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#LogService.Write">Write</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#LogWriteParams">LogWriteParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Agent

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#AgentListResponse">AgentListResponse</a>

Methods:

- <code title="get /agent">client.Agent.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#AgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#AgentListResponse">AgentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tui

Methods:

- <code title="post /tui/append-prompt">client.Tui.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#TuiService.AppendPrompt">AppendPrompt</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#TuiAppendPromptParams">TuiAppendPromptParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/clear-prompt">client.Tui.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#TuiService.ClearPrompt">ClearPrompt</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/execute-command">client.Tui.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#TuiService.ExecuteCommand">ExecuteCommand</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#TuiExecuteCommandParams">TuiExecuteCommandParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/open-help">client.Tui.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#TuiService.OpenHelp">OpenHelp</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/open-models">client.Tui.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#TuiService.OpenModels">OpenModels</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/open-sessions">client.Tui.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#TuiService.OpenSessions">OpenSessions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/open-themes">client.Tui.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#TuiService.OpenThemes">OpenThemes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/show-toast">client.Tui.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#TuiService.ShowToast">ShowToast</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#TuiShowToastParams">TuiShowToastParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/submit-prompt">client.Tui.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#TuiService.SubmitPrompt">SubmitPrompt</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Auth

Methods:

- <code title="put /auth/{id}">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#AuthService.SetCredentials">SetCredentials</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go">opencode</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/opencode-go#AuthSetCredentialsParams">AuthSetCredentialsParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
