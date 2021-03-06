// Code generated by cdpgen. DO NOT EDIT.

package browser

import (
	"encoding/json"
	"log"

	shared "github.com/4ydx/cdp/protocol"
	"github.com/4ydx/cdp/protocol/target"
)

const (
	CommandBrowserSetPermission         = "Browser.setPermission"
	CommandBrowserGrantPermissions      = "Browser.grantPermissions"
	CommandBrowserResetPermissions      = "Browser.resetPermissions"
	CommandBrowserSetDownloadBehavior   = "Browser.setDownloadBehavior"
	CommandBrowserClose                 = "Browser.close"
	CommandBrowserCrash                 = "Browser.crash"
	CommandBrowserCrashGpuProcess       = "Browser.crashGpuProcess"
	CommandBrowserGetVersion            = "Browser.getVersion"
	CommandBrowserGetBrowserCommandLine = "Browser.getBrowserCommandLine"
	CommandBrowserGetHistograms         = "Browser.getHistograms"
	CommandBrowserGetHistogram          = "Browser.getHistogram"
	CommandBrowserGetWindowBounds       = "Browser.getWindowBounds"
	CommandBrowserGetWindowForTarget    = "Browser.getWindowForTarget"
	CommandBrowserSetWindowBounds       = "Browser.setWindowBounds"
	CommandBrowserSetDockTile           = "Browser.setDockTile"
)

// SetPermissionArgs represents the arguments for SetPermission in the Browser domain.
type SetPermissionArgs struct {
	Origin           string               `json:"origin,omitempty"`           // Origin the permission applies to, all origins if not specified.
	Permission       PermissionDescriptor `json:"permission"`                 // Descriptor of permission to override.
	Setting          PermissionSetting    `json:"setting"`                    // Setting of the permission.
	BrowserContextID shared.ContextID     `json:"browserContextId,omitempty"` // Context to override. When omitted, default browser context is used.
}

