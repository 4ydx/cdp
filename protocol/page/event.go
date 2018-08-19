// Code generated by cdpgen. DO NOT EDIT.

package page

import (
	"encoding/json"

	"github.com/4ydx/cdp/protocol"
	"github.com/4ydx/cdp/protocol/network"
	"github.com/4ydx/cdp/protocol/runtime"
)

const EventPageDomContentEventFired = "Page.domContentEventFired"

// DOMContentEventFiredReply is the reply for DOMContentEventFired events.
type DOMContentEventFiredReply struct {
	Timestamp network.MonotonicTime `json:"timestamp"` // No description.
}

// Unmarshal the byte array into a return value for DOMContentEventFired in the Page domain.
func (a *DOMContentEventFiredReply) UnmarshalJSON(b []byte) error {
	type Copy DOMContentEventFiredReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DOMContentEventFiredReply(*c)
	return nil
}

const EventPageFrameAttached = "Page.frameAttached"

// FrameAttachedReply is the reply for FrameAttached events.
type FrameAttachedReply struct {
	FrameID       shared.FrameID     `json:"frameId"`         // Id of the frame that has been attached.
	ParentFrameID shared.FrameID     `json:"parentFrameId"`   // Parent frame identifier.
	Stack         runtime.StackTrace `json:"stack,omitempty"` // JavaScript stack trace of when frame was attached, only set if frame initiated from script.
}

// Unmarshal the byte array into a return value for FrameAttached in the Page domain.
func (a *FrameAttachedReply) UnmarshalJSON(b []byte) error {
	type Copy FrameAttachedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = FrameAttachedReply(*c)
	return nil
}

const EventPageFrameClearedScheduledNavigation = "Page.frameClearedScheduledNavigation"

// FrameClearedScheduledNavigationReply is the reply for FrameClearedScheduledNavigation events.
type FrameClearedScheduledNavigationReply struct {
	FrameID shared.FrameID `json:"frameId"` // Id of the frame that has cleared its scheduled navigation.
}

// Unmarshal the byte array into a return value for FrameClearedScheduledNavigation in the Page domain.
func (a *FrameClearedScheduledNavigationReply) UnmarshalJSON(b []byte) error {
	type Copy FrameClearedScheduledNavigationReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = FrameClearedScheduledNavigationReply(*c)
	return nil
}

const EventPageFrameDetached = "Page.frameDetached"

// FrameDetachedReply is the reply for FrameDetached events.
type FrameDetachedReply struct {
	FrameID shared.FrameID `json:"frameId"` // Id of the frame that has been detached.
}

// Unmarshal the byte array into a return value for FrameDetached in the Page domain.
func (a *FrameDetachedReply) UnmarshalJSON(b []byte) error {
	type Copy FrameDetachedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = FrameDetachedReply(*c)
	return nil
}

const EventPageFrameNavigated = "Page.frameNavigated"

// FrameNavigatedReply is the reply for FrameNavigated events.
type FrameNavigatedReply struct {
	Frame Frame `json:"frame"` // Frame object.
}

// Unmarshal the byte array into a return value for FrameNavigated in the Page domain.
func (a *FrameNavigatedReply) UnmarshalJSON(b []byte) error {
	type Copy FrameNavigatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = FrameNavigatedReply(*c)
	return nil
}

const EventPageFrameResized = "Page.frameResized"

// FrameResizedReply is the reply for FrameResized events.
type FrameResizedReply struct {
}

// Unmarshal the byte array into a return value for FrameResized in the Page domain.
func (a *FrameResizedReply) UnmarshalJSON(b []byte) error {
	type Copy FrameResizedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = FrameResizedReply(*c)
	return nil
}

const EventPageFrameScheduledNavigation = "Page.frameScheduledNavigation"

// FrameScheduledNavigationReply is the reply for FrameScheduledNavigation events.
type FrameScheduledNavigationReply struct {
	FrameID shared.FrameID `json:"frameId"` // Id of the frame that has scheduled a navigation.
	Delay   float64        `json:"delay"`   // Delay (in seconds) until the navigation is scheduled to begin. The navigation is not guaranteed to start.
	// Reason The reason for the navigation.
	//
	// Values: "formSubmissionGet", "formSubmissionPost", "httpHeaderRefresh", "scriptInitiated", "metaTagRefresh", "pageBlockInterstitial", "reload".
	Reason string `json:"reason"`
	URL    string `json:"url"` // The destination URL for the scheduled navigation.
}

