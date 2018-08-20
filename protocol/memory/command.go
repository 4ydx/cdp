// Code generated by cdpgen. DO NOT EDIT.

package memory

import (
	"encoding/json"
	"log"
)

const (
	CommandMemoryGetDOMCounters                     = "Memory.getDOMCounters"
	CommandMemoryPrepareForLeakDetection            = "Memory.prepareForLeakDetection"
	CommandMemorySetPressureNotificationsSuppressed = "Memory.setPressureNotificationsSuppressed"
	CommandMemorySimulatePressureNotification       = "Memory.simulatePressureNotification"
	CommandMemoryStartSampling                      = "Memory.startSampling"
	CommandMemoryStopSampling                       = "Memory.stopSampling"
	CommandMemoryGetAllTimeSamplingProfile          = "Memory.getAllTimeSamplingProfile"
	CommandMemoryGetBrowserSamplingProfile          = "Memory.getBrowserSamplingProfile"
	CommandMemoryGetSamplingProfile                 = "Memory.getSamplingProfile"
)

// GetDOMCountersArgs represents the arguments for GetDOMCounters in the Memory domain.
type GetDOMCountersArgs struct {
}

// Unmarshal the byte array into a return value for GetDOMCounters in the Memory domain.
func (a *GetDOMCountersArgs) UnmarshalJSON(b []byte) error {
	type Copy GetDOMCountersArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetDOMCountersArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetDOMCounters in the Memory domain.
func (a *GetDOMCountersArgs) MarshalJSON() ([]byte, error) {
	type Copy GetDOMCountersArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetDOMCountersReply represents the return values for GetDOMCounters in the Memory domain.
type GetDOMCountersReply struct {
	Documents        int `json:"documents"`        // No description.
	Nodes            int `json:"nodes"`            // No description.
	JsEventListeners int `json:"jsEventListeners"` // No description.
}

// GetDOMCountersReply returns whether or not the FrameID matches the reply value for GetDOMCounters in the Memory domain.
func (a *GetDOMCountersReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetDOMCountersReply", err)
	}
	return true
}

// GetDOMCountersReply returns the FrameID value for GetDOMCounters in the Memory domain.
func (a *GetDOMCountersReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetDOMCounters in the Memory domain.
func (a *GetDOMCountersReply) UnmarshalJSON(b []byte) error {
	type Copy GetDOMCountersReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetDOMCountersReply(*c)
	return nil
}

// PrepareForLeakDetectionArgs represents the arguments for PrepareForLeakDetection in the Memory domain.
type PrepareForLeakDetectionArgs struct {
}

// Unmarshal the byte array into a return value for PrepareForLeakDetection in the Memory domain.
func (a *PrepareForLeakDetectionArgs) UnmarshalJSON(b []byte) error {
	type Copy PrepareForLeakDetectionArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = PrepareForLeakDetectionArgs(*c)
	return nil
}

// Marshall the byte array into a return value for PrepareForLeakDetection in the Memory domain.
func (a *PrepareForLeakDetectionArgs) MarshalJSON() ([]byte, error) {
	type Copy PrepareForLeakDetectionArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// PrepareForLeakDetectionReply represents the return values for PrepareForLeakDetection in the Memory domain.
type PrepareForLeakDetectionReply struct {
}

// PrepareForLeakDetectionReply returns whether or not the FrameID matches the reply value for PrepareForLeakDetection in the Memory domain.
func (a *PrepareForLeakDetectionReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: PrepareForLeakDetectionReply", err)
	}
	return true
}

// PrepareForLeakDetectionReply returns the FrameID value for PrepareForLeakDetection in the Memory domain.
func (a *PrepareForLeakDetectionReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for PrepareForLeakDetection in the Memory domain.
func (a *PrepareForLeakDetectionReply) UnmarshalJSON(b []byte) error {
	type Copy PrepareForLeakDetectionReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = PrepareForLeakDetectionReply(*c)
	return nil
}

// SetPressureNotificationsSuppressedArgs represents the arguments for SetPressureNotificationsSuppressed in the Memory domain.
type SetPressureNotificationsSuppressedArgs struct {
	Suppressed bool `json:"suppressed"` // If true, memory pressure notifications will be suppressed.
}

// Unmarshal the byte array into a return value for SetPressureNotificationsSuppressed in the Memory domain.
func (a *SetPressureNotificationsSuppressedArgs) UnmarshalJSON(b []byte) error {
	type Copy SetPressureNotificationsSuppressedArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetPressureNotificationsSuppressedArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetPressureNotificationsSuppressed in the Memory domain.
func (a *SetPressureNotificationsSuppressedArgs) MarshalJSON() ([]byte, error) {
	type Copy SetPressureNotificationsSuppressedArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetPressureNotificationsSuppressedReply represents the return values for SetPressureNotificationsSuppressed in the Memory domain.
type SetPressureNotificationsSuppressedReply struct {
}

// SetPressureNotificationsSuppressedReply returns whether or not the FrameID matches the reply value for SetPressureNotificationsSuppressed in the Memory domain.
func (a *SetPressureNotificationsSuppressedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SetPressureNotificationsSuppressedReply", err)
	}
	return true
}

// SetPressureNotificationsSuppressedReply returns the FrameID value for SetPressureNotificationsSuppressed in the Memory domain.
func (a *SetPressureNotificationsSuppressedReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for SetPressureNotificationsSuppressed in the Memory domain.
func (a *SetPressureNotificationsSuppressedReply) UnmarshalJSON(b []byte) error {
	type Copy SetPressureNotificationsSuppressedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetPressureNotificationsSuppressedReply(*c)
	return nil
}

// SimulatePressureNotificationArgs represents the arguments for SimulatePressureNotification in the Memory domain.
type SimulatePressureNotificationArgs struct {
	Level PressureLevel `json:"level"` // Memory pressure level of the notification.
}

// Unmarshal the byte array into a return value for SimulatePressureNotification in the Memory domain.
func (a *SimulatePressureNotificationArgs) UnmarshalJSON(b []byte) error {
	type Copy SimulatePressureNotificationArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SimulatePressureNotificationArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SimulatePressureNotification in the Memory domain.
func (a *SimulatePressureNotificationArgs) MarshalJSON() ([]byte, error) {
	type Copy SimulatePressureNotificationArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SimulatePressureNotificationReply represents the return values for SimulatePressureNotification in the Memory domain.
type SimulatePressureNotificationReply struct {
}

// SimulatePressureNotificationReply returns whether or not the FrameID matches the reply value for SimulatePressureNotification in the Memory domain.
func (a *SimulatePressureNotificationReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SimulatePressureNotificationReply", err)
	}
	return true
}

// SimulatePressureNotificationReply returns the FrameID value for SimulatePressureNotification in the Memory domain.
func (a *SimulatePressureNotificationReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for SimulatePressureNotification in the Memory domain.
func (a *SimulatePressureNotificationReply) UnmarshalJSON(b []byte) error {
	type Copy SimulatePressureNotificationReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SimulatePressureNotificationReply(*c)
	return nil
}

// StartSamplingArgs represents the arguments for StartSampling in the Memory domain.
type StartSamplingArgs struct {
	SamplingInterval   int  `json:"samplingInterval,omitempty"`   // Average number of bytes between samples.
	SuppressRandomness bool `json:"suppressRandomness,omitempty"` // Do not randomize intervals between samples.
}

// Unmarshal the byte array into a return value for StartSampling in the Memory domain.
func (a *StartSamplingArgs) UnmarshalJSON(b []byte) error {
	type Copy StartSamplingArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StartSamplingArgs(*c)
	return nil
}

// Marshall the byte array into a return value for StartSampling in the Memory domain.
func (a *StartSamplingArgs) MarshalJSON() ([]byte, error) {
	type Copy StartSamplingArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// StartSamplingReply represents the return values for StartSampling in the Memory domain.
type StartSamplingReply struct {
}

// StartSamplingReply returns whether or not the FrameID matches the reply value for StartSampling in the Memory domain.
func (a *StartSamplingReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: StartSamplingReply", err)
	}
	return true
}

// StartSamplingReply returns the FrameID value for StartSampling in the Memory domain.
func (a *StartSamplingReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for StartSampling in the Memory domain.
func (a *StartSamplingReply) UnmarshalJSON(b []byte) error {
	type Copy StartSamplingReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StartSamplingReply(*c)
	return nil
}

// StopSamplingArgs represents the arguments for StopSampling in the Memory domain.
type StopSamplingArgs struct {
}

// Unmarshal the byte array into a return value for StopSampling in the Memory domain.
func (a *StopSamplingArgs) UnmarshalJSON(b []byte) error {
	type Copy StopSamplingArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StopSamplingArgs(*c)
	return nil
}

// Marshall the byte array into a return value for StopSampling in the Memory domain.
func (a *StopSamplingArgs) MarshalJSON() ([]byte, error) {
	type Copy StopSamplingArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// StopSamplingReply represents the return values for StopSampling in the Memory domain.
type StopSamplingReply struct {
}

// StopSamplingReply returns whether or not the FrameID matches the reply value for StopSampling in the Memory domain.
func (a *StopSamplingReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: StopSamplingReply", err)
	}
	return true
}

// StopSamplingReply returns the FrameID value for StopSampling in the Memory domain.
func (a *StopSamplingReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for StopSampling in the Memory domain.
func (a *StopSamplingReply) UnmarshalJSON(b []byte) error {
	type Copy StopSamplingReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StopSamplingReply(*c)
	return nil
}

// GetAllTimeSamplingProfileArgs represents the arguments for GetAllTimeSamplingProfile in the Memory domain.
type GetAllTimeSamplingProfileArgs struct {
}

// Unmarshal the byte array into a return value for GetAllTimeSamplingProfile in the Memory domain.
func (a *GetAllTimeSamplingProfileArgs) UnmarshalJSON(b []byte) error {
	type Copy GetAllTimeSamplingProfileArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetAllTimeSamplingProfileArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetAllTimeSamplingProfile in the Memory domain.
func (a *GetAllTimeSamplingProfileArgs) MarshalJSON() ([]byte, error) {
	type Copy GetAllTimeSamplingProfileArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetAllTimeSamplingProfileReply represents the return values for GetAllTimeSamplingProfile in the Memory domain.
type GetAllTimeSamplingProfileReply struct {
	Profile SamplingProfile `json:"profile"` // No description.
}

// GetAllTimeSamplingProfileReply returns whether or not the FrameID matches the reply value for GetAllTimeSamplingProfile in the Memory domain.
func (a *GetAllTimeSamplingProfileReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetAllTimeSamplingProfileReply", err)
	}
	return true
}

// GetAllTimeSamplingProfileReply returns the FrameID value for GetAllTimeSamplingProfile in the Memory domain.
func (a *GetAllTimeSamplingProfileReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetAllTimeSamplingProfile in the Memory domain.
func (a *GetAllTimeSamplingProfileReply) UnmarshalJSON(b []byte) error {
	type Copy GetAllTimeSamplingProfileReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetAllTimeSamplingProfileReply(*c)
	return nil
}

// GetBrowserSamplingProfileArgs represents the arguments for GetBrowserSamplingProfile in the Memory domain.
type GetBrowserSamplingProfileArgs struct {
}

// Unmarshal the byte array into a return value for GetBrowserSamplingProfile in the Memory domain.
func (a *GetBrowserSamplingProfileArgs) UnmarshalJSON(b []byte) error {
	type Copy GetBrowserSamplingProfileArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetBrowserSamplingProfileArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetBrowserSamplingProfile in the Memory domain.
func (a *GetBrowserSamplingProfileArgs) MarshalJSON() ([]byte, error) {
	type Copy GetBrowserSamplingProfileArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetBrowserSamplingProfileReply represents the return values for GetBrowserSamplingProfile in the Memory domain.
type GetBrowserSamplingProfileReply struct {
	Profile SamplingProfile `json:"profile"` // No description.
}

// GetBrowserSamplingProfileReply returns whether or not the FrameID matches the reply value for GetBrowserSamplingProfile in the Memory domain.
func (a *GetBrowserSamplingProfileReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetBrowserSamplingProfileReply", err)
	}
	return true
}

// GetBrowserSamplingProfileReply returns the FrameID value for GetBrowserSamplingProfile in the Memory domain.
func (a *GetBrowserSamplingProfileReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetBrowserSamplingProfile in the Memory domain.
func (a *GetBrowserSamplingProfileReply) UnmarshalJSON(b []byte) error {
	type Copy GetBrowserSamplingProfileReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetBrowserSamplingProfileReply(*c)
	return nil
}

// GetSamplingProfileArgs represents the arguments for GetSamplingProfile in the Memory domain.
type GetSamplingProfileArgs struct {
}

// Unmarshal the byte array into a return value for GetSamplingProfile in the Memory domain.
func (a *GetSamplingProfileArgs) UnmarshalJSON(b []byte) error {
	type Copy GetSamplingProfileArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetSamplingProfileArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetSamplingProfile in the Memory domain.
func (a *GetSamplingProfileArgs) MarshalJSON() ([]byte, error) {
	type Copy GetSamplingProfileArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetSamplingProfileReply represents the return values for GetSamplingProfile in the Memory domain.
type GetSamplingProfileReply struct {
	Profile SamplingProfile `json:"profile"` // No description.
}

// GetSamplingProfileReply returns whether or not the FrameID matches the reply value for GetSamplingProfile in the Memory domain.
func (a *GetSamplingProfileReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetSamplingProfileReply", err)
	}
	return true
}

// GetSamplingProfileReply returns the FrameID value for GetSamplingProfile in the Memory domain.
func (a *GetSamplingProfileReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetSamplingProfile in the Memory domain.
func (a *GetSamplingProfileReply) UnmarshalJSON(b []byte) error {
	type Copy GetSamplingProfileReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetSamplingProfileReply(*c)
	return nil
}
