// Code generated by cdpgen. DO NOT EDIT.

package audits

import (
	"encoding/json"
	"log"

	"github.com/4ydx/cdp/protocol/network"
)

const (
	CommandAuditsGetEncodedResponse = "Audits.getEncodedResponse"
	CommandAuditsDisable            = "Audits.disable"
	CommandAuditsEnable             = "Audits.enable"
)

// GetEncodedResponseArgs represents the arguments for GetEncodedResponse in the Audits domain.
type GetEncodedResponseArgs struct {
	RequestID network.RequestID `json:"requestId"` // Identifier of the network request to get content for.
	// Encoding The encoding to use.
	//
	// Values: "webp", "jpeg", "png".
	Encoding string  `json:"encoding"`
	Quality  float64 `json:"quality,omitempty"`  // The quality of the encoding (0-1). (defaults to 1)
	SizeOnly bool    `json:"sizeOnly,omitempty"` // Whether to only return the size information (defaults to false).
}

// Unmarshal the byte array into a return value for GetEncodedResponse in the Audits domain.
func (a *GetEncodedResponseArgs) UnmarshalJSON(b []byte) error {
	type Copy GetEncodedResponseArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetEncodedResponseArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetEncodedResponse in the Audits domain.
func (a *GetEncodedResponseArgs) MarshalJSON() ([]byte, error) {
	type Copy GetEncodedResponseArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetEncodedResponseReply represents the return values for GetEncodedResponse in the Audits domain.
type GetEncodedResponseReply struct {
	Body         string `json:"body,omitempty"` // The encoded body as a base64 string. Omitted if sizeOnly is true.
	OriginalSize int    `json:"originalSize"`   // Size before re-encoding.
	EncodedSize  int    `json:"encodedSize"`    // Size after re-encoding.
}

// GetEncodedResponseReply returns whether or not the FrameID matches the reply value for GetEncodedResponse in the Audits domain.
func (a *GetEncodedResponseReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetEncodedResponseReply %s", err)
		return false, err
	}
	return true, nil
}

// GetEncodedResponseReply returns the FrameID value for GetEncodedResponse in the Audits domain.
func (a *GetEncodedResponseReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetEncodedResponse in the Audits domain.
func (a *GetEncodedResponseReply) UnmarshalJSON(b []byte) error {
	type Copy GetEncodedResponseReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetEncodedResponseReply(*c)
	return nil
}

// DisableArgs represents the arguments for Disable in the Audits domain.
type DisableArgs struct {
}

// Unmarshal the byte array into a return value for Disable in the Audits domain.
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

// Marshall the byte array into a return value for Disable in the Audits domain.
func (a *DisableArgs) MarshalJSON() ([]byte, error) {
	type Copy DisableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DisableReply represents the return values for Disable in the Audits domain.
type DisableReply struct {
}

// DisableReply returns whether or not the FrameID matches the reply value for Disable in the Audits domain.
func (a *DisableReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: DisableReply %s", err)
		return false, err
	}
	return true, nil
}

// DisableReply returns the FrameID value for Disable in the Audits domain.
func (a *DisableReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Disable in the Audits domain.
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

// EnableArgs represents the arguments for Enable in the Audits domain.
type EnableArgs struct {
}

// Unmarshal the byte array into a return value for Enable in the Audits domain.
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

// Marshall the byte array into a return value for Enable in the Audits domain.
func (a *EnableArgs) MarshalJSON() ([]byte, error) {
	type Copy EnableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// EnableReply represents the return values for Enable in the Audits domain.
type EnableReply struct {
}

// EnableReply returns whether or not the FrameID matches the reply value for Enable in the Audits domain.
func (a *EnableReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: EnableReply %s", err)
		return false, err
	}
	return true, nil
}

// EnableReply returns the FrameID value for Enable in the Audits domain.
func (a *EnableReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Enable in the Audits domain.
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
