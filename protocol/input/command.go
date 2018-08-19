// Code generated by cdpgen. DO NOT EDIT.

package input

import (
	"encoding/json"
	"log"
)

const CommandInputDispatchKeyEvent = "Input.dispatchKeyEvent"

// DispatchKeyEventArgs represents the arguments for DispatchKeyEvent in the Input domain.
type DispatchKeyEventArgs struct {
	// Type Type of the key event.
	//
	// Values: "keyDown", "keyUp", "rawKeyDown", "char".
	Type                  string         `json:"type"`
	Modifiers             int            `json:"modifiers,omitempty"`             // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Timestamp             TimeSinceEpoch `json:"timestamp,omitempty"`             // Time at which the event occurred.
	Text                  string         `json:"text,omitempty"`                  // Text as generated by processing a virtual key code with a keyboard layout. Not needed for for `keyUp` and `rawKeyDown` events (default: "")
	UnmodifiedText        string         `json:"unmodifiedText,omitempty"`        // Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: "").
	KeyIdentifier         string         `json:"keyIdentifier,omitempty"`         // Unique key identifier (e.g., 'U+0041') (default: "").
	Code                  string         `json:"code,omitempty"`                  // Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: "").
	Key                   string         `json:"key,omitempty"`                   // Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr') (default: "").
	WindowsVirtualKeyCode int            `json:"windowsVirtualKeyCode,omitempty"` // Windows virtual key code (default: 0).
	NativeVirtualKeyCode  int            `json:"nativeVirtualKeyCode,omitempty"`  // Native virtual key code (default: 0).
	AutoRepeat            bool           `json:"autoRepeat,omitempty"`            // Whether the event was generated from auto repeat (default: false).
	IsKeypad              bool           `json:"isKeypad,omitempty"`              // Whether the event was generated from the keypad (default: false).
	IsSystemKey           bool           `json:"isSystemKey,omitempty"`           // Whether the event was a system key event (default: false).
	Location              int            `json:"location,omitempty"`              // Whether the event was from the left or right side of the keyboard. 1=Left, 2=Right (default: 0).
}

// Unmarshal the byte array into a return value for DispatchKeyEvent in the Input domain.
func (a *DispatchKeyEventArgs) UnmarshalJSON(b []byte) error {
	type Copy DispatchKeyEventArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DispatchKeyEventArgs(*c)
	return nil
}

