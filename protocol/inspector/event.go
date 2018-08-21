// Code generated by cdpgen. DO NOT EDIT.

package inspector

import (
	"encoding/json"
	"log"
)

const (
	EventInspectorDetached                 = "Inspector.detached"
	EventInspectorTargetCrashed            = "Inspector.targetCrashed"
	EventInspectorTargetReloadedAfterCrash = "Inspector.targetReloadedAfterCrash"
)

type Unmarshaler func() json.Unmarshaler

var EventConstants = map[string]Unmarshaler{
	EventInspectorDetached:                 func() json.Unmarshaler { return &DetachedReply{} },
	EventInspectorTargetCrashed:            func() json.Unmarshaler { return &TargetCrashedReply{} },
	EventInspectorTargetReloadedAfterCrash: func() json.Unmarshaler { return &TargetReloadedAfterCrashReply{} },
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	if ok {
		return e(), ok
	}
	return nil, ok
}

// DetachedReply is the reply for Detached events.
type DetachedReply struct {
	Reason string `json:"reason"` // The reason why connection has been terminated.
}

// Unmarshal the byte array into a return value for Detached in the Inspector domain.
func (a *DetachedReply) UnmarshalJSON(b []byte) error {
	type Copy DetachedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DetachedReply(*c)
	return nil
}

// DetachedReply returns whether or not the FrameID matches the reply value for Detached in the Inspector domain.
func (a *DetachedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: DetachedReply %s", err)
	}
	return true
}

// DetachedReply returns the FrameID for Detached in the Inspector domain.
func (a *DetachedReply) GetFrameID() string {
	return ""
}

// TargetCrashedReply is the reply for TargetCrashed events.
type TargetCrashedReply struct {
}

// Unmarshal the byte array into a return value for TargetCrashed in the Inspector domain.
func (a *TargetCrashedReply) UnmarshalJSON(b []byte) error {
	type Copy TargetCrashedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = TargetCrashedReply(*c)
	return nil
}

// TargetCrashedReply returns whether or not the FrameID matches the reply value for TargetCrashed in the Inspector domain.
func (a *TargetCrashedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: TargetCrashedReply %s", err)
	}
	return true
}

// TargetCrashedReply returns the FrameID for TargetCrashed in the Inspector domain.
func (a *TargetCrashedReply) GetFrameID() string {
	return ""
}

// TargetReloadedAfterCrashReply is the reply for TargetReloadedAfterCrash events.
type TargetReloadedAfterCrashReply struct {
}

// Unmarshal the byte array into a return value for TargetReloadedAfterCrash in the Inspector domain.
func (a *TargetReloadedAfterCrashReply) UnmarshalJSON(b []byte) error {
	type Copy TargetReloadedAfterCrashReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = TargetReloadedAfterCrashReply(*c)
	return nil
}

// TargetReloadedAfterCrashReply returns whether or not the FrameID matches the reply value for TargetReloadedAfterCrash in the Inspector domain.
func (a *TargetReloadedAfterCrashReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: TargetReloadedAfterCrashReply %s", err)
	}
	return true
}

// TargetReloadedAfterCrashReply returns the FrameID for TargetReloadedAfterCrash in the Inspector domain.
func (a *TargetReloadedAfterCrashReply) GetFrameID() string {
	return ""
}
