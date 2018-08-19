// Code generated by cdpgen. DO NOT EDIT.

package browser

import (
	"encoding/json"

	"github.com/4ydx/cdp/protocol/target"
)

const CommandBrowserClose = "Browser.close"

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

const CommandBrowserGetVersion = "Browser.getVersion"

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

const CommandBrowserGetBrowserCommandLine = "Browser.getBrowserCommandLine"

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

const CommandBrowserGetHistograms = "Browser.getHistograms"

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

const CommandBrowserGetHistogram = "Browser.getHistogram"

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

const CommandBrowserGetWindowBounds = "Browser.getWindowBounds"

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

const CommandBrowserGetWindowForTarget = "Browser.getWindowForTarget"

// GetWindowForTargetArgs represents the arguments for GetWindowForTarget in the Browser domain.
type GetWindowForTargetArgs struct {
	TargetID target.ID `json:"targetId"` // Devtools agent host id.
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

const CommandBrowserSetWindowBounds = "Browser.setWindowBounds"

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
