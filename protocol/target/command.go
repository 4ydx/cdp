// Code generated by cdpgen. DO NOT EDIT.

package target

import (
	"encoding/json"
	"log"

	shared "github.com/4ydx/cdp/protocol"
)

const (
	CommandTargetActivateTarget         = "Target.activateTarget"
	CommandTargetAttachToTarget         = "Target.attachToTarget"
	CommandTargetAttachToBrowserTarget  = "Target.attachToBrowserTarget"
	CommandTargetCloseTarget            = "Target.closeTarget"
	CommandTargetExposeDevToolsProtocol = "Target.exposeDevToolsProtocol"
	CommandTargetCreateBrowserContext   = "Target.createBrowserContext"
	CommandTargetGetBrowserContexts     = "Target.getBrowserContexts"
	CommandTargetCreateTarget           = "Target.createTarget"
	CommandTargetDetachFromTarget       = "Target.detachFromTarget"
	CommandTargetDisposeBrowserContext  = "Target.disposeBrowserContext"
	CommandTargetGetTargetInfo          = "Target.getTargetInfo"
	CommandTargetGetTargets             = "Target.getTargets"
	CommandTargetSendMessageToTarget    = "Target.sendMessageToTarget"
	CommandTargetSetAutoAttach          = "Target.setAutoAttach"
	CommandTargetSetDiscoverTargets     = "Target.setDiscoverTargets"
	CommandTargetSetRemoteLocations     = "Target.setRemoteLocations"
)

// ActivateTargetArgs represents the arguments for ActivateTarget in the Target domain.
type ActivateTargetArgs struct {
	TargetID ID `json:"targetId"` // No description.
}

// Unmarshal the byte array into a return value for ActivateTarget in the Target domain.
func (a *ActivateTargetArgs) UnmarshalJSON(b []byte) error {
	type Copy ActivateTargetArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ActivateTargetArgs(*c)
	return nil
}

// Marshall the byte array into a return value for ActivateTarget in the Target domain.
func (a *ActivateTargetArgs) MarshalJSON() ([]byte, error) {
	type Copy ActivateTargetArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// ActivateTargetReply represents the return values for ActivateTarget in the Target domain.
type ActivateTargetReply struct {
}

// ActivateTargetReply returns whether or not the FrameID matches the reply value for ActivateTarget in the Target domain.
func (a *ActivateTargetReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: ActivateTargetReply %s", err)
		return false, err
	}
	return true, nil
}

