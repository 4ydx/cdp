// Code generated by cdpgen. DO NOT EDIT.

package deviceorientation

import (
	"encoding/json"
	"log"
)

const (
	CommandDeviceOrientationClearDeviceOrientationOverride = "DeviceOrientation.clearDeviceOrientationOverride"
	CommandDeviceOrientationSetDeviceOrientationOverride   = "DeviceOrientation.setDeviceOrientationOverride"
)

// ClearDeviceOrientationOverrideArgs represents the arguments for ClearDeviceOrientationOverride in the DeviceOrientation domain.
type ClearDeviceOrientationOverrideArgs struct {
}

// Unmarshal the byte array into a return value for ClearDeviceOrientationOverride in the DeviceOrientation domain.
func (a *ClearDeviceOrientationOverrideArgs) UnmarshalJSON(b []byte) error {
	type Copy ClearDeviceOrientationOverrideArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ClearDeviceOrientationOverrideArgs(*c)
	return nil
}

// Marshall the byte array into a return value for ClearDeviceOrientationOverride in the DeviceOrientation domain.
func (a *ClearDeviceOrientationOverrideArgs) MarshalJSON() ([]byte, error) {
	type Copy ClearDeviceOrientationOverrideArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// ClearDeviceOrientationOverrideReply represents the return values for ClearDeviceOrientationOverride in the DeviceOrientation domain.
type ClearDeviceOrientationOverrideReply struct {
}

// ClearDeviceOrientationOverrideReply returns whether or not the FrameID matches the reply value for ClearDeviceOrientationOverride in the DeviceOrientation domain.
func (a *ClearDeviceOrientationOverrideReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: ClearDeviceOrientationOverrideReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for ClearDeviceOrientationOverride in the DeviceOrientation domain.
func (a *ClearDeviceOrientationOverrideReply) UnmarshalJSON(b []byte) error {
	type Copy ClearDeviceOrientationOverrideReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ClearDeviceOrientationOverrideReply(*c)
	return nil
}

// SetDeviceOrientationOverrideArgs represents the arguments for SetDeviceOrientationOverride in the DeviceOrientation domain.
type SetDeviceOrientationOverrideArgs struct {
	Alpha float64 `json:"alpha"` // Mock alpha
	Beta  float64 `json:"beta"`  // Mock beta
	Gamma float64 `json:"gamma"` // Mock gamma
}

// Unmarshal the byte array into a return value for SetDeviceOrientationOverride in the DeviceOrientation domain.
func (a *SetDeviceOrientationOverrideArgs) UnmarshalJSON(b []byte) error {
	type Copy SetDeviceOrientationOverrideArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetDeviceOrientationOverrideArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetDeviceOrientationOverride in the DeviceOrientation domain.
func (a *SetDeviceOrientationOverrideArgs) MarshalJSON() ([]byte, error) {
	type Copy SetDeviceOrientationOverrideArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetDeviceOrientationOverrideReply represents the return values for SetDeviceOrientationOverride in the DeviceOrientation domain.
type SetDeviceOrientationOverrideReply struct {
}

// SetDeviceOrientationOverrideReply returns whether or not the FrameID matches the reply value for SetDeviceOrientationOverride in the DeviceOrientation domain.
func (a *SetDeviceOrientationOverrideReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SetDeviceOrientationOverrideReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for SetDeviceOrientationOverride in the DeviceOrientation domain.
func (a *SetDeviceOrientationOverrideReply) UnmarshalJSON(b []byte) error {
	type Copy SetDeviceOrientationOverrideReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetDeviceOrientationOverrideReply(*c)
	return nil
}
