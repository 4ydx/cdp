// Code generated by cdpgen. DO NOT EDIT.

package console

// Message Console message.
type Message struct {
	// Source Message source.
	//
	// Values: "xml", "javascript", "network", "console-api", "storage", "appcache", "rendering", "security", "other", "deprecation", "worker".
	Source string `json:"source"`
	// Level Message severity.
	//
	// Values: "log", "warning", "error", "debug", "info".
	Level  string `json:"level"`
	Text   string `json:"text"`             // Message text.
	URL    string `json:"url,omitempty"`    // URL of the message origin.
	Line   int    `json:"line,omitempty"`   // Line number in the resource that generated this message (1-based).
	Column int    `json:"column,omitempty"` // Column number in the resource that generated this message (1-based).
}
