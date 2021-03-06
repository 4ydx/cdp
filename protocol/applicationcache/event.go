// Code generated by cdpgen. DO NOT EDIT.

package applicationcache

import (
	"encoding/json"
	"log"

	shared "github.com/4ydx/cdp/protocol"
)

const (
	EventApplicationCacheApplicationCacheStatusUpdated = "ApplicationCache.applicationCacheStatusUpdated"
	EventApplicationCacheNetworkStateUpdated           = "ApplicationCache.networkStateUpdated"
)

type Unmarshaler func() json.Unmarshaler

var EventConstants = map[string]Unmarshaler{
	EventApplicationCacheApplicationCacheStatusUpdated: func() json.Unmarshaler { return &StatusUpdatedReply{} },
	EventApplicationCacheNetworkStateUpdated:           func() json.Unmarshaler { return &NetworkStateUpdatedReply{} },
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	if ok {
		return e(), ok
	}
	return nil, ok
}

// StatusUpdatedReply is the reply for ApplicationCacheStatusUpdated events.
type StatusUpdatedReply struct {
	FrameID     shared.FrameID `json:"frameId"`     // Identifier of the frame containing document whose application cache updated status.
	ManifestURL string         `json:"manifestURL"` // Manifest URL.
	Status      int            `json:"status"`      // Updated application cache status.
}

// Unmarshal the byte array into a return value for ApplicationCacheStatusUpdated in the ApplicationCache domain.
func (a *StatusUpdatedReply) UnmarshalJSON(b []byte) error {
	type Copy StatusUpdatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StatusUpdatedReply(*c)
	return nil
}

// StatusUpdatedReply returns whether or not the FrameID matches the reply value for ApplicationCacheStatusUpdated in the ApplicationCache domain.
func (a *StatusUpdatedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	v := &StatusUpdatedReply{}
	err := v.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: StatusUpdatedReply %s", err)
		return false, err
	}
	if v.FrameID != shared.FrameID(frameID) {
		return false, nil
	}
	*a = *v
	return true, nil
}

// StatusUpdatedReply returns the FrameID for ApplicationCacheStatusUpdated in the ApplicationCache domain.
func (a *StatusUpdatedReply) GetFrameID() string {
	return string(a.FrameID)
}

// NetworkStateUpdatedReply is the reply for NetworkStateUpdated events.
type NetworkStateUpdatedReply struct {
	IsNowOnline bool `json:"isNowOnline"` // No description.
}

// Unmarshal the byte array into a return value for NetworkStateUpdated in the ApplicationCache domain.
func (a *NetworkStateUpdatedReply) UnmarshalJSON(b []byte) error {
	type Copy NetworkStateUpdatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = NetworkStateUpdatedReply(*c)
	return nil
}

// NetworkStateUpdatedReply returns whether or not the FrameID matches the reply value for NetworkStateUpdated in the ApplicationCache domain.
func (a *NetworkStateUpdatedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: NetworkStateUpdatedReply %s", err)
		return false, err
	}
	return true, nil
}

// NetworkStateUpdatedReply returns the FrameID for NetworkStateUpdated in the ApplicationCache domain.
func (a *NetworkStateUpdatedReply) GetFrameID() string {
	return ""
}