// ActivateTargetReply returns the FrameID value for ActivateTarget in the Target domain.
func (a *ActivateTargetReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for ActivateTarget in the Target domain.
func (a *ActivateTargetReply) UnmarshalJSON(b []byte) error {
	type Copy ActivateTargetReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ActivateTargetReply(*c)
	return nil
}

// AttachToTargetArgs represents the arguments for AttachToTarget in the Target domain.
type AttachToTargetArgs struct {
	TargetID ID   `json:"targetId"`          // No description.
	Flatten  bool `json:"flatten,omitempty"` // Enables "flat" access to the session via specifying sessionId attribute in the commands. We plan to make this the default, deprecate non-flattened mode, and eventually retire it. See crbug.com/991325.
}

// Unmarshal the byte array into a return value for AttachToTarget in the Target domain.
func (a *AttachToTargetArgs) UnmarshalJSON(b []byte) error {
	type Copy AttachToTargetArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = AttachToTargetArgs(*c)
	return nil
}

// Marshall the byte array into a return value for AttachToTarget in the Target domain.
func (a *AttachToTargetArgs) MarshalJSON() ([]byte, error) {
	type Copy AttachToTargetArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// AttachToTargetReply represents the return values for AttachToTarget in the Target domain.
type AttachToTargetReply struct {
	SessionID SessionID `json:"sessionId"` // Id assigned to the session.
}

// AttachToTargetReply returns whether or not the FrameID matches the reply value for AttachToTarget in the Target domain.
func (a *AttachToTargetReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: AttachToTargetReply %s", err)
		return false, err
	}
	return true, nil
}

// AttachToTargetReply returns the FrameID value for AttachToTarget in the Target domain.
func (a *AttachToTargetReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for AttachToTarget in the Target domain.
func (a *AttachToTargetReply) UnmarshalJSON(b []byte) error {
	type Copy AttachToTargetReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = AttachToTargetReply(*c)
	return nil
}

// AttachToBrowserTargetArgs represents the arguments for AttachToBrowserTarget in the Target domain.
type AttachToBrowserTargetArgs struct {
}

// Unmarshal the byte array into a return value for AttachToBrowserTarget in the Target domain.
func (a *AttachToBrowserTargetArgs) UnmarshalJSON(b []byte) error {
	type Copy AttachToBrowserTargetArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = AttachToBrowserTargetArgs(*c)
	return nil
}

// Marshall the byte array into a return value for AttachToBrowserTarget in the Target domain.
func (a *AttachToBrowserTargetArgs) MarshalJSON() ([]byte, error) {
	type Copy AttachToBrowserTargetArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// AttachToBrowserTargetReply represents the return values for AttachToBrowserTarget in the Target domain.
type AttachToBrowserTargetReply struct {
	SessionID SessionID `json:"sessionId"` // Id assigned to the session.
}

// AttachToBrowserTargetReply returns whether or not the FrameID matches the reply value for AttachToBrowserTarget in the Target domain.
func (a *AttachToBrowserTargetReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: AttachToBrowserTargetReply %s", err)
		return false, err
	}
	return true, nil
}

// AttachToBrowserTargetReply returns the FrameID value for AttachToBrowserTarget in the Target domain.
func (a *AttachToBrowserTargetReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for AttachToBrowserTarget in the Target domain.
func (a *AttachToBrowserTargetReply) UnmarshalJSON(b []byte) error {
	type Copy AttachToBrowserTargetReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = AttachToBrowserTargetReply(*c)
	return nil
}

// CloseTargetArgs represents the arguments for CloseTarget in the Target domain.
type CloseTargetArgs struct {
	TargetID ID `json:"targetId"` // No description.
}

// Unmarshal the byte array into a return value for CloseTarget in the Target domain.
func (a *CloseTargetArgs) UnmarshalJSON(b []byte) error {
	type Copy CloseTargetArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CloseTargetArgs(*c)
	return nil
}

// Marshall the byte array into a return value for CloseTarget in the Target domain.
func (a *CloseTargetArgs) MarshalJSON() ([]byte, error) {
	type Copy CloseTargetArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// CloseTargetReply represents the return values for CloseTarget in the Target domain.
type CloseTargetReply struct {
	Success bool `json:"success"` // No description.
}

// CloseTargetReply returns whether or not the FrameID matches the reply value for CloseTarget in the Target domain.
func (a *CloseTargetReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: CloseTargetReply %s", err)
		return false, err
	}
	return true, nil
}

// CloseTargetReply returns the FrameID value for CloseTarget in the Target domain.
func (a *CloseTargetReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for CloseTarget in the Target domain.
func (a *CloseTargetReply) UnmarshalJSON(b []byte) error {
	type Copy CloseTargetReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CloseTargetReply(*c)
	return nil
}

// ExposeDevToolsProtocolArgs represents the arguments for ExposeDevToolsProtocol in the Target domain.
type ExposeDevToolsProtocolArgs struct {
	TargetID    ID     `json:"targetId"`              // No description.
	BindingName string `json:"bindingName,omitempty"` // Binding name, 'cdp' if not specified.
}

// Unmarshal the byte array into a return value for ExposeDevToolsProtocol in the Target domain.
func (a *ExposeDevToolsProtocolArgs) UnmarshalJSON(b []byte) error {
	type Copy ExposeDevToolsProtocolArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ExposeDevToolsProtocolArgs(*c)
	return nil
}

// Marshall the byte array into a return value for ExposeDevToolsProtocol in the Target domain.
func (a *ExposeDevToolsProtocolArgs) MarshalJSON() ([]byte, error) {
	type Copy ExposeDevToolsProtocolArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// ExposeDevToolsProtocolReply represents the return values for ExposeDevToolsProtocol in the Target domain.
type ExposeDevToolsProtocolReply struct {
}

// ExposeDevToolsProtocolReply returns whether or not the FrameID matches the reply value for ExposeDevToolsProtocol in the Target domain.
func (a *ExposeDevToolsProtocolReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: ExposeDevToolsProtocolReply %s", err)
		return false, err
	}
	return true, nil
}

// ExposeDevToolsProtocolReply returns the FrameID value for ExposeDevToolsProtocol in the Target domain.
func (a *ExposeDevToolsProtocolReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for ExposeDevToolsProtocol in the Target domain.
func (a *ExposeDevToolsProtocolReply) UnmarshalJSON(b []byte) error {
	type Copy ExposeDevToolsProtocolReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ExposeDevToolsProtocolReply(*c)
	return nil
}

// CreateBrowserContextArgs represents the arguments for CreateBrowserContext in the Target domain.
type CreateBrowserContextArgs struct {
	DisposeOnDetach bool `json:"disposeOnDetach,omitempty"` // If specified, disposes this context when debugging session disconnects.
}

// Unmarshal the byte array into a return value for CreateBrowserContext in the Target domain.
func (a *CreateBrowserContextArgs) UnmarshalJSON(b []byte) error {
	type Copy CreateBrowserContextArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CreateBrowserContextArgs(*c)
	return nil
}

// Marshall the byte array into a return value for CreateBrowserContext in the Target domain.
func (a *CreateBrowserContextArgs) MarshalJSON() ([]byte, error) {
	type Copy CreateBrowserContextArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// CreateBrowserContextReply represents the return values for CreateBrowserContext in the Target domain.
type CreateBrowserContextReply struct {
	BrowserContextID shared.ContextID `json:"browserContextId"` // The id of the context created.
}

// CreateBrowserContextReply returns whether or not the FrameID matches the reply value for CreateBrowserContext in the Target domain.
func (a *CreateBrowserContextReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: CreateBrowserContextReply %s", err)
		return false, err
	}
	return true, nil
}

// CreateBrowserContextReply returns the FrameID value for CreateBrowserContext in the Target domain.
func (a *CreateBrowserContextReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for CreateBrowserContext in the Target domain.
func (a *CreateBrowserContextReply) UnmarshalJSON(b []byte) error {
	type Copy CreateBrowserContextReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CreateBrowserContextReply(*c)
	return nil
}

// GetBrowserContextsArgs represents the arguments for GetBrowserContexts in the Target domain.
type GetBrowserContextsArgs struct {
}

// Unmarshal the byte array into a return value for GetBrowserContexts in the Target domain.
func (a *GetBrowserContextsArgs) UnmarshalJSON(b []byte) error {
	type Copy GetBrowserContextsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetBrowserContextsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetBrowserContexts in the Target domain.
func (a *GetBrowserContextsArgs) MarshalJSON() ([]byte, error) {
	type Copy GetBrowserContextsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetBrowserContextsReply represents the return values for GetBrowserContexts in the Target domain.
type GetBrowserContextsReply struct {
	BrowserContextIDs []shared.ContextID `json:"browserContextIds"` // An array of browser context ids.
}

// GetBrowserContextsReply returns whether or not the FrameID matches the reply value for GetBrowserContexts in the Target domain.
func (a *GetBrowserContextsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetBrowserContextsReply %s", err)
		return false, err
	}
	return true, nil
}

// GetBrowserContextsReply returns the FrameID value for GetBrowserContexts in the Target domain.
func (a *GetBrowserContextsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetBrowserContexts in the Target domain.
func (a *GetBrowserContextsReply) UnmarshalJSON(b []byte) error {
	type Copy GetBrowserContextsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetBrowserContextsReply(*c)
	return nil
}

// CreateTargetArgs represents the arguments for CreateTarget in the Target domain.
type CreateTargetArgs struct {
	URL              string           `json:"url"`                        // The initial URL the page will be navigated to.
	Width            int              `json:"width,omitempty"`            // Frame width in DIP (headless chrome only).
	Height           int              `json:"height,omitempty"`           // Frame height in DIP (headless chrome only).
	BrowserContextID shared.ContextID `json:"browserContextId,omitempty"` // The browser context to create the page in.
	// EnableBeginFrameControl Whether BeginFrames for this target will be
	// controlled via DevTools (headless chrome only, not supported on
	// MacOS yet, false by default).
	//
	// Note: This property is experimental.
	EnableBeginFrameControl bool `json:"enableBeginFrameControl,omitempty"`
	NewWindow               bool `json:"newWindow,omitempty"`  // Whether to create a new Window or Tab (chrome-only, false by default).
	Background              bool `json:"background,omitempty"` // Whether to create the target in background or foreground (chrome-only, false by default).
}

// Unmarshal the byte array into a return value for CreateTarget in the Target domain.
func (a *CreateTargetArgs) UnmarshalJSON(b []byte) error {
	type Copy CreateTargetArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CreateTargetArgs(*c)
	return nil
}

// Marshall the byte array into a return value for CreateTarget in the Target domain.
func (a *CreateTargetArgs) MarshalJSON() ([]byte, error) {
	type Copy CreateTargetArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// CreateTargetReply represents the return values for CreateTarget in the Target domain.
type CreateTargetReply struct {
	TargetID ID `json:"targetId"` // The id of the page opened.
}

// CreateTargetReply returns whether or not the FrameID matches the reply value for CreateTarget in the Target domain.
func (a *CreateTargetReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: CreateTargetReply %s", err)
		return false, err
	}
	return true, nil
}

// CreateTargetReply returns the FrameID value for CreateTarget in the Target domain.
func (a *CreateTargetReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for CreateTarget in the Target domain.
func (a *CreateTargetReply) UnmarshalJSON(b []byte) error {
	type Copy CreateTargetReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CreateTargetReply(*c)
	return nil
}

// DetachFromTargetArgs represents the arguments for DetachFromTarget in the Target domain.
type DetachFromTargetArgs struct {
	SessionID SessionID `json:"sessionId,omitempty"` // Session to detach.
	// TargetID is deprecated.
	//
	// Deprecated: Deprecated.
	TargetID ID `json:"targetId,omitempty"`
}

// Unmarshal the byte array into a return value for DetachFromTarget in the Target domain.
func (a *DetachFromTargetArgs) UnmarshalJSON(b []byte) error {
	type Copy DetachFromTargetArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DetachFromTargetArgs(*c)
	return nil
}

// Marshall the byte array into a return value for DetachFromTarget in the Target domain.
func (a *DetachFromTargetArgs) MarshalJSON() ([]byte, error) {
	type Copy DetachFromTargetArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DetachFromTargetReply represents the return values for DetachFromTarget in the Target domain.
type DetachFromTargetReply struct {
}

// DetachFromTargetReply returns whether or not the FrameID matches the reply value for DetachFromTarget in the Target domain.
func (a *DetachFromTargetReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: DetachFromTargetReply %s", err)
		return false, err
	}
	return true, nil
}

// DetachFromTargetReply returns the FrameID value for DetachFromTarget in the Target domain.
func (a *DetachFromTargetReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for DetachFromTarget in the Target domain.
func (a *DetachFromTargetReply) UnmarshalJSON(b []byte) error {
	type Copy DetachFromTargetReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DetachFromTargetReply(*c)
	return nil
}

// DisposeBrowserContextArgs represents the arguments for DisposeBrowserContext in the Target domain.
type DisposeBrowserContextArgs struct {
	BrowserContextID shared.ContextID `json:"browserContextId"` // No description.
}

// Unmarshal the byte array into a return value for DisposeBrowserContext in the Target domain.
func (a *DisposeBrowserContextArgs) UnmarshalJSON(b []byte) error {
	type Copy DisposeBrowserContextArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DisposeBrowserContextArgs(*c)
	return nil
}

// Marshall the byte array into a return value for DisposeBrowserContext in the Target domain.
func (a *DisposeBrowserContextArgs) MarshalJSON() ([]byte, error) {
	type Copy DisposeBrowserContextArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DisposeBrowserContextReply represents the return values for DisposeBrowserContext in the Target domain.
type DisposeBrowserContextReply struct {
}

// DisposeBrowserContextReply returns whether or not the FrameID matches the reply value for DisposeBrowserContext in the Target domain.
func (a *DisposeBrowserContextReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: DisposeBrowserContextReply %s", err)
		return false, err
	}
	return true, nil
}

// DisposeBrowserContextReply returns the FrameID value for DisposeBrowserContext in the Target domain.
func (a *DisposeBrowserContextReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for DisposeBrowserContext in the Target domain.
func (a *DisposeBrowserContextReply) UnmarshalJSON(b []byte) error {
	type Copy DisposeBrowserContextReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DisposeBrowserContextReply(*c)
	return nil
}

// GetTargetInfoArgs represents the arguments for GetTargetInfo in the Target domain.
type GetTargetInfoArgs struct {
	TargetID ID `json:"targetId,omitempty"` // No description.
}

// Unmarshal the byte array into a return value for GetTargetInfo in the Target domain.
func (a *GetTargetInfoArgs) UnmarshalJSON(b []byte) error {
	type Copy GetTargetInfoArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetTargetInfoArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetTargetInfo in the Target domain.
func (a *GetTargetInfoArgs) MarshalJSON() ([]byte, error) {
	type Copy GetTargetInfoArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetTargetInfoReply represents the return values for GetTargetInfo in the Target domain.
type GetTargetInfoReply struct {
	TargetInfo Info `json:"targetInfo"` // No description.
}

// GetTargetInfoReply returns whether or not the FrameID matches the reply value for GetTargetInfo in the Target domain.
func (a *GetTargetInfoReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetTargetInfoReply %s", err)
		return false, err
	}
	return true, nil
}

// GetTargetInfoReply returns the FrameID value for GetTargetInfo in the Target domain.
func (a *GetTargetInfoReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetTargetInfo in the Target domain.
func (a *GetTargetInfoReply) UnmarshalJSON(b []byte) error {
	type Copy GetTargetInfoReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetTargetInfoReply(*c)
	return nil
}

// GetTargetsArgs represents the arguments for GetTargets in the Target domain.
type GetTargetsArgs struct {
}

// Unmarshal the byte array into a return value for GetTargets in the Target domain.
func (a *GetTargetsArgs) UnmarshalJSON(b []byte) error {
	type Copy GetTargetsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetTargetsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetTargets in the Target domain.
func (a *GetTargetsArgs) MarshalJSON() ([]byte, error) {
	type Copy GetTargetsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetTargetsReply represents the return values for GetTargets in the Target domain.
type GetTargetsReply struct {
	TargetInfos []Info `json:"targetInfos"` // The list of targets.
}

// GetTargetsReply returns whether or not the FrameID matches the reply value for GetTargets in the Target domain.
func (a *GetTargetsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetTargetsReply %s", err)
		return false, err
	}
	return true, nil
}

// GetTargetsReply returns the FrameID value for GetTargets in the Target domain.
func (a *GetTargetsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetTargets in the Target domain.
func (a *GetTargetsReply) UnmarshalJSON(b []byte) error {
	type Copy GetTargetsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetTargetsReply(*c)
	return nil
}

// SendMessageToTargetArgs represents the arguments for SendMessageToTarget in the Target domain.
type SendMessageToTargetArgs struct {
	Message   string    `json:"message"`             // No description.
	SessionID SessionID `json:"sessionId,omitempty"` // Identifier of the session.
	// TargetID is deprecated.
	//
	// Deprecated: Deprecated.
	TargetID ID `json:"targetId,omitempty"`
}

// Unmarshal the byte array into a return value for SendMessageToTarget in the Target domain.
func (a *SendMessageToTargetArgs) UnmarshalJSON(b []byte) error {
	type Copy SendMessageToTargetArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SendMessageToTargetArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SendMessageToTarget in the Target domain.
func (a *SendMessageToTargetArgs) MarshalJSON() ([]byte, error) {
	type Copy SendMessageToTargetArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SendMessageToTargetReply represents the return values for SendMessageToTarget in the Target domain.
type SendMessageToTargetReply struct {
}

// SendMessageToTargetReply returns whether or not the FrameID matches the reply value for SendMessageToTarget in the Target domain.
func (a *SendMessageToTargetReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: SendMessageToTargetReply %s", err)
		return false, err
	}
	return true, nil
}

// SendMessageToTargetReply returns the FrameID value for SendMessageToTarget in the Target domain.
func (a *SendMessageToTargetReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for SendMessageToTarget in the Target domain.
func (a *SendMessageToTargetReply) UnmarshalJSON(b []byte) error {
	type Copy SendMessageToTargetReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SendMessageToTargetReply(*c)
	return nil
}

// SetAutoAttachArgs represents the arguments for SetAutoAttach in the Target domain.
type SetAutoAttachArgs struct {
	AutoAttach             bool `json:"autoAttach"`             // Whether to auto-attach to related targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"` // Whether to pause new targets when attaching to them. Use `Runtime.runIfWaitingForDebugger` to run paused targets.
	Flatten                bool `json:"flatten,omitempty"`      // Enables "flat" access to the session via specifying sessionId attribute in the commands. We plan to make this the default, deprecate non-flattened mode, and eventually retire it. See crbug.com/991325.
}

// Unmarshal the byte array into a return value for SetAutoAttach in the Target domain.
func (a *SetAutoAttachArgs) UnmarshalJSON(b []byte) error {
	type Copy SetAutoAttachArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetAutoAttachArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetAutoAttach in the Target domain.
func (a *SetAutoAttachArgs) MarshalJSON() ([]byte, error) {
	type Copy SetAutoAttachArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetAutoAttachReply represents the return values for SetAutoAttach in the Target domain.
type SetAutoAttachReply struct {
}

// SetAutoAttachReply returns whether or not the FrameID matches the reply value for SetAutoAttach in the Target domain.
func (a *SetAutoAttachReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: SetAutoAttachReply %s", err)
		return false, err
	}
	return true, nil
}

// SetAutoAttachReply returns the FrameID value for SetAutoAttach in the Target domain.
func (a *SetAutoAttachReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for SetAutoAttach in the Target domain.
func (a *SetAutoAttachReply) UnmarshalJSON(b []byte) error {
	type Copy SetAutoAttachReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetAutoAttachReply(*c)
	return nil
}

// SetDiscoverTargetsArgs represents the arguments for SetDiscoverTargets in the Target domain.
type SetDiscoverTargetsArgs struct {
	Discover bool `json:"discover"` // Whether to discover available targets.
}

// Unmarshal the byte array into a return value for SetDiscoverTargets in the Target domain.
func (a *SetDiscoverTargetsArgs) UnmarshalJSON(b []byte) error {
	type Copy SetDiscoverTargetsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetDiscoverTargetsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetDiscoverTargets in the Target domain.
func (a *SetDiscoverTargetsArgs) MarshalJSON() ([]byte, error) {
	type Copy SetDiscoverTargetsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetDiscoverTargetsReply represents the return values for SetDiscoverTargets in the Target domain.
type SetDiscoverTargetsReply struct {
}

// SetDiscoverTargetsReply returns whether or not the FrameID matches the reply value for SetDiscoverTargets in the Target domain.
func (a *SetDiscoverTargetsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: SetDiscoverTargetsReply %s", err)
		return false, err
	}
	return true, nil
}

// SetDiscoverTargetsReply returns the FrameID value for SetDiscoverTargets in the Target domain.
func (a *SetDiscoverTargetsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for SetDiscoverTargets in the Target domain.
func (a *SetDiscoverTargetsReply) UnmarshalJSON(b []byte) error {
	type Copy SetDiscoverTargetsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetDiscoverTargetsReply(*c)
	return nil
}

// SetRemoteLocationsArgs represents the arguments for SetRemoteLocations in the Target domain.
type SetRemoteLocationsArgs struct {
	Locations []RemoteLocation `json:"locations"` // List of remote locations.
}

// Unmarshal the byte array into a return value for SetRemoteLocations in the Target domain.
func (a *SetRemoteLocationsArgs) UnmarshalJSON(b []byte) error {
	type Copy SetRemoteLocationsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetRemoteLocationsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetRemoteLocations in the Target domain.
func (a *SetRemoteLocationsArgs) MarshalJSON() ([]byte, error) {
	type Copy SetRemoteLocationsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetRemoteLocationsReply represents the return values for SetRemoteLocations in the Target domain.
type SetRemoteLocationsReply struct {
}

// SetRemoteLocationsReply returns whether or not the FrameID matches the reply value for SetRemoteLocations in the Target domain.
func (a *SetRemoteLocationsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: SetRemoteLocationsReply %s", err)
		return false, err
	}
	return true, nil
}

// SetRemoteLocationsReply returns the FrameID value for SetRemoteLocations in the Target domain.
func (a *SetRemoteLocationsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for SetRemoteLocations in the Target domain.
func (a *SetRemoteLocationsReply) UnmarshalJSON(b []byte) error {
	type Copy SetRemoteLocationsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetRemoteLocationsReply(*c)
	return nil
}
