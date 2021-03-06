// Code generated by cdpgen. DO NOT EDIT.

package storage

import (
	"encoding/json"
	"log"
)

const (
	EventStorageCacheStorageContentUpdated = "Storage.cacheStorageContentUpdated"
	EventStorageCacheStorageListUpdated    = "Storage.cacheStorageListUpdated"
	EventStorageIndexedDBContentUpdated    = "Storage.indexedDBContentUpdated"
	EventStorageIndexedDBListUpdated       = "Storage.indexedDBListUpdated"
)

type Unmarshaler func() json.Unmarshaler

var EventConstants = map[string]Unmarshaler{
	EventStorageCacheStorageContentUpdated: func() json.Unmarshaler { return &CacheStorageContentUpdatedReply{} },
	EventStorageCacheStorageListUpdated:    func() json.Unmarshaler { return &CacheStorageListUpdatedReply{} },
	EventStorageIndexedDBContentUpdated:    func() json.Unmarshaler { return &IndexedDBContentUpdatedReply{} },
	EventStorageIndexedDBListUpdated:       func() json.Unmarshaler { return &IndexedDBListUpdatedReply{} },
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	if ok {
		return e(), ok
	}
	return nil, ok
}

// CacheStorageContentUpdatedReply is the reply for CacheStorageContentUpdated events.
type CacheStorageContentUpdatedReply struct {
	Origin    string `json:"origin"`    // Origin to update.
	CacheName string `json:"cacheName"` // Name of cache in origin.
}

// Unmarshal the byte array into a return value for CacheStorageContentUpdated in the Storage domain.
func (a *CacheStorageContentUpdatedReply) UnmarshalJSON(b []byte) error {
	type Copy CacheStorageContentUpdatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CacheStorageContentUpdatedReply(*c)
	return nil
}

// CacheStorageContentUpdatedReply returns whether or not the FrameID matches the reply value for CacheStorageContentUpdated in the Storage domain.
func (a *CacheStorageContentUpdatedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: CacheStorageContentUpdatedReply %s", err)
		return false, err
	}
	return true, nil
}

// CacheStorageContentUpdatedReply returns the FrameID for CacheStorageContentUpdated in the Storage domain.
func (a *CacheStorageContentUpdatedReply) GetFrameID() string {
	return ""
}

// CacheStorageListUpdatedReply is the reply for CacheStorageListUpdated events.
type CacheStorageListUpdatedReply struct {
	Origin string `json:"origin"` // Origin to update.
}

// Unmarshal the byte array into a return value for CacheStorageListUpdated in the Storage domain.
func (a *CacheStorageListUpdatedReply) UnmarshalJSON(b []byte) error {
	type Copy CacheStorageListUpdatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CacheStorageListUpdatedReply(*c)
	return nil
}

// CacheStorageListUpdatedReply returns whether or not the FrameID matches the reply value for CacheStorageListUpdated in the Storage domain.
func (a *CacheStorageListUpdatedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: CacheStorageListUpdatedReply %s", err)
		return false, err
	}
	return true, nil
}

// CacheStorageListUpdatedReply returns the FrameID for CacheStorageListUpdated in the Storage domain.
func (a *CacheStorageListUpdatedReply) GetFrameID() string {
	return ""
}

// IndexedDBContentUpdatedReply is the reply for IndexedDBContentUpdated events.
type IndexedDBContentUpdatedReply struct {
	Origin          string `json:"origin"`          // Origin to update.
	DatabaseName    string `json:"databaseName"`    // Database to update.
	ObjectStoreName string `json:"objectStoreName"` // ObjectStore to update.
}

// Unmarshal the byte array into a return value for IndexedDBContentUpdated in the Storage domain.
func (a *IndexedDBContentUpdatedReply) UnmarshalJSON(b []byte) error {
	type Copy IndexedDBContentUpdatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = IndexedDBContentUpdatedReply(*c)
	return nil
}

// IndexedDBContentUpdatedReply returns whether or not the FrameID matches the reply value for IndexedDBContentUpdated in the Storage domain.
func (a *IndexedDBContentUpdatedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: IndexedDBContentUpdatedReply %s", err)
		return false, err
	}
	return true, nil
}

// IndexedDBContentUpdatedReply returns the FrameID for IndexedDBContentUpdated in the Storage domain.
func (a *IndexedDBContentUpdatedReply) GetFrameID() string {
	return ""
}

// IndexedDBListUpdatedReply is the reply for IndexedDBListUpdated events.
type IndexedDBListUpdatedReply struct {
	Origin string `json:"origin"` // Origin to update.
}

// Unmarshal the byte array into a return value for IndexedDBListUpdated in the Storage domain.
func (a *IndexedDBListUpdatedReply) UnmarshalJSON(b []byte) error {
	type Copy IndexedDBListUpdatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = IndexedDBListUpdatedReply(*c)
	return nil
}

// IndexedDBListUpdatedReply returns whether or not the FrameID matches the reply value for IndexedDBListUpdated in the Storage domain.
func (a *IndexedDBListUpdatedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: IndexedDBListUpdatedReply %s", err)
		return false, err
	}
	return true, nil
}

// IndexedDBListUpdatedReply returns the FrameID for IndexedDBListUpdated in the Storage domain.
func (a *IndexedDBListUpdatedReply) GetFrameID() string {
	return ""
}
