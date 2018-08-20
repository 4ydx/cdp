// Code generated by cdpgen. DO NOT EDIT.

package headlessexperimental

import (
	"encoding/json"
	"log"
)

const (
	CommandHeadlessExperimentalBeginFrame = "HeadlessExperimental.beginFrame"
	CommandHeadlessExperimentalDisable    = "HeadlessExperimental.disable"
	CommandHeadlessExperimentalEnable     = "HeadlessExperimental.enable"
)

// BeginFrameArgs represents the arguments for BeginFrame in the HeadlessExperimental domain.
type BeginFrameArgs struct {
	FrameTimeTicks   float64          `json:"frameTimeTicks,omitempty"`   // Timestamp of this BeginFrame in Renderer TimeTicks (milliseconds of uptime). If not set, the current time will be used.
	Interval         float64          `json:"interval,omitempty"`         // The interval between BeginFrames that is reported to the compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	NoDisplayUpdates bool             `json:"noDisplayUpdates,omitempty"` // Whether updates should not be committed and drawn onto the display. False by default. If true, only side effects of the BeginFrame will be run, such as layout and animations, but any visual updates may not be visible on the display or in screenshots.
	Screenshot       ScreenshotParams `json:"screenshot,omitempty"`       // If set, a screenshot of the frame will be captured and returned in the response. Otherwise, no screenshot will be captured. Note that capturing a screenshot can fail, for example, during renderer initialization. In such a case, no screenshot data will be returned.
}

// Unmarshal the byte array into a return value for BeginFrame in the HeadlessExperimental domain.
func (a *BeginFrameArgs) UnmarshalJSON(b []byte) error {
	type Copy BeginFrameArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = BeginFrameArgs(*c)
	return nil
}

// Marshall the byte array into a return value for BeginFrame in the HeadlessExperimental domain.
func (a *BeginFrameArgs) MarshalJSON() ([]byte, error) {
	type Copy BeginFrameArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// BeginFrameReply represents the return values for BeginFrame in the HeadlessExperimental domain.
type BeginFrameReply struct {
	HasDamage      bool   `json:"hasDamage"`                // Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display. Reported for diagnostic uses, may be removed in the future.
	ScreenshotData []byte `json:"screenshotData,omitempty"` // Base64-encoded image data of the screenshot, if one was requested and successfully taken.
}

// BeginFrameReply returns whether or not the FrameID matches the reply value for BeginFrame in the HeadlessExperimental domain.
func (a *BeginFrameReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: BeginFrameReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for BeginFrame in the HeadlessExperimental domain.
func (a *BeginFrameReply) UnmarshalJSON(b []byte) error {
	type Copy BeginFrameReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = BeginFrameReply(*c)
	return nil
}

// DisableArgs represents the arguments for Disable in the HeadlessExperimental domain.
type DisableArgs struct {
}

// Unmarshal the byte array into a return value for Disable in the HeadlessExperimental domain.
func (a *DisableArgs) UnmarshalJSON(b []byte) error {
	type Copy DisableArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DisableArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Disable in the HeadlessExperimental domain.
func (a *DisableArgs) MarshalJSON() ([]byte, error) {
	type Copy DisableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DisableReply represents the return values for Disable in the HeadlessExperimental domain.
type DisableReply struct {
}

// DisableReply returns whether or not the FrameID matches the reply value for Disable in the HeadlessExperimental domain.
func (a *DisableReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: DisableReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for Disable in the HeadlessExperimental domain.
func (a *DisableReply) UnmarshalJSON(b []byte) error {
	type Copy DisableReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DisableReply(*c)
	return nil
}

// EnableArgs represents the arguments for Enable in the HeadlessExperimental domain.
type EnableArgs struct {
}

// Unmarshal the byte array into a return value for Enable in the HeadlessExperimental domain.
func (a *EnableArgs) UnmarshalJSON(b []byte) error {
	type Copy EnableArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = EnableArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Enable in the HeadlessExperimental domain.
func (a *EnableArgs) MarshalJSON() ([]byte, error) {
	type Copy EnableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// EnableReply represents the return values for Enable in the HeadlessExperimental domain.
type EnableReply struct {
}

// EnableReply returns whether or not the FrameID matches the reply value for Enable in the HeadlessExperimental domain.
func (a *EnableReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: EnableReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for Enable in the HeadlessExperimental domain.
func (a *EnableReply) UnmarshalJSON(b []byte) error {
	type Copy EnableReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = EnableReply(*c)
	return nil
}
