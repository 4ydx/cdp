// Code generated by cdpgen. DO NOT EDIT.

package log

import (
	"encoding/json"
	"log"
)

const (
	EventLogEntryAdded = "Log.entryAdded"
)

type Unmarshaler func() json.Unmarshaler

var EventConstants = map[string]Unmarshaler{
	EventLogEntryAdded: func() json.Unmarshaler { return &EntryAddedReply{} },
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	if ok {
		return e(), ok
	}
	return nil, ok
}

// EntryAddedReply is the reply for EntryAdded events.
type EntryAddedReply struct {
	Entry Entry `json:"entry"` // The entry.
}

// Unmarshal the byte array into a return value for EntryAdded in the Log domain.
func (a *EntryAddedReply) UnmarshalJSON(b []byte) error {
	type Copy EntryAddedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = EntryAddedReply(*c)
	return nil
}

// EntryAddedReply returns whether or not the FrameID matches the reply value for EntryAdded in the Log domain.
func (a *EntryAddedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: EntryAddedReply %s", err)
		return false, err
	}
	return true, nil
}

// EntryAddedReply returns the FrameID for EntryAdded in the Log domain.
func (a *EntryAddedReply) GetFrameID() string {
	return ""
}