// Unmarshal the byte array into a return value for FrameScheduledNavigation in the Page domain.
func (a *FrameScheduledNavigationReply) UnmarshalJSON(b []byte) error {
	type Copy FrameScheduledNavigationReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = FrameScheduledNavigationReply(*c)
	return nil
}

const EventPageFrameStartedLoading = "Page.frameStartedLoading"

// FrameStartedLoadingReply is the reply for FrameStartedLoading events.
type FrameStartedLoadingReply struct {
	FrameID shared.FrameID `json:"frameId"` // Id of the frame that has started loading.
}

// Unmarshal the byte array into a return value for FrameStartedLoading in the Page domain.
func (a *FrameStartedLoadingReply) UnmarshalJSON(b []byte) error {
	type Copy FrameStartedLoadingReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = FrameStartedLoadingReply(*c)
	return nil
}

const EventPageFrameStoppedLoading = "Page.frameStoppedLoading"

// FrameStoppedLoadingReply is the reply for FrameStoppedLoading events.
type FrameStoppedLoadingReply struct {
	FrameID shared.FrameID `json:"frameId"` // Id of the frame that has stopped loading.
}

// Unmarshal the byte array into a return value for FrameStoppedLoading in the Page domain.
func (a *FrameStoppedLoadingReply) UnmarshalJSON(b []byte) error {
	type Copy FrameStoppedLoadingReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = FrameStoppedLoadingReply(*c)
	return nil
}

const EventPageInterstitialHidden = "Page.interstitialHidden"

// InterstitialHiddenReply is the reply for InterstitialHidden events.
type InterstitialHiddenReply struct {
}

// Unmarshal the byte array into a return value for InterstitialHidden in the Page domain.
func (a *InterstitialHiddenReply) UnmarshalJSON(b []byte) error {
	type Copy InterstitialHiddenReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = InterstitialHiddenReply(*c)
	return nil
}

const EventPageInterstitialShown = "Page.interstitialShown"

// InterstitialShownReply is the reply for InterstitialShown events.
type InterstitialShownReply struct {
}

// Unmarshal the byte array into a return value for InterstitialShown in the Page domain.
func (a *InterstitialShownReply) UnmarshalJSON(b []byte) error {
	type Copy InterstitialShownReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = InterstitialShownReply(*c)
	return nil
}

const EventPageJavascriptDialogClosed = "Page.javascriptDialogClosed"

// JavascriptDialogClosedReply is the reply for JavascriptDialogClosed events.
type JavascriptDialogClosedReply struct {
	Result    bool   `json:"result"`    // Whether dialog was confirmed.
	UserInput string `json:"userInput"` // User input in case of prompt.
}

// Unmarshal the byte array into a return value for JavascriptDialogClosed in the Page domain.
func (a *JavascriptDialogClosedReply) UnmarshalJSON(b []byte) error {
	type Copy JavascriptDialogClosedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = JavascriptDialogClosedReply(*c)
	return nil
}

const EventPageJavascriptDialogOpening = "Page.javascriptDialogOpening"

// JavascriptDialogOpeningReply is the reply for JavascriptDialogOpening events.
type JavascriptDialogOpeningReply struct {
	URL               string     `json:"url"`                     // Frame url.
	Message           string     `json:"message"`                 // Message that will be displayed by the dialog.
	Type              DialogType `json:"type"`                    // Dialog type.
	HasBrowserHandler bool       `json:"hasBrowserHandler"`       // True iff browser is capable showing or acting on the given dialog. When browser has no dialog handler for given target, calling alert while Page domain is engaged will stall the page execution. Execution can be resumed via calling Page.handleJavaScriptDialog.
	DefaultPrompt     string     `json:"defaultPrompt,omitempty"` // Default dialog prompt.
}

