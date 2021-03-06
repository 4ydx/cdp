// Code generated by cdpgen. DO NOT EDIT.

package layertree

import (
	"encoding/json"
	"log"

	"github.com/4ydx/cdp/protocol/dom"
)

const (
	EventLayerTreeLayerPainted       = "LayerTree.layerPainted"
	EventLayerTreeLayerTreeDidChange = "LayerTree.layerTreeDidChange"
)

type Unmarshaler func() json.Unmarshaler

var EventConstants = map[string]Unmarshaler{
	EventLayerTreeLayerPainted:       func() json.Unmarshaler { return &LayerPaintedReply{} },
	EventLayerTreeLayerTreeDidChange: func() json.Unmarshaler { return &DidChangeReply{} },
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	if ok {
		return e(), ok
	}
	return nil, ok
}

// LayerPaintedReply is the reply for LayerPainted events.
type LayerPaintedReply struct {
	LayerID LayerID  `json:"layerId"` // The id of the painted layer.
	Clip    dom.Rect `json:"clip"`    // Clip rectangle.
}

// Unmarshal the byte array into a return value for LayerPainted in the LayerTree domain.
func (a *LayerPaintedReply) UnmarshalJSON(b []byte) error {
	type Copy LayerPaintedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = LayerPaintedReply(*c)
	return nil
}

// LayerPaintedReply returns whether or not the FrameID matches the reply value for LayerPainted in the LayerTree domain.
func (a *LayerPaintedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: LayerPaintedReply %s", err)
		return false, err
	}
	return true, nil
}

// LayerPaintedReply returns the FrameID for LayerPainted in the LayerTree domain.
func (a *LayerPaintedReply) GetFrameID() string {
	return ""
}

// DidChangeReply is the reply for LayerTreeDidChange events.
type DidChangeReply struct {
	Layers *[]Layer `json:"layers,omitempty"` // Layer tree, absent if not in the comspositing mode.
}

// Unmarshal the byte array into a return value for LayerTreeDidChange in the LayerTree domain.
func (a *DidChangeReply) UnmarshalJSON(b []byte) error {
	type Copy DidChangeReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DidChangeReply(*c)
	return nil
}

// DidChangeReply returns whether or not the FrameID matches the reply value for LayerTreeDidChange in the LayerTree domain.
func (a *DidChangeReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: DidChangeReply %s", err)
		return false, err
	}
	return true, nil
}

// DidChangeReply returns the FrameID for LayerTreeDidChange in the LayerTree domain.
func (a *DidChangeReply) GetFrameID() string {
	return ""
}
