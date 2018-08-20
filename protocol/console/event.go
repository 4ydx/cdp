// Code generated by cdpgen. DO NOT EDIT.

package console

import (
	"encoding/json"
)

const (
	EventConsoleMessageAdded = "Console.messageAdded"
)

var EventConstants = map[string]json.Unmarshaler{
	EventConsoleMessageAdded: &MessageAddedReply{},
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	return e, ok
}

// MessageAddedReply is the reply for MessageAdded events.
type MessageAddedReply struct {
	Message Message `json:"message"` // Console message that has been added.
}

// Unmarshal the byte array into a return value for MessageAdded in the Console domain.
func (a *MessageAddedReply) UnmarshalJSON(b []byte) error {
	type Copy MessageAddedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = MessageAddedReply(*c)
	return nil
}
