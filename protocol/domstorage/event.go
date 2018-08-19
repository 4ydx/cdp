// Code generated by cdpgen. DO NOT EDIT.

package domstorage

import (
	"encoding/json"
)

const EventDOMStorageDomStorageItemAdded = "DOMStorage.domStorageItemAdded"

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

const EventDOMStorageDomStorageItemRemoved = "DOMStorage.domStorageItemRemoved"

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

const EventDOMStorageDomStorageItemUpdated = "DOMStorage.domStorageItemUpdated"

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

const EventDOMStorageDomStorageItemsCleared = "DOMStorage.domStorageItemsCleared"

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