// Unmarshal the byte array into a return value for SetPermission in the Browser domain.
func (a *SetPermissionArgs) UnmarshalJSON(b []byte) error {
	type Copy SetPermissionArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetPermissionArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetPermission in the Browser domain.
func (a *SetPermissionArgs) MarshalJSON() ([]byte, error) {
	type Copy SetPermissionArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetPermissionReply represents the return values for SetPermission in the Browser domain.
type SetPermissionReply struct {
}

// SetPermissionReply returns whether or not the FrameID matches the reply value for SetPermission in the Browser domain.
func (a *SetPermissionReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: SetPermissionReply %s", err)
		return false, err
	}
	return true, nil
}

// SetPermissionReply returns the FrameID value for SetPermission in the Browser domain.
func (a *SetPermissionReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for SetPermission in the Browser domain.
func (a *SetPermissionReply) UnmarshalJSON(b []byte) error {
	type Copy SetPermissionReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetPermissionReply(*c)
	return nil
}

// GrantPermissionsArgs represents the arguments for GrantPermissions in the Browser domain.
type GrantPermissionsArgs struct {
	Origin           string           `json:"origin,omitempty"`           // Origin the permission applies to, all origins if not specified.
	Permissions      []PermissionType `json:"permissions"`                // No description.
	BrowserContextID shared.ContextID `json:"browserContextId,omitempty"` // BrowserContext to override permissions. When omitted, default browser context is used.
}

// Unmarshal the byte array into a return value for GrantPermissions in the Browser domain.
func (a *GrantPermissionsArgs) UnmarshalJSON(b []byte) error {
	type Copy GrantPermissionsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GrantPermissionsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GrantPermissions in the Browser domain.
func (a *GrantPermissionsArgs) MarshalJSON() ([]byte, error) {
	type Copy GrantPermissionsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GrantPermissionsReply represents the return values for GrantPermissions in the Browser domain.
type GrantPermissionsReply struct {
}

// GrantPermissionsReply returns whether or not the FrameID matches the reply value for GrantPermissions in the Browser domain.
func (a *GrantPermissionsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GrantPermissionsReply %s", err)
		return false, err
	}
	return true, nil
}

// GrantPermissionsReply returns the FrameID value for GrantPermissions in the Browser domain.
func (a *GrantPermissionsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GrantPermissions in the Browser domain.
func (a *GrantPermissionsReply) UnmarshalJSON(b []byte) error {
	type Copy GrantPermissionsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GrantPermissionsReply(*c)
	return nil
}

// ResetPermissionsArgs represents the arguments for ResetPermissions in the Browser domain.
type ResetPermissionsArgs struct {
	BrowserContextID shared.ContextID `json:"browserContextId,omitempty"` // BrowserContext to reset permissions. When omitted, default browser context is used.
}

// Unmarshal the byte array into a return value for ResetPermissions in the Browser domain.
func (a *ResetPermissionsArgs) UnmarshalJSON(b []byte) error {
	type Copy ResetPermissionsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ResetPermissionsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for ResetPermissions in the Browser domain.
func (a *ResetPermissionsArgs) MarshalJSON() ([]byte, error) {
	type Copy ResetPermissionsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// ResetPermissionsReply represents the return values for ResetPermissions in the Browser domain.
type ResetPermissionsReply struct {
}

// ResetPermissionsReply returns whether or not the FrameID matches the reply value for ResetPermissions in the Browser domain.
func (a *ResetPermissionsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: ResetPermissionsReply %s", err)
		return false, err
	}
	return true, nil
}

// ResetPermissionsReply returns the FrameID value for ResetPermissions in the Browser domain.
func (a *ResetPermissionsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for ResetPermissions in the Browser domain.
func (a *ResetPermissionsReply) UnmarshalJSON(b []byte) error {
	type Copy ResetPermissionsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ResetPermissionsReply(*c)
	return nil
}

// SetDownloadBehaviorArgs represents the arguments for SetDownloadBehavior in the Browser domain.
type SetDownloadBehaviorArgs struct {
	// Behavior Whether to allow all or deny all download requests, or use
	// default Chrome behavior if available (otherwise deny).
	// |allowAndName| allows download and names files according to their
	// dowmload guids.
	//
	// Values: "deny", "allow", "allowAndName", "default".
	Behavior         string           `json:"behavior"`
	BrowserContextID shared.ContextID `json:"browserContextId,omitempty"` // BrowserContext to set download behavior. When omitted, default browser context is used.
	DownloadPath     string           `json:"downloadPath,omitempty"`     // The default path to save downloaded files to. This is required if behavior is set to 'allow' or 'allowAndName'.
}

// Unmarshal the byte array into a return value for SetDownloadBehavior in the Browser domain.
func (a *SetDownloadBehaviorArgs) UnmarshalJSON(b []byte) error {
	type Copy SetDownloadBehaviorArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetDownloadBehaviorArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetDownloadBehavior in the Browser domain.
func (a *SetDownloadBehaviorArgs) MarshalJSON() ([]byte, error) {
	type Copy SetDownloadBehaviorArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetDownloadBehaviorReply represents the return values for SetDownloadBehavior in the Browser domain.
type SetDownloadBehaviorReply struct {
}

// SetDownloadBehaviorReply returns whether or not the FrameID matches the reply value for SetDownloadBehavior in the Browser domain.
func (a *SetDownloadBehaviorReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: SetDownloadBehaviorReply %s", err)
		return false, err
	}
	return true, nil
}

// SetDownloadBehaviorReply returns the FrameID value for SetDownloadBehavior in the Browser domain.
func (a *SetDownloadBehaviorReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for SetDownloadBehavior in the Browser domain.
func (a *SetDownloadBehaviorReply) UnmarshalJSON(b []byte) error {
	type Copy SetDownloadBehaviorReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetDownloadBehaviorReply(*c)
	return nil
}

// CloseArgs represents the arguments for Close in the Browser domain.
type CloseArgs struct {
}

// Unmarshal the byte array into a return value for Close in the Browser domain.
func (a *CloseArgs) UnmarshalJSON(b []byte) error {
	type Copy CloseArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CloseArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Close in the Browser domain.
func (a *CloseArgs) MarshalJSON() ([]byte, error) {
	type Copy CloseArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// CloseReply represents the return values for Close in the Browser domain.
type CloseReply struct {
}

// CloseReply returns whether or not the FrameID matches the reply value for Close in the Browser domain.
func (a *CloseReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: CloseReply %s", err)
		return false, err
	}
	return true, nil
}

// CloseReply returns the FrameID value for Close in the Browser domain.
func (a *CloseReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Close in the Browser domain.
func (a *CloseReply) UnmarshalJSON(b []byte) error {
	type Copy CloseReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CloseReply(*c)
	return nil
}

// CrashArgs represents the arguments for Crash in the Browser domain.
type CrashArgs struct {
}

// Unmarshal the byte array into a return value for Crash in the Browser domain.
func (a *CrashArgs) UnmarshalJSON(b []byte) error {
	type Copy CrashArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CrashArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Crash in the Browser domain.
func (a *CrashArgs) MarshalJSON() ([]byte, error) {
	type Copy CrashArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// CrashReply represents the return values for Crash in the Browser domain.
type CrashReply struct {
}

// CrashReply returns whether or not the FrameID matches the reply value for Crash in the Browser domain.
func (a *CrashReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: CrashReply %s", err)
		return false, err
	}
	return true, nil
}

// CrashReply returns the FrameID value for Crash in the Browser domain.
func (a *CrashReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Crash in the Browser domain.
func (a *CrashReply) UnmarshalJSON(b []byte) error {
	type Copy CrashReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CrashReply(*c)
	return nil
}

// CrashGPUProcessArgs represents the arguments for CrashGPUProcess in the Browser domain.
type CrashGPUProcessArgs struct {
}

// Unmarshal the byte array into a return value for CrashGPUProcess in the Browser domain.
func (a *CrashGPUProcessArgs) UnmarshalJSON(b []byte) error {
	type Copy CrashGPUProcessArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CrashGPUProcessArgs(*c)
	return nil
}

// Marshall the byte array into a return value for CrashGPUProcess in the Browser domain.
func (a *CrashGPUProcessArgs) MarshalJSON() ([]byte, error) {
	type Copy CrashGPUProcessArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// CrashGPUProcessReply represents the return values for CrashGPUProcess in the Browser domain.
type CrashGPUProcessReply struct {
}

// CrashGPUProcessReply returns whether or not the FrameID matches the reply value for CrashGPUProcess in the Browser domain.
func (a *CrashGPUProcessReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: CrashGPUProcessReply %s", err)
		return false, err
	}
	return true, nil
}

// CrashGPUProcessReply returns the FrameID value for CrashGPUProcess in the Browser domain.
func (a *CrashGPUProcessReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for CrashGPUProcess in the Browser domain.
func (a *CrashGPUProcessReply) UnmarshalJSON(b []byte) error {
	type Copy CrashGPUProcessReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CrashGPUProcessReply(*c)
	return nil
}

// GetVersionArgs represents the arguments for GetVersion in the Browser domain.
type GetVersionArgs struct {
}

// Unmarshal the byte array into a return value for GetVersion in the Browser domain.
func (a *GetVersionArgs) UnmarshalJSON(b []byte) error {
	type Copy GetVersionArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetVersionArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetVersion in the Browser domain.
func (a *GetVersionArgs) MarshalJSON() ([]byte, error) {
	type Copy GetVersionArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetVersionReply represents the return values for GetVersion in the Browser domain.
type GetVersionReply struct {
	ProtocolVersion string `json:"protocolVersion"` // Protocol version.
	Product         string `json:"product"`         // Product name.
	Revision        string `json:"revision"`        // Product revision.
	UserAgent       string `json:"userAgent"`       // User-Agent.
	JsVersion       string `json:"jsVersion"`       // V8 version.
}

// GetVersionReply returns whether or not the FrameID matches the reply value for GetVersion in the Browser domain.
func (a *GetVersionReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetVersionReply %s", err)
		return false, err
	}
	return true, nil
}

// GetVersionReply returns the FrameID value for GetVersion in the Browser domain.
func (a *GetVersionReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetVersion in the Browser domain.
func (a *GetVersionReply) UnmarshalJSON(b []byte) error {
	type Copy GetVersionReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetVersionReply(*c)
	return nil
}

// GetBrowserCommandLineArgs represents the arguments for GetBrowserCommandLine in the Browser domain.
type GetBrowserCommandLineArgs struct {
}

// Unmarshal the byte array into a return value for GetBrowserCommandLine in the Browser domain.
func (a *GetBrowserCommandLineArgs) UnmarshalJSON(b []byte) error {
	type Copy GetBrowserCommandLineArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetBrowserCommandLineArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetBrowserCommandLine in the Browser domain.
func (a *GetBrowserCommandLineArgs) MarshalJSON() ([]byte, error) {
	type Copy GetBrowserCommandLineArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetBrowserCommandLineReply represents the return values for GetBrowserCommandLine in the Browser domain.
type GetBrowserCommandLineReply struct {
	Arguments []string `json:"arguments"` // Commandline parameters
}

// GetBrowserCommandLineReply returns whether or not the FrameID matches the reply value for GetBrowserCommandLine in the Browser domain.
func (a *GetBrowserCommandLineReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetBrowserCommandLineReply %s", err)
		return false, err
	}
	return true, nil
}

// GetBrowserCommandLineReply returns the FrameID value for GetBrowserCommandLine in the Browser domain.
func (a *GetBrowserCommandLineReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetBrowserCommandLine in the Browser domain.
func (a *GetBrowserCommandLineReply) UnmarshalJSON(b []byte) error {
	type Copy GetBrowserCommandLineReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetBrowserCommandLineReply(*c)
	return nil
}

// GetHistogramsArgs represents the arguments for GetHistograms in the Browser domain.
type GetHistogramsArgs struct {
	Query string `json:"query,omitempty"` // Requested substring in name. Only histograms which have query as a substring in their name are extracted. An empty or absent query returns all histograms.
	Delta bool   `json:"delta,omitempty"` // If true, retrieve delta since last call.
}

// Unmarshal the byte array into a return value for GetHistograms in the Browser domain.
func (a *GetHistogramsArgs) UnmarshalJSON(b []byte) error {
	type Copy GetHistogramsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetHistogramsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetHistograms in the Browser domain.
func (a *GetHistogramsArgs) MarshalJSON() ([]byte, error) {
	type Copy GetHistogramsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetHistogramsReply represents the return values for GetHistograms in the Browser domain.
type GetHistogramsReply struct {
	Histograms []Histogram `json:"histograms"` // Histograms.
}

// GetHistogramsReply returns whether or not the FrameID matches the reply value for GetHistograms in the Browser domain.
func (a *GetHistogramsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetHistogramsReply %s", err)
		return false, err
	}
	return true, nil
}

// GetHistogramsReply returns the FrameID value for GetHistograms in the Browser domain.
func (a *GetHistogramsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetHistograms in the Browser domain.
func (a *GetHistogramsReply) UnmarshalJSON(b []byte) error {
	type Copy GetHistogramsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetHistogramsReply(*c)
	return nil
}

// GetHistogramArgs represents the arguments for GetHistogram in the Browser domain.
type GetHistogramArgs struct {
	Name  string `json:"name"`            // Requested histogram name.
	Delta bool   `json:"delta,omitempty"` // If true, retrieve delta since last call.
}

// Unmarshal the byte array into a return value for GetHistogram in the Browser domain.
func (a *GetHistogramArgs) UnmarshalJSON(b []byte) error {
	type Copy GetHistogramArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetHistogramArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetHistogram in the Browser domain.
func (a *GetHistogramArgs) MarshalJSON() ([]byte, error) {
	type Copy GetHistogramArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetHistogramReply represents the return values for GetHistogram in the Browser domain.
type GetHistogramReply struct {
	Histogram Histogram `json:"histogram"` // Histogram.
}

// GetHistogramReply returns whether or not the FrameID matches the reply value for GetHistogram in the Browser domain.
func (a *GetHistogramReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetHistogramReply %s", err)
		return false, err
	}
	return true, nil
}

// GetHistogramReply returns the FrameID value for GetHistogram in the Browser domain.
func (a *GetHistogramReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetHistogram in the Browser domain.
func (a *GetHistogramReply) UnmarshalJSON(b []byte) error {
	type Copy GetHistogramReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetHistogramReply(*c)
	return nil
}

// GetWindowBoundsArgs represents the arguments for GetWindowBounds in the Browser domain.
type GetWindowBoundsArgs struct {
	WindowID WindowID `json:"windowId"` // Browser window id.
}

// Unmarshal the byte array into a return value for GetWindowBounds in the Browser domain.
func (a *GetWindowBoundsArgs) UnmarshalJSON(b []byte) error {
	type Copy GetWindowBoundsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetWindowBoundsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetWindowBounds in the Browser domain.
func (a *GetWindowBoundsArgs) MarshalJSON() ([]byte, error) {
	type Copy GetWindowBoundsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetWindowBoundsReply represents the return values for GetWindowBounds in the Browser domain.
type GetWindowBoundsReply struct {
	Bounds Bounds `json:"bounds"` // Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
}

// GetWindowBoundsReply returns whether or not the FrameID matches the reply value for GetWindowBounds in the Browser domain.
func (a *GetWindowBoundsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetWindowBoundsReply %s", err)
		return false, err
	}
	return true, nil
}

// GetWindowBoundsReply returns the FrameID value for GetWindowBounds in the Browser domain.
func (a *GetWindowBoundsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetWindowBounds in the Browser domain.
func (a *GetWindowBoundsReply) UnmarshalJSON(b []byte) error {
	type Copy GetWindowBoundsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetWindowBoundsReply(*c)
	return nil
}

// GetWindowForTargetArgs represents the arguments for GetWindowForTarget in the Browser domain.
type GetWindowForTargetArgs struct {
	TargetID target.ID `json:"targetId,omitempty"` // Devtools agent host id. If called as a part of the session, associated targetId is used.
}

// Unmarshal the byte array into a return value for GetWindowForTarget in the Browser domain.
func (a *GetWindowForTargetArgs) UnmarshalJSON(b []byte) error {
	type Copy GetWindowForTargetArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetWindowForTargetArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetWindowForTarget in the Browser domain.
func (a *GetWindowForTargetArgs) MarshalJSON() ([]byte, error) {
	type Copy GetWindowForTargetArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetWindowForTargetReply represents the return values for GetWindowForTarget in the Browser domain.
type GetWindowForTargetReply struct {
	WindowID WindowID `json:"windowId"` // Browser window id.
	Bounds   Bounds   `json:"bounds"`   // Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
}

// GetWindowForTargetReply returns whether or not the FrameID matches the reply value for GetWindowForTarget in the Browser domain.
func (a *GetWindowForTargetReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetWindowForTargetReply %s", err)
		return false, err
	}
	return true, nil
}

// GetWindowForTargetReply returns the FrameID value for GetWindowForTarget in the Browser domain.
func (a *GetWindowForTargetReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetWindowForTarget in the Browser domain.
func (a *GetWindowForTargetReply) UnmarshalJSON(b []byte) error {
	type Copy GetWindowForTargetReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetWindowForTargetReply(*c)
	return nil
}

// SetWindowBoundsArgs represents the arguments for SetWindowBounds in the Browser domain.
type SetWindowBoundsArgs struct {
	WindowID WindowID `json:"windowId"` // Browser window id.
	Bounds   Bounds   `json:"bounds"`   // New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
}

// Unmarshal the byte array into a return value for SetWindowBounds in the Browser domain.
func (a *SetWindowBoundsArgs) UnmarshalJSON(b []byte) error {
	type Copy SetWindowBoundsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetWindowBoundsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetWindowBounds in the Browser domain.
func (a *SetWindowBoundsArgs) MarshalJSON() ([]byte, error) {
	type Copy SetWindowBoundsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetWindowBoundsReply represents the return values for SetWindowBounds in the Browser domain.
type SetWindowBoundsReply struct {
}

// SetWindowBoundsReply returns whether or not the FrameID matches the reply value for SetWindowBounds in the Browser domain.
func (a *SetWindowBoundsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: SetWindowBoundsReply %s", err)
		return false, err
	}
	return true, nil
}

// SetWindowBoundsReply returns the FrameID value for SetWindowBounds in the Browser domain.
func (a *SetWindowBoundsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for SetWindowBounds in the Browser domain.
func (a *SetWindowBoundsReply) UnmarshalJSON(b []byte) error {
	type Copy SetWindowBoundsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetWindowBoundsReply(*c)
	return nil
}

// SetDockTileArgs represents the arguments for SetDockTile in the Browser domain.
type SetDockTileArgs struct {
	BadgeLabel string `json:"badgeLabel,omitempty"` // No description.
	Image      string `json:"image,omitempty"`      // Png encoded image.
}

// Unmarshal the byte array into a return value for SetDockTile in the Browser domain.
func (a *SetDockTileArgs) UnmarshalJSON(b []byte) error {
	type Copy SetDockTileArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetDockTileArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetDockTile in the Browser domain.
func (a *SetDockTileArgs) MarshalJSON() ([]byte, error) {
	type Copy SetDockTileArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetDockTileReply represents the return values for SetDockTile in the Browser domain.
type SetDockTileReply struct {
}

// SetDockTileReply returns whether or not the FrameID matches the reply value for SetDockTile in the Browser domain.
func (a *SetDockTileReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: SetDockTileReply %s", err)
		return false, err
	}
	return true, nil
}

// SetDockTileReply returns the FrameID value for SetDockTile in the Browser domain.
func (a *SetDockTileReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for SetDockTile in the Browser domain.
func (a *SetDockTileReply) UnmarshalJSON(b []byte) error {
	type Copy SetDockTileReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetDockTileReply(*c)
	return nil
}
