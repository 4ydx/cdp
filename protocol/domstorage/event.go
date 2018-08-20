// Code generated by cdpgen. DO NOT EDIT.

package domstorage

import (
	"encoding/json"
	"log"
)

const (
	EventDOMStorageDomStorageItemAdded    = "DOMStorage.domStorageItemAdded"
	EventDOMStorageDomStorageItemRemoved  = "DOMStorage.domStorageItemRemoved"
	EventDOMStorageDomStorageItemUpdated  = "DOMStorage.domStorageItemUpdated"
	EventDOMStorageDomStorageItemsCleared = "DOMStorage.domStorageItemsCleared"
)

var EventConstants = map[string]json.Unmarshaler{
	EventDOMStorageDomStorageItemAdded:    &ItemAddedReply{},
	EventDOMStorageDomStorageItemRemoved:  &ItemRemovedReply{},
	EventDOMStorageDomStorageItemUpdated:  &ItemUpdatedReply{},
	EventDOMStorageDomStorageItemsCleared: &ItemsClearedReply{},
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	return e, ok
}

// ItemAddedReply is the reply for DOMStorageItemAdded events.
type ItemAddedReply struct {
	StorageID StorageID `json:"storageId"` // No description.
	Key       string    `json:"key"`       // No description.
	NewValue  string    `json:"newValue"`  // No description.
}

// Unmarshal the byte array into a return value for DOMStorageItemAdded in the DOMStorage domain.
func (a *ItemAddedReply) UnmarshalJSON(b []byte) error {
	type Copy ItemAddedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ItemAddedReply(*c)
	return nil
}

// ItemAddedReply returns whether or not the FrameID matches the reply value for DOMStorageItemAdded in the DOMStorageItemAdded domain.
func (a *ItemAddedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: %!s(MISSING)", err)
	}
	return true
}

// ItemAddedReply returns the FrameID for DOMStorageItemAdded in the DOMStorageItemAdded domain.
func (a *ItemAddedReply) GetFrameID() string {
	return ""
}

// ItemRemovedReply is the reply for DOMStorageItemRemoved events.
type ItemRemovedReply struct {
	StorageID StorageID `json:"storageId"` // No description.
	Key       string    `json:"key"`       // No description.
}

// Unmarshal the byte array into a return value for DOMStorageItemRemoved in the DOMStorage domain.
func (a *ItemRemovedReply) UnmarshalJSON(b []byte) error {
	type Copy ItemRemovedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ItemRemovedReply(*c)
	return nil
}

// ItemRemovedReply returns whether or not the FrameID matches the reply value for DOMStorageItemRemoved in the DOMStorageItemRemoved domain.
func (a *ItemRemovedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: %!s(MISSING)", err)
	}
	return true
}

// ItemRemovedReply returns the FrameID for DOMStorageItemRemoved in the DOMStorageItemRemoved domain.
func (a *ItemRemovedReply) GetFrameID() string {
	return ""
}

// ItemUpdatedReply is the reply for DOMStorageItemUpdated events.
type ItemUpdatedReply struct {
	StorageID StorageID `json:"storageId"` // No description.
	Key       string    `json:"key"`       // No description.
	OldValue  string    `json:"oldValue"`  // No description.
	NewValue  string    `json:"newValue"`  // No description.
}

// Unmarshal the byte array into a return value for DOMStorageItemUpdated in the DOMStorage domain.
func (a *ItemUpdatedReply) UnmarshalJSON(b []byte) error {
	type Copy ItemUpdatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ItemUpdatedReply(*c)
	return nil
}

// ItemUpdatedReply returns whether or not the FrameID matches the reply value for DOMStorageItemUpdated in the DOMStorageItemUpdated domain.
func (a *ItemUpdatedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: %!s(MISSING)", err)
	}
	return true
}

// ItemUpdatedReply returns the FrameID for DOMStorageItemUpdated in the DOMStorageItemUpdated domain.
func (a *ItemUpdatedReply) GetFrameID() string {
	return ""
}

// ItemsClearedReply is the reply for DOMStorageItemsCleared events.
type ItemsClearedReply struct {
	StorageID StorageID `json:"storageId"` // No description.
}

// Unmarshal the byte array into a return value for DOMStorageItemsCleared in the DOMStorage domain.
func (a *ItemsClearedReply) UnmarshalJSON(b []byte) error {
	type Copy ItemsClearedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ItemsClearedReply(*c)
	return nil
}

// ItemsClearedReply returns whether or not the FrameID matches the reply value for DOMStorageItemsCleared in the DOMStorageItemsCleared domain.
func (a *ItemsClearedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: %!s(MISSING)", err)
	}
	return true
}

// ItemsClearedReply returns the FrameID for DOMStorageItemsCleared in the DOMStorageItemsCleared domain.
func (a *ItemsClearedReply) GetFrameID() string {
	return ""
}
