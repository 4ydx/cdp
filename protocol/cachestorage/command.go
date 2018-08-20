// Code generated by cdpgen. DO NOT EDIT.

package cachestorage

import (
	"encoding/json"
	"log"
)

const (
	CommandCacheStorageDeleteCache           = "CacheStorage.deleteCache"
	CommandCacheStorageDeleteEntry           = "CacheStorage.deleteEntry"
	CommandCacheStorageRequestCacheNames     = "CacheStorage.requestCacheNames"
	CommandCacheStorageRequestCachedResponse = "CacheStorage.requestCachedResponse"
	CommandCacheStorageRequestEntries        = "CacheStorage.requestEntries"
)

// DeleteCacheArgs represents the arguments for DeleteCache in the CacheStorage domain.
type DeleteCacheArgs struct {
	CacheID CacheID `json:"cacheId"` // Id of cache for deletion.
}

// Unmarshal the byte array into a return value for DeleteCache in the CacheStorage domain.
func (a *DeleteCacheArgs) UnmarshalJSON(b []byte) error {
	type Copy DeleteCacheArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DeleteCacheArgs(*c)
	return nil
}

// Marshall the byte array into a return value for DeleteCache in the CacheStorage domain.
func (a *DeleteCacheArgs) MarshalJSON() ([]byte, error) {
	type Copy DeleteCacheArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DeleteCacheReply represents the return values for DeleteCache in the CacheStorage domain.
type DeleteCacheReply struct {
}

// DeleteCacheReply returns whether or not the FrameID matches the reply value for DeleteCache in the CacheStorage domain.
func (a *DeleteCacheReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: DeleteCacheReply", err)
	}
	return true
}

