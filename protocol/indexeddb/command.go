// Code generated by cdpgen. DO NOT EDIT.

package indexeddb

import (
	"encoding/json"
	"log"
)

const (
	CommandIndexedDBClearObjectStore         = "IndexedDB.clearObjectStore"
	CommandIndexedDBDeleteDatabase           = "IndexedDB.deleteDatabase"
	CommandIndexedDBDeleteObjectStoreEntries = "IndexedDB.deleteObjectStoreEntries"
	CommandIndexedDBDisable                  = "IndexedDB.disable"
	CommandIndexedDBEnable                   = "IndexedDB.enable"
	CommandIndexedDBRequestData              = "IndexedDB.requestData"
	CommandIndexedDBRequestDatabase          = "IndexedDB.requestDatabase"
	CommandIndexedDBRequestDatabaseNames     = "IndexedDB.requestDatabaseNames"
)

// ClearObjectStoreArgs represents the arguments for ClearObjectStore in the IndexedDB domain.
type ClearObjectStoreArgs struct {
	SecurityOrigin  string `json:"securityOrigin"`  // Security origin.
	DatabaseName    string `json:"databaseName"`    // Database name.
	ObjectStoreName string `json:"objectStoreName"` // Object store name.
}

// Unmarshal the byte array into a return value for ClearObjectStore in the IndexedDB domain.
func (a *ClearObjectStoreArgs) UnmarshalJSON(b []byte) error {
	type Copy ClearObjectStoreArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ClearObjectStoreArgs(*c)
	return nil
}

// Marshall the byte array into a return value for ClearObjectStore in the IndexedDB domain.
func (a *ClearObjectStoreArgs) MarshalJSON() ([]byte, error) {
	type Copy ClearObjectStoreArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// ClearObjectStoreReply represents the return values for ClearObjectStore in the IndexedDB domain.
type ClearObjectStoreReply struct {
}

// ClearObjectStoreReply returns whether or not the FrameID matches the reply value for ClearObjectStore in the IndexedDB domain.
func (a *ClearObjectStoreReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: ClearObjectStoreReply %s", err)
		return false, err
	}
	return true, nil
}

// ClearObjectStoreReply returns the FrameID value for ClearObjectStore in the IndexedDB domain.
func (a *ClearObjectStoreReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for ClearObjectStore in the IndexedDB domain.
func (a *ClearObjectStoreReply) UnmarshalJSON(b []byte) error {
	type Copy ClearObjectStoreReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ClearObjectStoreReply(*c)
	return nil
}

// DeleteDatabaseArgs represents the arguments for DeleteDatabase in the IndexedDB domain.
type DeleteDatabaseArgs struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin.
	DatabaseName   string `json:"databaseName"`   // Database name.
}

// Unmarshal the byte array into a return value for DeleteDatabase in the IndexedDB domain.
func (a *DeleteDatabaseArgs) UnmarshalJSON(b []byte) error {
	type Copy DeleteDatabaseArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DeleteDatabaseArgs(*c)
	return nil
}

// Marshall the byte array into a return value for DeleteDatabase in the IndexedDB domain.
func (a *DeleteDatabaseArgs) MarshalJSON() ([]byte, error) {
	type Copy DeleteDatabaseArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DeleteDatabaseReply represents the return values for DeleteDatabase in the IndexedDB domain.
type DeleteDatabaseReply struct {
}

// DeleteDatabaseReply returns whether or not the FrameID matches the reply value for DeleteDatabase in the IndexedDB domain.
func (a *DeleteDatabaseReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: DeleteDatabaseReply %s", err)
		return false, err
	}
	return true, nil
}

