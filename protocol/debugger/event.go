// Code generated by cdpgen. DO NOT EDIT.

package debugger

import (
	"encoding/json"
	"log"

	"github.com/4ydx/cdp/protocol/runtime"
)

const (
	EventDebuggerBreakpointResolved  = "Debugger.breakpointResolved"
	EventDebuggerPaused              = "Debugger.paused"
	EventDebuggerResumed             = "Debugger.resumed"
	EventDebuggerScriptFailedToParse = "Debugger.scriptFailedToParse"
	EventDebuggerScriptParsed        = "Debugger.scriptParsed"
)

var EventConstants = map[string]json.Unmarshaler{
	EventDebuggerBreakpointResolved:  &BreakpointResolvedReply{},
	EventDebuggerPaused:              &PausedReply{},
	EventDebuggerResumed:             &ResumedReply{},
	EventDebuggerScriptFailedToParse: &ScriptFailedToParseReply{},
	EventDebuggerScriptParsed:        &ScriptParsedReply{},
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	return e, ok
}

// BreakpointResolvedReply is the reply for BreakpointResolved events.
type BreakpointResolvedReply struct {
	BreakpointID BreakpointID `json:"breakpointId"` // Breakpoint unique identifier.
	Location     Location     `json:"location"`     // Actual breakpoint location.
}

// Unmarshal the byte array into a return value for BreakpointResolved in the Debugger domain.
func (a *BreakpointResolvedReply) UnmarshalJSON(b []byte) error {
	type Copy BreakpointResolvedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = BreakpointResolvedReply(*c)
	return nil
}

// BreakpointResolvedReply returns whether or not the FrameID matches the reply value for BreakpointResolved in the BreakpointResolved domain.
func (a *BreakpointResolvedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: BreakpointResolved", err)
	}
	return true
}

// BreakpointResolvedReply returns the FrameID for BreakpointResolved in the BreakpointResolved domain.
func (a *BreakpointResolvedReply) GetFrameID() string {
	return ""
}

// PausedReply is the reply for Paused events.
type PausedReply struct {
	CallFrames []CallFrame `json:"callFrames"` // Call stack the virtual machine stopped on.
	// Reason Pause reason.
	//
	// Values: "XHR", "DOM", "EventListener", "exception", "assert", "debugCommand", "promiseRejection", "OOM", "other", "ambiguous".
	Reason          string             `json:"reason"`
	Data            json.RawMessage    `json:"data,omitempty"`            // Object containing break-specific auxiliary properties.
	HitBreakpoints  []string           `json:"hitBreakpoints,omitempty"`  // Hit breakpoints IDs
	AsyncStackTrace runtime.StackTrace `json:"asyncStackTrace,omitempty"` // Async stack trace, if any.
	// AsyncStackTraceID Async stack trace, if any.
	//
	// Note: This property is experimental.
	AsyncStackTraceID runtime.StackTraceID `json:"asyncStackTraceId,omitempty"`
	// AsyncCallStackTraceID Just scheduled async call will have this
	// stack trace as parent stack during async execution. This field is
	// available only after `Debugger.stepInto` call with `breakOnAsynCall`
	// flag.
	//
	// Note: This property is experimental.
	AsyncCallStackTraceID runtime.StackTraceID `json:"asyncCallStackTraceId,omitempty"`
}

// Unmarshal the byte array into a return value for Paused in the Debugger domain.
func (a *PausedReply) UnmarshalJSON(b []byte) error {
	type Copy PausedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = PausedReply(*c)
	return nil
}

// PausedReply returns whether or not the FrameID matches the reply value for Paused in the Paused domain.
func (a *PausedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: Paused", err)
	}
	return true
}

// PausedReply returns the FrameID for Paused in the Paused domain.
func (a *PausedReply) GetFrameID() string {
	return ""
}

// ResumedReply is the reply for Resumed events.
type ResumedReply struct {
}

// Unmarshal the byte array into a return value for Resumed in the Debugger domain.
func (a *ResumedReply) UnmarshalJSON(b []byte) error {
	type Copy ResumedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ResumedReply(*c)
	return nil
}

// ResumedReply returns whether or not the FrameID matches the reply value for Resumed in the Resumed domain.
func (a *ResumedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: Resumed", err)
	}
	return true
}

// ResumedReply returns the FrameID for Resumed in the Resumed domain.
func (a *ResumedReply) GetFrameID() string {
	return ""
}

