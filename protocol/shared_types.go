// Code generated by cdpgen. DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"time"
)

// ContextID
//
// Note: This type is experimental.
type ContextID string

// TimeSinceEpoch UTC time in seconds, counted from January 1, 1970.
type TimeSinceEpoch float64

// String calls (time.Time).String().
func (t TimeSinceEpoch) String() string {
	return t.Time().String()
}

// Time parses the Unix time with millisecond accuracy.
func (t TimeSinceEpoch) Time() time.Time {
	secs := int64(t)
	// The Unix time in t only has ms accuracy.
	ms := int64((float64(t) - float64(secs)) * 1000000)
	return time.Unix(secs, ms*1000)
}

// MarshalJSON implements json.Marshaler. Encodes to null if t is zero.
func (t TimeSinceEpoch) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	f := float64(t)
	return json.Marshal(&f)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TimeSinceEpoch) UnmarshalJSON(data []byte) error {
	*t = 0
	if len(data) == 0 {
		return nil
	}
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return errors.New("shared.TimeSinceEpoch: " + err.Error())
	}
	*t = TimeSinceEpoch(f)
	return nil
}

var _ json.Marshaler = (*TimeSinceEpoch)(nil)
var _ json.Unmarshaler = (*TimeSinceEpoch)(nil)

// ResourceType Resource type as it was perceived by the rendering engine.
type ResourceType string

// ResourceType as enums.
const (
	ResourceTypeNotSet             ResourceType = ""
	ResourceTypeDocument           ResourceType = "Document"
	ResourceTypeStylesheet         ResourceType = "Stylesheet"
	ResourceTypeImage              ResourceType = "Image"
	ResourceTypeMedia              ResourceType = "Media"
	ResourceTypeFont               ResourceType = "Font"
	ResourceTypeScript             ResourceType = "Script"
	ResourceTypeTextTrack          ResourceType = "TextTrack"
	ResourceTypeXHR                ResourceType = "XHR"
	ResourceTypeFetch              ResourceType = "Fetch"
	ResourceTypeEventSource        ResourceType = "EventSource"
	ResourceTypeWebSocket          ResourceType = "WebSocket"
	ResourceTypeManifest           ResourceType = "Manifest"
	ResourceTypeSignedExchange     ResourceType = "SignedExchange"
	ResourceTypePing               ResourceType = "Ping"
	ResourceTypeCSPViolationReport ResourceType = "CSPViolationReport"
	ResourceTypeOther              ResourceType = "Other"
)

func (e ResourceType) Valid() bool {
	switch e {
	case "Document", "Stylesheet", "Image", "Media", "Font", "Script", "TextTrack", "XHR", "Fetch", "EventSource", "WebSocket", "Manifest", "SignedExchange", "Ping", "CSPViolationReport", "Other":
		return true
	default:
		return false
	}
}

func (e ResourceType) String() string {
	return string(e)
}

// FrameID Unique frame identifier.
type FrameID string

// RemoteObjectID Unique object identifier.
type RemoteObjectID string
