// Code generated by cdpgen. DO NOT EDIT.

package tethering

import (
	"encoding/json"
	"log"
)

const (
	CommandTetheringBind   = "Tethering.bind"
	CommandTetheringUnbind = "Tethering.unbind"
)

// BindArgs represents the arguments for Bind in the Tethering domain.
type BindArgs struct {
	Port int `json:"port"` // Port number to bind.
}

// Unmarshal the byte array into a return value for Bind in the Tethering domain.
func (a *BindArgs) UnmarshalJSON(b []byte) error {
	type Copy BindArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = BindArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Bind in the Tethering domain.
func (a *BindArgs) MarshalJSON() ([]byte, error) {
	type Copy BindArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// BindReply represents the return values for Bind in the Tethering domain.
type BindReply struct {
}

// BindReply returns whether or not the FrameID matches the reply value for Bind in the Tethering domain.
func (a *BindReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: BindReply %s", err)
		return false, err
	}
	return true, nil
}

// BindReply returns the FrameID value for Bind in the Tethering domain.
func (a *BindReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Bind in the Tethering domain.
func (a *BindReply) UnmarshalJSON(b []byte) error {
	type Copy BindReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = BindReply(*c)
	return nil
}

// UnbindArgs represents the arguments for Unbind in the Tethering domain.
type UnbindArgs struct {
	Port int `json:"port"` // Port number to unbind.
}

// Unmarshal the byte array into a return value for Unbind in the Tethering domain.
func (a *UnbindArgs) UnmarshalJSON(b []byte) error {
	type Copy UnbindArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = UnbindArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Unbind in the Tethering domain.
func (a *UnbindArgs) MarshalJSON() ([]byte, error) {
	type Copy UnbindArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// UnbindReply represents the return values for Unbind in the Tethering domain.
type UnbindReply struct {
}

// UnbindReply returns whether or not the FrameID matches the reply value for Unbind in the Tethering domain.
func (a *UnbindReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: UnbindReply %s", err)
		return false, err
	}
	return true, nil
}

// UnbindReply returns the FrameID value for Unbind in the Tethering domain.
func (a *UnbindReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Unbind in the Tethering domain.
func (a *UnbindReply) UnmarshalJSON(b []byte) error {
	type Copy UnbindReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = UnbindReply(*c)
	return nil
}
