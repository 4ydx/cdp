// Code generated by cdpgen. DO NOT EDIT.

package cast

import (
	"encoding/json"
	"log"
)

const (
	EventCastSinksUpdated = "Cast.sinksUpdated"
	EventCastIssueUpdated = "Cast.issueUpdated"
)

type Unmarshaler func() json.Unmarshaler

var EventConstants = map[string]Unmarshaler{
	EventCastSinksUpdated: func() json.Unmarshaler { return &SinksUpdatedReply{} },
	EventCastIssueUpdated: func() json.Unmarshaler { return &IssueUpdatedReply{} },
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	if ok {
		return e(), ok
	}
	return nil, ok
}

// SinksUpdatedReply is the reply for SinksUpdated events.
type SinksUpdatedReply struct {
	Sinks []Sink `json:"sinks"` // No description.
}

// Unmarshal the byte array into a return value for SinksUpdated in the Cast domain.
func (a *SinksUpdatedReply) UnmarshalJSON(b []byte) error {
	type Copy SinksUpdatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SinksUpdatedReply(*c)
	return nil
}

// SinksUpdatedReply returns whether or not the FrameID matches the reply value for SinksUpdated in the Cast domain.
func (a *SinksUpdatedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: SinksUpdatedReply %s", err)
		return false, err
	}
	return true, nil
}

// SinksUpdatedReply returns the FrameID for SinksUpdated in the Cast domain.
func (a *SinksUpdatedReply) GetFrameID() string {
	return ""
}

// IssueUpdatedReply is the reply for IssueUpdated events.
type IssueUpdatedReply struct {
	IssueMessage string `json:"issueMessage"` // No description.
}

// Unmarshal the byte array into a return value for IssueUpdated in the Cast domain.
func (a *IssueUpdatedReply) UnmarshalJSON(b []byte) error {
	type Copy IssueUpdatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = IssueUpdatedReply(*c)
	return nil
}

// IssueUpdatedReply returns whether or not the FrameID matches the reply value for IssueUpdated in the Cast domain.
func (a *IssueUpdatedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: IssueUpdatedReply %s", err)
		return false, err
	}
	return true, nil
}

// IssueUpdatedReply returns the FrameID for IssueUpdated in the Cast domain.
func (a *IssueUpdatedReply) GetFrameID() string {
	return ""
}
