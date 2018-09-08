// Code generated by cdpgen. DO NOT EDIT.

package animation

import (
	"encoding/json"
	"log"
)

const (
	EventAnimationAnimationCanceled = "Animation.animationCanceled"
	EventAnimationAnimationCreated  = "Animation.animationCreated"
	EventAnimationAnimationStarted  = "Animation.animationStarted"
)

type Unmarshaler func() json.Unmarshaler

var EventConstants = map[string]Unmarshaler{
	EventAnimationAnimationCanceled: func() json.Unmarshaler { return &CanceledReply{} },
	EventAnimationAnimationCreated:  func() json.Unmarshaler { return &CreatedReply{} },
	EventAnimationAnimationStarted:  func() json.Unmarshaler { return &StartedReply{} },
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	if ok {
		return e(), ok
	}
	return nil, ok
}

// CanceledReply is the reply for AnimationCanceled events.
type CanceledReply struct {
	ID string `json:"id"` // Id of the animation that was canceled.
}

// Unmarshal the byte array into a return value for AnimationCanceled in the Animation domain.
func (a *CanceledReply) UnmarshalJSON(b []byte) error {
	type Copy CanceledReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CanceledReply(*c)
	return nil
}

// CanceledReply returns whether or not the FrameID matches the reply value for AnimationCanceled in the Animation domain.
func (a *CanceledReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: CanceledReply %s", err)
		return false, err
	}
	return true, nil
}

// CanceledReply returns the FrameID for AnimationCanceled in the Animation domain.
func (a *CanceledReply) GetFrameID() string {
	return ""
}

// CreatedReply is the reply for AnimationCreated events.
type CreatedReply struct {
	ID string `json:"id"` // Id of the animation that was created.
}

// Unmarshal the byte array into a return value for AnimationCreated in the Animation domain.
func (a *CreatedReply) UnmarshalJSON(b []byte) error {
	type Copy CreatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CreatedReply(*c)
	return nil
}

// CreatedReply returns whether or not the FrameID matches the reply value for AnimationCreated in the Animation domain.
func (a *CreatedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: CreatedReply %s", err)
		return false, err
	}
	return true, nil
}

// CreatedReply returns the FrameID for AnimationCreated in the Animation domain.
func (a *CreatedReply) GetFrameID() string {
	return ""
}

// StartedReply is the reply for AnimationStarted events.
type StartedReply struct {
	Animation Animation `json:"animation"` // Animation that was started.
}

// Unmarshal the byte array into a return value for AnimationStarted in the Animation domain.
func (a *StartedReply) UnmarshalJSON(b []byte) error {
	type Copy StartedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StartedReply(*c)
	return nil
}

// StartedReply returns whether or not the FrameID matches the reply value for AnimationStarted in the Animation domain.
func (a *StartedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: StartedReply %s", err)
		return false, err
	}
	return true, nil
}

// StartedReply returns the FrameID for AnimationStarted in the Animation domain.
func (a *StartedReply) GetFrameID() string {
	return ""
}
