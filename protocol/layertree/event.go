// Code generated by cdpgen. DO NOT EDIT.

package layertree

import (
	"encoding/json"

	"github.com/4ydx/cdp/protocol/dom"
)

const (
	EventLayerTreeLayerPainted       = "LayerTree.layerPainted"
	EventLayerTreeLayerTreeDidChange = "LayerTree.layerTreeDidChange"
)

var EventConstants = map[string]json.Unmarshaler{
	EventLayerTreeLayerPainted:       &LayerPaintedReply{},
	EventLayerTreeLayerTreeDidChange: &DidChangeReply{},
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	return e, ok
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

// DidChangeReply is the reply for LayerTreeDidChange events.
type DidChangeReply struct {
	Layers []Layer `json:"layers,omitempty"` // Layer tree, absent if not in the comspositing mode.
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
