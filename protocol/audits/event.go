// Code generated by cdpgen. DO NOT EDIT.

package audits

import (
	"encoding/json"
	"log"
)

const (
	EventAuditsIssueAdded = "Audits.issueAdded"
)

type Unmarshaler func() json.Unmarshaler

var EventConstants = map[string]Unmarshaler{
	EventAuditsIssueAdded: func() json.Unmarshaler { return &IssueAddedReply{} },
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	if ok {
		return e(), ok
	}
	return nil, ok
}

// IssueAddedReply is the reply for IssueAdded events.
type IssueAddedReply struct {
	Issue InspectorIssue `json:"issue"` // No description.
}

// Unmarshal the byte array into a return value for IssueAdded in the Audits domain.
func (a *IssueAddedReply) UnmarshalJSON(b []byte) error {
	type Copy IssueAddedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = IssueAddedReply(*c)
	return nil
}

// IssueAddedReply returns whether or not the FrameID matches the reply value for IssueAdded in the Audits domain.
func (a *IssueAddedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: IssueAddedReply %s", err)
		return false, err
	}
	return true, nil
}

// IssueAddedReply returns the FrameID for IssueAdded in the Audits domain.
func (a *IssueAddedReply) GetFrameID() string {
	return ""
}
