// Code generated by cdpgen. DO NOT EDIT.

package domstorage

import (
	"encoding/json"
	"log"
)

const (
	CommandDOMStorageClear                = "DOMStorage.clear"
	CommandDOMStorageDisable              = "DOMStorage.disable"
	CommandDOMStorageEnable               = "DOMStorage.enable"
	CommandDOMStorageGetDOMStorageItems   = "DOMStorage.getDOMStorageItems"
	CommandDOMStorageRemoveDOMStorageItem = "DOMStorage.removeDOMStorageItem"
	CommandDOMStorageSetDOMStorageItem    = "DOMStorage.setDOMStorageItem"
)

// ClearArgs represents the arguments for Clear in the DOMStorage domain.
type ClearArgs struct {
	StorageID StorageID `json:"storageId"` // No description.
}

// Unmarshal the byte array into a return value for Clear in the DOMStorage domain.
func (a *ClearArgs) UnmarshalJSON(b []byte) error {
	type Copy ClearArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ClearArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Clear in the DOMStorage domain.
func (a *ClearArgs) MarshalJSON() ([]byte, error) {
	type Copy ClearArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// ClearReply represents the return values for Clear in the DOMStorage domain.
type ClearReply struct {
}

// ClearReply returns whether or not the FrameID matches the reply value for Clear in the DOMStorage domain.
func (a *ClearReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: ClearReply %s", err)
		return false, err
	}
	return true, nil
}

// ClearReply returns the FrameID value for Clear in the DOMStorage domain.
func (a *ClearReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Clear in the DOMStorage domain.
func (a *ClearReply) UnmarshalJSON(b []byte) error {
	type Copy ClearReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ClearReply(*c)
	return nil
}

// DisableArgs represents the arguments for Disable in the DOMStorage domain.
type DisableArgs struct {
}

// Unmarshal the byte array into a return value for Disable in the DOMStorage domain.
func (a *DisableArgs) UnmarshalJSON(b []byte) error {
	type Copy DisableArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DisableArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Disable in the DOMStorage domain.
func (a *DisableArgs) MarshalJSON() ([]byte, error) {
	type Copy DisableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DisableReply represents the return values for Disable in the DOMStorage domain.
type DisableReply struct {
}

// DisableReply returns whether or not the FrameID matches the reply value for Disable in the DOMStorage domain.
func (a *DisableReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: DisableReply %s", err)
		return false, err
	}
	return true, nil
}

// DisableReply returns the FrameID value for Disable in the DOMStorage domain.
func (a *DisableReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Disable in the DOMStorage domain.
func (a *DisableReply) UnmarshalJSON(b []byte) error {
	type Copy DisableReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DisableReply(*c)
	return nil
}

// EnableArgs represents the arguments for Enable in the DOMStorage domain.
type EnableArgs struct {
}

// Unmarshal the byte array into a return value for Enable in the DOMStorage domain.
func (a *EnableArgs) UnmarshalJSON(b []byte) error {
	type Copy EnableArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = EnableArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Enable in the DOMStorage domain.
func (a *EnableArgs) MarshalJSON() ([]byte, error) {
	type Copy EnableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// EnableReply represents the return values for Enable in the DOMStorage domain.
type EnableReply struct {
}

// EnableReply returns whether or not the FrameID matches the reply value for Enable in the DOMStorage domain.
func (a *EnableReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: EnableReply %s", err)
		return false, err
	}
	return true, nil
}

// EnableReply returns the FrameID value for Enable in the DOMStorage domain.
func (a *EnableReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Enable in the DOMStorage domain.
func (a *EnableReply) UnmarshalJSON(b []byte) error {
	type Copy EnableReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = EnableReply(*c)
	return nil
}

// GetDOMStorageItemsArgs represents the arguments for GetDOMStorageItems in the DOMStorage domain.
type GetDOMStorageItemsArgs struct {
	StorageID StorageID `json:"storageId"` // No description.
}

// Unmarshal the byte array into a return value for GetDOMStorageItems in the DOMStorage domain.
func (a *GetDOMStorageItemsArgs) UnmarshalJSON(b []byte) error {
	type Copy GetDOMStorageItemsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetDOMStorageItemsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetDOMStorageItems in the DOMStorage domain.
func (a *GetDOMStorageItemsArgs) MarshalJSON() ([]byte, error) {
	type Copy GetDOMStorageItemsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetDOMStorageItemsReply represents the return values for GetDOMStorageItems in the DOMStorage domain.
type GetDOMStorageItemsReply struct {
	Entries []Item `json:"entries"` // No description.
}

// GetDOMStorageItemsReply returns whether or not the FrameID matches the reply value for GetDOMStorageItems in the DOMStorage domain.
func (a *GetDOMStorageItemsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetDOMStorageItemsReply %s", err)
		return false, err
	}
	return true, nil
}

// GetDOMStorageItemsReply returns the FrameID value for GetDOMStorageItems in the DOMStorage domain.
func (a *GetDOMStorageItemsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetDOMStorageItems in the DOMStorage domain.
func (a *GetDOMStorageItemsReply) UnmarshalJSON(b []byte) error {
	type Copy GetDOMStorageItemsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetDOMStorageItemsReply(*c)
	return nil
}

// RemoveDOMStorageItemArgs represents the arguments for RemoveDOMStorageItem in the DOMStorage domain.
type RemoveDOMStorageItemArgs struct {
	StorageID StorageID `json:"storageId"` // No description.
	Key       string    `json:"key"`       // No description.
}

// Unmarshal the byte array into a return value for RemoveDOMStorageItem in the DOMStorage domain.
func (a *RemoveDOMStorageItemArgs) UnmarshalJSON(b []byte) error {
	type Copy RemoveDOMStorageItemArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RemoveDOMStorageItemArgs(*c)
	return nil
}

// Marshall the byte array into a return value for RemoveDOMStorageItem in the DOMStorage domain.
func (a *RemoveDOMStorageItemArgs) MarshalJSON() ([]byte, error) {
	type Copy RemoveDOMStorageItemArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// RemoveDOMStorageItemReply represents the return values for RemoveDOMStorageItem in the DOMStorage domain.
type RemoveDOMStorageItemReply struct {
}

// RemoveDOMStorageItemReply returns whether or not the FrameID matches the reply value for RemoveDOMStorageItem in the DOMStorage domain.
func (a *RemoveDOMStorageItemReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: RemoveDOMStorageItemReply %s", err)
		return false, err
	}
	return true, nil
}

// RemoveDOMStorageItemReply returns the FrameID value for RemoveDOMStorageItem in the DOMStorage domain.
func (a *RemoveDOMStorageItemReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for RemoveDOMStorageItem in the DOMStorage domain.
func (a *RemoveDOMStorageItemReply) UnmarshalJSON(b []byte) error {
	type Copy RemoveDOMStorageItemReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RemoveDOMStorageItemReply(*c)
	return nil
}

// SetDOMStorageItemArgs represents the arguments for SetDOMStorageItem in the DOMStorage domain.
type SetDOMStorageItemArgs struct {
	StorageID StorageID `json:"storageId"` // No description.
	Key       string    `json:"key"`       // No description.
	Value     string    `json:"value"`     // No description.
}

// Unmarshal the byte array into a return value for SetDOMStorageItem in the DOMStorage domain.
func (a *SetDOMStorageItemArgs) UnmarshalJSON(b []byte) error {
	type Copy SetDOMStorageItemArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetDOMStorageItemArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetDOMStorageItem in the DOMStorage domain.
func (a *SetDOMStorageItemArgs) MarshalJSON() ([]byte, error) {
	type Copy SetDOMStorageItemArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetDOMStorageItemReply represents the return values for SetDOMStorageItem in the DOMStorage domain.
type SetDOMStorageItemReply struct {
}

// SetDOMStorageItemReply returns whether or not the FrameID matches the reply value for SetDOMStorageItem in the DOMStorage domain.
func (a *SetDOMStorageItemReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: SetDOMStorageItemReply %s", err)
		return false, err
	}
	return true, nil
}

// SetDOMStorageItemReply returns the FrameID value for SetDOMStorageItem in the DOMStorage domain.
func (a *SetDOMStorageItemReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for SetDOMStorageItem in the DOMStorage domain.
func (a *SetDOMStorageItemReply) UnmarshalJSON(b []byte) error {
	type Copy SetDOMStorageItemReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetDOMStorageItemReply(*c)
	return nil
}
