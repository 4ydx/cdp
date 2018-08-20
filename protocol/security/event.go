// Code generated by cdpgen. DO NOT EDIT.

package security

import (
	"encoding/json"
	"log"
)

const (
	EventSecurityCertificateError     = "Security.certificateError"
	EventSecuritySecurityStateChanged = "Security.securityStateChanged"
)

var EventConstants = map[string]json.Unmarshaler{
	EventSecurityCertificateError:     &CertificateErrorReply{},
	EventSecuritySecurityStateChanged: &StateChangedReply{},
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	return e, ok
}

// CertificateErrorReply is the reply for CertificateError events.
type CertificateErrorReply struct {
	EventID    int    `json:"eventId"`    // The ID of the event.
	ErrorType  string `json:"errorType"`  // The type of the error.
	RequestURL string `json:"requestURL"` // The url that was requested.
}

// Unmarshal the byte array into a return value for CertificateError in the Security domain.
func (a *CertificateErrorReply) UnmarshalJSON(b []byte) error {
	type Copy CertificateErrorReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CertificateErrorReply(*c)
	return nil
}

// CertificateErrorReply returns whether or not the FrameID matches the reply value for CertificateError in the CertificateError domain.
func (a *CertificateErrorReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: %!s(MISSING)", err)
	}
	return true
}

// CertificateErrorReply returns the FrameID for CertificateError in the CertificateError domain.
func (a *CertificateErrorReply) GetFrameID() string {
	return ""
}

// StateChangedReply is the reply for SecurityStateChanged events.
type StateChangedReply struct {
	SecurityState         State                 `json:"securityState"`         // Security state.
	SchemeIsCryptographic bool                  `json:"schemeIsCryptographic"` // True if the page was loaded over cryptographic transport such as HTTPS.
	Explanations          []StateExplanation    `json:"explanations"`          // List of explanations for the security state. If the overall security state is `insecure` or `warning`, at least one corresponding explanation should be included.
	InsecureContentStatus InsecureContentStatus `json:"insecureContentStatus"` // Information about insecure content on the page.
	Summary               string                `json:"summary,omitempty"`     // Overrides user-visible description of the state.
}

// Unmarshal the byte array into a return value for SecurityStateChanged in the Security domain.
func (a *StateChangedReply) UnmarshalJSON(b []byte) error {
	type Copy StateChangedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StateChangedReply(*c)
	return nil
}

// StateChangedReply returns whether or not the FrameID matches the reply value for SecurityStateChanged in the SecurityStateChanged domain.
func (a *StateChangedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: %!s(MISSING)", err)
	}
	return true
}

// StateChangedReply returns the FrameID for SecurityStateChanged in the SecurityStateChanged domain.
func (a *StateChangedReply) GetFrameID() string {
	return ""
}
