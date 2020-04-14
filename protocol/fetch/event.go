// Code generated by cdpgen. DO NOT EDIT.

package fetch

import (
	"encoding/json"
	"log"

	shared "github.com/4ydx/cdp/protocol"
	"github.com/4ydx/cdp/protocol/network"
)

const (
	EventFetchRequestPaused = "Fetch.requestPaused"
	EventFetchAuthRequired  = "Fetch.authRequired"
)

type Unmarshaler func() json.Unmarshaler

var EventConstants = map[string]Unmarshaler{
	EventFetchRequestPaused: func() json.Unmarshaler { return &RequestPausedReply{} },
	EventFetchAuthRequired:  func() json.Unmarshaler { return &AuthRequiredReply{} },
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	if ok {
		return e(), ok
	}
	return nil, ok
}

// RequestPausedReply is the reply for RequestPaused events.
type RequestPausedReply struct {
	RequestID           RequestID            `json:"requestId"`                     // Each request the page makes will have a unique id.
	Request             network.Request      `json:"request"`                       // The details of the request.
	FrameID             shared.FrameID       `json:"frameId"`                       // The id of the frame that initiated the request.
	ResourceType        shared.ResourceType  `json:"resourceType"`                  // How the requested resource will be used.
	ResponseErrorReason *network.ErrorReason `json:"responseErrorReason,omitempty"` // Response error if intercepted at response stage.
	ResponseStatusCode  int                  `json:"responseStatusCode,omitempty"`  // Response code if intercepted at response stage.
	ResponseHeaders     *[]HeaderEntry       `json:"responseHeaders,omitempty"`     // Response headers if intercepted at the response stage.
	NetworkID           RequestID            `json:"networkId,omitempty"`           // If the intercepted request had a corresponding Network.requestWillBeSent event fired for it, then this networkId will be the same as the requestId present in the requestWillBeSent event.
}

// Unmarshal the byte array into a return value for RequestPaused in the Fetch domain.
func (a *RequestPausedReply) UnmarshalJSON(b []byte) error {
	type Copy RequestPausedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestPausedReply(*c)
	return nil
}

// RequestPausedReply returns whether or not the FrameID matches the reply value for RequestPaused in the Fetch domain.
func (a *RequestPausedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	v := &RequestPausedReply{}
	err := v.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: RequestPausedReply %s", err)
		return false, err
	}
	if v.FrameID != shared.FrameID(frameID) {
		return false, nil
	}
	*a = *v
	return true, nil
}

// RequestPausedReply returns the FrameID for RequestPaused in the Fetch domain.
func (a *RequestPausedReply) GetFrameID() string {
	return string(a.FrameID)
}

// AuthRequiredReply is the reply for AuthRequired events.
type AuthRequiredReply struct {
	RequestID     RequestID           `json:"requestId"`     // Each request the page makes will have a unique id.
	Request       network.Request     `json:"request"`       // The details of the request.
	FrameID       shared.FrameID      `json:"frameId"`       // The id of the frame that initiated the request.
	ResourceType  shared.ResourceType `json:"resourceType"`  // How the requested resource will be used.
	AuthChallenge AuthChallenge       `json:"authChallenge"` // Details of the Authorization Challenge encountered. If this is set, client should respond with continueRequest that contains AuthChallengeResponse.
}

// Unmarshal the byte array into a return value for AuthRequired in the Fetch domain.
func (a *AuthRequiredReply) UnmarshalJSON(b []byte) error {
	type Copy AuthRequiredReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = AuthRequiredReply(*c)
	return nil
}

// AuthRequiredReply returns whether or not the FrameID matches the reply value for AuthRequired in the Fetch domain.
func (a *AuthRequiredReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	v := &AuthRequiredReply{}
	err := v.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: AuthRequiredReply %s", err)
		return false, err
	}
	if v.FrameID != shared.FrameID(frameID) {
		return false, nil
	}
	*a = *v
	return true, nil
}

// AuthRequiredReply returns the FrameID for AuthRequired in the Fetch domain.
func (a *AuthRequiredReply) GetFrameID() string {
	return string(a.FrameID)
}