// Unmarshal the byte array into a return value for JavascriptDialogOpening in the Page domain.
func (a *JavascriptDialogOpeningReply) UnmarshalJSON(b []byte) error {
	type Copy JavascriptDialogOpeningReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = JavascriptDialogOpeningReply(*c)
	return nil
}

const EventPageLifecycleEvent = "Page.lifecycleEvent"

// LifecycleEventReply is the reply for LifecycleEvent events.
type LifecycleEventReply struct {
	FrameID   shared.FrameID        `json:"frameId"`   // Id of the frame.
	LoaderID  network.LoaderID      `json:"loaderId"`  // Loader identifier. Empty string if the request is fetched from worker.
	Name      string                `json:"name"`      // No description.
	Timestamp network.MonotonicTime `json:"timestamp"` // No description.
}

// Unmarshal the byte array into a return value for LifecycleEvent in the Page domain.
func (a *LifecycleEventReply) UnmarshalJSON(b []byte) error {
	type Copy LifecycleEventReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = LifecycleEventReply(*c)
	return nil
}

const EventPageLoadEventFired = "Page.loadEventFired"

// LoadEventFiredReply is the reply for LoadEventFired events.
type LoadEventFiredReply struct {
	Timestamp network.MonotonicTime `json:"timestamp"` // No description.
}

// Unmarshal the byte array into a return value for LoadEventFired in the Page domain.
func (a *LoadEventFiredReply) UnmarshalJSON(b []byte) error {
	type Copy LoadEventFiredReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = LoadEventFiredReply(*c)
	return nil
}

const EventPageNavigatedWithinDocument = "Page.navigatedWithinDocument"

// NavigatedWithinDocumentReply is the reply for NavigatedWithinDocument events.
type NavigatedWithinDocumentReply struct {
	FrameID shared.FrameID `json:"frameId"` // Id of the frame.
	URL     string         `json:"url"`     // Frame's new url.
}

// Unmarshal the byte array into a return value for NavigatedWithinDocument in the Page domain.
func (a *NavigatedWithinDocumentReply) UnmarshalJSON(b []byte) error {
	type Copy NavigatedWithinDocumentReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = NavigatedWithinDocumentReply(*c)
	return nil
}

const EventPageScreencastFrame = "Page.screencastFrame"

// ScreencastFrameReply is the reply for ScreencastFrame events.
type ScreencastFrameReply struct {
	Data      []byte                  `json:"data"`      // Base64-encoded compressed image.
	Metadata  ScreencastFrameMetadata `json:"metadata"`  // Screencast frame metadata.
	SessionID int                     `json:"sessionId"` // Frame number.
}

// Unmarshal the byte array into a return value for ScreencastFrame in the Page domain.
func (a *ScreencastFrameReply) UnmarshalJSON(b []byte) error {
	type Copy ScreencastFrameReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ScreencastFrameReply(*c)
	return nil
}

const EventPageScreencastVisibilityChanged = "Page.screencastVisibilityChanged"

// ScreencastVisibilityChangedReply is the reply for ScreencastVisibilityChanged events.
type ScreencastVisibilityChangedReply struct {
	Visible bool `json:"visible"` // True if the page is visible.
}

// Unmarshal the byte array into a return value for ScreencastVisibilityChanged in the Page domain.
func (a *ScreencastVisibilityChangedReply) UnmarshalJSON(b []byte) error {
	type Copy ScreencastVisibilityChangedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ScreencastVisibilityChangedReply(*c)
	return nil
}

const EventPageWindowOpen = "Page.windowOpen"

// WindowOpenReply is the reply for WindowOpen events.
type WindowOpenReply struct {
	URL            string   `json:"url"`            // The URL for the new window.
	WindowName     string   `json:"windowName"`     // Window name.
	WindowFeatures []string `json:"windowFeatures"` // An array of enabled window features.
	UserGesture    bool     `json:"userGesture"`    // Whether or not it was triggered by user gesture.
}

// Unmarshal the byte array into a return value for WindowOpen in the Page domain.
func (a *WindowOpenReply) UnmarshalJSON(b []byte) error {
	type Copy WindowOpenReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = WindowOpenReply(*c)
	return nil
}