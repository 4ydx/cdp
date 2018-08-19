// Code generated by cdpgen. DO NOT EDIT.

package profiler

import (
	"encoding/json"

	"github.com/4ydx/cdp/protocol/debugger"
)

const EventProfilerConsoleProfileFinished = "Profiler.consoleProfileFinished"

// ConsoleProfileFinishedReply is the reply for ConsoleProfileFinished events.
type ConsoleProfileFinishedReply struct {
	ID       string            `json:"id"`              // No description.
	Location debugger.Location `json:"location"`        // Location of console.profileEnd().
	Profile  Profile           `json:"profile"`         // No description.
	Title    string            `json:"title,omitempty"` // Profile title passed as an argument to console.profile().
}

// Unmarshal the byte array into a return value for ConsoleProfileFinished in the Profiler domain.
func (a *ConsoleProfileFinishedReply) UnmarshalJSON(b []byte) error {
	type Copy ConsoleProfileFinishedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ConsoleProfileFinishedReply(*c)
	return nil
}

const EventProfilerConsoleProfileStarted = "Profiler.consoleProfileStarted"

// ConsoleProfileStartedReply is the reply for ConsoleProfileStarted events.
type ConsoleProfileStartedReply struct {
	ID       string            `json:"id"`              // No description.
	Location debugger.Location `json:"location"`        // Location of console.profile().
	Title    string            `json:"title,omitempty"` // Profile title passed as an argument to console.profile().
}

// Unmarshal the byte array into a return value for ConsoleProfileStarted in the Profiler domain.
func (a *ConsoleProfileStartedReply) UnmarshalJSON(b []byte) error {
	type Copy ConsoleProfileStartedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ConsoleProfileStartedReply(*c)
	return nil
}