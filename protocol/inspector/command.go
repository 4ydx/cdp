// Code generated by cdpgen. DO NOT EDIT.

package inspector

import (
	"encoding/json"
	"log"
)

const (
	CommandInspectorDisable = "Inspector.disable"
	CommandInspectorEnable  = "Inspector.enable"
)

// DisableArgs represents the arguments for Disable in the Inspector domain.
type DisableArgs struct {
}

// Unmarshal the byte array into a return value for Disable in the Inspector domain.
func (a *DisableArgs) UnmarshalJSON(b []byte) error {
	type Copy DisableArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DisableArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Disable in the Inspector domain.
func (a *DisableArgs) MarshalJSON() ([]byte, error) {
	type Copy DisableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DisableReply represents the return values for Disable in the Inspector domain.
type DisableReply struct {
}

// DisableReply returns whether or not the FrameID matches the reply value for Disable in the Inspector domain.
func (a *DisableReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: DisableReply %s", err)
	}
	return true
}

// DisableReply returns the FrameID value for Disable in the Inspector domain.
func (a *DisableReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Disable in the Inspector domain.
func (a *DisableReply) UnmarshalJSON(b []byte) error {
	type Copy DisableReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DisableReply(*c)
	return nil
}

// EnableArgs represents the arguments for Enable in the Inspector domain.
type EnableArgs struct {
}

// Unmarshal the byte array into a return value for Enable in the Inspector domain.
func (a *EnableArgs) UnmarshalJSON(b []byte) error {
	type Copy EnableArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = EnableArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Enable in the Inspector domain.
func (a *EnableArgs) MarshalJSON() ([]byte, error) {
	type Copy EnableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// EnableReply represents the return values for Enable in the Inspector domain.
type EnableReply struct {
}

// EnableReply returns whether or not the FrameID matches the reply value for Enable in the Inspector domain.
func (a *EnableReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: EnableReply %s", err)
	}
	return true
}

// EnableReply returns the FrameID value for Enable in the Inspector domain.
func (a *EnableReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Enable in the Inspector domain.
func (a *EnableReply) UnmarshalJSON(b []byte) error {
	type Copy EnableReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = EnableReply(*c)
	return nil
}