// ScriptFailedToParseReply is the reply for ScriptFailedToParse events.
type ScriptFailedToParseReply struct {
	ScriptID                runtime.ScriptID           `json:"scriptId"`                          // Identifier of the script parsed.
	URL                     string                     `json:"url"`                               // URL or name of the script parsed (if any).
	StartLine               int                        `json:"startLine"`                         // Line offset of the script within the resource with given URL (for script tags).
	StartColumn             int                        `json:"startColumn"`                       // Column offset of the script within the resource with given URL.
	EndLine                 int                        `json:"endLine"`                           // Last line of the script.
	EndColumn               int                        `json:"endColumn"`                         // Length of the last line of the script.
	ExecutionContextID      runtime.ExecutionContextID `json:"executionContextId"`                // Specifies script creation context.
	Hash                    string                     `json:"hash"`                              // Content hash of the script.
	ExecutionContextAuxData json.RawMessage            `json:"executionContextAuxData,omitempty"` // Embedder-specific auxiliary data.
	SourceMapURL            string                     `json:"sourceMapURL,omitempty"`            // URL of source map associated with script (if any).
	HasSourceURL            bool                       `json:"hasSourceURL,omitempty"`            // True, if this script has sourceURL.
	IsModule                bool                       `json:"isModule,omitempty"`                // True, if this script is ES6 module.
	Length                  int                        `json:"length,omitempty"`                  // This script length.
	// StackTrace JavaScript top stack frame of where the script parsed
	// event was triggered if available.
	//
	// Note: This property is experimental.
	StackTrace runtime.StackTrace `json:"stackTrace,omitempty"`
}

// Unmarshal the byte array into a return value for ScriptFailedToParse in the Debugger domain.
func (a *ScriptFailedToParseReply) UnmarshalJSON(b []byte) error {
	type Copy ScriptFailedToParseReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ScriptFailedToParseReply(*c)
	return nil
}

// ScriptFailedToParseReply returns whether or not the FrameID matches the reply value for ScriptFailedToParse in the ScriptFailedToParse domain.
func (a *ScriptFailedToParseReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: ScriptFailedToParse", err)
	}
	return true
}

// ScriptFailedToParseReply returns the FrameID for ScriptFailedToParse in the ScriptFailedToParse domain.
func (a *ScriptFailedToParseReply) GetFrameID() string {
	return ""
}

// ScriptParsedReply is the reply for ScriptParsed events.
type ScriptParsedReply struct {
	ScriptID                runtime.ScriptID           `json:"scriptId"`                          // Identifier of the script parsed.
	URL                     string                     `json:"url"`                               // URL or name of the script parsed (if any).
	StartLine               int                        `json:"startLine"`                         // Line offset of the script within the resource with given URL (for script tags).
	StartColumn             int                        `json:"startColumn"`                       // Column offset of the script within the resource with given URL.
	EndLine                 int                        `json:"endLine"`                           // Last line of the script.
	EndColumn               int                        `json:"endColumn"`                         // Length of the last line of the script.
	ExecutionContextID      runtime.ExecutionContextID `json:"executionContextId"`                // Specifies script creation context.
	Hash                    string                     `json:"hash"`                              // Content hash of the script.
	ExecutionContextAuxData json.RawMessage            `json:"executionContextAuxData,omitempty"` // Embedder-specific auxiliary data.
	// IsLiveEdit True, if this script is generated as a result of the
	// live edit operation.
	//
	// Note: This property is experimental.
	IsLiveEdit   bool   `json:"isLiveEdit,omitempty"`
	SourceMapURL string `json:"sourceMapURL,omitempty"` // URL of source map associated with script (if any).
	HasSourceURL bool   `json:"hasSourceURL,omitempty"` // True, if this script has sourceURL.
	IsModule     bool   `json:"isModule,omitempty"`     // True, if this script is ES6 module.
	Length       int    `json:"length,omitempty"`       // This script length.
	// StackTrace JavaScript top stack frame of where the script parsed
	// event was triggered if available.
	//
	// Note: This property is experimental.
	StackTrace runtime.StackTrace `json:"stackTrace,omitempty"`
}

// Unmarshal the byte array into a return value for ScriptParsed in the Debugger domain.
func (a *ScriptParsedReply) UnmarshalJSON(b []byte) error {
	type Copy ScriptParsedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ScriptParsedReply(*c)
	return nil
}

// ScriptParsedReply returns whether or not the FrameID matches the reply value for ScriptParsed in the ScriptParsed domain.
func (a *ScriptParsedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: ScriptParsed", err)
	}
	return true
}

// ScriptParsedReply returns the FrameID for ScriptParsed in the ScriptParsed domain.
func (a *ScriptParsedReply) GetFrameID() string {
	return ""
}
