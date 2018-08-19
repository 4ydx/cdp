// Code generated by cdpgen. DO NOT EDIT.

package log

import (
	"encoding/json"
)

const EventLogEntryAdded = "Log.entryAdded"

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