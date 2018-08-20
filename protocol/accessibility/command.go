// Code generated by cdpgen. DO NOT EDIT.

package accessibility

import (
	"encoding/json"
	"log"

	"github.com/4ydx/cdp/protocol"
	"github.com/4ydx/cdp/protocol/dom"
)

const (
	CommandAccessibilityGetPartialAXTree = "Accessibility.getPartialAXTree"
)

// GetPartialAXTreeArgs represents the arguments for GetPartialAXTree in the Accessibility domain.
type GetPartialAXTreeArgs struct {
	NodeID         dom.NodeID            `json:"nodeId,omitempty"`         // Identifier of the node to get the partial accessibility tree for.
	BackendNodeID  dom.BackendNodeID     `json:"backendNodeId,omitempty"`  // Identifier of the backend node to get the partial accessibility tree for.
	ObjectID       shared.RemoteObjectID `json:"objectId,omitempty"`       // JavaScript object id of the node wrapper to get the partial accessibility tree for.
	FetchRelatives bool                  `json:"fetchRelatives,omitempty"` // Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
}

// Unmarshal the byte array into a return value for GetPartialAXTree in the Accessibility domain.
func (a *GetPartialAXTreeArgs) UnmarshalJSON(b []byte) error {
	type Copy GetPartialAXTreeArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetPartialAXTreeArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetPartialAXTree in the Accessibility domain.
func (a *GetPartialAXTreeArgs) MarshalJSON() ([]byte, error) {
	type Copy GetPartialAXTreeArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetPartialAXTreeReply represents the return values for GetPartialAXTree in the Accessibility domain.
type GetPartialAXTreeReply struct {
	Nodes []AXNode `json:"nodes"` // The `Accessibility.AXNode` for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
}

// GetPartialAXTreeReply returns whether or not the FrameID matches the reply value for GetPartialAXTree in the Accessibility domain.
func (a *GetPartialAXTreeReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetPartialAXTreeReply", err)
	}
	return true
}

// GetPartialAXTreeReply returns the FrameID value for GetPartialAXTree in the Accessibility domain.
func (a *GetPartialAXTreeReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetPartialAXTree in the Accessibility domain.
func (a *GetPartialAXTreeReply) UnmarshalJSON(b []byte) error {
	type Copy GetPartialAXTreeReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetPartialAXTreeReply(*c)
	return nil
}
