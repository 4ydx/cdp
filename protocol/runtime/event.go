// Code generated by cdpgen. DO NOT EDIT.

package runtime

import (
	"encoding/json"
	"log"
)

const (
	EventRuntimeBindingCalled             = "Runtime.bindingCalled"
	EventRuntimeConsoleAPICalled          = "Runtime.consoleAPICalled"
	EventRuntimeExceptionRevoked          = "Runtime.exceptionRevoked"
	EventRuntimeExceptionThrown           = "Runtime.exceptionThrown"
	EventRuntimeExecutionContextCreated   = "Runtime.executionContextCreated"
	EventRuntimeExecutionContextDestroyed = "Runtime.executionContextDestroyed"
	EventRuntimeExecutionContextsCleared  = "Runtime.executionContextsCleared"
	EventRuntimeInspectRequested          = "Runtime.inspectRequested"
)

var EventConstants = map[string]json.Unmarshaler{
	EventRuntimeBindingCalled:             &BindingCalledReply{},
	EventRuntimeConsoleAPICalled:          &ConsoleAPICalledReply{},
	EventRuntimeExceptionRevoked:          &ExceptionRevokedReply{},
	EventRuntimeExceptionThrown:           &ExceptionThrownReply{},
	EventRuntimeExecutionContextCreated:   &ExecutionContextCreatedReply{},
	EventRuntimeExecutionContextDestroyed: &ExecutionContextDestroyedReply{},
	EventRuntimeExecutionContextsCleared:  &ExecutionContextsClearedReply{},
	EventRuntimeInspectRequested:          &InspectRequestedReply{},
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	return e, ok
}

// BindingCalledReply is the reply for BindingCalled events.
type BindingCalledReply struct {
	Name               string             `json:"name"`               // No description.
	Payload            string             `json:"payload"`            // No description.
	ExecutionContextID ExecutionContextID `json:"executionContextId"` // Identifier of the context where the call was made.
}

// Unmarshal the byte array into a return value for BindingCalled in the Runtime domain.
func (a *BindingCalledReply) UnmarshalJSON(b []byte) error {
	type Copy BindingCalledReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = BindingCalledReply(*c)
	return nil
}

// BindingCalledReply returns whether or not the FrameID matches the reply value for BindingCalled in the BindingCalled domain.
func (a *BindingCalledReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: BindingCalled", err)
	}
	return true
}

// BindingCalledReply returns the FrameID for BindingCalled in the BindingCalled domain.
func (a *BindingCalledReply) GetFrameID() string {
	return ""
}

// ConsoleAPICalledReply is the reply for ConsoleAPICalled events.
type ConsoleAPICalledReply struct {
	// Type Type of the call.
	//
	// Values: "log", "debug", "info", "error", "warning", "dir", "dirxml", "table", "trace", "clear", "startGroup", "startGroupCollapsed", "endGroup", "assert", "profile", "profileEnd", "count", "timeEnd".
	Type               string             `json:"type"`
	Args               []RemoteObject     `json:"args"`                 // Call arguments.
	ExecutionContextID ExecutionContextID `json:"executionContextId"`   // Identifier of the context where the call was made.
	Timestamp          Timestamp          `json:"timestamp"`            // Call timestamp.
	StackTrace         StackTrace         `json:"stackTrace,omitempty"` // Stack trace captured when the call was made.
	// Context Console context descriptor for calls on non-default console
	// context (not console.*): 'anonymous#unique-logger-id' for call on
	// unnamed context, 'name#unique-logger-id' for call on named context.
	//
	// Note: This property is experimental.
	Context string `json:"context,omitempty"`
}

// Unmarshal the byte array into a return value for ConsoleAPICalled in the Runtime domain.
func (a *ConsoleAPICalledReply) UnmarshalJSON(b []byte) error {
	type Copy ConsoleAPICalledReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ConsoleAPICalledReply(*c)
	return nil
}

// ConsoleAPICalledReply returns whether or not the FrameID matches the reply value for ConsoleAPICalled in the ConsoleAPICalled domain.
func (a *ConsoleAPICalledReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: ConsoleAPICalled", err)
	}
	return true
}

// ConsoleAPICalledReply returns the FrameID for ConsoleAPICalled in the ConsoleAPICalled domain.
func (a *ConsoleAPICalledReply) GetFrameID() string {
	return ""
}

// ExceptionRevokedReply is the reply for ExceptionRevoked events.
type ExceptionRevokedReply struct {
	Reason      string `json:"reason"`      // Reason describing why exception was revoked.
	ExceptionID int    `json:"exceptionId"` // The id of revoked exception, as reported in `exceptionThrown`.
}

// Unmarshal the byte array into a return value for ExceptionRevoked in the Runtime domain.
func (a *ExceptionRevokedReply) UnmarshalJSON(b []byte) error {
	type Copy ExceptionRevokedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ExceptionRevokedReply(*c)
	return nil
}