// DeleteCacheReply returns the FrameID value for DeleteCache in the CacheStorage domain.
func (a *DeleteCacheReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for DeleteCache in the CacheStorage domain.
func (a *DeleteCacheReply) UnmarshalJSON(b []byte) error {
	type Copy DeleteCacheReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DeleteCacheReply(*c)
	return nil
}

// DeleteEntryArgs represents the arguments for DeleteEntry in the CacheStorage domain.
type DeleteEntryArgs struct {
	CacheID CacheID `json:"cacheId"` // Id of cache where the entry will be deleted.
	Request string  `json:"request"` // URL spec of the request.
}

// Unmarshal the byte array into a return value for DeleteEntry in the CacheStorage domain.
func (a *DeleteEntryArgs) UnmarshalJSON(b []byte) error {
	type Copy DeleteEntryArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DeleteEntryArgs(*c)
	return nil
}

// Marshall the byte array into a return value for DeleteEntry in the CacheStorage domain.
func (a *DeleteEntryArgs) MarshalJSON() ([]byte, error) {
	type Copy DeleteEntryArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DeleteEntryReply represents the return values for DeleteEntry in the CacheStorage domain.
type DeleteEntryReply struct {
}

// DeleteEntryReply returns whether or not the FrameID matches the reply value for DeleteEntry in the CacheStorage domain.
func (a *DeleteEntryReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: DeleteEntryReply", err)
	}
	return true
}

// DeleteEntryReply returns the FrameID value for DeleteEntry in the CacheStorage domain.
func (a *DeleteEntryReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for DeleteEntry in the CacheStorage domain.
func (a *DeleteEntryReply) UnmarshalJSON(b []byte) error {
	type Copy DeleteEntryReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DeleteEntryReply(*c)
	return nil
}

// RequestCacheNamesArgs represents the arguments for RequestCacheNames in the CacheStorage domain.
type RequestCacheNamesArgs struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin.
}

// Unmarshal the byte array into a return value for RequestCacheNames in the CacheStorage domain.
func (a *RequestCacheNamesArgs) UnmarshalJSON(b []byte) error {
	type Copy RequestCacheNamesArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestCacheNamesArgs(*c)
	return nil
}

// Marshall the byte array into a return value for RequestCacheNames in the CacheStorage domain.
func (a *RequestCacheNamesArgs) MarshalJSON() ([]byte, error) {
	type Copy RequestCacheNamesArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// RequestCacheNamesReply represents the return values for RequestCacheNames in the CacheStorage domain.
type RequestCacheNamesReply struct {
	Caches []Cache `json:"caches"` // Caches for the security origin.
}

// RequestCacheNamesReply returns whether or not the FrameID matches the reply value for RequestCacheNames in the CacheStorage domain.
func (a *RequestCacheNamesReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: RequestCacheNamesReply", err)
	}
	return true
}

// RequestCacheNamesReply returns the FrameID value for RequestCacheNames in the CacheStorage domain.
func (a *RequestCacheNamesReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for RequestCacheNames in the CacheStorage domain.
func (a *RequestCacheNamesReply) UnmarshalJSON(b []byte) error {
	type Copy RequestCacheNamesReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestCacheNamesReply(*c)
	return nil
}

// RequestCachedResponseArgs represents the arguments for RequestCachedResponse in the CacheStorage domain.
type RequestCachedResponseArgs struct {
	CacheID    CacheID `json:"cacheId"`    // Id of cache that contains the enty.
	RequestURL string  `json:"requestURL"` // URL spec of the request.
}

// Unmarshal the byte array into a return value for RequestCachedResponse in the CacheStorage domain.
func (a *RequestCachedResponseArgs) UnmarshalJSON(b []byte) error {
	type Copy RequestCachedResponseArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestCachedResponseArgs(*c)
	return nil
}

// Marshall the byte array into a return value for RequestCachedResponse in the CacheStorage domain.
func (a *RequestCachedResponseArgs) MarshalJSON() ([]byte, error) {
	type Copy RequestCachedResponseArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// RequestCachedResponseReply represents the return values for RequestCachedResponse in the CacheStorage domain.
type RequestCachedResponseReply struct {
	Response CachedResponse `json:"response"` // Response read from the cache.
}

// RequestCachedResponseReply returns whether or not the FrameID matches the reply value for RequestCachedResponse in the CacheStorage domain.
func (a *RequestCachedResponseReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: RequestCachedResponseReply", err)
	}
	return true
}

// RequestCachedResponseReply returns the FrameID value for RequestCachedResponse in the CacheStorage domain.
func (a *RequestCachedResponseReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for RequestCachedResponse in the CacheStorage domain.
func (a *RequestCachedResponseReply) UnmarshalJSON(b []byte) error {
	type Copy RequestCachedResponseReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestCachedResponseReply(*c)
	return nil
}

// RequestEntriesArgs represents the arguments for RequestEntries in the CacheStorage domain.
type RequestEntriesArgs struct {
	CacheID   CacheID `json:"cacheId"`   // ID of cache to get entries from.
	SkipCount int     `json:"skipCount"` // Number of records to skip.
	PageSize  int     `json:"pageSize"`  // Number of records to fetch.
}

// Unmarshal the byte array into a return value for RequestEntries in the CacheStorage domain.
func (a *RequestEntriesArgs) UnmarshalJSON(b []byte) error {
	type Copy RequestEntriesArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestEntriesArgs(*c)
	return nil
}

// Marshall the byte array into a return value for RequestEntries in the CacheStorage domain.
func (a *RequestEntriesArgs) MarshalJSON() ([]byte, error) {
	type Copy RequestEntriesArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// RequestEntriesReply represents the return values for RequestEntries in the CacheStorage domain.
type RequestEntriesReply struct {
	CacheDataEntries []DataEntry `json:"cacheDataEntries"` // Array of object store data entries.
	HasMore          bool        `json:"hasMore"`          // If true, there are more entries to fetch in the given range.
}

// RequestEntriesReply returns whether or not the FrameID matches the reply value for RequestEntries in the CacheStorage domain.
func (a *RequestEntriesReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: RequestEntriesReply", err)
	}
	return true
}

// RequestEntriesReply returns the FrameID value for RequestEntries in the CacheStorage domain.
func (a *RequestEntriesReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for RequestEntries in the CacheStorage domain.
func (a *RequestEntriesReply) UnmarshalJSON(b []byte) error {
	type Copy RequestEntriesReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestEntriesReply(*c)
	return nil
}
