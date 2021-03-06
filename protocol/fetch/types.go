// Code generated by cdpgen. DO NOT EDIT.

package fetch

import shared "github.com/4ydx/cdp/protocol"

// RequestID Unique request identifier.
type RequestID string

// RequestStage Stages of the request to handle. Request will intercept before
// the request is sent. Response will intercept after the response is received
// (but before response body is received.
//
// Note: This type is experimental.
type RequestStage string

// RequestStage as enums.
const (
	RequestStageNotSet   RequestStage = ""
	RequestStageRequest  RequestStage = "Request"
	RequestStageResponse RequestStage = "Response"
)

func (e RequestStage) Valid() bool {
	switch e {
	case "Request", "Response":
		return true
	default:
		return false
	}
}

func (e RequestStage) String() string {
	return string(e)
}

// RequestPattern
//
// Note: This type is experimental.
type RequestPattern struct {
	URLPattern   string               `json:"urlPattern,omitempty"`   // Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed. Escape character is backslash. Omitting is equivalent to "*".
	ResourceType *shared.ResourceType `json:"resourceType,omitempty"` // If set, only requests for matching resource types will be intercepted.
	RequestStage *RequestStage        `json:"requestStage,omitempty"` // Stage at which to begin intercepting requests. Default is Request.
}

// HeaderEntry Response HTTP header entry
type HeaderEntry struct {
	Name  string `json:"name"`  // No description.
	Value string `json:"value"` // No description.
}

// AuthChallenge Authorization challenge for HTTP status code 401 or 407.
//
// Note: This type is experimental.
type AuthChallenge struct {
	// Source Source of the authentication challenge.
	//
	// Values: "Server", "Proxy".
	Source string `json:"source,omitempty"`
	Origin string `json:"origin"` // Origin of the challenger.
	Scheme string `json:"scheme"` // The authentication scheme used, such as basic or digest
	Realm  string `json:"realm"`  // The realm of the challenge. May be empty.
}

// AuthChallengeResponse Response to an AuthChallenge.
//
// Note: This type is experimental.
type AuthChallengeResponse struct {
	// Response The decision on what to do in response to the
	// authorization challenge. Default means deferring to the default
	// behavior of the net stack, which will likely either the Cancel
	// authentication or display a popup dialog box.
	//
	// Values: "Default", "CancelAuth", "ProvideCredentials".
	Response string `json:"response"`
	Username string `json:"username,omitempty"` // The username to provide, possibly empty. Should only be set if response is ProvideCredentials.
	Password string `json:"password,omitempty"` // The password to provide, possibly empty. Should only be set if response is ProvideCredentials.
}