// Marshall the byte array into a return value for DispatchKeyEvent in the Input domain.
func (a *DispatchKeyEventArgs) MarshalJSON() ([]byte, error) {
	type Copy DispatchKeyEventArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DispatchKeyEventReply represents the return values for DispatchKeyEvent in the Input domain.
type DispatchKeyEventReply struct {
}

// DispatchKeyEventReply returns whether or not the FrameID matches the reply value for DispatchKeyEvent in the Input domain.
func (a *DispatchKeyEventReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: DispatchKeyEventReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for DispatchKeyEvent in the Input domain.
func (a *DispatchKeyEventReply) UnmarshalJSON(b []byte) error {
	type Copy DispatchKeyEventReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DispatchKeyEventReply(*c)
	return nil
}

const CommandInputDispatchMouseEvent = "Input.dispatchMouseEvent"

// DispatchMouseEventArgs represents the arguments for DispatchMouseEvent in the Input domain.
type DispatchMouseEventArgs struct {
	// Type Type of the mouse event.
	//
	// Values: "mousePressed", "mouseReleased", "mouseMoved", "mouseWheel".
	Type      string         `json:"type"`
	X         float64        `json:"x"`                   // X coordinate of the event relative to the main frame's viewport in CSS pixels.
	Y         float64        `json:"y"`                   // Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Modifiers int            `json:"modifiers,omitempty"` // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Timestamp TimeSinceEpoch `json:"timestamp,omitempty"` // Time at which the event occurred.
	// Button Mouse button (default: "none").
	//
	// Values: "none", "left", "middle", "right".
	Button     string  `json:"button,omitempty"`
	ClickCount int     `json:"clickCount,omitempty"` // Number of times the mouse button was clicked (default: 0).
	DeltaX     float64 `json:"deltaX,omitempty"`     // X delta in CSS pixels for mouse wheel event (default: 0).
	DeltaY     float64 `json:"deltaY,omitempty"`     // Y delta in CSS pixels for mouse wheel event (default: 0).
}

// Unmarshal the byte array into a return value for DispatchMouseEvent in the Input domain.
func (a *DispatchMouseEventArgs) UnmarshalJSON(b []byte) error {
	type Copy DispatchMouseEventArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DispatchMouseEventArgs(*c)
	return nil
}

// Marshall the byte array into a return value for DispatchMouseEvent in the Input domain.
func (a *DispatchMouseEventArgs) MarshalJSON() ([]byte, error) {
	type Copy DispatchMouseEventArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DispatchMouseEventReply represents the return values for DispatchMouseEvent in the Input domain.
type DispatchMouseEventReply struct {
}

// DispatchMouseEventReply returns whether or not the FrameID matches the reply value for DispatchMouseEvent in the Input domain.
func (a *DispatchMouseEventReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: DispatchMouseEventReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for DispatchMouseEvent in the Input domain.
func (a *DispatchMouseEventReply) UnmarshalJSON(b []byte) error {
	type Copy DispatchMouseEventReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DispatchMouseEventReply(*c)
	return nil
}

const CommandInputDispatchTouchEvent = "Input.dispatchTouchEvent"

// DispatchTouchEventArgs represents the arguments for DispatchTouchEvent in the Input domain.
type DispatchTouchEventArgs struct {
	// Type Type of the touch event. TouchEnd and TouchCancel must not
	// contain any touch points, while TouchStart and TouchMove must
	// contains at least one.
	//
	// Values: "touchStart", "touchEnd", "touchMove", "touchCancel".
	Type        string         `json:"type"`
	TouchPoints []TouchPoint   `json:"touchPoints"`         // Active touch points on the touch device. One event per any changed point (compared to previous touch event in a sequence) is generated, emulating pressing/moving/releasing points one by one.
	Modifiers   int            `json:"modifiers,omitempty"` // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Timestamp   TimeSinceEpoch `json:"timestamp,omitempty"` // Time at which the event occurred.
}

// Unmarshal the byte array into a return value for DispatchTouchEvent in the Input domain.
func (a *DispatchTouchEventArgs) UnmarshalJSON(b []byte) error {
	type Copy DispatchTouchEventArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DispatchTouchEventArgs(*c)
	return nil
}

// Marshall the byte array into a return value for DispatchTouchEvent in the Input domain.
func (a *DispatchTouchEventArgs) MarshalJSON() ([]byte, error) {
	type Copy DispatchTouchEventArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DispatchTouchEventReply represents the return values for DispatchTouchEvent in the Input domain.
type DispatchTouchEventReply struct {
}

// DispatchTouchEventReply returns whether or not the FrameID matches the reply value for DispatchTouchEvent in the Input domain.
func (a *DispatchTouchEventReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: DispatchTouchEventReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for DispatchTouchEvent in the Input domain.
func (a *DispatchTouchEventReply) UnmarshalJSON(b []byte) error {
	type Copy DispatchTouchEventReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DispatchTouchEventReply(*c)
	return nil
}

const CommandInputEmulateTouchFromMouseEvent = "Input.emulateTouchFromMouseEvent"

// EmulateTouchFromMouseEventArgs represents the arguments for EmulateTouchFromMouseEvent in the Input domain.
type EmulateTouchFromMouseEventArgs struct {
	// Type Type of the mouse event.
	//
	// Values: "mousePressed", "mouseReleased", "mouseMoved", "mouseWheel".
	Type string `json:"type"`
	X    int    `json:"x"` // X coordinate of the mouse pointer in DIP.
	Y    int    `json:"y"` // Y coordinate of the mouse pointer in DIP.
	// Button Mouse button.
	//
	// Values: "none", "left", "middle", "right".
	Button     string         `json:"button"`
	Timestamp  TimeSinceEpoch `json:"timestamp,omitempty"`  // Time at which the event occurred (default: current time).
	DeltaX     float64        `json:"deltaX,omitempty"`     // X delta in DIP for mouse wheel event (default: 0).
	DeltaY     float64        `json:"deltaY,omitempty"`     // Y delta in DIP for mouse wheel event (default: 0).
	Modifiers  int            `json:"modifiers,omitempty"`  // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	ClickCount int            `json:"clickCount,omitempty"` // Number of times the mouse button was clicked (default: 0).
}

// Unmarshal the byte array into a return value for EmulateTouchFromMouseEvent in the Input domain.
func (a *EmulateTouchFromMouseEventArgs) UnmarshalJSON(b []byte) error {
	type Copy EmulateTouchFromMouseEventArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = EmulateTouchFromMouseEventArgs(*c)
	return nil
}

// Marshall the byte array into a return value for EmulateTouchFromMouseEvent in the Input domain.
func (a *EmulateTouchFromMouseEventArgs) MarshalJSON() ([]byte, error) {
	type Copy EmulateTouchFromMouseEventArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// EmulateTouchFromMouseEventReply represents the return values for EmulateTouchFromMouseEvent in the Input domain.
type EmulateTouchFromMouseEventReply struct {
}

// EmulateTouchFromMouseEventReply returns whether or not the FrameID matches the reply value for EmulateTouchFromMouseEvent in the Input domain.
func (a *EmulateTouchFromMouseEventReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: EmulateTouchFromMouseEventReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for EmulateTouchFromMouseEvent in the Input domain.
func (a *EmulateTouchFromMouseEventReply) UnmarshalJSON(b []byte) error {
	type Copy EmulateTouchFromMouseEventReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = EmulateTouchFromMouseEventReply(*c)
	return nil
}

const CommandInputSetIgnoreInputEvents = "Input.setIgnoreInputEvents"

// SetIgnoreInputEventsArgs represents the arguments for SetIgnoreInputEvents in the Input domain.
type SetIgnoreInputEventsArgs struct {
	Ignore bool `json:"ignore"` // Ignores input events processing when set to true.
}

// Unmarshal the byte array into a return value for SetIgnoreInputEvents in the Input domain.
func (a *SetIgnoreInputEventsArgs) UnmarshalJSON(b []byte) error {
	type Copy SetIgnoreInputEventsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetIgnoreInputEventsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetIgnoreInputEvents in the Input domain.
func (a *SetIgnoreInputEventsArgs) MarshalJSON() ([]byte, error) {
	type Copy SetIgnoreInputEventsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetIgnoreInputEventsReply represents the return values for SetIgnoreInputEvents in the Input domain.
type SetIgnoreInputEventsReply struct {
}

// SetIgnoreInputEventsReply returns whether or not the FrameID matches the reply value for SetIgnoreInputEvents in the Input domain.
func (a *SetIgnoreInputEventsReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SetIgnoreInputEventsReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for SetIgnoreInputEvents in the Input domain.
func (a *SetIgnoreInputEventsReply) UnmarshalJSON(b []byte) error {
	type Copy SetIgnoreInputEventsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetIgnoreInputEventsReply(*c)
	return nil
}

const CommandInputSynthesizePinchGesture = "Input.synthesizePinchGesture"

// SynthesizePinchGestureArgs represents the arguments for SynthesizePinchGesture in the Input domain.
type SynthesizePinchGestureArgs struct {
	X                 float64           `json:"x"`                           // X coordinate of the start of the gesture in CSS pixels.
	Y                 float64           `json:"y"`                           // Y coordinate of the start of the gesture in CSS pixels.
	ScaleFactor       float64           `json:"scaleFactor"`                 // Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
	RelativeSpeed     int               `json:"relativeSpeed,omitempty"`     // Relative pointer speed in pixels per second (default: 800).
	GestureSourceType GestureSourceType `json:"gestureSourceType,omitempty"` // Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
}

// Unmarshal the byte array into a return value for SynthesizePinchGesture in the Input domain.
func (a *SynthesizePinchGestureArgs) UnmarshalJSON(b []byte) error {
	type Copy SynthesizePinchGestureArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SynthesizePinchGestureArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SynthesizePinchGesture in the Input domain.
func (a *SynthesizePinchGestureArgs) MarshalJSON() ([]byte, error) {
	type Copy SynthesizePinchGestureArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SynthesizePinchGestureReply represents the return values for SynthesizePinchGesture in the Input domain.
type SynthesizePinchGestureReply struct {
}

// SynthesizePinchGestureReply returns whether or not the FrameID matches the reply value for SynthesizePinchGesture in the Input domain.
func (a *SynthesizePinchGestureReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SynthesizePinchGestureReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for SynthesizePinchGesture in the Input domain.
func (a *SynthesizePinchGestureReply) UnmarshalJSON(b []byte) error {
	type Copy SynthesizePinchGestureReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SynthesizePinchGestureReply(*c)
	return nil
}

const CommandInputSynthesizeScrollGesture = "Input.synthesizeScrollGesture"

// SynthesizeScrollGestureArgs represents the arguments for SynthesizeScrollGesture in the Input domain.
type SynthesizeScrollGestureArgs struct {
	X                     float64           `json:"x"`                               // X coordinate of the start of the gesture in CSS pixels.
	Y                     float64           `json:"y"`                               // Y coordinate of the start of the gesture in CSS pixels.
	XDistance             float64           `json:"xDistance,omitempty"`             // The distance to scroll along the X axis (positive to scroll left).
	YDistance             float64           `json:"yDistance,omitempty"`             // The distance to scroll along the Y axis (positive to scroll up).
	XOverscroll           float64           `json:"xOverscroll,omitempty"`           // The number of additional pixels to scroll back along the X axis, in addition to the given distance.
	YOverscroll           float64           `json:"yOverscroll,omitempty"`           // The number of additional pixels to scroll back along the Y axis, in addition to the given distance.
	PreventFling          bool              `json:"preventFling,omitempty"`          // Prevent fling (default: true).
	Speed                 int               `json:"speed,omitempty"`                 // Swipe speed in pixels per second (default: 800).
	GestureSourceType     GestureSourceType `json:"gestureSourceType,omitempty"`     // Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
	RepeatCount           int               `json:"repeatCount,omitempty"`           // The number of times to repeat the gesture (default: 0).
	RepeatDelayMs         int               `json:"repeatDelayMs,omitempty"`         // The number of milliseconds delay between each repeat. (default: 250).
	InteractionMarkerName string            `json:"interactionMarkerName,omitempty"` // The name of the interaction markers to generate, if not empty (default: "").
}

// Unmarshal the byte array into a return value for SynthesizeScrollGesture in the Input domain.
func (a *SynthesizeScrollGestureArgs) UnmarshalJSON(b []byte) error {
	type Copy SynthesizeScrollGestureArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SynthesizeScrollGestureArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SynthesizeScrollGesture in the Input domain.
func (a *SynthesizeScrollGestureArgs) MarshalJSON() ([]byte, error) {
	type Copy SynthesizeScrollGestureArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SynthesizeScrollGestureReply represents the return values for SynthesizeScrollGesture in the Input domain.
type SynthesizeScrollGestureReply struct {
}

// SynthesizeScrollGestureReply returns whether or not the FrameID matches the reply value for SynthesizeScrollGesture in the Input domain.
func (a *SynthesizeScrollGestureReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SynthesizeScrollGestureReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for SynthesizeScrollGesture in the Input domain.
func (a *SynthesizeScrollGestureReply) UnmarshalJSON(b []byte) error {
	type Copy SynthesizeScrollGestureReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SynthesizeScrollGestureReply(*c)
	return nil
}

const CommandInputSynthesizeTapGesture = "Input.synthesizeTapGesture"

// SynthesizeTapGestureArgs represents the arguments for SynthesizeTapGesture in the Input domain.
type SynthesizeTapGestureArgs struct {
	X                 float64           `json:"x"`                           // X coordinate of the start of the gesture in CSS pixels.
	Y                 float64           `json:"y"`                           // Y coordinate of the start of the gesture in CSS pixels.
	Duration          int               `json:"duration,omitempty"`          // Duration between touchdown and touchup events in ms (default: 50).
	TapCount          int               `json:"tapCount,omitempty"`          // Number of times to perform the tap (e.g. 2 for double tap, default: 1).
	GestureSourceType GestureSourceType `json:"gestureSourceType,omitempty"` // Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
}

// Unmarshal the byte array into a return value for SynthesizeTapGesture in the Input domain.
func (a *SynthesizeTapGestureArgs) UnmarshalJSON(b []byte) error {
	type Copy SynthesizeTapGestureArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SynthesizeTapGestureArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SynthesizeTapGesture in the Input domain.
func (a *SynthesizeTapGestureArgs) MarshalJSON() ([]byte, error) {
	type Copy SynthesizeTapGestureArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SynthesizeTapGestureReply represents the return values for SynthesizeTapGesture in the Input domain.
type SynthesizeTapGestureReply struct {
}

// SynthesizeTapGestureReply returns whether or not the FrameID matches the reply value for SynthesizeTapGesture in the Input domain.
func (a *SynthesizeTapGestureReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SynthesizeTapGestureReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for SynthesizeTapGesture in the Input domain.
func (a *SynthesizeTapGestureReply) UnmarshalJSON(b []byte) error {
	type Copy SynthesizeTapGestureReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SynthesizeTapGestureReply(*c)
	return nil
}
