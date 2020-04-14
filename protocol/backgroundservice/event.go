// Code generated by cdpgen. DO NOT EDIT.

package backgroundservice

import (
	"encoding/json"
	"log"
)

const (
	EventBackgroundServiceRecordingStateChanged          = "BackgroundService.recordingStateChanged"
	EventBackgroundServiceBackgroundServiceEventReceived = "BackgroundService.backgroundServiceEventReceived"
)

type Unmarshaler func() json.Unmarshaler

var EventConstants = map[string]Unmarshaler{
	EventBackgroundServiceRecordingStateChanged:          func() json.Unmarshaler { return &RecordingStateChangedReply{} },
	EventBackgroundServiceBackgroundServiceEventReceived: func() json.Unmarshaler { return &EventReceivedReply{} },
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	if ok {
		return e(), ok
	}
	return nil, ok
}

// RecordingStateChangedReply is the reply for RecordingStateChanged events.
type RecordingStateChangedReply struct {
	IsRecording bool        `json:"isRecording"` // No description.
	Service     ServiceName `json:"service"`     // No description.
}

// Unmarshal the byte array into a return value for RecordingStateChanged in the BackgroundService domain.
func (a *RecordingStateChangedReply) UnmarshalJSON(b []byte) error {
	type Copy RecordingStateChangedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RecordingStateChangedReply(*c)
	return nil
}

// RecordingStateChangedReply returns whether or not the FrameID matches the reply value for RecordingStateChanged in the BackgroundService domain.
func (a *RecordingStateChangedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: RecordingStateChangedReply %s", err)
		return false, err
	}
	return true, nil
}

// RecordingStateChangedReply returns the FrameID for RecordingStateChanged in the BackgroundService domain.
func (a *RecordingStateChangedReply) GetFrameID() string {
	return ""
}

// EventReceivedReply is the reply for BackgroundServiceEventReceived events.
type EventReceivedReply struct {
	BackgroundServiceEvent Event `json:"backgroundServiceEvent"` // No description.
}

// Unmarshal the byte array into a return value for BackgroundServiceEventReceived in the BackgroundService domain.
func (a *EventReceivedReply) UnmarshalJSON(b []byte) error {
	type Copy EventReceivedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = EventReceivedReply(*c)
	return nil
}

// EventReceivedReply returns whether or not the FrameID matches the reply value for BackgroundServiceEventReceived in the BackgroundService domain.
func (a *EventReceivedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: EventReceivedReply %s", err)
		return false, err
	}
	return true, nil
}

// EventReceivedReply returns the FrameID for BackgroundServiceEventReceived in the BackgroundService domain.
func (a *EventReceivedReply) GetFrameID() string {
	return ""
}