// ExceptionRevokedReply returns whether or not the FrameID matches the reply value for ExceptionRevoked in the ExceptionRevoked domain.
func (a *ExceptionRevokedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: ExceptionRevoked", err)
	}
	return true
}

// ExceptionRevokedReply returns the FrameID for ExceptionRevoked in the ExceptionRevoked domain.
func (a *ExceptionRevokedReply) GetFrameID() string {
	return ""
}

// ExceptionThrownReply is the reply for ExceptionThrown events.
type ExceptionThrownReply struct {
	Timestamp        Timestamp        `json:"timestamp"`        // Timestamp of the exception.
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"` // No description.
}

// Unmarshal the byte array into a return value for ExceptionThrown in the Runtime domain.
func (a *ExceptionThrownReply) UnmarshalJSON(b []byte) error {
	type Copy ExceptionThrownReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ExceptionThrownReply(*c)
	return nil
}

// ExceptionThrownReply returns whether or not the FrameID matches the reply value for ExceptionThrown in the ExceptionThrown domain.
func (a *ExceptionThrownReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: ExceptionThrown", err)
	}
	return true
}

// ExceptionThrownReply returns the FrameID for ExceptionThrown in the ExceptionThrown domain.
func (a *ExceptionThrownReply) GetFrameID() string {
	return ""
}

// ExecutionContextCreatedReply is the reply for ExecutionContextCreated events.
type ExecutionContextCreatedReply struct {
	Context ExecutionContextDescription `json:"context"` // A newly created execution context.
}

// Unmarshal the byte array into a return value for ExecutionContextCreated in the Runtime domain.
func (a *ExecutionContextCreatedReply) UnmarshalJSON(b []byte) error {
	type Copy ExecutionContextCreatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ExecutionContextCreatedReply(*c)
	return nil
}

// ExecutionContextCreatedReply returns whether or not the FrameID matches the reply value for ExecutionContextCreated in the ExecutionContextCreated domain.
func (a *ExecutionContextCreatedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: ExecutionContextCreated", err)
	}
	return true
}

// ExecutionContextCreatedReply returns the FrameID for ExecutionContextCreated in the ExecutionContextCreated domain.
func (a *ExecutionContextCreatedReply) GetFrameID() string {
	return ""
}

// ExecutionContextDestroyedReply is the reply for ExecutionContextDestroyed events.
type ExecutionContextDestroyedReply struct {
	ExecutionContextID ExecutionContextID `json:"executionContextId"` // Id of the destroyed context
}

// Unmarshal the byte array into a return value for ExecutionContextDestroyed in the Runtime domain.
func (a *ExecutionContextDestroyedReply) UnmarshalJSON(b []byte) error {
	type Copy ExecutionContextDestroyedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ExecutionContextDestroyedReply(*c)
	return nil
}

// ExecutionContextDestroyedReply returns whether or not the FrameID matches the reply value for ExecutionContextDestroyed in the ExecutionContextDestroyed domain.
func (a *ExecutionContextDestroyedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: ExecutionContextDestroyed", err)
	}
	return true
}

// ExecutionContextDestroyedReply returns the FrameID for ExecutionContextDestroyed in the ExecutionContextDestroyed domain.
func (a *ExecutionContextDestroyedReply) GetFrameID() string {
	return ""
}

// ExecutionContextsClearedReply is the reply for ExecutionContextsCleared events.
type ExecutionContextsClearedReply struct {
}

// Unmarshal the byte array into a return value for ExecutionContextsCleared in the Runtime domain.
func (a *ExecutionContextsClearedReply) UnmarshalJSON(b []byte) error {
	type Copy ExecutionContextsClearedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ExecutionContextsClearedReply(*c)
	return nil
}

// ExecutionContextsClearedReply returns whether or not the FrameID matches the reply value for ExecutionContextsCleared in the ExecutionContextsCleared domain.
func (a *ExecutionContextsClearedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: ExecutionContextsCleared", err)
	}
	return true
}

// ExecutionContextsClearedReply returns the FrameID for ExecutionContextsCleared in the ExecutionContextsCleared domain.
func (a *ExecutionContextsClearedReply) GetFrameID() string {
	return ""
}

// InspectRequestedReply is the reply for InspectRequested events.
type InspectRequestedReply struct {
	Object RemoteObject    `json:"object"` // No description.
	Hints  json.RawMessage `json:"hints"`  // No description.
}

// Unmarshal the byte array into a return value for InspectRequested in the Runtime domain.
func (a *InspectRequestedReply) UnmarshalJSON(b []byte) error {
	type Copy InspectRequestedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = InspectRequestedReply(*c)
	return nil
}

// InspectRequestedReply returns whether or not the FrameID matches the reply value for InspectRequested in the InspectRequested domain.
func (a *InspectRequestedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: InspectRequested", err)
	}
	return true
}

// InspectRequestedReply returns the FrameID for InspectRequested in the InspectRequested domain.
func (a *InspectRequestedReply) GetFrameID() string {
	return ""
}