// DeleteDatabaseReply returns the FrameID value for DeleteDatabase in the IndexedDB domain.
func (a *DeleteDatabaseReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for DeleteDatabase in the IndexedDB domain.
func (a *DeleteDatabaseReply) UnmarshalJSON(b []byte) error {
	type Copy DeleteDatabaseReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DeleteDatabaseReply(*c)
	return nil
}

// DeleteObjectStoreEntriesArgs represents the arguments for DeleteObjectStoreEntries in the IndexedDB domain.
type DeleteObjectStoreEntriesArgs struct {
	SecurityOrigin  string   `json:"securityOrigin"`  // No description.
	DatabaseName    string   `json:"databaseName"`    // No description.
	ObjectStoreName string   `json:"objectStoreName"` // No description.
	KeyRange        KeyRange `json:"keyRange"`        // Range of entry keys to delete
}

// Unmarshal the byte array into a return value for DeleteObjectStoreEntries in the IndexedDB domain.
func (a *DeleteObjectStoreEntriesArgs) UnmarshalJSON(b []byte) error {
	type Copy DeleteObjectStoreEntriesArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DeleteObjectStoreEntriesArgs(*c)
	return nil
}

// Marshall the byte array into a return value for DeleteObjectStoreEntries in the IndexedDB domain.
func (a *DeleteObjectStoreEntriesArgs) MarshalJSON() ([]byte, error) {
	type Copy DeleteObjectStoreEntriesArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DeleteObjectStoreEntriesReply represents the return values for DeleteObjectStoreEntries in the IndexedDB domain.
type DeleteObjectStoreEntriesReply struct {
}

// DeleteObjectStoreEntriesReply returns whether or not the FrameID matches the reply value for DeleteObjectStoreEntries in the IndexedDB domain.
func (a *DeleteObjectStoreEntriesReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: DeleteObjectStoreEntriesReply %s", err)
		return false, err
	}
	return true, nil
}

// DeleteObjectStoreEntriesReply returns the FrameID value for DeleteObjectStoreEntries in the IndexedDB domain.
func (a *DeleteObjectStoreEntriesReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for DeleteObjectStoreEntries in the IndexedDB domain.
func (a *DeleteObjectStoreEntriesReply) UnmarshalJSON(b []byte) error {
	type Copy DeleteObjectStoreEntriesReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DeleteObjectStoreEntriesReply(*c)
	return nil
}

// DisableArgs represents the arguments for Disable in the IndexedDB domain.
type DisableArgs struct {
}

// Unmarshal the byte array into a return value for Disable in the IndexedDB domain.
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

// Marshall the byte array into a return value for Disable in the IndexedDB domain.
func (a *DisableArgs) MarshalJSON() ([]byte, error) {
	type Copy DisableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DisableReply represents the return values for Disable in the IndexedDB domain.
type DisableReply struct {
}

// DisableReply returns whether or not the FrameID matches the reply value for Disable in the IndexedDB domain.
func (a *DisableReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: DisableReply %s", err)
		return false, err
	}
	return true, nil
}

// DisableReply returns the FrameID value for Disable in the IndexedDB domain.
func (a *DisableReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Disable in the IndexedDB domain.
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

// EnableArgs represents the arguments for Enable in the IndexedDB domain.
type EnableArgs struct {
}

// Unmarshal the byte array into a return value for Enable in the IndexedDB domain.
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

// Marshall the byte array into a return value for Enable in the IndexedDB domain.
func (a *EnableArgs) MarshalJSON() ([]byte, error) {
	type Copy EnableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// EnableReply represents the return values for Enable in the IndexedDB domain.
type EnableReply struct {
}

// EnableReply returns whether or not the FrameID matches the reply value for Enable in the IndexedDB domain.
func (a *EnableReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: EnableReply %s", err)
		return false, err
	}
	return true, nil
}

// EnableReply returns the FrameID value for Enable in the IndexedDB domain.
func (a *EnableReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Enable in the IndexedDB domain.
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

// RequestDataArgs represents the arguments for RequestData in the IndexedDB domain.
type RequestDataArgs struct {
	SecurityOrigin  string    `json:"securityOrigin"`     // Security origin.
	DatabaseName    string    `json:"databaseName"`       // Database name.
	ObjectStoreName string    `json:"objectStoreName"`    // Object store name.
	IndexName       string    `json:"indexName"`          // Index name, empty string for object store data requests.
	SkipCount       int       `json:"skipCount"`          // Number of records to skip.
	PageSize        int       `json:"pageSize"`           // Number of records to fetch.
	KeyRange        *KeyRange `json:"keyRange,omitempty"` // Key range.
}

// Unmarshal the byte array into a return value for RequestData in the IndexedDB domain.
func (a *RequestDataArgs) UnmarshalJSON(b []byte) error {
	type Copy RequestDataArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestDataArgs(*c)
	return nil
}

// Marshall the byte array into a return value for RequestData in the IndexedDB domain.
func (a *RequestDataArgs) MarshalJSON() ([]byte, error) {
	type Copy RequestDataArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// RequestDataReply represents the return values for RequestData in the IndexedDB domain.
type RequestDataReply struct {
	ObjectStoreDataEntries []DataEntry `json:"objectStoreDataEntries"` // Array of object store data entries.
	HasMore                bool        `json:"hasMore"`                // If true, there are more entries to fetch in the given range.
}

// RequestDataReply returns whether or not the FrameID matches the reply value for RequestData in the IndexedDB domain.
func (a *RequestDataReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: RequestDataReply %s", err)
		return false, err
	}
	return true, nil
}

// RequestDataReply returns the FrameID value for RequestData in the IndexedDB domain.
func (a *RequestDataReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for RequestData in the IndexedDB domain.
func (a *RequestDataReply) UnmarshalJSON(b []byte) error {
	type Copy RequestDataReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestDataReply(*c)
	return nil
}

// RequestDatabaseArgs represents the arguments for RequestDatabase in the IndexedDB domain.
type RequestDatabaseArgs struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin.
	DatabaseName   string `json:"databaseName"`   // Database name.
}

// Unmarshal the byte array into a return value for RequestDatabase in the IndexedDB domain.
func (a *RequestDatabaseArgs) UnmarshalJSON(b []byte) error {
	type Copy RequestDatabaseArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestDatabaseArgs(*c)
	return nil
}

// Marshall the byte array into a return value for RequestDatabase in the IndexedDB domain.
func (a *RequestDatabaseArgs) MarshalJSON() ([]byte, error) {
	type Copy RequestDatabaseArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// RequestDatabaseReply represents the return values for RequestDatabase in the IndexedDB domain.
type RequestDatabaseReply struct {
	DatabaseWithObjectStores DatabaseWithObjectStores `json:"databaseWithObjectStores"` // Database with an array of object stores.
}

// RequestDatabaseReply returns whether or not the FrameID matches the reply value for RequestDatabase in the IndexedDB domain.
func (a *RequestDatabaseReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: RequestDatabaseReply %s", err)
		return false, err
	}
	return true, nil
}

// RequestDatabaseReply returns the FrameID value for RequestDatabase in the IndexedDB domain.
func (a *RequestDatabaseReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for RequestDatabase in the IndexedDB domain.
func (a *RequestDatabaseReply) UnmarshalJSON(b []byte) error {
	type Copy RequestDatabaseReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestDatabaseReply(*c)
	return nil
}

// RequestDatabaseNamesArgs represents the arguments for RequestDatabaseNames in the IndexedDB domain.
type RequestDatabaseNamesArgs struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin.
}

// Unmarshal the byte array into a return value for RequestDatabaseNames in the IndexedDB domain.
func (a *RequestDatabaseNamesArgs) UnmarshalJSON(b []byte) error {
	type Copy RequestDatabaseNamesArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestDatabaseNamesArgs(*c)
	return nil
}

// Marshall the byte array into a return value for RequestDatabaseNames in the IndexedDB domain.
func (a *RequestDatabaseNamesArgs) MarshalJSON() ([]byte, error) {
	type Copy RequestDatabaseNamesArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// RequestDatabaseNamesReply represents the return values for RequestDatabaseNames in the IndexedDB domain.
type RequestDatabaseNamesReply struct {
	DatabaseNames []string `json:"databaseNames"` // Database names for origin.
}

// RequestDatabaseNamesReply returns whether or not the FrameID matches the reply value for RequestDatabaseNames in the IndexedDB domain.
func (a *RequestDatabaseNamesReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: RequestDatabaseNamesReply %s", err)
		return false, err
	}
	return true, nil
}

// RequestDatabaseNamesReply returns the FrameID value for RequestDatabaseNames in the IndexedDB domain.
func (a *RequestDatabaseNamesReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for RequestDatabaseNames in the IndexedDB domain.
func (a *RequestDatabaseNamesReply) UnmarshalJSON(b []byte) error {
	type Copy RequestDatabaseNamesReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RequestDatabaseNamesReply(*c)
	return nil
}
