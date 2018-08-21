// Code generated by cdpgen. DO NOT EDIT.

package tracing

import (
	"encoding/json"
	"log"

	"github.com/4ydx/cdp/protocol/io"
)

const (
	EventTracingBufferUsage     = "Tracing.bufferUsage"
	EventTracingDataCollected   = "Tracing.dataCollected"
	EventTracingTracingComplete = "Tracing.tracingComplete"
)

var EventConstants = map[string]json.Unmarshaler{
	EventTracingBufferUsage:     &BufferUsageReply{},
	EventTracingDataCollected:   &DataCollectedReply{},
	EventTracingTracingComplete: &CompleteReply{},
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	return e, ok
}

// BufferUsageReply is the reply for BufferUsage events.
type BufferUsageReply struct {
	PercentFull float64 `json:"percentFull,omitempty"` // A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
	EventCount  float64 `json:"eventCount,omitempty"`  // An approximate number of events in the trace log.
	Value       float64 `json:"value,omitempty"`       // A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
}

// Unmarshal the byte array into a return value for BufferUsage in the Tracing domain.
func (a *BufferUsageReply) UnmarshalJSON(b []byte) error {
	type Copy BufferUsageReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = BufferUsageReply(*c)
	return nil
}

// BufferUsageReply returns whether or not the FrameID matches the reply value for BufferUsage in the BufferUsage domain.
func (a *BufferUsageReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: BufferUsage", err)
	}
	return true
}

// BufferUsageReply returns the FrameID for BufferUsage in the BufferUsage domain.
func (a *BufferUsageReply) GetFrameID() string {
	return ""
}

// DataCollectedReply is the reply for DataCollected events.
type DataCollectedReply struct {
	Value []json.RawMessage `json:"value"` // No description.
}

// Unmarshal the byte array into a return value for DataCollected in the Tracing domain.
func (a *DataCollectedReply) UnmarshalJSON(b []byte) error {
	type Copy DataCollectedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DataCollectedReply(*c)
	return nil
}

// DataCollectedReply returns whether or not the FrameID matches the reply value for DataCollected in the DataCollected domain.
func (a *DataCollectedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: DataCollected", err)
	}
	return true
}

// DataCollectedReply returns the FrameID for DataCollected in the DataCollected domain.
func (a *DataCollectedReply) GetFrameID() string {
	return ""
}

// CompleteReply is the reply for TracingComplete events.
type CompleteReply struct {
	Stream            io.StreamHandle   `json:"stream,omitempty"`            // A handle of the stream that holds resulting trace data.
	StreamCompression StreamCompression `json:"streamCompression,omitempty"` // Compression format of returned stream.
}

// Unmarshal the byte array into a return value for TracingComplete in the Tracing domain.
func (a *CompleteReply) UnmarshalJSON(b []byte) error {
	type Copy CompleteReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CompleteReply(*c)
	return nil
}

// CompleteReply returns whether or not the FrameID matches the reply value for TracingComplete in the TracingComplete domain.
func (a *CompleteReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: TracingComplete", err)
	}
	return true
}

// CompleteReply returns the FrameID for TracingComplete in the TracingComplete domain.
func (a *CompleteReply) GetFrameID() string {
	return ""
}
