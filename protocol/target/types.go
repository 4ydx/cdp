// Code generated by cdpgen. DO NOT EDIT.

package target

// ID
type ID string

// SessionID Unique identifier of attached debugging session.
type SessionID string

// BrowserContextID
//
// Note: This type is experimental.
type BrowserContextID string

// Info
type Info struct {
	TargetID ID     `json:"targetId"`           // No description.
	Type     string `json:"type"`               // No description.
	Title    string `json:"title"`              // No description.
	URL      string `json:"url"`                // No description.
	Attached bool   `json:"attached"`           // Whether the target has an attached client.
	OpenerID ID     `json:"openerId,omitempty"` // Opener target Id
	// BrowserContextID
	//
	// Note: This property is experimental.
	BrowserContextID BrowserContextID `json:"browserContextId,omitempty"`
}

// RemoteLocation
//
// Note: This type is experimental.
type RemoteLocation struct {
	Host string `json:"host"` // No description.
	Port int    `json:"port"` // No description.
}
